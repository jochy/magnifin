import 'package:flutter/material.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:go_router/go_router.dart';

import 'navigation_title.dart';

class NavigationAppBar extends StatelessWidget implements PreferredSizeWidget {
  const NavigationAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    return AppBar(
      backgroundColor: Theme.of(context).colorScheme.surfaceContainer,
      title: const NavigationTitle(),
      centerTitle: false,
      scrolledUnderElevation: 0,
      actions: [
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8),
          child: PopupMenuButton<void>(
            itemBuilder: (context) => [
              PopupMenuItem(
                child: Text(S.of(context).signOut),
                onTap: () {
                  var goRouter = GoRouter.of(context);
                  AuthCubit.of(context).logout();
                  goRouter.go('/');
                },
              ),
            ],
            child: const Icon(Icons.account_circle_outlined),
          ),
        ),
        const SizedBox(height: 8),
      ],
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(48.0);
}
