// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'connections_cubit.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$ConnectionsState {
  List<Connection> get connections => throw _privateConstructorUsedError;
  bool get isLoading => throw _privateConstructorUsedError;
  bool get hasLoaded => throw _privateConstructorUsedError;
  String? get error => throw _privateConstructorUsedError;

  /// Create a copy of ConnectionsState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $ConnectionsStateCopyWith<ConnectionsState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ConnectionsStateCopyWith<$Res> {
  factory $ConnectionsStateCopyWith(
          ConnectionsState value, $Res Function(ConnectionsState) then) =
      _$ConnectionsStateCopyWithImpl<$Res, ConnectionsState>;
  @useResult
  $Res call(
      {List<Connection> connections,
      bool isLoading,
      bool hasLoaded,
      String? error});
}

/// @nodoc
class _$ConnectionsStateCopyWithImpl<$Res, $Val extends ConnectionsState>
    implements $ConnectionsStateCopyWith<$Res> {
  _$ConnectionsStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of ConnectionsState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? connections = null,
    Object? isLoading = null,
    Object? hasLoaded = null,
    Object? error = freezed,
  }) {
    return _then(_value.copyWith(
      connections: null == connections
          ? _value.connections
          : connections // ignore: cast_nullable_to_non_nullable
              as List<Connection>,
      isLoading: null == isLoading
          ? _value.isLoading
          : isLoading // ignore: cast_nullable_to_non_nullable
              as bool,
      hasLoaded: null == hasLoaded
          ? _value.hasLoaded
          : hasLoaded // ignore: cast_nullable_to_non_nullable
              as bool,
      error: freezed == error
          ? _value.error
          : error // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$ConnectionsStateImplCopyWith<$Res>
    implements $ConnectionsStateCopyWith<$Res> {
  factory _$$ConnectionsStateImplCopyWith(_$ConnectionsStateImpl value,
          $Res Function(_$ConnectionsStateImpl) then) =
      __$$ConnectionsStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {List<Connection> connections,
      bool isLoading,
      bool hasLoaded,
      String? error});
}

/// @nodoc
class __$$ConnectionsStateImplCopyWithImpl<$Res>
    extends _$ConnectionsStateCopyWithImpl<$Res, _$ConnectionsStateImpl>
    implements _$$ConnectionsStateImplCopyWith<$Res> {
  __$$ConnectionsStateImplCopyWithImpl(_$ConnectionsStateImpl _value,
      $Res Function(_$ConnectionsStateImpl) _then)
      : super(_value, _then);

  /// Create a copy of ConnectionsState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? connections = null,
    Object? isLoading = null,
    Object? hasLoaded = null,
    Object? error = freezed,
  }) {
    return _then(_$ConnectionsStateImpl(
      connections: null == connections
          ? _value._connections
          : connections // ignore: cast_nullable_to_non_nullable
              as List<Connection>,
      isLoading: null == isLoading
          ? _value.isLoading
          : isLoading // ignore: cast_nullable_to_non_nullable
              as bool,
      hasLoaded: null == hasLoaded
          ? _value.hasLoaded
          : hasLoaded // ignore: cast_nullable_to_non_nullable
              as bool,
      error: freezed == error
          ? _value.error
          : error // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc

class _$ConnectionsStateImpl implements _ConnectionsState {
  const _$ConnectionsStateImpl(
      {required final List<Connection> connections,
      required this.isLoading,
      required this.hasLoaded,
      this.error})
      : _connections = connections;

  final List<Connection> _connections;
  @override
  List<Connection> get connections {
    if (_connections is EqualUnmodifiableListView) return _connections;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_connections);
  }

  @override
  final bool isLoading;
  @override
  final bool hasLoaded;
  @override
  final String? error;

  @override
  String toString() {
    return 'ConnectionsState(connections: $connections, isLoading: $isLoading, hasLoaded: $hasLoaded, error: $error)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ConnectionsStateImpl &&
            const DeepCollectionEquality()
                .equals(other._connections, _connections) &&
            (identical(other.isLoading, isLoading) ||
                other.isLoading == isLoading) &&
            (identical(other.hasLoaded, hasLoaded) ||
                other.hasLoaded == hasLoaded) &&
            (identical(other.error, error) || other.error == error));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(_connections),
      isLoading,
      hasLoaded,
      error);

  /// Create a copy of ConnectionsState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$ConnectionsStateImplCopyWith<_$ConnectionsStateImpl> get copyWith =>
      __$$ConnectionsStateImplCopyWithImpl<_$ConnectionsStateImpl>(
          this, _$identity);
}

abstract class _ConnectionsState implements ConnectionsState {
  const factory _ConnectionsState(
      {required final List<Connection> connections,
      required final bool isLoading,
      required final bool hasLoaded,
      final String? error}) = _$ConnectionsStateImpl;

  @override
  List<Connection> get connections;
  @override
  bool get isLoading;
  @override
  bool get hasLoaded;
  @override
  String? get error;

  /// Create a copy of ConnectionsState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$ConnectionsStateImplCopyWith<_$ConnectionsStateImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

Connection _$ConnectionFromJson(Map<String, dynamic> json) {
  return _Connection.fromJson(json);
}

/// @nodoc
mixin _$Connection {
  int get id => throw _privateConstructorUsedError;
  String get status => throw _privateConstructorUsedError;
  @JsonKey(name: "renew_consent_before")
  DateTime? get renewConsentBefore => throw _privateConstructorUsedError;
  @JsonKey(name: "error_message")
  String? get errorMessage => throw _privateConstructorUsedError;
  @JsonKey(name: "last_successful_sync")
  DateTime? get lastSuccessfulSync => throw _privateConstructorUsedError;
  List<Account> get accounts => throw _privateConstructorUsedError;
  Connector get connector => throw _privateConstructorUsedError;

  /// Serializes this Connection to a JSON map.
  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;

  /// Create a copy of Connection
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $ConnectionCopyWith<Connection> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ConnectionCopyWith<$Res> {
  factory $ConnectionCopyWith(
          Connection value, $Res Function(Connection) then) =
      _$ConnectionCopyWithImpl<$Res, Connection>;
  @useResult
  $Res call(
      {int id,
      String status,
      @JsonKey(name: "renew_consent_before") DateTime? renewConsentBefore,
      @JsonKey(name: "error_message") String? errorMessage,
      @JsonKey(name: "last_successful_sync") DateTime? lastSuccessfulSync,
      List<Account> accounts,
      Connector connector});

  $ConnectorCopyWith<$Res> get connector;
}

/// @nodoc
class _$ConnectionCopyWithImpl<$Res, $Val extends Connection>
    implements $ConnectionCopyWith<$Res> {
  _$ConnectionCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of Connection
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? status = null,
    Object? renewConsentBefore = freezed,
    Object? errorMessage = freezed,
    Object? lastSuccessfulSync = freezed,
    Object? accounts = null,
    Object? connector = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String,
      renewConsentBefore: freezed == renewConsentBefore
          ? _value.renewConsentBefore
          : renewConsentBefore // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      errorMessage: freezed == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as String?,
      lastSuccessfulSync: freezed == lastSuccessfulSync
          ? _value.lastSuccessfulSync
          : lastSuccessfulSync // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      accounts: null == accounts
          ? _value.accounts
          : accounts // ignore: cast_nullable_to_non_nullable
              as List<Account>,
      connector: null == connector
          ? _value.connector
          : connector // ignore: cast_nullable_to_non_nullable
              as Connector,
    ) as $Val);
  }

  /// Create a copy of Connection
  /// with the given fields replaced by the non-null parameter values.
  @override
  @pragma('vm:prefer-inline')
  $ConnectorCopyWith<$Res> get connector {
    return $ConnectorCopyWith<$Res>(_value.connector, (value) {
      return _then(_value.copyWith(connector: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$ConnectionImplCopyWith<$Res>
    implements $ConnectionCopyWith<$Res> {
  factory _$$ConnectionImplCopyWith(
          _$ConnectionImpl value, $Res Function(_$ConnectionImpl) then) =
      __$$ConnectionImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {int id,
      String status,
      @JsonKey(name: "renew_consent_before") DateTime? renewConsentBefore,
      @JsonKey(name: "error_message") String? errorMessage,
      @JsonKey(name: "last_successful_sync") DateTime? lastSuccessfulSync,
      List<Account> accounts,
      Connector connector});

  @override
  $ConnectorCopyWith<$Res> get connector;
}

/// @nodoc
class __$$ConnectionImplCopyWithImpl<$Res>
    extends _$ConnectionCopyWithImpl<$Res, _$ConnectionImpl>
    implements _$$ConnectionImplCopyWith<$Res> {
  __$$ConnectionImplCopyWithImpl(
      _$ConnectionImpl _value, $Res Function(_$ConnectionImpl) _then)
      : super(_value, _then);

  /// Create a copy of Connection
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? status = null,
    Object? renewConsentBefore = freezed,
    Object? errorMessage = freezed,
    Object? lastSuccessfulSync = freezed,
    Object? accounts = null,
    Object? connector = null,
  }) {
    return _then(_$ConnectionImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String,
      renewConsentBefore: freezed == renewConsentBefore
          ? _value.renewConsentBefore
          : renewConsentBefore // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      errorMessage: freezed == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as String?,
      lastSuccessfulSync: freezed == lastSuccessfulSync
          ? _value.lastSuccessfulSync
          : lastSuccessfulSync // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      accounts: null == accounts
          ? _value._accounts
          : accounts // ignore: cast_nullable_to_non_nullable
              as List<Account>,
      connector: null == connector
          ? _value.connector
          : connector // ignore: cast_nullable_to_non_nullable
              as Connector,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$ConnectionImpl implements _Connection {
  const _$ConnectionImpl(
      {required this.id,
      required this.status,
      @JsonKey(name: "renew_consent_before") this.renewConsentBefore,
      @JsonKey(name: "error_message") this.errorMessage,
      @JsonKey(name: "last_successful_sync") this.lastSuccessfulSync,
      required final List<Account> accounts,
      required this.connector})
      : _accounts = accounts;

  factory _$ConnectionImpl.fromJson(Map<String, dynamic> json) =>
      _$$ConnectionImplFromJson(json);

  @override
  final int id;
  @override
  final String status;
  @override
  @JsonKey(name: "renew_consent_before")
  final DateTime? renewConsentBefore;
  @override
  @JsonKey(name: "error_message")
  final String? errorMessage;
  @override
  @JsonKey(name: "last_successful_sync")
  final DateTime? lastSuccessfulSync;
  final List<Account> _accounts;
  @override
  List<Account> get accounts {
    if (_accounts is EqualUnmodifiableListView) return _accounts;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_accounts);
  }

  @override
  final Connector connector;

  @override
  String toString() {
    return 'Connection(id: $id, status: $status, renewConsentBefore: $renewConsentBefore, errorMessage: $errorMessage, lastSuccessfulSync: $lastSuccessfulSync, accounts: $accounts, connector: $connector)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ConnectionImpl &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.status, status) || other.status == status) &&
            (identical(other.renewConsentBefore, renewConsentBefore) ||
                other.renewConsentBefore == renewConsentBefore) &&
            (identical(other.errorMessage, errorMessage) ||
                other.errorMessage == errorMessage) &&
            (identical(other.lastSuccessfulSync, lastSuccessfulSync) ||
                other.lastSuccessfulSync == lastSuccessfulSync) &&
            const DeepCollectionEquality().equals(other._accounts, _accounts) &&
            (identical(other.connector, connector) ||
                other.connector == connector));
  }

  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      id,
      status,
      renewConsentBefore,
      errorMessage,
      lastSuccessfulSync,
      const DeepCollectionEquality().hash(_accounts),
      connector);

  /// Create a copy of Connection
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$ConnectionImplCopyWith<_$ConnectionImpl> get copyWith =>
      __$$ConnectionImplCopyWithImpl<_$ConnectionImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$ConnectionImplToJson(
      this,
    );
  }
}

abstract class _Connection implements Connection {
  const factory _Connection(
      {required final int id,
      required final String status,
      @JsonKey(name: "renew_consent_before") final DateTime? renewConsentBefore,
      @JsonKey(name: "error_message") final String? errorMessage,
      @JsonKey(name: "last_successful_sync") final DateTime? lastSuccessfulSync,
      required final List<Account> accounts,
      required final Connector connector}) = _$ConnectionImpl;

  factory _Connection.fromJson(Map<String, dynamic> json) =
      _$ConnectionImpl.fromJson;

  @override
  int get id;
  @override
  String get status;
  @override
  @JsonKey(name: "renew_consent_before")
  DateTime? get renewConsentBefore;
  @override
  @JsonKey(name: "error_message")
  String? get errorMessage;
  @override
  @JsonKey(name: "last_successful_sync")
  DateTime? get lastSuccessfulSync;
  @override
  List<Account> get accounts;
  @override
  Connector get connector;

  /// Create a copy of Connection
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$ConnectionImplCopyWith<_$ConnectionImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

Account _$AccountFromJson(Map<String, dynamic> json) {
  return _Account.fromJson(json);
}

/// @nodoc
mixin _$Account {
  int get id => throw _privateConstructorUsedError;
  @JsonKey(name: "bank_account_id")
  String? get bankAccountId => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;
  String? get type => throw _privateConstructorUsedError;
  String? get currency => throw _privateConstructorUsedError;
  @JsonKey(name: "account_number")
  String? get accountNumber => throw _privateConstructorUsedError;
  double get balance => throw _privateConstructorUsedError;

  /// Serializes this Account to a JSON map.
  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;

  /// Create a copy of Account
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $AccountCopyWith<Account> get copyWith => throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $AccountCopyWith<$Res> {
  factory $AccountCopyWith(Account value, $Res Function(Account) then) =
      _$AccountCopyWithImpl<$Res, Account>;
  @useResult
  $Res call(
      {int id,
      @JsonKey(name: "bank_account_id") String? bankAccountId,
      String name,
      String? type,
      String? currency,
      @JsonKey(name: "account_number") String? accountNumber,
      double balance});
}

/// @nodoc
class _$AccountCopyWithImpl<$Res, $Val extends Account>
    implements $AccountCopyWith<$Res> {
  _$AccountCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of Account
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? bankAccountId = freezed,
    Object? name = null,
    Object? type = freezed,
    Object? currency = freezed,
    Object? accountNumber = freezed,
    Object? balance = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      bankAccountId: freezed == bankAccountId
          ? _value.bankAccountId
          : bankAccountId // ignore: cast_nullable_to_non_nullable
              as String?,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      type: freezed == type
          ? _value.type
          : type // ignore: cast_nullable_to_non_nullable
              as String?,
      currency: freezed == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String?,
      accountNumber: freezed == accountNumber
          ? _value.accountNumber
          : accountNumber // ignore: cast_nullable_to_non_nullable
              as String?,
      balance: null == balance
          ? _value.balance
          : balance // ignore: cast_nullable_to_non_nullable
              as double,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$AccountImplCopyWith<$Res> implements $AccountCopyWith<$Res> {
  factory _$$AccountImplCopyWith(
          _$AccountImpl value, $Res Function(_$AccountImpl) then) =
      __$$AccountImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {int id,
      @JsonKey(name: "bank_account_id") String? bankAccountId,
      String name,
      String? type,
      String? currency,
      @JsonKey(name: "account_number") String? accountNumber,
      double balance});
}

/// @nodoc
class __$$AccountImplCopyWithImpl<$Res>
    extends _$AccountCopyWithImpl<$Res, _$AccountImpl>
    implements _$$AccountImplCopyWith<$Res> {
  __$$AccountImplCopyWithImpl(
      _$AccountImpl _value, $Res Function(_$AccountImpl) _then)
      : super(_value, _then);

  /// Create a copy of Account
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? bankAccountId = freezed,
    Object? name = null,
    Object? type = freezed,
    Object? currency = freezed,
    Object? accountNumber = freezed,
    Object? balance = null,
  }) {
    return _then(_$AccountImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      bankAccountId: freezed == bankAccountId
          ? _value.bankAccountId
          : bankAccountId // ignore: cast_nullable_to_non_nullable
              as String?,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      type: freezed == type
          ? _value.type
          : type // ignore: cast_nullable_to_non_nullable
              as String?,
      currency: freezed == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String?,
      accountNumber: freezed == accountNumber
          ? _value.accountNumber
          : accountNumber // ignore: cast_nullable_to_non_nullable
              as String?,
      balance: null == balance
          ? _value.balance
          : balance // ignore: cast_nullable_to_non_nullable
              as double,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$AccountImpl implements _Account {
  const _$AccountImpl(
      {required this.id,
      @JsonKey(name: "bank_account_id") this.bankAccountId,
      required this.name,
      this.type,
      this.currency,
      @JsonKey(name: "account_number") this.accountNumber,
      required this.balance});

  factory _$AccountImpl.fromJson(Map<String, dynamic> json) =>
      _$$AccountImplFromJson(json);

  @override
  final int id;
  @override
  @JsonKey(name: "bank_account_id")
  final String? bankAccountId;
  @override
  final String name;
  @override
  final String? type;
  @override
  final String? currency;
  @override
  @JsonKey(name: "account_number")
  final String? accountNumber;
  @override
  final double balance;

  @override
  String toString() {
    return 'Account(id: $id, bankAccountId: $bankAccountId, name: $name, type: $type, currency: $currency, accountNumber: $accountNumber, balance: $balance)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$AccountImpl &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.bankAccountId, bankAccountId) ||
                other.bankAccountId == bankAccountId) &&
            (identical(other.name, name) || other.name == name) &&
            (identical(other.type, type) || other.type == type) &&
            (identical(other.currency, currency) ||
                other.currency == currency) &&
            (identical(other.accountNumber, accountNumber) ||
                other.accountNumber == accountNumber) &&
            (identical(other.balance, balance) || other.balance == balance));
  }

  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  int get hashCode => Object.hash(runtimeType, id, bankAccountId, name, type,
      currency, accountNumber, balance);

  /// Create a copy of Account
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$AccountImplCopyWith<_$AccountImpl> get copyWith =>
      __$$AccountImplCopyWithImpl<_$AccountImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$AccountImplToJson(
      this,
    );
  }
}

abstract class _Account implements Account {
  const factory _Account(
      {required final int id,
      @JsonKey(name: "bank_account_id") final String? bankAccountId,
      required final String name,
      final String? type,
      final String? currency,
      @JsonKey(name: "account_number") final String? accountNumber,
      required final double balance}) = _$AccountImpl;

  factory _Account.fromJson(Map<String, dynamic> json) = _$AccountImpl.fromJson;

  @override
  int get id;
  @override
  @JsonKey(name: "bank_account_id")
  String? get bankAccountId;
  @override
  String get name;
  @override
  String? get type;
  @override
  String? get currency;
  @override
  @JsonKey(name: "account_number")
  String? get accountNumber;
  @override
  double get balance;

  /// Create a copy of Account
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$AccountImplCopyWith<_$AccountImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

Connector _$ConnectorFromJson(Map<String, dynamic> json) {
  return _Connector.fromJson(json);
}

/// @nodoc
mixin _$Connector {
  int get id => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;
  @JsonKey(name: "logo_url")
  String get logoUrl => throw _privateConstructorUsedError;

  /// Serializes this Connector to a JSON map.
  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;

  /// Create a copy of Connector
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $ConnectorCopyWith<Connector> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $ConnectorCopyWith<$Res> {
  factory $ConnectorCopyWith(Connector value, $Res Function(Connector) then) =
      _$ConnectorCopyWithImpl<$Res, Connector>;
  @useResult
  $Res call({int id, String name, @JsonKey(name: "logo_url") String logoUrl});
}

/// @nodoc
class _$ConnectorCopyWithImpl<$Res, $Val extends Connector>
    implements $ConnectorCopyWith<$Res> {
  _$ConnectorCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of Connector
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? name = null,
    Object? logoUrl = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      logoUrl: null == logoUrl
          ? _value.logoUrl
          : logoUrl // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$ConnectorImplCopyWith<$Res>
    implements $ConnectorCopyWith<$Res> {
  factory _$$ConnectorImplCopyWith(
          _$ConnectorImpl value, $Res Function(_$ConnectorImpl) then) =
      __$$ConnectorImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({int id, String name, @JsonKey(name: "logo_url") String logoUrl});
}

/// @nodoc
class __$$ConnectorImplCopyWithImpl<$Res>
    extends _$ConnectorCopyWithImpl<$Res, _$ConnectorImpl>
    implements _$$ConnectorImplCopyWith<$Res> {
  __$$ConnectorImplCopyWithImpl(
      _$ConnectorImpl _value, $Res Function(_$ConnectorImpl) _then)
      : super(_value, _then);

  /// Create a copy of Connector
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? name = null,
    Object? logoUrl = null,
  }) {
    return _then(_$ConnectorImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      logoUrl: null == logoUrl
          ? _value.logoUrl
          : logoUrl // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$ConnectorImpl implements _Connector {
  const _$ConnectorImpl(
      {required this.id,
      required this.name,
      @JsonKey(name: "logo_url") required this.logoUrl});

  factory _$ConnectorImpl.fromJson(Map<String, dynamic> json) =>
      _$$ConnectorImplFromJson(json);

  @override
  final int id;
  @override
  final String name;
  @override
  @JsonKey(name: "logo_url")
  final String logoUrl;

  @override
  String toString() {
    return 'Connector(id: $id, name: $name, logoUrl: $logoUrl)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ConnectorImpl &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.name, name) || other.name == name) &&
            (identical(other.logoUrl, logoUrl) || other.logoUrl == logoUrl));
  }

  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  int get hashCode => Object.hash(runtimeType, id, name, logoUrl);

  /// Create a copy of Connector
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$ConnectorImplCopyWith<_$ConnectorImpl> get copyWith =>
      __$$ConnectorImplCopyWithImpl<_$ConnectorImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$ConnectorImplToJson(
      this,
    );
  }
}

abstract class _Connector implements Connector {
  const factory _Connector(
          {required final int id,
          required final String name,
          @JsonKey(name: "logo_url") required final String logoUrl}) =
      _$ConnectorImpl;

  factory _Connector.fromJson(Map<String, dynamic> json) =
      _$ConnectorImpl.fromJson;

  @override
  int get id;
  @override
  String get name;
  @override
  @JsonKey(name: "logo_url")
  String get logoUrl;

  /// Create a copy of Connector
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$ConnectorImplCopyWith<_$ConnectorImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
