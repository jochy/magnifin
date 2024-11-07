import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/components/empty_state.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/generated/l10n.dart';

class DistributionComponent extends StatelessWidget {
  const DistributionComponent({super.key});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Card.outlined(
        child: Padding(
          padding: const EdgeInsets.all(12.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              StyledText.titleMedium(S.of(context).distribution),
              SizedBox(
                width: double.infinity,
                height: 200,
                child: BlocBuilder<ConnectionsCubit, ConnectionsState>(
                  builder: (context, state) {
                    if (state.connections.isEmpty) {
                      return const Center(
                        child: EmptyStateComponent(),
                      );
                    }

                    return PieChart(
                      PieChartData(
                        borderData: FlBorderData(
                          show: false,
                        ),
                        sectionsSpace: 2,
                        centerSpaceRadius: 45,
                        sections: showingSections(context),
                      ),
                    );
                  },
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  List<PieChartSectionData> showingSections(context) {
    ColorScheme colors = Theme.of(context).colorScheme;

    var accounts = _compute5Accounts(context).entries.toList();
    var pieColors = [
      Colors.blue,
      Colors.orange,
      Colors.purple,
      Colors.green,
      Colors.red,
    ];

    return List.generate(accounts.length, (i) {
      const fontSize = 12.0;
      const radius = 55.0;
      final shadows = [Shadow(color: colors.onInverseSurface, blurRadius: 2)];
      return PieChartSectionData(
        color: pieColors[i],
        value: accounts[i].value,
        title: "${accounts[i].key} (${accounts[i].value.round()}%)",
        radius: radius,
        titleStyle: TextStyle(
          fontSize: fontSize,
          fontWeight: FontWeight.bold,
          color: colors.onSurface,
          shadows: shadows,
        ),
      );
    });
  }

  Map<String, double> _compute5Accounts(BuildContext context) {
    var connections = ConnectionsCubit.of(context).state.connections;
    var accounts = <String, double>{};
    var maxMoney = 0.0;
    try {
      maxMoney = connections
          .expand((e) => e.accounts)
          .map((e) => e.balance)
          .reduce((value, element) => value + element);
    } catch (e) {
      maxMoney = 0.0;
    }

    // Extract the 5 accounts with the highest balance
    var sortedAccounts = connections.expand((e) => e.accounts).toList();
    var takenAccounts = [];

    sortedAccounts.sort((a, b) => b.balance.compareTo(a.balance));

    sortedAccounts.take(4).forEach((element) {
      takenAccounts.add(element);
      var connection = connections.firstWhere(
        (c) => c.accounts.contains(element),
      );
      accounts["${connection.connector.name} - ${element.name}"] =
          (element.balance / maxMoney) * 100;
    });

    if (sortedAccounts.length > 4) {
      var otherBalance = sortedAccounts
          .skip(4)
          .map((e) => e.balance)
          .reduce((value, element) => value + element);
      accounts[S.of(context).other] = (otherBalance / maxMoney) * 100;
    }

    return accounts;
  }
}
