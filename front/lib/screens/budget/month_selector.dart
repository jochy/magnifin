import 'package:flutter/material.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';

class MonthSelector extends StatelessWidget {
  const MonthSelector({
    super.key,
    required this.allMonths,
    required this.onMonthSelected,
    required this.monthSelect,
  });

  final List<MonthlyBudget> allMonths;
  final MonthlyBudget monthSelect;
  final Function(MonthlyBudget) onMonthSelected;

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        TextButton.icon(
          onPressed: allMonths.contains(monthSelect.previous())
              ? () => onMonthSelected(monthSelect.previous())
              : null,
          icon: const Icon(MdiIcons.chevronLeft),
          style: TextButton.styleFrom(
            visualDensity: const VisualDensity(horizontal: -4, vertical: -4),
            padding: const EdgeInsets.all(0),
          ),
          label: const SizedBox(),
        ),
        DropdownButton<MonthlyBudget>(
            value: monthSelect,
            focusColor: Colors.transparent,
            items: allMonths.map((e) {
              return DropdownMenuItem<MonthlyBudget>(
                value: e,
                child: Text("${e.month}/${e.year}"),
              );
            }).toList(),
            underline: const SizedBox(),
            onChanged: (val) => onMonthSelected(val!)),
        TextButton.icon(
          onPressed: allMonths.contains(monthSelect.next())
              ? () => onMonthSelected(monthSelect.next())
              : null,
          icon: const Icon(MdiIcons.chevronRight),
          style: TextButton.styleFrom(
            visualDensity: const VisualDensity(horizontal: -4, vertical: -4),
            padding: const EdgeInsets.all(0),
          ),
          label: const SizedBox(),
        ),
      ],
    );
  }
}
