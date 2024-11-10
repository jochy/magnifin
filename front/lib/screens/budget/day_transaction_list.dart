import 'package:flutter/material.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:moment_dart/moment_dart.dart';
import 'package:money2/money2.dart';

class DayTransactionListWidget extends StatelessWidget {
  const DayTransactionListWidget({
    super.key,
    required this.date,
    required this.transactions,
  });

  final DateTime date;
  final List<Transaction> transactions;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 8.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Padding(
            padding: const EdgeInsets.only(left: 4.0),
            child: StyledText.bodySmall(Moment(date).formatDate(true)),
          ),
          SizedBox(
            width: double.infinity,
            child: Card.outlined(
              child: Column(
                children: transactions
                    .map((e) => TransactionRow(transaction: e))
                    .toList(),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class TransactionRow extends StatelessWidget {
  const TransactionRow({
    super.key,
    required this.transaction,
  });

  final Transaction transaction;

  @override
  Widget build(BuildContext context) {
    var accountLogo = ConnectionsCubit.of(context)
        .state
        .accountLogoById(transaction.accountId);

    return Row(
      children: [
        Padding(
          padding: const EdgeInsets.all(4.0),
          child: ClipOval(
            child: CircleAvatar(
              maxRadius: 18,
              backgroundColor: transaction.direction == 'CREDIT'
                  ? getGreenColorScheme(context).onPrimaryContainer
                  : Theme.of(context).colorScheme.onErrorContainer,
              child: transaction.counterpartyLogoUrl != null
                  ? Image.network(transaction.counterpartyLogoUrl!)
                  : Icon(
                      transaction.direction == 'CREDIT'
                          ? Icons.arrow_downward
                          : Icons.arrow_upward,
                      color: transaction.direction == 'CREDIT'
                          ? getGreenColorScheme(context).onPrimary
                          : Theme.of(context).colorScheme.onError,
                    ),
            ),
          ),
        ),
        Expanded(
          child: Padding(
            padding: const EdgeInsets.symmetric(vertical: 8.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(transaction.counterpartyName ??
                    transaction.reference ??
                    S.of(context).transaction),
                Row(
                  children: [
                    CircleAvatar(
                      maxRadius: 10,
                      child: accountLogo == null
                          ? Image.network(accountLogo!)
                          : const Icon(
                              MdiIcons.bank,
                              size: 10,
                            ),
                    ),
                    const SizedBox(width: 4),
                    StyledText.hintText(
                      ConnectionsCubit.of(context)
                          .state
                          .accountNameById(transaction.accountId),
                      size: 11,
                    ),
                    const SizedBox(width: 12),
                    Container(
                      decoration: BoxDecoration(
                        border: Border.all(
                            color: Theme.of(context).colorScheme.onSurface),
                        borderRadius: BorderRadius.circular(100),
                      ),
                      child: Padding(
                        padding: const EdgeInsets.symmetric(
                            horizontal: 2, vertical: 1),
                        child: Text(
                          "Unknown",
                          style: TextStyle(fontSize: 10),
                        ),
                      ),
                    )
                  ],
                ),
              ],
            ),
          ),
        ),
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: Text(
            Money.fromNum(
                    transaction.amount *
                        (transaction.direction == 'CREDIT' ? 1 : -1),
                    isoCode: transaction.currency)
                .toString(),
            style: TextStyle(
              color: transaction.direction == 'CREDIT'
                  ? getGreenColorScheme(context).primary
                  : Theme.of(context).colorScheme.error,
            ),
          ),
        ),
      ],
    );
  }
}

final ColorScheme greenColorScheme =
    ColorScheme.fromSeed(seedColor: Colors.greenAccent);
final ColorScheme greenColorSchemeDark = ColorScheme.fromSeed(
    seedColor: Colors.greenAccent, brightness: Brightness.dark);

ColorScheme getGreenColorScheme(BuildContext context) {
  return Theme.of(context).brightness == Brightness.light
      ? greenColorScheme
      : greenColorSchemeDark;
}
