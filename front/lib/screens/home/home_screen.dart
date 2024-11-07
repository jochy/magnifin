import 'package:flutter/material.dart';
import 'package:front/screens/home/distribution_component.dart';
import 'package:front/screens/home/worth_component.dart';
import 'package:responsive_framework/responsive_framework.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final breakpoint = ResponsiveBreakpoints.of(context).breakpoint;

    return SingleChildScrollView(
      child: switch (breakpoint.name) {
        MOBILE || TABLET => _HomeScreenMobile(),
        (_) => _HomeScreenDesktop(),
      },
    );
  }
}

class _HomeScreenDesktop extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return const Row(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      mainAxisSize: MainAxisSize.max,
      children: [
        Expanded(flex: 2, child: WorthComponent()),
        Expanded(child: DistributionComponent()),
      ],
    );
  }
}

class _HomeScreenMobile extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return const Column(
      children: [
        WorthComponent(),
        const SizedBox(height: 16),
        DistributionComponent(),
      ],
    );
  }
}
