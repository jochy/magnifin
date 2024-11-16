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
        "autoTransport":
            MessageLookupByLibrary.simpleMessage("Auto & Transport"),
        "backendUrl": MessageLookupByLibrary.simpleMessage("Backend url"),
        "balance": MessageLookupByLibrary.simpleMessage("Balance"),
        "bankName": MessageLookupByLibrary.simpleMessage("Bank name"),
        "businessWork": MessageLookupByLibrary.simpleMessage("Business & Work"),
        "cancel": MessageLookupByLibrary.simpleMessage("Cancel"),
        "cashChecks": MessageLookupByLibrary.simpleMessage("Cash & Checks"),
        "category": MessageLookupByLibrary.simpleMessage("Category:"),
        "connect": MessageLookupByLibrary.simpleMessage("Connect"),
        "createANewAccount":
            MessageLookupByLibrary.simpleMessage("Create a new account"),
        "createNewAccount":
            MessageLookupByLibrary.simpleMessage("Create new account"),
        "createRuleError":
            MessageLookupByLibrary.simpleMessage("Create rule error"),
        "createUserFailed":
            MessageLookupByLibrary.simpleMessage("Create user failed"),
        "delete": MessageLookupByLibrary.simpleMessage("Delete"),
        "deleteConnection":
            MessageLookupByLibrary.simpleMessage("Delete connection"),
        "distribution": MessageLookupByLibrary.simpleMessage("Distribution"),
        "doYouWantToCreateASmartRule": MessageLookupByLibrary.simpleMessage(
            "Do you want to create a smart rule?"),
        "essentialNeeds":
            MessageLookupByLibrary.simpleMessage("Essential Needs"),
        "expenses": MessageLookupByLibrary.simpleMessage("Expenses"),
        "foodDrink": MessageLookupByLibrary.simpleMessage("Food & Drink"),
        "health": MessageLookupByLibrary.simpleMessage("Health"),
        "income": MessageLookupByLibrary.simpleMessage("Income"),
        "investment": MessageLookupByLibrary.simpleMessage("Investment"),
        "lastSuccessfulSync": m0,
        "loanRepayment": MessageLookupByLibrary.simpleMessage("Loan Repayment"),
        "login": MessageLookupByLibrary.simpleMessage("Login"),
        "loginError": MessageLookupByLibrary.simpleMessage("Login error"),
        "loginFailedPleaseTryAgain": MessageLookupByLibrary.simpleMessage(
            "Login failed. Please try again."),
        "newAccount": MessageLookupByLibrary.simpleMessage("New account"),
        "newSmartRule": MessageLookupByLibrary.simpleMessage("New smart rule"),
        "no": MessageLookupByLibrary.simpleMessage("No"),
        "noDataToDisplay":
            MessageLookupByLibrary.simpleMessage("No data to display"),
        "other": MessageLookupByLibrary.simpleMessage("Other"),
        "password": MessageLookupByLibrary.simpleMessage("Password"),
        "rateLimited": MessageLookupByLibrary.simpleMessage("Rate Limited"),
        "selectTheDataUsedToCategorizeFutureTransactions":
            MessageLookupByLibrary.simpleMessage(
                "Select the data used to categorize future transactions."),
        "signOut": MessageLookupByLibrary.simpleMessage("Sign out"),
        "subscriptionsAndBills":
            MessageLookupByLibrary.simpleMessage("Subscriptions and Bills"),
        "suspended": MessageLookupByLibrary.simpleMessage("Suspended"),
        "switchToDark": MessageLookupByLibrary.simpleMessage("Switch to dark"),
        "switchToLight":
            MessageLookupByLibrary.simpleMessage("Switch to light"),
        "syncInProgress":
            MessageLookupByLibrary.simpleMessage("Sync in Progress"),
        "synced": MessageLookupByLibrary.simpleMessage("Synced"),
        "taxes": MessageLookupByLibrary.simpleMessage("Taxes"),
        "thisWillAutomaticallyCategorizeSimilarTransactionsInTheFuture":
            MessageLookupByLibrary.simpleMessage(
                "This will automatically categorize similar transactions in the future."),
        "transaction": MessageLookupByLibrary.simpleMessage("Transaction"),
        "transfers": MessageLookupByLibrary.simpleMessage("Transfers"),
        "tryAgain": MessageLookupByLibrary.simpleMessage("Try again"),
        "unableToConnectToConnectorname": m1,
        "unableToDeleteConnection":
            MessageLookupByLibrary.simpleMessage("Unable to delete connection"),
        "unableToLoadApp":
            MessageLookupByLibrary.simpleMessage("Unable to load app"),
        "unknown": MessageLookupByLibrary.simpleMessage("Unknown"),
        "updateTransactionFailed":
            MessageLookupByLibrary.simpleMessage("Update transaction failed"),
        "url": MessageLookupByLibrary.simpleMessage("Url"),
        "username": MessageLookupByLibrary.simpleMessage("Username"),
        "validate": MessageLookupByLibrary.simpleMessage("Validate"),
        "yes": MessageLookupByLibrary.simpleMessage("Yes")
      };
}
