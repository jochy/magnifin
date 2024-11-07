import 'package:flutter/material.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';
import 'package:front/components/navigation_items.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/generated/l10n.dart';
import 'package:go_router/go_router.dart';

class EmptyStateComponent extends StatelessWidget {
  const EmptyStateComponent({super.key});

  @override
  Widget build(BuildContext context) {
    return Wrap(
      runAlignment: WrapAlignment.center,
      alignment: WrapAlignment.center,
      spacing: 12,
      runSpacing: 8,
      crossAxisAlignment: WrapCrossAlignment.center,
      children: [
        const Icon(MdiIcons.bankOff, size: 72),
        Column(
          mainAxisAlignment: MainAxisAlignment.center,
          mainAxisSize: MainAxisSize.min,
          children: [
            StyledText.titleMedium(S.of(context).noDataToDisplay),
            const SizedBox(height: 4),
            TextButton(
              onPressed: () =>
                  GoRouter.of(context).go(NavigationItem.accounts.route),
              child: Text(S.of(context).addAConnectionToGetStarted),
            ),
          ],
        ),
      ],
    );
  }
}
