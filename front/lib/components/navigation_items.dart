import 'package:flutter/material.dart';

enum NavigationItem {
  dashboard(iconData: Icons.dashboard_outlined, route: "/"),
  accounts(iconData: Icons.account_balance, route: "/accounts"),
  transactions(iconData: Icons.list, route: "/transactions"),
  ;

  const NavigationItem({required this.iconData, required this.route});

  final IconData iconData;
  final String route;

  String get label => name[0].toUpperCase() + name.substring(1).toLowerCase();
}
