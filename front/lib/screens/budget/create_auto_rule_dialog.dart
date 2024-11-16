import 'package:flutter/material.dart';
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
      title: Text(S.of(context).newSmartRule),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(S.of(context).selectTheDataUsedToCategorizeFutureTransactions),
          const SizedBox(height: 16),
          Wrap(
            spacing: 4,
            runSpacing: 4,
            children: list.map((e) => _chip(e)).toList(),
          ),
        ],
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: Text(S.of(context).cancel),
        ),
        FilledButton(
          onPressed: () => _createRule(context),
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
        labelPadding: const EdgeInsets.symmetric(horizontal: 4),
        label: Text(label),
        labelStyle: TextStyle(
          color: (selected[label] ?? false)
              ? Theme.of(context).colorScheme.primary
              : Theme.of(context).colorScheme.onSurface,
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
      await cubit.createCategoryRule(widget.category, keywords);

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
