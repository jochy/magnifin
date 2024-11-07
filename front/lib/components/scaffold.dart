import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:front/components/navigation_appbar.dart';
import 'package:front/components/navigation_items.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/components/theme_mode_button.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:go_router/go_router.dart';
import 'package:responsive_framework/responsive_framework.dart';

class ScaffoldWithNavigation extends StatelessWidget {
  const ScaffoldWithNavigation({
    super.key,
    required this.child,
    required this.state,
  });

  final GoRouterState state;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    final breakpoint = ResponsiveBreakpoints.of(context).breakpoint;
    return BlocBuilder<ConnectionsCubit, ConnectionsState>(
      builder: (context, s) {
        if (s.isLoading) {
          return const Center(child: CircularProgressIndicator());
        }
        if (s.error != null) {
          return Material(
            child: Column(
              mainAxisSize: MainAxisSize.max,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  S.of(context).unableToLoadApp,
                  style: TextStyle(
                    color: Theme.of(context).colorScheme.error,
                    fontWeight: FontWeight.bold,
                    fontSize: 30,
                  ),
                ),
                const SizedBox(height: 8),
                StyledText.onError(s.error!),
                const SizedBox(height: 24),
                OutlinedButton(
                  style: OutlinedButton.styleFrom(
                    foregroundColor: Theme.of(context).colorScheme.error,
                  ),
                  onPressed: () =>
                      ConnectionsCubit.of(context).loadConnections(),
                  child: Text(S.of(context).tryAgain),
                ),
              ],
            ),
          );
        }

        return switch (breakpoint.name) {
          MOBILE || TABLET => _ScaffoldWithDrawer(child, state),
          (_) => _ScaffoldWithNavigationRail(child, state),
        };
      },
    );
  }
}

class _ScaffoldWithNavigationRail extends StatelessWidget {
  const _ScaffoldWithNavigationRail(this.child, this.state);

  final GoRouterState state;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final colorScheme = theme.colorScheme;
    return Scaffold(
      appBar: const NavigationAppBar(),
      backgroundColor: colorScheme.surfaceContainer,
      body: Row(
        children: [
          Container(
            color: colorScheme.surfaceContainer,
            child: Column(
              children: [
                Expanded(
                  child: _NavigationRail(
                    state: state,
                    expand: false,
                  ),
                ),
                const Padding(
                  padding: EdgeInsets.all(16),
                  child: ThemeModeButton.icon(),
                ),
              ],
            ),
          ),
          Expanded(
            child: Padding(
              padding: const EdgeInsets.only(top: 4.0, left: 4, right: 4),
              child: Column(
                mainAxisSize: MainAxisSize.max,
                children: [
                  Expanded(
                    child: SizedBox(
                      width: double.infinity,
                      child: Padding(
                        padding: const EdgeInsets.all(12.0),
                        child: child,
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class _ScaffoldWithDrawer extends StatelessWidget {
  const _ScaffoldWithDrawer(this.child, this.state);

  final GoRouterState state;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Scaffold(
      appBar: const NavigationAppBar(),
      backgroundColor: theme.colorScheme.surfaceContainer,
      body: SizedBox(
        width: double.infinity,
        child: Padding(
          padding: const EdgeInsets.only(top: 8, left: 8, right: 8),
          child: child,
        ),
      ),
      drawer: Drawer(
        backgroundColor: theme.colorScheme.surfaceContainer,
        child: Column(
          children: [
            DrawerHeader(
              decoration: const BoxDecoration(border: Border()),
              margin: EdgeInsets.zero,
              child: Center(
                child: Text(
                  "MaginFin",
                  style: theme.textTheme.bodyMedium!
                      .copyWith(fontWeight: FontWeight.w600),
                ),
              ),
            ),
            Expanded(
              child: _NavigationRail(
                state: state,
                expand: true,
              ),
            ),
            const Padding(
              padding: EdgeInsets.all(16),
              child: ThemeModeButton.outlined(),
            ),
          ],
        ),
      ),
    );
  }
}

class _NavigationRail extends StatelessWidget {
  const _NavigationRail({required this.expand, required this.state});

  final GoRouterState state;
  final bool expand;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    var currentRoute = state.fullPath;
    var navigationItems = NavigationItem.values;

    int? selectedIndex;
    for (var i = 0; i < navigationItems.length; i++) {
      if (navigationItems[i].route == currentRoute) {
        selectedIndex = i;
        break;
      }
    }

    return NavigationRail(
      backgroundColor: theme.colorScheme.surfaceContainer,
      extended: expand,
      selectedIndex: selectedIndex,
      unselectedLabelTextStyle: theme.textTheme.bodyMedium,
      selectedLabelTextStyle: theme.textTheme.bodyMedium!.copyWith(
        fontWeight: FontWeight.bold,
      ),
      onDestinationSelected: (index) {
        GoRouter.of(context).go(navigationItems[index].route);
        if (Scaffold.of(context).isDrawerOpen) {
          Scaffold.of(context).closeDrawer();
        }
      },
      destinations: [
        for (final item in NavigationItem.values)
          NavigationRailDestination(
            icon: Icon(item.iconData),
            label: Text(item.label),
          ),
      ],
    );
  }
}
