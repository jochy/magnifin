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
                    .mapIndexed((i, e) => TransactionRow(
                          transaction: e,
                          isFirst: i == 0,
                        ))
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
    required this.isFirst,
  });

  final Transaction transaction;
  final bool isFirst;

  @override
  Widget build(BuildContext context) {
    var accountLogo = ConnectionsCubit.of(context)
        .state
        .accountLogoById(transaction.accountId);

    return Container(
      decoration: BoxDecoration(
        border: Border(
          top: BorderSide(
            color: isFirst
                ? Colors.transparent
                : Theme.of(context).colorScheme.onSurface.withOpacity(0.1),
          ),
        ),
      ),
      child: Row(
        children: [
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: ClipOval(
              child: CircleAvatar(
                maxRadius: 18,
                backgroundColor: transaction.direction == 'CREDIT'
                    ? getGreenColorScheme(context).onPrimaryContainer
                    : getGreenColorScheme(context).onErrorContainer,
                child: transaction.counterpartyLogoUrl != null
                    ? Image.network(transaction.counterpartyLogoUrl!)
                    : Icon(
                        transaction.direction == 'CREDIT'
                            ? Icons.arrow_downward
                            : Icons.arrow_upward,
                        color: transaction.direction == 'CREDIT'
                            ? getGreenColorScheme(context).onPrimary
                            : getGreenColorScheme(context).onError,
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
                  Text(
                    "${(transaction.method ?? S.of(context).other).capitalize()} - ${transaction.counterpartyName?.getOrNull() ?? transaction.reference?.getOrNull() ?? S.of(context).transaction}",
                    style: const TextStyle(fontSize: 12),
                  ),
                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 4.0),
                    child: Text(
                      transaction.reference?.getOrNull() ?? '',
                      style: TextStyle(
                          fontSize: 9, color: Theme.of(context).hintColor),
                    ),
                  ),
                  Row(
                    children: [
                      CircleAvatar(
                        maxRadius: 10,
                        child: accountLogo != null
                            ? Image.network(accountLogo)
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
      ),
    );
  }
}

final ColorScheme greenColorScheme =
    ColorScheme.fromSeed(seedColor: Colors.greenAccent);
final ColorScheme greenColorSchemeDark = ColorScheme.fromSeed(
    seedColor: Colors.greenAccent, brightness: Brightness.dark);

ColorScheme getGreenColorScheme(BuildContext context) {
  return greenColorSchemeDark;
}

extension IterableIndexed<E> on Iterable<E> {
  Iterable<T> mapIndexed<T>(T Function(int index, E e) f) sync* {
    var index = 0;
    for (var element in this) {
      yield f(index, element);
      index++;
    }
  }
}

extension StringExt on String {
  String? getOrNull() {
    return isEmpty ? null : this;
  }

  String capitalize() {
    return "${this[0].toUpperCase()}${substring(1).toLowerCase()}";
  }
}
