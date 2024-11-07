// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'connections_cubit.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$ConnectionImpl _$$ConnectionImplFromJson(Map<String, dynamic> json) =>
    _$ConnectionImpl(
      id: (json['id'] as num).toInt(),
      status: json['status'] as String,
      renewConsentBefore: json['renew_consent_before'] == null
          ? null
          : DateTime.parse(json['renew_consent_before'] as String),
      errorMessage: json['error_message'] as String?,
      lastSuccessfulSync: json['last_successful_sync'] == null
          ? null
          : DateTime.parse(json['last_successful_sync'] as String),
      accounts: (json['accounts'] as List<dynamic>)
          .map((e) => Account.fromJson(e as Map<String, dynamic>))
          .toList(),
      connector: Connector.fromJson(json['connector'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$ConnectionImplToJson(_$ConnectionImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'status': instance.status,
      'renew_consent_before': instance.renewConsentBefore?.toIso8601String(),
      'error_message': instance.errorMessage,
      'last_successful_sync': instance.lastSuccessfulSync?.toIso8601String(),
      'accounts': instance.accounts,
      'connector': instance.connector,
    };

_$AccountImpl _$$AccountImplFromJson(Map<String, dynamic> json) =>
    _$AccountImpl(
      id: (json['id'] as num).toInt(),
      bankAccountId: json['bank_account_id'] as String?,
      name: json['name'] as String,
      type: json['type'] as String?,
      currency: json['currency'] as String?,
      accountNumber: json['account_number'] as String?,
      balance: (json['balance'] as num).toDouble(),
    );

Map<String, dynamic> _$$AccountImplToJson(_$AccountImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'bank_account_id': instance.bankAccountId,
      'name': instance.name,
      'type': instance.type,
      'currency': instance.currency,
      'account_number': instance.accountNumber,
      'balance': instance.balance,
    };

_$ConnectorImpl _$$ConnectorImplFromJson(Map<String, dynamic> json) =>
    _$ConnectorImpl(
      id: (json['id'] as num).toInt(),
      name: json['name'] as String,
      logoUrl: json['logo_url'] as String,
    );

Map<String, dynamic> _$$ConnectorImplToJson(_$ConnectorImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'logo_url': instance.logoUrl,
    };
