// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'transactions_cubit.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$TransactionsState {
  List<Transaction> get transactions => throw _privateConstructorUsedError;
  DateTime? get minDate => throw _privateConstructorUsedError;
  DateTime? get maxDate => throw _privateConstructorUsedError;
  bool get isLoading => throw _privateConstructorUsedError;
  bool get hasLoaded => throw _privateConstructorUsedError;
  String? get error => throw _privateConstructorUsedError;
  @JsonKey(includeToJson: false)
  List<MonthlyBudget> get loadedMonths => throw _privateConstructorUsedError;

  /// Create a copy of TransactionsState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $TransactionsStateCopyWith<TransactionsState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TransactionsStateCopyWith<$Res> {
  factory $TransactionsStateCopyWith(
          TransactionsState value, $Res Function(TransactionsState) then) =
      _$TransactionsStateCopyWithImpl<$Res, TransactionsState>;
  @useResult
  $Res call(
      {List<Transaction> transactions,
      DateTime? minDate,
      DateTime? maxDate,
      bool isLoading,
      bool hasLoaded,
      String? error,
      @JsonKey(includeToJson: false) List<MonthlyBudget> loadedMonths});
}

/// @nodoc
class _$TransactionsStateCopyWithImpl<$Res, $Val extends TransactionsState>
    implements $TransactionsStateCopyWith<$Res> {
  _$TransactionsStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of TransactionsState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? transactions = null,
    Object? minDate = freezed,
    Object? maxDate = freezed,
    Object? isLoading = null,
    Object? hasLoaded = null,
    Object? error = freezed,
    Object? loadedMonths = null,
  }) {
    return _then(_value.copyWith(
      transactions: null == transactions
          ? _value.transactions
          : transactions // ignore: cast_nullable_to_non_nullable
              as List<Transaction>,
      minDate: freezed == minDate
          ? _value.minDate
          : minDate // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      maxDate: freezed == maxDate
          ? _value.maxDate
          : maxDate // ignore: cast_nullable_to_non_nullable
              as DateTime?,
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
      loadedMonths: null == loadedMonths
          ? _value.loadedMonths
          : loadedMonths // ignore: cast_nullable_to_non_nullable
              as List<MonthlyBudget>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$TransactionsStateImplCopyWith<$Res>
    implements $TransactionsStateCopyWith<$Res> {
  factory _$$TransactionsStateImplCopyWith(_$TransactionsStateImpl value,
          $Res Function(_$TransactionsStateImpl) then) =
      __$$TransactionsStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {List<Transaction> transactions,
      DateTime? minDate,
      DateTime? maxDate,
      bool isLoading,
      bool hasLoaded,
      String? error,
      @JsonKey(includeToJson: false) List<MonthlyBudget> loadedMonths});
}

/// @nodoc
class __$$TransactionsStateImplCopyWithImpl<$Res>
    extends _$TransactionsStateCopyWithImpl<$Res, _$TransactionsStateImpl>
    implements _$$TransactionsStateImplCopyWith<$Res> {
  __$$TransactionsStateImplCopyWithImpl(_$TransactionsStateImpl _value,
      $Res Function(_$TransactionsStateImpl) _then)
      : super(_value, _then);

  /// Create a copy of TransactionsState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? transactions = null,
    Object? minDate = freezed,
    Object? maxDate = freezed,
    Object? isLoading = null,
    Object? hasLoaded = null,
    Object? error = freezed,
    Object? loadedMonths = null,
  }) {
    return _then(_$TransactionsStateImpl(
      transactions: null == transactions
          ? _value._transactions
          : transactions // ignore: cast_nullable_to_non_nullable
              as List<Transaction>,
      minDate: freezed == minDate
          ? _value.minDate
          : minDate // ignore: cast_nullable_to_non_nullable
              as DateTime?,
      maxDate: freezed == maxDate
          ? _value.maxDate
          : maxDate // ignore: cast_nullable_to_non_nullable
              as DateTime?,
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
      loadedMonths: null == loadedMonths
          ? _value._loadedMonths
          : loadedMonths // ignore: cast_nullable_to_non_nullable
              as List<MonthlyBudget>,
    ));
  }
}

/// @nodoc

class _$TransactionsStateImpl implements _TransactionsState {
  const _$TransactionsStateImpl(
      {required final List<Transaction> transactions,
      this.minDate,
      this.maxDate,
      required this.isLoading,
      required this.hasLoaded,
      this.error,
      @JsonKey(includeToJson: false)
      required final List<MonthlyBudget> loadedMonths})
      : _transactions = transactions,
        _loadedMonths = loadedMonths;

  final List<Transaction> _transactions;
  @override
  List<Transaction> get transactions {
    if (_transactions is EqualUnmodifiableListView) return _transactions;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_transactions);
  }

  @override
  final DateTime? minDate;
  @override
  final DateTime? maxDate;
  @override
  final bool isLoading;
  @override
  final bool hasLoaded;
  @override
  final String? error;
  final List<MonthlyBudget> _loadedMonths;
  @override
  @JsonKey(includeToJson: false)
  List<MonthlyBudget> get loadedMonths {
    if (_loadedMonths is EqualUnmodifiableListView) return _loadedMonths;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_loadedMonths);
  }

  @override
  String toString() {
    return 'TransactionsState(transactions: $transactions, minDate: $minDate, maxDate: $maxDate, isLoading: $isLoading, hasLoaded: $hasLoaded, error: $error, loadedMonths: $loadedMonths)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$TransactionsStateImpl &&
            const DeepCollectionEquality()
                .equals(other._transactions, _transactions) &&
            (identical(other.minDate, minDate) || other.minDate == minDate) &&
            (identical(other.maxDate, maxDate) || other.maxDate == maxDate) &&
            (identical(other.isLoading, isLoading) ||
                other.isLoading == isLoading) &&
            (identical(other.hasLoaded, hasLoaded) ||
                other.hasLoaded == hasLoaded) &&
            (identical(other.error, error) || other.error == error) &&
            const DeepCollectionEquality()
                .equals(other._loadedMonths, _loadedMonths));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(_transactions),
      minDate,
      maxDate,
      isLoading,
      hasLoaded,
      error,
      const DeepCollectionEquality().hash(_loadedMonths));

  /// Create a copy of TransactionsState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$TransactionsStateImplCopyWith<_$TransactionsStateImpl> get copyWith =>
      __$$TransactionsStateImplCopyWithImpl<_$TransactionsStateImpl>(
          this, _$identity);
}

abstract class _TransactionsState implements TransactionsState {
  const factory _TransactionsState(
          {required final List<Transaction> transactions,
          final DateTime? minDate,
          final DateTime? maxDate,
          required final bool isLoading,
          required final bool hasLoaded,
          final String? error,
          @JsonKey(includeToJson: false)
          required final List<MonthlyBudget> loadedMonths}) =
      _$TransactionsStateImpl;

  @override
  List<Transaction> get transactions;
  @override
  DateTime? get minDate;
  @override
  DateTime? get maxDate;
  @override
  bool get isLoading;
  @override
  bool get hasLoaded;
  @override
  String? get error;
  @override
  @JsonKey(includeToJson: false)
  List<MonthlyBudget> get loadedMonths;

  /// Create a copy of TransactionsState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$TransactionsStateImplCopyWith<_$TransactionsStateImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

Transaction _$TransactionFromJson(Map<String, dynamic> json) {
  return _Transaction.fromJson(json);
}

/// @nodoc
mixin _$Transaction {
  @JsonKey(name: 'id')
  int get id => throw _privateConstructorUsedError;
  @JsonKey(name: 'aid')
  int get accountId => throw _privateConstructorUsedError;
  @JsonKey(name: 'bid')
  String? get bankTransactionId => throw _privateConstructorUsedError;
  @JsonKey(name: 'a')
  double get amount => throw _privateConstructorUsedError;
  @JsonKey(name: 'c')
  String get currency => throw _privateConstructorUsedError;
  @JsonKey(name: 'd')
  String get direction => throw _privateConstructorUsedError;
  @JsonKey(name: 's')
  String get status => throw _privateConstructorUsedError;
  @JsonKey(name: 'at')
  DateTime get operationAt => throw _privateConstructorUsedError;
  @JsonKey(name: 'name')
  String? get counterpartyName => throw _privateConstructorUsedError;
  @JsonKey(name: 'acc')
  String? get counterpartyAccount => throw _privateConstructorUsedError;
  @JsonKey(name: 'ref')
  String? get reference => throw _privateConstructorUsedError;
  @JsonKey(name: 'logo')
  String? get counterpartyLogoUrl => throw _privateConstructorUsedError;
  @JsonKey(name: 'ca')
  int? get category => throw _privateConstructorUsedError;
  @JsonKey(name: 'm')
  String? get method => throw _privateConstructorUsedError;

  /// Serializes this Transaction to a JSON map.
  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;

  /// Create a copy of Transaction
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $TransactionCopyWith<Transaction> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TransactionCopyWith<$Res> {
  factory $TransactionCopyWith(
          Transaction value, $Res Function(Transaction) then) =
      _$TransactionCopyWithImpl<$Res, Transaction>;
  @useResult
  $Res call(
      {@JsonKey(name: 'id') int id,
      @JsonKey(name: 'aid') int accountId,
      @JsonKey(name: 'bid') String? bankTransactionId,
      @JsonKey(name: 'a') double amount,
      @JsonKey(name: 'c') String currency,
      @JsonKey(name: 'd') String direction,
      @JsonKey(name: 's') String status,
      @JsonKey(name: 'at') DateTime operationAt,
      @JsonKey(name: 'name') String? counterpartyName,
      @JsonKey(name: 'acc') String? counterpartyAccount,
      @JsonKey(name: 'ref') String? reference,
      @JsonKey(name: 'logo') String? counterpartyLogoUrl,
      @JsonKey(name: 'ca') int? category,
      @JsonKey(name: 'm') String? method});
}

/// @nodoc
class _$TransactionCopyWithImpl<$Res, $Val extends Transaction>
    implements $TransactionCopyWith<$Res> {
  _$TransactionCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of Transaction
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? accountId = null,
    Object? bankTransactionId = freezed,
    Object? amount = null,
    Object? currency = null,
    Object? direction = null,
    Object? status = null,
    Object? operationAt = null,
    Object? counterpartyName = freezed,
    Object? counterpartyAccount = freezed,
    Object? reference = freezed,
    Object? counterpartyLogoUrl = freezed,
    Object? category = freezed,
    Object? method = freezed,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      accountId: null == accountId
          ? _value.accountId
          : accountId // ignore: cast_nullable_to_non_nullable
              as int,
      bankTransactionId: freezed == bankTransactionId
          ? _value.bankTransactionId
          : bankTransactionId // ignore: cast_nullable_to_non_nullable
              as String?,
      amount: null == amount
          ? _value.amount
          : amount // ignore: cast_nullable_to_non_nullable
              as double,
      currency: null == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String,
      direction: null == direction
          ? _value.direction
          : direction // ignore: cast_nullable_to_non_nullable
              as String,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String,
      operationAt: null == operationAt
          ? _value.operationAt
          : operationAt // ignore: cast_nullable_to_non_nullable
              as DateTime,
      counterpartyName: freezed == counterpartyName
          ? _value.counterpartyName
          : counterpartyName // ignore: cast_nullable_to_non_nullable
              as String?,
      counterpartyAccount: freezed == counterpartyAccount
          ? _value.counterpartyAccount
          : counterpartyAccount // ignore: cast_nullable_to_non_nullable
              as String?,
      reference: freezed == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String?,
      counterpartyLogoUrl: freezed == counterpartyLogoUrl
          ? _value.counterpartyLogoUrl
          : counterpartyLogoUrl // ignore: cast_nullable_to_non_nullable
              as String?,
      category: freezed == category
          ? _value.category
          : category // ignore: cast_nullable_to_non_nullable
              as int?,
      method: freezed == method
          ? _value.method
          : method // ignore: cast_nullable_to_non_nullable
              as String?,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$TransactionImplCopyWith<$Res>
    implements $TransactionCopyWith<$Res> {
  factory _$$TransactionImplCopyWith(
          _$TransactionImpl value, $Res Function(_$TransactionImpl) then) =
      __$$TransactionImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {@JsonKey(name: 'id') int id,
      @JsonKey(name: 'aid') int accountId,
      @JsonKey(name: 'bid') String? bankTransactionId,
      @JsonKey(name: 'a') double amount,
      @JsonKey(name: 'c') String currency,
      @JsonKey(name: 'd') String direction,
      @JsonKey(name: 's') String status,
      @JsonKey(name: 'at') DateTime operationAt,
      @JsonKey(name: 'name') String? counterpartyName,
      @JsonKey(name: 'acc') String? counterpartyAccount,
      @JsonKey(name: 'ref') String? reference,
      @JsonKey(name: 'logo') String? counterpartyLogoUrl,
      @JsonKey(name: 'ca') int? category,
      @JsonKey(name: 'm') String? method});
}

/// @nodoc
class __$$TransactionImplCopyWithImpl<$Res>
    extends _$TransactionCopyWithImpl<$Res, _$TransactionImpl>
    implements _$$TransactionImplCopyWith<$Res> {
  __$$TransactionImplCopyWithImpl(
      _$TransactionImpl _value, $Res Function(_$TransactionImpl) _then)
      : super(_value, _then);

  /// Create a copy of Transaction
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? accountId = null,
    Object? bankTransactionId = freezed,
    Object? amount = null,
    Object? currency = null,
    Object? direction = null,
    Object? status = null,
    Object? operationAt = null,
    Object? counterpartyName = freezed,
    Object? counterpartyAccount = freezed,
    Object? reference = freezed,
    Object? counterpartyLogoUrl = freezed,
    Object? category = freezed,
    Object? method = freezed,
  }) {
    return _then(_$TransactionImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      accountId: null == accountId
          ? _value.accountId
          : accountId // ignore: cast_nullable_to_non_nullable
              as int,
      bankTransactionId: freezed == bankTransactionId
          ? _value.bankTransactionId
          : bankTransactionId // ignore: cast_nullable_to_non_nullable
              as String?,
      amount: null == amount
          ? _value.amount
          : amount // ignore: cast_nullable_to_non_nullable
              as double,
      currency: null == currency
          ? _value.currency
          : currency // ignore: cast_nullable_to_non_nullable
              as String,
      direction: null == direction
          ? _value.direction
          : direction // ignore: cast_nullable_to_non_nullable
              as String,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as String,
      operationAt: null == operationAt
          ? _value.operationAt
          : operationAt // ignore: cast_nullable_to_non_nullable
              as DateTime,
      counterpartyName: freezed == counterpartyName
          ? _value.counterpartyName
          : counterpartyName // ignore: cast_nullable_to_non_nullable
              as String?,
      counterpartyAccount: freezed == counterpartyAccount
          ? _value.counterpartyAccount
          : counterpartyAccount // ignore: cast_nullable_to_non_nullable
              as String?,
      reference: freezed == reference
          ? _value.reference
          : reference // ignore: cast_nullable_to_non_nullable
              as String?,
      counterpartyLogoUrl: freezed == counterpartyLogoUrl
          ? _value.counterpartyLogoUrl
          : counterpartyLogoUrl // ignore: cast_nullable_to_non_nullable
              as String?,
      category: freezed == category
          ? _value.category
          : category // ignore: cast_nullable_to_non_nullable
              as int?,
      method: freezed == method
          ? _value.method
          : method // ignore: cast_nullable_to_non_nullable
              as String?,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$TransactionImpl implements _Transaction {
  const _$TransactionImpl(
      {@JsonKey(name: 'id') required this.id,
      @JsonKey(name: 'aid') required this.accountId,
      @JsonKey(name: 'bid') this.bankTransactionId,
      @JsonKey(name: 'a') required this.amount,
      @JsonKey(name: 'c') required this.currency,
      @JsonKey(name: 'd') required this.direction,
      @JsonKey(name: 's') required this.status,
      @JsonKey(name: 'at') required this.operationAt,
      @JsonKey(name: 'name') this.counterpartyName,
      @JsonKey(name: 'acc') this.counterpartyAccount,
      @JsonKey(name: 'ref') this.reference,
      @JsonKey(name: 'logo') this.counterpartyLogoUrl,
      @JsonKey(name: 'ca') this.category,
      @JsonKey(name: 'm') this.method});

  factory _$TransactionImpl.fromJson(Map<String, dynamic> json) =>
      _$$TransactionImplFromJson(json);

  @override
  @JsonKey(name: 'id')
  final int id;
  @override
  @JsonKey(name: 'aid')
  final int accountId;
  @override
  @JsonKey(name: 'bid')
  final String? bankTransactionId;
  @override
  @JsonKey(name: 'a')
  final double amount;
  @override
  @JsonKey(name: 'c')
  final String currency;
  @override
  @JsonKey(name: 'd')
  final String direction;
  @override
  @JsonKey(name: 's')
  final String status;
  @override
  @JsonKey(name: 'at')
  final DateTime operationAt;
  @override
  @JsonKey(name: 'name')
  final String? counterpartyName;
  @override
  @JsonKey(name: 'acc')
  final String? counterpartyAccount;
  @override
  @JsonKey(name: 'ref')
  final String? reference;
  @override
  @JsonKey(name: 'logo')
  final String? counterpartyLogoUrl;
  @override
  @JsonKey(name: 'ca')
  final int? category;
  @override
  @JsonKey(name: 'm')
  final String? method;

  @override
  String toString() {
    return 'Transaction(id: $id, accountId: $accountId, bankTransactionId: $bankTransactionId, amount: $amount, currency: $currency, direction: $direction, status: $status, operationAt: $operationAt, counterpartyName: $counterpartyName, counterpartyAccount: $counterpartyAccount, reference: $reference, counterpartyLogoUrl: $counterpartyLogoUrl, category: $category, method: $method)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$TransactionImpl &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.accountId, accountId) ||
                other.accountId == accountId) &&
            (identical(other.bankTransactionId, bankTransactionId) ||
                other.bankTransactionId == bankTransactionId) &&
            (identical(other.amount, amount) || other.amount == amount) &&
            (identical(other.currency, currency) ||
                other.currency == currency) &&
            (identical(other.direction, direction) ||
                other.direction == direction) &&
            (identical(other.status, status) || other.status == status) &&
            (identical(other.operationAt, operationAt) ||
                other.operationAt == operationAt) &&
            (identical(other.counterpartyName, counterpartyName) ||
                other.counterpartyName == counterpartyName) &&
            (identical(other.counterpartyAccount, counterpartyAccount) ||
                other.counterpartyAccount == counterpartyAccount) &&
            (identical(other.reference, reference) ||
                other.reference == reference) &&
            (identical(other.counterpartyLogoUrl, counterpartyLogoUrl) ||
                other.counterpartyLogoUrl == counterpartyLogoUrl) &&
            (identical(other.category, category) ||
                other.category == category) &&
            (identical(other.method, method) || other.method == method));
  }

  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      id,
      accountId,
      bankTransactionId,
      amount,
      currency,
      direction,
      status,
      operationAt,
      counterpartyName,
      counterpartyAccount,
      reference,
      counterpartyLogoUrl,
      category,
      method);

  /// Create a copy of Transaction
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$TransactionImplCopyWith<_$TransactionImpl> get copyWith =>
      __$$TransactionImplCopyWithImpl<_$TransactionImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$TransactionImplToJson(
      this,
    );
  }
}

abstract class _Transaction implements Transaction {
  const factory _Transaction(
      {@JsonKey(name: 'id') required final int id,
      @JsonKey(name: 'aid') required final int accountId,
      @JsonKey(name: 'bid') final String? bankTransactionId,
      @JsonKey(name: 'a') required final double amount,
      @JsonKey(name: 'c') required final String currency,
      @JsonKey(name: 'd') required final String direction,
      @JsonKey(name: 's') required final String status,
      @JsonKey(name: 'at') required final DateTime operationAt,
      @JsonKey(name: 'name') final String? counterpartyName,
      @JsonKey(name: 'acc') final String? counterpartyAccount,
      @JsonKey(name: 'ref') final String? reference,
      @JsonKey(name: 'logo') final String? counterpartyLogoUrl,
      @JsonKey(name: 'ca') final int? category,
      @JsonKey(name: 'm') final String? method}) = _$TransactionImpl;

  factory _Transaction.fromJson(Map<String, dynamic> json) =
      _$TransactionImpl.fromJson;

  @override
  @JsonKey(name: 'id')
  int get id;
  @override
  @JsonKey(name: 'aid')
  int get accountId;
  @override
  @JsonKey(name: 'bid')
  String? get bankTransactionId;
  @override
  @JsonKey(name: 'a')
  double get amount;
  @override
  @JsonKey(name: 'c')
  String get currency;
  @override
  @JsonKey(name: 'd')
  String get direction;
  @override
  @JsonKey(name: 's')
  String get status;
  @override
  @JsonKey(name: 'at')
  DateTime get operationAt;
  @override
  @JsonKey(name: 'name')
  String? get counterpartyName;
  @override
  @JsonKey(name: 'acc')
  String? get counterpartyAccount;
  @override
  @JsonKey(name: 'ref')
  String? get reference;
  @override
  @JsonKey(name: 'logo')
  String? get counterpartyLogoUrl;
  @override
  @JsonKey(name: 'ca')
  int? get category;
  @override
  @JsonKey(name: 'm')
  String? get method;

  /// Create a copy of Transaction
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$TransactionImplCopyWith<_$TransactionImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

Category _$CategoryFromJson(Map<String, dynamic> json) {
  return _Category.fromJson(json);
}

/// @nodoc
mixin _$Category {
  int get id => throw _privateConstructorUsedError;
  String get name => throw _privateConstructorUsedError;
  String get type => throw _privateConstructorUsedError;
  String get color => throw _privateConstructorUsedError;

  /// Serializes this Category to a JSON map.
  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;

  /// Create a copy of Category
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $CategoryCopyWith<Category> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $CategoryCopyWith<$Res> {
  factory $CategoryCopyWith(Category value, $Res Function(Category) then) =
      _$CategoryCopyWithImpl<$Res, Category>;
  @useResult
  $Res call({int id, String name, String type, String color});
}

/// @nodoc
class _$CategoryCopyWithImpl<$Res, $Val extends Category>
    implements $CategoryCopyWith<$Res> {
  _$CategoryCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of Category
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? name = null,
    Object? type = null,
    Object? color = null,
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
      type: null == type
          ? _value.type
          : type // ignore: cast_nullable_to_non_nullable
              as String,
      color: null == color
          ? _value.color
          : color // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$CategoryImplCopyWith<$Res>
    implements $CategoryCopyWith<$Res> {
  factory _$$CategoryImplCopyWith(
          _$CategoryImpl value, $Res Function(_$CategoryImpl) then) =
      __$$CategoryImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({int id, String name, String type, String color});
}

/// @nodoc
class __$$CategoryImplCopyWithImpl<$Res>
    extends _$CategoryCopyWithImpl<$Res, _$CategoryImpl>
    implements _$$CategoryImplCopyWith<$Res> {
  __$$CategoryImplCopyWithImpl(
      _$CategoryImpl _value, $Res Function(_$CategoryImpl) _then)
      : super(_value, _then);

  /// Create a copy of Category
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? name = null,
    Object? type = null,
    Object? color = null,
  }) {
    return _then(_$CategoryImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as int,
      name: null == name
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      type: null == type
          ? _value.type
          : type // ignore: cast_nullable_to_non_nullable
              as String,
      color: null == color
          ? _value.color
          : color // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$CategoryImpl implements _Category {
  const _$CategoryImpl(
      {required this.id,
      required this.name,
      required this.type,
      required this.color});

  factory _$CategoryImpl.fromJson(Map<String, dynamic> json) =>
      _$$CategoryImplFromJson(json);

  @override
  final int id;
  @override
  final String name;
  @override
  final String type;
  @override
  final String color;

  @override
  String toString() {
    return 'Category(id: $id, name: $name, type: $type, color: $color)';
  }

  /// Create a copy of Category
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$CategoryImplCopyWith<_$CategoryImpl> get copyWith =>
      __$$CategoryImplCopyWithImpl<_$CategoryImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$CategoryImplToJson(
      this,
    );
  }
}

abstract class _Category implements Category {
  const factory _Category(
      {required final int id,
      required final String name,
      required final String type,
      required final String color}) = _$CategoryImpl;

  factory _Category.fromJson(Map<String, dynamic> json) =
      _$CategoryImpl.fromJson;

  @override
  int get id;
  @override
  String get name;
  @override
  String get type;
  @override
  String get color;

  /// Create a copy of Category
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$CategoryImplCopyWith<_$CategoryImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
