import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:front/config.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:http/http.dart' as http;
import 'package:moment_dart/moment_dart.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

part 'transactions_state.dart';

part 'transactions_cubit.freezed.dart';

part 'transactions_cubit.g.dart';

class TransactionsCubit extends Cubit<TransactionsState> {
  final AuthCubit authCubit;
  WebSocketChannel? channel;

  static TransactionsCubit of(context) =>
      BlocProvider.of<TransactionsCubit>(context);

  TransactionsCubit(this.authCubit)
      : super(const TransactionsState(
          transactions: [],
          categories: [],
          isLoading: false,
          hasLoaded: false,
          error: null,
          loadedMonths: [],
        ));

  @override
  Future<void> close() {
    if (channel != null) {
      channel!.sink.close();
    }
    return super.close();
  }

  void loadRecentTransactions() async {
    if (channel == null) {
      channel = WebSocketChannel.connect(
        Uri.parse(
            "${Configuration.instance.baseUrl.replaceFirst("http", "ws")}/v1/ws?token=${authCubit.state.token}"),
      );
      channel!.stream.listen((d) {
        var data = jsonDecode(d);
        var transaction = Transaction.fromJson(data);
        emit(state.copyWith(
          transactions: cleanTransactions([
            ...state.transactions,
            transaction,
          ]),
        ));
      });
    }
    emit(state.copyWith(isLoading: true, error: null));
    // Load min-max

    final response = await http.get(
      Uri.parse("${Configuration.instance.baseUrl}/v1/transactions/minmax"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
        'Authorization': authCubit.state.token ?? "",
      },
    );

    if (response.statusCode != 200) {
      emit(
        state.copyWith(
          isLoading: false,
          hasLoaded: false,
          error: 'Failed to load transactions: ${response.statusCode}',
        ),
      );
      return;
    }

    final Map<String, dynamic> data = jsonDecode(response.body);
    var minDate = data['min'];
    var maxDate = data['max'];
    if (minDate == null || maxDate == null) {
      emit(
        state.copyWith(
          isLoading: false,
          hasLoaded: true,
        ),
      );
    } else {
      emit(
        state.copyWith(
          minDate: DateTime.parse(minDate),
          maxDate: DateTime.parse(maxDate),
        ),
      );

      for (var i = 0; i < 3; i++) {
        var year = state.maxDate!.month == 1
            ? state.maxDate!.year - 1
            : state.maxDate!.year;
        var month = state.maxDate!.month == 1 ? 12 : state.maxDate!.month - i;

        await loadMonthTransactions(year, month);
        if (state.error != null) {
          break;
        }
      }

      await loadCategories();
    }
  }

  Future<void> loadCategories() async {
    emit(state.copyWith(isLoading: true, error: null));

    try {
      final response = await http.get(
        Uri.parse("${Configuration.instance.baseUrl}/v1/categories"),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
          'Accept': 'application/json',
          'Authorization': authCubit.state.token ?? "",
        },
      );

      if (response.statusCode != 200) {
        emit(
          state.copyWith(
            isLoading: false,
            hasLoaded: false,
            error: 'Failed to load categories: ${response.statusCode}',
          ),
        );
        return;
      }

      final Map<String, dynamic> data = jsonDecode(response.body);
      var categories = data["categories"]
          .map<Category>((c) => Category.fromJson(c))
          .toList();

      emit(
        state.copyWith(
          categories: categories,
          isLoading: false,
          hasLoaded: true,
        ),
      );
    } catch (e) {
      emit(state.copyWith(
        isLoading: false,
        error: e.toString(),
      ));
    }
  }

  Future<void> loadMonthTransactions(int year, int month) async {
    var from = DateTime.utc(year, month, 1);
    var toYear = month == 12 ? year + 1 : year;
    var toMonth = month == 12 ? 1 : month + 1;
    var to = DateTime.utc(toYear, toMonth, 1)
        .subtract(const Duration(milliseconds: 1));

    emit(state.copyWith(isLoading: true, error: null));

    try {
      final response = await http.get(
        Uri.parse(
            "${Configuration.instance.baseUrl}/v1/transactions?from=${from.toIso8601String()}&to=${to.toIso8601String()}"),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
          'Accept': 'application/json',
          'Authorization': authCubit.state.token ?? "",
        },
      );

      if (response.statusCode != 200) {
        emit(
          state.copyWith(
            isLoading: false,
            hasLoaded: false,
            error: 'Failed to load transactions: ${response.statusCode}',
          ),
        );
        return;
      }

      final Map<String, dynamic> data = jsonDecode(response.body);
      var allTransactions = <Transaction>[
        ...state.transactions,
        ...data['transactions']
            .map<Transaction>((t) => Transaction.fromJson(t)),
      ];

      emit(
        state.copyWith(
          transactions: cleanTransactions(allTransactions),
          isLoading: false,
          hasLoaded: true,
          loadedMonths: [
            ...state.loadedMonths,
            MonthlyBudget(month: month, year: year),
          ],
        ),
      );
    } catch (e) {
      emit(state.copyWith(
        isLoading: false,
        error: e.toString(),
      ));
    }
  }

  void updateTransaction(
    Transaction transaction,
    Category? category,
    String counterparty,
  ) async {
    final response = await http.patch(
      Uri.parse(
          "${Configuration.instance.baseUrl}/v1/transactions/${transaction.id}"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
        'Authorization': authCubit.state.token ?? "",
      },
      body: jsonEncode({
        'id': transaction.id,
        'category_id': category?.id,
        'user_counterparty': counterparty,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to update transaction: ${response.statusCode}');
    }

    final Map<String, dynamic> data = jsonDecode(response.body);
    var updatedTransaction = Transaction.fromJson(data);

    emit(state.copyWith(
      transactions: [
        for (var t in state.transactions)
          if (t.id == transaction.id) updatedTransaction else t
      ],
    ));
  }

  Future<bool> createCategoryRule(
      Category category, List<String> keywords) async {
    final response = await http.post(
      Uri.parse(
          "${Configuration.instance.baseUrl}/v1/categories/${category.id}/rule"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
        'Authorization': authCubit.state.token ?? "",
      },
      body: jsonEncode({
        'category_id': category.id,
        'keywords': keywords,
      }),
    );

    return response.statusCode == 204;
  }
}

List<Transaction> cleanTransactions(List<Transaction> transactions) {
  var mapTrs = <int, Transaction>{};
  for (var t in transactions) {
    mapTrs[t.id] = t;
  }

  var trs = mapTrs.values.toList();
  trs.sort((a, b) {
    if (a.operationAt.isAtSameMomentAs(b.operationAt)) {
      return a.id.compareTo(b.id);
    }

    return a.operationAt.compareTo(b.operationAt);
  });

  return trs;
}
