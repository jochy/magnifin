
part of 'connections_cubit.dart';

@freezed
class ConnectionsState with _$ConnectionsState {
  const factory ConnectionsState({
    required List<Connection> connections,
    required bool isLoading,
    required bool hasLoaded,
    String? error,
  }) = _ConnectionsState;
}

@freezed
class Connection with _$Connection {
  const factory Connection({
    required int id,
    required String status,
    @JsonKey(name: "renew_consent_before") DateTime? renewConsentBefore,
    @JsonKey(name: "error_message") String? errorMessage,
    @JsonKey(name: "last_successful_sync") DateTime? lastSuccessfulSync,
    required List<Account> accounts,
    required Connector connector,
  }) = _Connection;

  factory Connection.fromJson(Map<String, dynamic> json) =>
      _$ConnectionFromJson(json);
}

@freezed
class Account with _$Account {
  const factory Account({
    required int id,
    @JsonKey(name: "bank_account_id") String? bankAccountId,
    required String name,
    String? type,
    String? currency,
    @JsonKey(name: "account_number") String? accountNumber,
    required double balance,
  }) = _Account;

  factory Account.fromJson(Map<String, dynamic> json) =>
      _$AccountFromJson(json);
}

@freezed
class Connector with _$Connector {
  const factory Connector({
    required int id,
    required String name,
    @JsonKey(name: "logo_url") required String logoUrl,
  }) = _Connector;

  factory Connector.fromJson(Map<String, dynamic> json) =>
      _$ConnectorFromJson(json);
}

extension ConnectionExtension on ConnectionsState {
  String accountNameById(int id) {
    return connections.expand((e) => e.accounts).firstWhere((element) => element.id == id).name;
  }

  String? accountLogoById(int id) {
    for (var c in connections) {
      for (var a in c.accounts) {
        if (a.id == id) {
          return c.connector.logoUrl;
        }
      }
    }
    return null;
  }
}