import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:front/config.dart';
import 'package:http/http.dart' as http;

part 'auth_state.dart';

part 'auth_cubit.g.dart';

part 'auth_cubit.freezed.dart';

class AuthCubit extends Cubit<AuthState> {
  AuthCubit() : super(AuthState());

  static AuthCubit of(context) => BlocProvider.of<AuthCubit>(context);

  Future<bool> ping(String url) async {
    final response = await http.get(
      Uri.parse("$url/v1/ping"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
      },
    );

    if (response.statusCode == 200) {
      return true;
    } else {
      return false;
    }
  }

  Future<bool> login(String username, String password) async {
    final response = await http.post(
      Uri.parse("${Configuration.instance.baseUrl}/v1/login"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
      },
      body: jsonEncode(<String, String>{
        'username': username,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      final Map<String, dynamic> data = jsonDecode(response.body);
      emit(AuthState(token: data['token']));
    } else {
      return false;
    }

    return true;
  }

  Future<bool> createUser(String username, String password) async {
    final response = await http.post(
      Uri.parse("${Configuration.instance.baseUrl}/v1/users"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
        'Accept': 'application/json',
      },
      body: jsonEncode(<String, String>{
        'username': username,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      final Map<String, dynamic> data = jsonDecode(response.body);
      emit(AuthState(token: data['token']));
    } else {
      return false;
    }

    return true;
  }

  void logout() {
    emit(AuthState());
  }
}
