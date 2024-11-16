// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'transactions_cubit.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$TransactionImpl _$$TransactionImplFromJson(Map<String, dynamic> json) =>
    _$TransactionImpl(
      id: (json['id'] as num).toInt(),
      accountId: (json['aid'] as num).toInt(),
      bankTransactionId: json['bid'] as String?,
      amount: (json['a'] as num).toDouble(),
      currency: json['c'] as String,
      direction: json['d'] as String,
      status: json['s'] as String,
      operationAt: DateTime.parse(json['at'] as String),
      counterpartyName: json['name'] as String?,
      counterpartyAccount: json['acc'] as String?,
      reference: json['ref'] as String?,
      counterpartyLogoUrl: json['logo'] as String?,
      category: (json['ca'] as num?)?.toInt(),
      method: json['m'] as String?,
    );

Map<String, dynamic> _$$TransactionImplToJson(_$TransactionImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'aid': instance.accountId,
      'bid': instance.bankTransactionId,
      'a': instance.amount,
      'c': instance.currency,
      'd': instance.direction,
      's': instance.status,
      'at': instance.operationAt.toIso8601String(),
      'name': instance.counterpartyName,
      'acc': instance.counterpartyAccount,
      'ref': instance.reference,
      'logo': instance.counterpartyLogoUrl,
      'ca': instance.category,
      'm': instance.method,
    };

_$CategoryImpl _$$CategoryImplFromJson(Map<String, dynamic> json) =>
    _$CategoryImpl(
      id: (json['id'] as num).toInt(),
      name: json['name'] as String,
      userId: (json['uid'] as num?)?.toInt(),
      icon: json['icon'] as String,
      color: json['color'] as String,
      includeInBudget: json['include_in_budget'] as bool,
    );

Map<String, dynamic> _$$CategoryImplToJson(_$CategoryImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'uid': instance.userId,
      'icon': instance.icon,
      'color': instance.color,
      'include_in_budget': instance.includeInBudget,
    };
