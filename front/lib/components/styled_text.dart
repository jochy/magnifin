import 'package:flutter/material.dart';

class StyledText extends StatelessWidget {
  final String text;
  final int? maxLines;
  final double? size;
  late final TextStyle? Function(BuildContext context)? getStyle;

  StyledText.headlineLarge(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.headlineLarge;
    };
  }

  StyledText.headlineMedium(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.headlineMedium;
    };
  }

  StyledText.headlineSmall(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.headlineSmall;
    };
  }

  StyledText.labelSmall(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.labelSmall;
    };
  }

  StyledText.labelMedium(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.labelMedium;
    };
  }

  StyledText.labelLarge(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.labelLarge;
    };
  }

  StyledText.titleSmall(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.titleSmall;
    };
  }

  StyledText.titleMedium(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.titleMedium;
    };
  }

  StyledText.titleLarge(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.titleLarge;
    };
  }

  StyledText.bodySmall(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.bodySmall;
    };
  }

  StyledText.bodyMedium(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.bodyMedium;
    };
  }

  StyledText.bodyLarge(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return Theme.of(context).textTheme.bodyLarge;
    };
  }

  StyledText.text(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = null;
  }

  StyledText.onError(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return TextStyle(color: Theme.of(context).colorScheme.error);
    };
  }

  StyledText.errorText(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return TextStyle(
          color: Theme.of(context).colorScheme.error, fontSize: size);
    };
  }

  StyledText.hintText(this.text, {super.key, this.maxLines, this.size}) {
    getStyle = (context) {
      return TextStyle(color: Theme.of(context).hintColor, fontSize: size);
    };
  }

  @override
  Widget build(BuildContext context) {
    return FittedBox(
      fit: BoxFit.scaleDown,
      child: Text(text, style: getStyle?.call(context), maxLines: maxLines),
    );
  }
}
