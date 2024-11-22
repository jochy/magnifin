import 'package:adaptive_theme/adaptive_theme.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:front/components/scaffold.dart';
import 'package:front/config.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:front/screens/accounts/accounts_screen.dart';
import 'package:front/screens/accounts/add_account_screen.dart';
import 'package:front/screens/auth/signin_screen.dart';
import 'package:front/screens/budget/budget_screen.dart';
import 'package:front/screens/home/home_screen.dart';
import 'package:front/screens/auth/login_screen.dart';
import 'package:front/screens/url_screen.dart';
import 'package:go_router/go_router.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:toastification/toastification.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final savedThemeMode = await AdaptiveTheme.getThemeMode();
  var authCubit = AuthCubit();
  await authCubit.loadFromStorage();
  var transactionsCubit = TransactionsCubit(authCubit);
  var connectionsCubit = ConnectionsCubit(authCubit, transactionsCubit);
  runApp(MyApp(
    authCubit: authCubit,
    connectionsCubit: connectionsCubit,
    transactionsCubit: transactionsCubit,
    savedThemeMode: savedThemeMode,
  ));
}

final GoRouter _router = GoRouter(
  routes: [
    ShellRoute(
      routes: [
        GoRoute(
          path: '/',
          builder: (BuildContext context, GoRouterState state) {
            return const HomeScreen();
          },
        ),
        GoRoute(
          path: '/budget',
          builder: (BuildContext context, GoRouterState state) {
            return const BudgetScreen();
          },
        ),
        GoRoute(
          path: '/accounts',
          builder: (BuildContext context, GoRouterState state) {
            return const AccountsScreen();
          },
        ),
        GoRoute(
          path: '/add-account',
          builder: (BuildContext context, GoRouterState state) {
            return const AddAccountScreen();
          },
        ),
      ],
      builder: (context, state, child) {
        return ScaffoldWithNavigation(
          state: state,
          child: child,
        );
      },
    ),
    GoRoute(
      path: '/login',
      builder: (BuildContext context, GoRouterState state) {
        return const LoginScreen();
      },
    ),
    GoRoute(
      path: '/sign-in',
      builder: (BuildContext context, GoRouterState state) {
        return const SigninScreen();
      },
    ),
    GoRoute(
      path: "/url",
      builder: (context, state) {
        return const UrlScreen();
      },
    ),
  ],
  redirect: (context, state) {
    if (Configuration.instance.baseUrl == "") {
      return '/url';
    }

    var auth = context.read<AuthCubit>().state;
    if (auth.token == null &&
        (state.fullPath != '/login' && state.fullPath != '/sign-in')) {
      return '/login';
    }

    return null;
  },
);

class MyApp extends StatelessWidget {
  final AuthCubit authCubit;
  final ConnectionsCubit connectionsCubit;
  final TransactionsCubit transactionsCubit;
  final AdaptiveThemeMode? savedThemeMode;

  const MyApp({
    super.key,
    required this.authCubit,
    required this.connectionsCubit,
    required this.transactionsCubit,
    this.savedThemeMode,
  });

  @override
  Widget build(BuildContext context) {
    if (kIsWeb) {
      Configuration.instance.baseUrl = webUri();
      SharedPreferencesAsync().setString("url", Configuration.instance.baseUrl);
    }

    return ToastificationWrapper(
      child: MultiBlocProvider(
        providers: [
          BlocProvider<AuthCubit>(create: (context) => authCubit),
          BlocProvider<ConnectionsCubit>(create: (context) => connectionsCubit),
          BlocProvider<TransactionsCubit>(
              create: (context) => transactionsCubit),
        ],
        child: AdaptiveTheme(
          light: ThemeData(
            colorScheme: ColorScheme.fromSeed(
              seedColor: Colors.teal,
              brightness: Brightness.light,
            ),
            useMaterial3: true,
          ),
          dark: ThemeData(
            colorScheme: ColorScheme.fromSeed(
              seedColor: Colors.teal,
              brightness: Brightness.dark,
            ),
            useMaterial3: true,
          ),
          initial: savedThemeMode ?? AdaptiveThemeMode.system,
          builder: (light, dark) => ResponsiveBreakpoints(
            breakpoints: const [
              Breakpoint(start: 0, end: 450, name: MOBILE),
              Breakpoint(start: 451, end: 960, name: TABLET),
              Breakpoint(start: 961, end: double.infinity, name: DESKTOP),
            ],
            child: MaterialApp.router(
              routerConfig: _router,
              title: 'MagniFin',
              theme: light,
              localizationsDelegates: const [
                S.delegate,
                GlobalCupertinoLocalizations.delegate,
                GlobalMaterialLocalizations.delegate,
                GlobalWidgetsLocalizations.delegate,
              ],
              locale: const Locale('en', ''),
              supportedLocales: S.delegate.supportedLocales,
              debugShowCheckedModeBanner: false,
            ),
          ),
        ),
      ),
    );
  }

  String webUri() {
    var currentUri = Uri.base.toString();
    var apiUri = currentUri;

    if (apiUri.contains("#")) {
      apiUri = apiUri.substring(0, apiUri.indexOf("#"));
    }

    if (apiUri.endsWith("/")) {
      apiUri = apiUri.substring(0, apiUri.length - 1);
    }

    return "$apiUri/api";
  }
}
