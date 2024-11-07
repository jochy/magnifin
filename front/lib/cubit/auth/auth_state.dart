part of 'auth_cubit.dart';

@freezed
class AuthState with _$AuthState {
  factory AuthState({
    @Default(null) String? token,
  }) = _AuthState;

  factory AuthState.fromJson(Map<String, dynamic> json) =>
      _$AuthStateFromJson(json);
}
