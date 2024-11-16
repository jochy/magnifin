part of 'transactions_cubit.dart';

@freezed
class TransactionsState with _$TransactionsState {
  const factory TransactionsState({
    required List<Transaction> transactions,
    required List<Category> categories,
    DateTime? minDate,
    DateTime? maxDate,
    required bool isLoading,
    required bool hasLoaded,
    String? error,
    @JsonKey(includeToJson: false) required List<MonthlyBudget> loadedMonths,
  }) = _TransactionsState;
}

@freezed
class Transaction with _$Transaction {
  const factory Transaction({
    @JsonKey(name: 'id') required int id,
    @JsonKey(name: 'aid') required int accountId,
    @JsonKey(name: 'bid') String? bankTransactionId,
    @JsonKey(name: 'a') required double amount,
    @JsonKey(name: 'c') required String currency,
    @JsonKey(name: 'd') required String direction,
    @JsonKey(name: 's') required String status,
    @JsonKey(name: 'at') required DateTime operationAt,
    @JsonKey(name: 'name') String? counterpartyName,
    @JsonKey(name: 'acc') String? counterpartyAccount,
    @JsonKey(name: 'ref') String? reference,
    @JsonKey(name: 'logo') String? counterpartyLogoUrl,
    @JsonKey(name: 'ca') int? category,
    @JsonKey(name: 'm') String? method,
  }) = _Transaction;

  factory Transaction.fromJson(Map<String, dynamic> json) =>
      _$TransactionFromJson(json);
}

@freezed
class Category with _$Category {
  const factory Category({
    required int id,
    required String name,
    @JsonKey(name: "uid") int? userId,
    required String icon,
    required String color,
    @JsonKey(name: "include_in_budget") required bool includeInBudget,
  }) = _Category;

  factory Category.fromJson(Map<String, dynamic> json) =>
      _$CategoryFromJson(json);

  @override
  bool operator ==(Object other) {
    if (other is Category) {
      return id == other.id;
    }
    return super == other;
  }
}

class MonthlyBudget {
  final int month;
  final int year;

  MonthlyBudget({required this.month, required this.year});

  static List<MonthlyBudget> fromMinMax(DateTime min, DateTime max) {
    var result = <MonthlyBudget>[];
    var current = DateTime(max.year, max.month);
    while (current.isAfter(min)) {
      result.add(MonthlyBudget(month: current.month, year: current.year));
      current = DateTime(current.year, current.month - 1);
    }
    return result;
  }

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;

    return other is MonthlyBudget && other.month == month && other.year == year;
  }

  MonthlyBudget next() => MonthlyBudget(
      month: month == 12 ? 1 : month + 1, year: month == 12 ? year + 1 : year);

  MonthlyBudget previous() => MonthlyBudget(
      month: month == 1 ? 12 : month - 1, year: month == 1 ? year - 1 : year);

  MonthlyBudgetSummary loadIndicatorsFromTransactions(List<Transaction> trs) {
    var transactions = trs
        .where(
            (t) => t.operationAt.month == month && t.operationAt.year == year)
        .toList();
    transactions.sort((t, t2) => t.operationAt.compareTo(t2.operationAt));

    return MonthlyBudgetSummary(month: this, transactions: transactions);
  }

  @override
  int get hashCode => "$month-$year".hashCode;
}

class MonthlyBudgetSummary {
  final MonthlyBudget month;
  final List<Transaction> transactions;
  final Map<DateTime, List<Transaction>> groupedTransactions = {};
  final Map<Category, double> categoryTotals = {};
  late final double income;
  late final double expenses;
  late final double remaining;

  MonthlyBudgetSummary({
    required this.month,
    required this.transactions,
  }) {
    try {
      income = transactions
          .where((t) => t.direction == 'CREDIT')
          .map((t) => t.amount)
          .fold(0, (a, b) => a + b);
    } catch (e) {
      income = 0;
    }
    try {
      expenses = transactions
          .where((t) => t.direction == 'DEBIT')
          .map((t) => t.amount)
          .fold(0, (a, b) => a + b);
    } catch (e) {
      expenses = 0;
    }
    remaining = income - expenses;

    transactions.sort((a, b) => a.operationAt.compareTo(b.operationAt));

    DateTime lastDate = DateTime.utc(1970);
    for (var t in transactions) {
      if (t.operationAt.isAtSameDayAs(lastDate)) {
        groupedTransactions[lastDate]!.add(t);
      } else {
        lastDate = t.operationAt;
        groupedTransactions[lastDate] = [t];
      }
    }
  }

  List<MapEntry<DateTime, List<Transaction>>> get groupedTransactionsEntries {
    var l = groupedTransactions.entries.toList();
    l.sort((a, b) => a.key.compareTo(b.key));
    return l;
  }
}

extension TransactionExt on Transaction {
  Category? getCategory(TransactionsState state) => state.categories
      .where((c) => c.id == category).firstOrNull;
}