import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:front/config.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:front/cubit/transactions/transactions_cubit.dart';
import 'package:http/http.dart' as http;

part 'connections_state.dart';

part 'connections_cubit.freezed.dart';

part 'connections_cubit.g.dart';

class ConnectionsCubit extends Cubit<ConnectionsState> {
  final AuthCubit authCubit;
  final TransactionsCubit transactionsCubit;

  static ConnectionsCubit of(context) =>
      BlocProvider.of<ConnectionsCubit>(context);

  ConnectionsCubit(this.authCubit, this.transactionsCubit)
      : super(const ConnectionsState(
          isLoading: false,
          hasLoaded: false,
          connections: [],
        )) {
    authCubit.stream.listen((state) {
      if (state.token != null && !this.state.hasLoaded) {
        loadConnections();
      }
    });
  }

  Future<void> loadConnections() async {
    emit(state.copyWith(isLoading: true));

    try {
      final response = await http.get(
        Uri.parse("${Configuration.instance.baseUrl}/v1/connections"),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
          'Accept': 'application/json',
          'Authorization': authCubit.state.token ?? "",
        },
      );

      if (response.statusCode == 200) {
        final Map<String, dynamic> data = jsonDecode(response.body);
        emit(ConnectionsState(
          connections: data['connections']
              .map<Connection>((c) => Connection.fromJson(c))
              .toList(),
          isLoading: false,
          hasLoaded: true,
        ));

        transactionsCubit.loadRecentTransactions();
      } else {
        emit(
          state.copyWith(
            isLoading: false,
            hasLoaded: false,
            error: 'Failed to load connections: ${response.statusCode}',
          ),
        );
        return;
      }
    } catch (e) {
      emit(state.copyWith(
        isLoading: false,
        hasLoaded: false,
        error: e.toString(),
      ));
    }
  }

  /// Search for connectors
  /// No change of state
  Future<List<Connector>> searchConnectors(String query) async {
    final response = await http.get(
      Uri.parse("${Configuration.instance.baseUrl}/v1/connectors?name=$query"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
        'Authorization': authCubit.state.token ?? "",
      },
    );

    if (response.statusCode == 200) {
      final Map<String, dynamic> data = jsonDecode(response.body);
      return data['connectors']
          .map<Connector>((c) => Connector.fromJson(c))
          .toList();
    } else {
      throw Exception('Failed to load connectors: ${response.statusCode}');
    }
  }

  Future<String> connect(Connector connector) async {
    final response = await http.post(
      Uri.parse(
          "${Configuration.instance.baseUrl}/v1/connectors/${connector.id}/connect"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
        'Authorization': authCubit.state.token ?? "",
      },
      body: jsonEncode(<String, String>{
        'success_url': Configuration.instance.baseUrl,
        'error_url': Configuration.instance.baseUrl,
      }),
    );

    if (response.statusCode == 200) {
      final Map<String, dynamic> data = jsonDecode(response.body);
      return data["redirect_url"];
    } else {
      throw Exception('Failed to connect: ${response.statusCode}');
    }
  }

  Future<void> deleteConnection(Connection connection) async {
    final response = await http.delete(
      Uri.parse(
          "${Configuration.instance.baseUrl}/v1/connections/${connection.id}"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
        'Authorization': authCubit.state.token ?? "",
      },
    );

    if (response.statusCode == 200) {
      loadConnections();
    } else {
      throw Exception('Failed to delete connection: ${response.statusCode}');
    }
  }
}
