// GENERATED CODE - DO NOT MODIFY BY HAND
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'intl/messages_all.dart';

// **************************************************************************
// Generator: Flutter Intl IDE plugin
// Made by Localizely
// **************************************************************************

// ignore_for_file: non_constant_identifier_names, lines_longer_than_80_chars
// ignore_for_file: join_return_with_assignment, prefer_final_in_for_each
// ignore_for_file: avoid_redundant_argument_values, avoid_escaping_inner_quotes

class S {
  S();

  static S? _current;

  static S get current {
    assert(_current != null,
        'No instance of S was loaded. Try to initialize the S delegate before accessing S.current.');
    return _current!;
  }

  static const AppLocalizationDelegate delegate = AppLocalizationDelegate();

  static Future<S> load(Locale locale) {
    final name = (locale.countryCode?.isEmpty ?? false)
        ? locale.languageCode
        : locale.toString();
    final localeName = Intl.canonicalizedLocale(name);
    return initializeMessages(localeName).then((_) {
      Intl.defaultLocale = localeName;
      final instance = S();
      S._current = instance;

      return instance;
    });
  }

  static S of(BuildContext context) {
    final instance = S.maybeOf(context);
    assert(instance != null,
        'No instance of S present in the widget tree. Did you add S.delegate in localizationsDelegates?');
    return instance!;
  }

  static S? maybeOf(BuildContext context) {
    return Localizations.of<S>(context, S);
  }

  /// `Login`
  String get login {
    return Intl.message(
      'Login',
      name: 'login',
      desc: '',
      args: [],
    );
  }

  /// `Username`
  String get username {
    return Intl.message(
      'Username',
      name: 'username',
      desc: '',
      args: [],
    );
  }

  /// `Password`
  String get password {
    return Intl.message(
      'Password',
      name: 'password',
      desc: '',
      args: [],
    );
  }

  /// `Create a new account`
  String get createANewAccount {
    return Intl.message(
      'Create a new account',
      name: 'createANewAccount',
      desc: '',
      args: [],
    );
  }

  /// `Login failed. Please try again.`
  String get loginFailedPleaseTryAgain {
    return Intl.message(
      'Login failed. Please try again.',
      name: 'loginFailedPleaseTryAgain',
      desc: '',
      args: [],
    );
  }

  /// `Login error`
  String get loginError {
    return Intl.message(
      'Login error',
      name: 'loginError',
      desc: '',
      args: [],
    );
  }

  /// `Create new account`
  String get createNewAccount {
    return Intl.message(
      'Create new account',
      name: 'createNewAccount',
      desc: '',
      args: [],
    );
  }

  /// `Create user failed`
  String get createUserFailed {
    return Intl.message(
      'Create user failed',
      name: 'createUserFailed',
      desc: '',
      args: [],
    );
  }

  /// `Url`
  String get url {
    return Intl.message(
      'Url',
      name: 'url',
      desc: '',
      args: [],
    );
  }

  /// `Validate`
  String get validate {
    return Intl.message(
      'Validate',
      name: 'validate',
      desc: '',
      args: [],
    );
  }

  /// `Backend url`
  String get backendUrl {
    return Intl.message(
      'Backend url',
      name: 'backendUrl',
      desc: '',
      args: [],
    );
  }

  /// `An error occurred while trying to contact the server. \nPlease make sure the URL is correct and try again.`
  String get anErrorOccurredWhileTryingToContactTheServerNplease {
    return Intl.message(
      'An error occurred while trying to contact the server. \nPlease make sure the URL is correct and try again.',
      name: 'anErrorOccurredWhileTryingToContactTheServerNplease',
      desc: '',
      args: [],
    );
  }

  /// `Sign out`
  String get signOut {
    return Intl.message(
      'Sign out',
      name: 'signOut',
      desc: '',
      args: [],
    );
  }

  /// `Switch to dark`
  String get switchToDark {
    return Intl.message(
      'Switch to dark',
      name: 'switchToDark',
      desc: '',
      args: [],
    );
  }

  /// `Switch to light`
  String get switchToLight {
    return Intl.message(
      'Switch to light',
      name: 'switchToLight',
      desc: '',
      args: [],
    );
  }

  /// `New account`
  String get newAccount {
    return Intl.message(
      'New account',
      name: 'newAccount',
      desc: '',
      args: [],
    );
  }

  /// `Try again`
  String get tryAgain {
    return Intl.message(
      'Try again',
      name: 'tryAgain',
      desc: '',
      args: [],
    );
  }

  /// `Unable to load app`
  String get unableToLoadApp {
    return Intl.message(
      'Unable to load app',
      name: 'unableToLoadApp',
      desc: '',
      args: [],
    );
  }

  /// `No data to display`
  String get noDataToDisplay {
    return Intl.message(
      'No data to display',
      name: 'noDataToDisplay',
      desc: '',
      args: [],
    );
  }

  /// `Add a connection to get started`
  String get addAConnectionToGetStarted {
    return Intl.message(
      'Add a connection to get started',
      name: 'addAConnectionToGetStarted',
      desc: '',
      args: [],
    );
  }

  /// `Bank name`
  String get bankName {
    return Intl.message(
      'Bank name',
      name: 'bankName',
      desc: '',
      args: [],
    );
  }

  /// `Connect`
  String get connect {
    return Intl.message(
      'Connect',
      name: 'connect',
      desc: '',
      args: [],
    );
  }

  /// `Unable to connect to {connectorname}`
  String unableToConnectToConnectorname(Object connectorname) {
    return Intl.message(
      'Unable to connect to $connectorname',
      name: 'unableToConnectToConnectorname',
      desc: '',
      args: [connectorname],
    );
  }

  /// `Distribution`
  String get distribution {
    return Intl.message(
      'Distribution',
      name: 'distribution',
      desc: '',
      args: [],
    );
  }

  /// `Other`
  String get other {
    return Intl.message(
      'Other',
      name: 'other',
      desc: '',
      args: [],
    );
  }

  /// `Suspended`
  String get suspended {
    return Intl.message(
      'Suspended',
      name: 'suspended',
      desc: '',
      args: [],
    );
  }

  /// `Unknown`
  String get unknown {
    return Intl.message(
      'Unknown',
      name: 'unknown',
      desc: '',
      args: [],
    );
  }

  /// `Sync in Progress`
  String get syncInProgress {
    return Intl.message(
      'Sync in Progress',
      name: 'syncInProgress',
      desc: '',
      args: [],
    );
  }

  /// `Rate Limited`
  String get rateLimited {
    return Intl.message(
      'Rate Limited',
      name: 'rateLimited',
      desc: '',
      args: [],
    );
  }

  /// `Synced`
  String get synced {
    return Intl.message(
      'Synced',
      name: 'synced',
      desc: '',
      args: [],
    );
  }

  /// `Delete`
  String get delete {
    return Intl.message(
      'Delete',
      name: 'delete',
      desc: '',
      args: [],
    );
  }

  /// `Last sync {time} ago`
  String lastSuccessfulSync(Object time) {
    return Intl.message(
      'Last sync $time ago',
      name: 'lastSuccessfulSync',
      desc: '',
      args: [time],
    );
  }

  /// `Cancel`
  String get cancel {
    return Intl.message(
      'Cancel',
      name: 'cancel',
      desc: '',
      args: [],
    );
  }

  /// `Delete connection`
  String get deleteConnection {
    return Intl.message(
      'Delete connection',
      name: 'deleteConnection',
      desc: '',
      args: [],
    );
  }

  /// `Are you sure you want to delete this connection?\nThis can not be undone and all data will be lost.`
  String get areYouSureYouWantToDeleteThisConnectionnthisCan {
    return Intl.message(
      'Are you sure you want to delete this connection?\nThis can not be undone and all data will be lost.',
      name: 'areYouSureYouWantToDeleteThisConnectionnthisCan',
      desc: '',
      args: [],
    );
  }

  /// `Unable to delete connection`
  String get unableToDeleteConnection {
    return Intl.message(
      'Unable to delete connection',
      name: 'unableToDeleteConnection',
      desc: '',
      args: [],
    );
  }

  /// `Add a new Account`
  String get addANewAccount {
    return Intl.message(
      'Add a new Account',
      name: 'addANewAccount',
      desc: '',
      args: [],
    );
  }
}

class AppLocalizationDelegate extends LocalizationsDelegate<S> {
  const AppLocalizationDelegate();

  List<Locale> get supportedLocales {
    return const <Locale>[
      Locale.fromSubtags(languageCode: 'en'),
    ];
  }

  @override
  bool isSupported(Locale locale) => _isSupported(locale);
  @override
  Future<S> load(Locale locale) => S.load(locale);
  @override
  bool shouldReload(AppLocalizationDelegate old) => false;

  bool _isSupported(Locale locale) {
    for (var supportedLocale in supportedLocales) {
      if (supportedLocale.languageCode == locale.languageCode) {
        return true;
      }
    }
    return false;
  }
}
