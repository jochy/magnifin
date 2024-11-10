// DO NOT EDIT. This is code generated via package:intl/generate_localized.dart
// This is a library that provides messages for a en locale. All the
// messages from the main program should be duplicated here with the same
// function name.

// Ignore issues from commonly used lints in this file.
// ignore_for_file:unnecessary_brace_in_string_interps, unnecessary_new
// ignore_for_file:prefer_single_quotes,comment_references, directives_ordering
// ignore_for_file:annotate_overrides,prefer_generic_function_type_aliases
// ignore_for_file:unused_import, file_names, avoid_escaping_inner_quotes
// ignore_for_file:unnecessary_string_interpolations, unnecessary_string_escapes

import 'package:intl/intl.dart';
import 'package:intl/message_lookup_by_library.dart';

final messages = new MessageLookup();

typedef String MessageIfAbsent(String messageStr, List<dynamic> args);

class MessageLookup extends MessageLookupByLibrary {
  String get localeName => 'en';

  static String m0(time) => "Last sync ${time} ago";

  static String m1(connectorname) => "Unable to connect to ${connectorname}";

  final messages = _notInlinedMessages(_notInlinedMessages);
  static Map<String, Function> _notInlinedMessages(_) => <String, Function>{
        "addAConnectionToGetStarted": MessageLookupByLibrary.simpleMessage(
            "Add a connection to get started"),
        "addANewAccount":
            MessageLookupByLibrary.simpleMessage("Add a new Account"),
        "anErrorOccurredWhileTryingToContactTheServerNplease":
            MessageLookupByLibrary.simpleMessage(
                "An error occurred while trying to contact the server. \nPlease make sure the URL is correct and try again."),
        "areYouSureYouWantToDeleteThisConnectionnthisCan":
            MessageLookupByLibrary.simpleMessage(
                "Are you sure you want to delete this connection?\nThis can not be undone and all data will be lost."),
        "backendUrl": MessageLookupByLibrary.simpleMessage("Backend url"),
        "balance": MessageLookupByLibrary.simpleMessage("Balance"),
        "bankName": MessageLookupByLibrary.simpleMessage("Bank name"),
        "cancel": MessageLookupByLibrary.simpleMessage("Cancel"),
        "connect": MessageLookupByLibrary.simpleMessage("Connect"),
        "createANewAccount":
            MessageLookupByLibrary.simpleMessage("Create a new account"),
        "createNewAccount":
            MessageLookupByLibrary.simpleMessage("Create new account"),
        "createUserFailed":
            MessageLookupByLibrary.simpleMessage("Create user failed"),
        "delete": MessageLookupByLibrary.simpleMessage("Delete"),
        "deleteConnection":
            MessageLookupByLibrary.simpleMessage("Delete connection"),
        "distribution": MessageLookupByLibrary.simpleMessage("Distribution"),
        "expenses": MessageLookupByLibrary.simpleMessage("Expenses"),
        "income": MessageLookupByLibrary.simpleMessage("Income"),
        "lastSuccessfulSync": m0,
        "login": MessageLookupByLibrary.simpleMessage("Login"),
        "loginError": MessageLookupByLibrary.simpleMessage("Login error"),
        "loginFailedPleaseTryAgain": MessageLookupByLibrary.simpleMessage(
            "Login failed. Please try again."),
        "newAccount": MessageLookupByLibrary.simpleMessage("New account"),
        "noDataToDisplay":
            MessageLookupByLibrary.simpleMessage("No data to display"),
        "other": MessageLookupByLibrary.simpleMessage("Other"),
        "password": MessageLookupByLibrary.simpleMessage("Password"),
        "rateLimited": MessageLookupByLibrary.simpleMessage("Rate Limited"),
        "signOut": MessageLookupByLibrary.simpleMessage("Sign out"),
        "suspended": MessageLookupByLibrary.simpleMessage("Suspended"),
        "switchToDark": MessageLookupByLibrary.simpleMessage("Switch to dark"),
        "switchToLight":
            MessageLookupByLibrary.simpleMessage("Switch to light"),
        "syncInProgress":
            MessageLookupByLibrary.simpleMessage("Sync in Progress"),
        "synced": MessageLookupByLibrary.simpleMessage("Synced"),
        "transaction": MessageLookupByLibrary.simpleMessage("Transaction"),
        "tryAgain": MessageLookupByLibrary.simpleMessage("Try again"),
        "unableToConnectToConnectorname": m1,
        "unableToDeleteConnection":
            MessageLookupByLibrary.simpleMessage("Unable to delete connection"),
        "unableToLoadApp":
            MessageLookupByLibrary.simpleMessage("Unable to load app"),
        "unknown": MessageLookupByLibrary.simpleMessage("Unknown"),
        "url": MessageLookupByLibrary.simpleMessage("Url"),
        "username": MessageLookupByLibrary.simpleMessage("Username"),
        "validate": MessageLookupByLibrary.simpleMessage("Validate")
      };
}
