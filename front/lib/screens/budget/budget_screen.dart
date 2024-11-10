import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/components/empty_state.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:front/screens/accounts/accounts_screen.dart';
import 'package:front/screens/budget/day_transaction_list.dart';
import 'package:front/screens/budget/month_selector.dart';
import 'package:moment_dart/moment_dart.dart';
import 'package:money2/money2.dart';

class BudgetScreen extends StatefulWidget {
  const BudgetScreen({super.key});

  @override
  State<BudgetScreen> createState() => _BudgetScreenState();
}

class _BudgetScreenState extends State<BudgetScreen> {
  MonthlyBudget? _monthSelect;

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<TransactionsCubit, TransactionsState>(
      builder: (context, state) {
        var allMonths =
            MonthlyBudget.fromMinMax(state.minDate!, state.maxDate!);

        if (state.isLoading) {
          return const Center(child: CircularProgressIndicator());
        }

        if (state.minDate == null || state.maxDate == null) {
          return const EmptyStateComponent();
        }

        if (_monthSelect == null) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            setState(() {
              _monthSelect = allMonths.first;
            });
          });
          return const SizedBox();
        }

        if (!state.loadedMonths.contains(_monthSelect!)) {
          TransactionsCubit.of(context).loadMonthTransactions(
            _monthSelect!.year,
            _monthSelect!.month,
          );
          return const SizedBox();
        }

        // TODO: handle error message

        var budget =
            _monthSelect!.loadIndicatorsFromTransactions(state.transactions);

        return SingleChildScrollView(
          child: Column(
            mainAxisSize: MainAxisSize.max,
            children: [
              Card.outlined(
                child: SizedBox(
                  width: double.infinity,
                  child: Column(
                    children: [
                      MonthSelector(
                        allMonths: allMonths,
                        onMonthSelected: (m) => setState(() {
                          _monthSelect = m;
                        }),
                        monthSelect: _monthSelect!,
                      ),
                      const SizedBox(height: 12),
                      BudgetSummaryWidget(budget: budget),
                      const SizedBox(height: 12),
                      // TODO: display Sankey Diagram
                    ],
                  ),
                ),
              ),
              const SizedBox(height: 20),
              ...budget.groupedTransactionsEntries.map((e) {
                return DayTransactionListWidget(
                  date: e.key,
                  transactions: e.value,
                );
              })
            ],
          ),
        );
      },
    );
  }
}

class BudgetSummaryWidget extends StatelessWidget {
  const BudgetSummaryWidget({
    super.key,
    required this.budget,
  });

  final MonthlyBudgetSummary budget;

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [
        Column(
          children: [
            StyledText.bodyLarge(S.of(context).income),
            StyledText.bodySmall(
              budget.income.toStringAsFixed(2),
            ),
          ],
        ),
        Column(
          children: [
            StyledText.bodyLarge(S.of(context).expenses),
            StyledText.bodySmall(
              budget.expenses.toStringAsFixed(2),
            ),
          ],
        ),
        Column(
          children: [
            StyledText.bodyLarge(S.of(context).balance),
            StyledText.bodySmall(
              budget.remaining.toStringAsFixed(2),
            ),
          ],
        ),
      ],
    );
  }
}
