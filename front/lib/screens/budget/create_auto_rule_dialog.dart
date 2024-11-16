import 'package:flutter/material.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:toastification/toastification.dart';

class CreateAutoRuleDialog extends StatefulWidget {
  const CreateAutoRuleDialog({
    super.key,
    required this.transaction,
    required this.category,
    required this.counterparty,
  });

  final Transaction transaction;
  final Category category;
  final String counterparty;

  @override
  State<CreateAutoRuleDialog> createState() => _CreateAutoRuleDialogState();
}

class _CreateAutoRuleDialogState extends State<CreateAutoRuleDialog> {
  List<String> selectedElements = [];
  Map<String, bool> selected = {};
  bool isLoading = false;
  bool applyToAll = false;

  @override
  Widget build(BuildContext context) {
    Set<String> list = {
      widget.transaction.amount.toString(),
      widget.transaction.currency,
      ConnectionsCubit.of(context)
          .state
          .accountNameById(widget.transaction.accountId),
      widget.counterparty.trim(),
      if (widget.transaction.method != null) widget.transaction.method!,
      widget.transaction.direction,
      if (widget.transaction.reference != null)
        ...widget.transaction.reference!.split(" "),
    };

    list = list.map((e) => e.trim()).where((e) => e.isNotEmpty).toSet();

    return AlertDialog(
      titlePadding: const EdgeInsets.all(16),
      insetPadding: const EdgeInsets.all(16),
      buttonPadding: const EdgeInsets.all(8),
      actionsPadding: const EdgeInsets.all(16),
      contentPadding: const EdgeInsets.all(12),
      title: Text(S.of(context).newSmartRule),
      content: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(S.of(context).selectTheDataUsedToCategorizeFutureTransactions),
            const SizedBox(height: 8),
            Wrap(
              spacing: 2,
              runSpacing: 2,
              children: list.map((e) => _chip(e)).toList(),
            ),
            const SizedBox(height: 8),
            CheckboxListTile(
              value: applyToAll,
              onChanged: (v) {
                setState(() {
                  applyToAll = v ?? false;
                });
              },
              controlAffinity: ListTileControlAffinity.leading,
              dense: true,
              visualDensity: const VisualDensity(horizontal: -4, vertical: -4),
              contentPadding: const EdgeInsets.all(0),
              title: Text(S.of(context).applyRuleToExistingTransactions),
              subtitle: Align(
                alignment: Alignment.centerLeft,
                child: StyledText.hintText(S.of(context).thisMayImpactAllTransactions),
              ),
            ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: Text(
            S.of(context).cancel,
            style: const TextStyle(fontSize: 12),
          ),
        ),
        FilledButton(
          onPressed: () => _createRule(context),
          style: FilledButton.styleFrom(
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          ),
          child: Text(S.of(context).validate),
        ),
      ],
    );
  }

  Widget _chip(String label) {
    return InkWell(
      borderRadius: BorderRadius.circular(100),
      onTap: () {
        setState(() {
          if (selected[label] == true) {
            selectedElements.remove(label);
          } else {
            selectedElements.add(label);
          }
          selected[label] = !(selected[label] ?? false);
        });
      },
      child: Chip(
        visualDensity: const VisualDensity(horizontal: -4, vertical: -4),
        padding: const EdgeInsets.all(0),
        labelPadding: const EdgeInsets.symmetric(horizontal: 4, vertical: 0),
        label: Text(label),
        labelStyle: TextStyle(
          color: (selected[label] ?? false)
              ? Theme.of(context).colorScheme.primary
              : Theme.of(context).colorScheme.onSurface,
          fontSize: 10,
        ),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(100),
          side: BorderSide(
            color: (selected[label] ?? false)
                ? Theme.of(context).colorScheme.primary
                : Theme.of(context).colorScheme.onSurface.withOpacity(0.5),
          ),
        ),
        elevation: (selected[label] ?? false) ? 2 : 0,
      ),
    );
  }

  void _createRule(BuildContext context) async {
    var keywords = [...selectedElements];
    var accountName = ConnectionsCubit.of(context)
        .state
        .accountNameById(widget.transaction.accountId);
    if (keywords.isEmpty) {
      Navigator.of(context).pop();
      return;
    }

    if (keywords.contains(accountName)) {
      keywords.remove(accountName);
      keywords.add(widget.transaction.accountId.toString());
    }

    var cubit = TransactionsCubit.of(context);
    setState(() {
      isLoading = true;
    });

    try {
      await cubit.createCategoryRule(widget.category, keywords, applyToAll);

      if (this.context.mounted) {
        Navigator.of(context).pop();
      } else {
        WidgetsBinding.instance.addPostFrameCallback((_) {
          if (this.context.mounted) {
            Navigator.of(context).pop();
          }
        });
        setState(() {
          isLoading = false;
        });
      }
    } catch (e) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        toastification.show(
          context: context,
          title: Text(S.of(context).createRuleError),
          description: Text(e.toString()),
          autoCloseDuration: const Duration(seconds: 5),
          type: ToastificationType.error,
          style: ToastificationStyle.fillColored,
          showProgressBar: false,
        );
      });

      setState(() {
        isLoading = false;
      });
    }
  }
}
