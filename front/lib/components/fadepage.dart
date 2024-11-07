import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:go_router/go_router.dart';

class FadedPage extends CustomTransitionPage {
  FadedPage({required super.child, required GoRouterState state})
      : super(
          key: state.pageKey,
          restorationId: state.pageKey.value,
          transitionsBuilder: (context, animation, secondaryAnimation, child) {
            if (kIsWeb) {
              return child;
            }
            return FadeTransition(opacity: animation, child: child);
          },
          transitionDuration: const Duration(milliseconds: 5000),
        );
}
