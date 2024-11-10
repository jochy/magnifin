import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:front/config.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:front/screens/accounts/accounts_screen.dart';
import 'package:http/http.dart' as http;
import 'package:moment_dart/moment_dart.dart';

part 'transactions_state.dart';

part 'transactions_cubit.freezed.dart';

part 'transactions_cubit.g.dart';

class TransactionsCubit extends Cubit<TransactionsState> {
  final AuthCubit authCubit;

  static TransactionsCubit of(context) =>
      BlocProvider.of<TransactionsCubit>(context);

  TransactionsCubit(this.authCubit)
      : super(const TransactionsState(
          transactions: [],
          isLoading: false,
          hasLoaded: false,
          error: null,
          loadedMonths: [],
        ));

  void loadRecentTransactions() async {
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
      var allTransactions = [
        ...state.transactions,
        ...data['transactions']
            .map<Transaction>((t) => Transaction.fromJson(t)),
      ];

      // Let's remove duplicated transactions
      var mapTrs = <int, Transaction>{};
      for (var t in allTransactions) {
        mapTrs[t.id] = t;
      }

      emit(
        state.copyWith(
          transactions: mapTrs.values.toList(),
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
}
