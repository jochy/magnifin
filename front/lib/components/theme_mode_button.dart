import 'package:adaptive_theme/adaptive_theme.dart';
import 'package:flutter/material.dart';
import 'package:front/generated/l10n.dart';

enum _Variant {
  icon,
  outlined,
}

class ThemeModeButton extends StatelessWidget {
  const ThemeModeButton._(this.variant);

  const ThemeModeButton.icon() : this._(_Variant.icon);

  const ThemeModeButton.outlined() : this._(_Variant.outlined);

  final _Variant variant;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final colorScheme = theme.colorScheme;
    final brightness = Theme.of(context).brightness;
    final (iconData, action, actionLabel) = switch (brightness) {
      Brightness.light => (
          Icons.dark_mode_outlined,
          AdaptiveTheme.of(context).setDark,
          S.of(context).switchToDark
        ),
      Brightness.dark => (
          Icons.light_mode_outlined,
          AdaptiveTheme.of(context).setLight,
          S.of(context).switchToLight
        )
    };

    return switch (variant) {
      _Variant.icon => IconButton(
          icon: Icon(iconData),
          onPressed: action,
        ),
      _Variant.outlined => OutlinedButton.icon(
          onPressed: action,
          icon: Icon(iconData),
          label: Text(actionLabel),
          style: OutlinedButton.styleFrom(
            side: BorderSide(
              color: colorScheme.primary.withOpacity(0.5),
            ),
          ),
        )
    };
  }
}
