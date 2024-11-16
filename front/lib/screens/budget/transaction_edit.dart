import 'package:flutter/material.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:front/screens/budget/create_auto_rule_dialog.dart';
import 'package:front/screens/budget/day_transaction_list.dart';
import 'package:toastification/toastification.dart';

class TransactionEdit extends StatefulWidget {
  const TransactionEdit({
    super.key,
    required this.transaction,
  });

  final Transaction transaction;

  @override
  State<TransactionEdit> createState() => _TransactionEditState();
}

class _TransactionEditState extends State<TransactionEdit> {
  Category? category;
  final counterparty = TextEditingController();
  TransactionsCubit? transactionsCubit;

  @override
  void initState() {
    counterparty.text = widget.transaction.counterpartyName ?? '';
    super.initState();
  }

  @override
  void dispose() {
    updateTransaction();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    category ??=
        widget.transaction.getCategory(TransactionsCubit.of(context).state);
    transactionsCubit ??= TransactionsCubit.of(context);

    return SizedBox(
      width: double.infinity,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Stack(
          children: [
            Align(
              alignment: Alignment.topRight,
              child: IconButton(
                icon: const Icon(Icons.close),
                onPressed: () {
                  updateTransaction();
                  Navigator.of(context).pop();
                },
              ),
            ),
            Column(
              mainAxisSize: MainAxisSize.max,
              children: [
                const SizedBox(width: double.infinity),
                StyledText.headlineSmall(S.of(context).transaction),
                const SizedBox(height: 8),
                Row(
                  children: [
                    TransactionIcon(transaction: widget.transaction),
                    Expanded(
                      child: TextField(
                        controller: counterparty,
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 8),
                Wrap(
                  spacing: 8,
                  runSpacing: 4,
                  alignment: WrapAlignment.center,
                  crossAxisAlignment: WrapCrossAlignment.center,
                  children: [
                    StyledText.labelLarge(S.of(context).category),
                    DropdownButton<Category>(
                        focusColor: Colors.transparent,
                        underline: const SizedBox(),
                        items: TransactionsCubit.of(context)
                            .state
                            .categories
                            .map((e) {
                          return DropdownMenuItem(
                              value: e,
                              child: CategoryChip.fromCategory(category: e));
                        }).toList(),
                        value: category,
                        onChanged: (c) {
                          if (c != category) {
                            askCreateAutoRule(context, c);
                            setState(() {
                              category = c;
                            });
                          }
                        })
                  ],
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  void askCreateAutoRule(BuildContext context, Category? category) async {
    if (category == null) {
      return;
    }

    late BuildContext dialogContext;

    Widget cancelButton = TextButton(
      child: Text(S.of(context).no),
      onPressed: () {
        Navigator.of(dialogContext).pop(false);
      },
    );
    Widget continueButton = FilledButton(
      child: Text(S.of(context).yes),
      onPressed: () {
        Navigator.of(dialogContext).pop(true);
      },
    );
    AlertDialog alert = AlertDialog(
      title: Text(S.of(context).doYouWantToCreateASmartRule),
      content: Text(S
          .of(context)
          .thisWillAutomaticallyCategorizeSimilarTransactionsInTheFuture),
      actions: [
        cancelButton,
        continueButton,
      ],
    );

    final result = await showDialog<bool?>(
      context: context,
      builder: (BuildContext context) {
        dialogContext = context;
        return alert;
      },
    );

    if (result == true && this.context.mounted) {
      showDialog(
          context: this.context,
          builder: (context) {
            return CreateAutoRuleDialog(
              transaction: widget.transaction,
              category: category,
              counterparty: counterparty.text,
            );
          });
    }
  }

  void updateTransaction() {
    try {
      if (widget.transaction.counterpartyName != counterparty.text ||
          widget.transaction.category != category?.id) {
        transactionsCubit!.updateTransaction(
          widget.transaction,
          category,
          counterparty.text.trim(),
        );
      }
    } catch (e) {
      if (context.mounted) {
        toastification.show(
          context: context,
          title: Text(S.of(context).updateTransactionFailed),
          description: Text(e.toString()),
          autoCloseDuration: const Duration(seconds: 5),
          type: ToastificationType.error,
          style: ToastificationStyle.fillColored,
          showProgressBar: false,
        );
      }
    }
  }
}
