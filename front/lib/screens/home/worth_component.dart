import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/components/empty_state.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/connections/connections_cubit.dart';

class WorthComponent extends StatefulWidget {
  const WorthComponent({super.key});

  @override
  State<WorthComponent> createState() => _WorthComponentState();
}

class _WorthComponentState extends State<WorthComponent> {
  @override
  Widget build(BuildContext context) {
    var colorScheme = Theme
        .of(context)
        .colorScheme;
    List<Color> gradientColors = [
      colorScheme.primary,
      colorScheme.tertiary,
    ];

    return SizedBox(
      width: double.infinity,
      child: Card.outlined(
        child: Padding(
          padding: const EdgeInsets.all(12.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              StyledText.titleMedium("Worth"),
              StyledText.onError("TODO"),
              SizedBox(
                width: double.infinity,
                height: 200,
                child: Padding(
                  padding: const EdgeInsets.only(right: 12, top: 12),
                  child: BlocBuilder<ConnectionsCubit, ConnectionsState>(
                    builder: (context, state) {
                      if (state.connections.isEmpty) {
                        return const Center(
                          child: EmptyStateComponent(),
                        );
                      }

                      return LineChart(
                        avgData(gradientColors),
                      );
                    },
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget bottomTitleWidgets(double value, TitleMeta meta) {
    Widget text;
    switch (value.toInt()) {
      case 2:
        text = const Text('MAR');
        break;
      case 5:
        text = const Text('JUN');
        break;
      case 8:
        text = const Text('SEP');
        break;
      default:
        text = const Text('');
        break;
    }

    return SideTitleWidget(
      axisSide: meta.axisSide,
      child: text,
    );
  }

  Widget leftTitleWidgets(double value, TitleMeta meta) {
    String text;
    switch (value.toInt()) {
      case 1:
        text = '10K';
        break;
      case 5:
        text = '50k';
        break;
      default:
        return Container();
    }

    return Text(text, textAlign: TextAlign.left);
  }

  LineChartData avgData(List<Color> colors) {
    return LineChartData(
      lineTouchData: const LineTouchData(enabled: true),
      gridData: FlGridData(
        show: false,
        drawHorizontalLine: true,
        verticalInterval: 1,
        horizontalInterval: 1,
        getDrawingVerticalLine: (value) {
          return const FlLine(
            color: Color(0xff37434d),
            strokeWidth: 1,
          );
        },
        getDrawingHorizontalLine: (value) {
          return const FlLine(
            color: Color(0xff37434d),
            strokeWidth: 1,
          );
        },
      ),
      titlesData: FlTitlesData(
        show: true,
        bottomTitles: AxisTitles(
          sideTitles: SideTitles(
            showTitles: true,
            reservedSize: 40,
            getTitlesWidget: bottomTitleWidgets,
            interval: 1,
          ),
        ),
        leftTitles: AxisTitles(
          sideTitles: SideTitles(
            showTitles: true,
            getTitlesWidget: leftTitleWidgets,
            reservedSize: 40,
            interval: 1,
          ),
        ),
        topTitles: const AxisTitles(
          sideTitles: SideTitles(showTitles: false),
        ),
        rightTitles: const AxisTitles(
          sideTitles: SideTitles(showTitles: false),
        ),
      ),
      borderData: FlBorderData(show: false),
      minX: 0,
      maxX: 11,
      minY: 0,
      maxY: 6,
      lineBarsData: [
        LineChartBarData(
          spots: const [
            FlSpot(0, 1.44),
            FlSpot(2.6, 2.44),
            FlSpot(4.9, 5.48),
            FlSpot(6.8, 3.44),
            FlSpot(8, 3.44),
            FlSpot(9.5, 3.44),
            FlSpot(11, 4.44),
          ],
          isCurved: true,
          gradient: LinearGradient(
            colors: [
              ColorTween(begin: colors[0], end: colors[1]).lerp(0.2)!,
              ColorTween(begin: colors[0], end: colors[1]).lerp(0.2)!,
            ],
          ),
          isStrokeCapRound: true,
          dotData: const FlDotData(show: false),
          belowBarData: BarAreaData(
            show: true,
            gradient: LinearGradient(
              colors: [
                ColorTween(begin: colors[0], end: colors[1])
                    .lerp(0.2)!
                    .withOpacity(0.1),
                ColorTween(begin: colors[0], end: colors[1])
                    .lerp(0.2)!
                    .withOpacity(0.2),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
