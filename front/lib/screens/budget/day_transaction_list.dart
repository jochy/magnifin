import 'package:flutter/material.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:front/screens/budget/transaction_edit.dart';
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
                          isLast: i == transactions.length - 1,
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
    required this.isLast,
  });

  final Transaction transaction;
  final bool isFirst;
  final bool isLast;

  @override
  Widget build(BuildContext context) {
    var accountLogo = ConnectionsCubit.of(context)
        .state
        .accountLogoById(transaction.accountId);
    var state = TransactionsCubit.of(context).state;

    return InkWell(
      borderRadius: BorderRadius.vertical(
        top: isFirst ? const Radius.circular(12) : Radius.zero,
        bottom: isLast ? const Radius.circular(12) : Radius.zero,
      ),
      onTap: () => _showDetails(context),
      child: Container(
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
            TransactionIcon(transaction: transaction),
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
                    Wrap(
                      spacing: 8,
                      runSpacing: 4,
                      crossAxisAlignment: WrapCrossAlignment.center,
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
                        StyledText.hintText(
                          ConnectionsCubit.of(context)
                              .state
                              .accountNameById(transaction.accountId),
                          size: 11,
                        ),
                        CategoryChip(transaction: transaction, state: state)
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
      ),
    );
  }

  void _showDetails(BuildContext context) {
    final counterparty = TextEditingController();
    counterparty.value =
        TextEditingValue(text: transaction.counterpartyName ?? '');

    showModalBottomSheet(
      context: context,
      builder: (context) {
        return TransactionEdit(transaction: transaction);
      },
    );
  }

  void updateTransaction(BuildContext context, String counterparty) {
    if (transaction.counterpartyName != counterparty) {
      print(counterparty);
    }
  }
}

class TransactionIcon extends StatelessWidget {
  const TransactionIcon({
    super.key,
    required this.transaction,
  });

  final Transaction transaction;

  @override
  Widget build(BuildContext context) {
    return Padding(
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
    );
  }
}

class CategoryChip extends StatelessWidget {
  CategoryChip({
    super.key,
    required Transaction transaction,
    required TransactionsState state,
  }) {
    category = transaction.getCategory(state);
  }

  CategoryChip.fromCategory({
    super.key,
    required this.category,
  });

  late Category? category;

  @override
  Widget build(BuildContext context) {
    var color = Colors.white;
    if (category != null) {
      color = HexColor.fromHex(category!.color);
    }

    var colorTheme = ColorScheme.fromSeed(
        seedColor: color, brightness: Theme.of(context).brightness);

    return Container(
      decoration: BoxDecoration(
        color: colorTheme.primaryFixed.withAlpha(30),
        border: Border.all(color: colorTheme.primary),
        borderRadius: BorderRadius.circular(100),
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 2, vertical: 1),
        child: Text(
          translateCategory(context, category?.name ?? S.of(context).unknown),
          style: TextStyle(fontSize: 10, color: colorTheme.primary),
        ),
      ),
    );
  }

  String translateCategory(BuildContext context, String name) =>
      switch (name.toLowerCase()) {
        "auto & transport" => S.of(context).autoTransport,
        "subscriptions and bills" => S.of(context).subscriptionsAndBills,
        "cash & checks" => S.of(context).cashChecks,
        "business & work" => S.of(context).businessWork,
        "food & drink" => S.of(context).foodDrink,
        "investment" => S.of(context).investment,
        "health" => S.of(context).health,
        "loan repayment" => S.of(context).loanRepayment,
        "income" => S.of(context).income,
        "taxes" => S.of(context).taxes,
        "transfers" => S.of(context).transfers,
        "essential needs" => S.of(context).essentialNeeds,
        "unknown" => S.of(context).unknown,
        _ => name,
      };
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

extension HexColor on Color {
  static Color fromHex(String hexString) {
    final buffer = StringBuffer();
    if (hexString.length == 6 || hexString.length == 7) buffer.write('ff');
    buffer.write(hexString.replaceFirst('#', ''));
    return Color(int.parse(buffer.toString(), radix: 16));
  }
}
