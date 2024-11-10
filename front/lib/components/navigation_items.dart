import 'package:flutter/material.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';

enum NavigationItem {
  dashboard(iconData: Icons.dashboard_outlined, route: "/"),
  accounts(iconData: Icons.account_balance, route: "/accounts"),
  budget(iconData: MdiIcons.handCoin, route: "/budget"),
  ;

  const NavigationItem({required this.iconData, required this.route});

  final IconData iconData;
  final String route;

  String get label => name[0].toUpperCase() + name.substring(1).toLowerCase(); // Will have to translate it
}
