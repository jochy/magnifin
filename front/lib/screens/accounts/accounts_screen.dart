import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:front/screens/accounts/add_account_screen.dart';
import 'package:go_router/go_router.dart';
import 'package:moment_dart/moment_dart.dart';
import 'package:money2/money2.dart';
import 'package:responsive_framework/responsive_framework.dart';
import 'package:toastification/toastification.dart';

class AccountsScreen extends StatelessWidget {
  const AccountsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<ConnectionsCubit, ConnectionsState>(
      builder: (context, state) {
        if (state.connections.isEmpty) {
          return const AddAccountScreen();
        }

        return SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              TextButton.icon(
                onPressed: () => GoRouter.of(context).go('/add-account'),
                label: Text(S.of(context).addANewAccount),
                icon: const Icon(MdiIcons.plus),
              ),
              const SizedBox(height: 16),
              Wrap(alignment: WrapAlignment.start, spacing: 16, children: [
                ...state.connections.map((connection) {
                  return connectionItem(connection, context);
                }),
              ]),
            ],
          ),
        );
      },
    );
  }

  Widget connectionItem(Connection connection, BuildContext context) {
    return SizedBox(
      width: ResponsiveBreakpoints.of(context).breakpoint.name == MOBILE
          ? double.infinity
          : 500,
      child: Card.outlined(
        child: Padding(
          padding: const EdgeInsets.all(12.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  CircleAvatar(
                    maxRadius: 18,
                    child: connection.connector.logoUrl != ""
                        ? Image.network(connection.connector.logoUrl)
                        : const Icon(MdiIcons.bank),
                  ),
                  const SizedBox(width: 12),
                  StyledText.headlineSmall(connection.connector.name),
                  Expanded(
                    child: Align(
                      alignment: Alignment.centerRight,
                      child: _statusToBadge(context, connection.status),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 4),
              const Divider(thickness: 0.2, indent: 8, endIndent: 8),
              const SizedBox(height: 4),
              ...connection.accounts.map(
                (a) => Padding(
                  padding: const EdgeInsets.only(
                      left: 12, top: 4, bottom: 4, right: 4),
                  child: Row(
                    children: [
                      const CircleAvatar(
                        maxRadius: 14,
                        child: Icon(
                          MdiIcons.bank,
                          size: 16,
                        ),
                      ),
                      const SizedBox(width: 8),
                      Expanded(
                        child: Align(
                          alignment: Alignment.centerLeft,
                          child: StyledText.bodyLarge(a.name),
                        ),
                      ),
                      Flexible(
                        child: Align(
                          alignment: Alignment.topRight,
                          child: StyledText.text(a.currency == null
                              ? a.balance.toString()
                              : Money.fromFixed(Fixed.fromNum(a.balance),
                                      isoCode: a.currency!)
                                  .toString()),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              const SizedBox(height: 4),
              const Divider(thickness: 0.2, indent: 8, endIndent: 8),
              const SizedBox(height: 4),
              Padding(
                padding: const EdgeInsets.only(left: 4.0),
                child: SizedBox(
                  width: double.infinity,
                  child: Wrap(
                    runAlignment: WrapAlignment.spaceEvenly,
                    alignment: WrapAlignment.spaceBetween,
                    crossAxisAlignment: WrapCrossAlignment.center,
                    children: [
                      if (connection.lastSuccessfulSync != null)
                        StyledText.hintText(
                          S.of(context).lastSuccessfulSync(Moment.now().from(
                                connection.lastSuccessfulSync!,
                                dropPrefixOrSuffix: true,
                              )),
                          size: 12,
                        ),
                      TextButton(
                        onPressed: () => _deleteConnection(context, connection),
                        style: TextButton.styleFrom(
                          foregroundColor: Theme.of(context).colorScheme.error,
                          visualDensity:
                              const VisualDensity(horizontal: -4, vertical: -4),
                        ),
                        child: Text(S.of(context).delete),
                      ),
                    ],
                  ),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }

  void _deleteConnection(BuildContext context, Connection connection) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: Text(S.of(context).deleteConnection),
          content: Text(
              S.of(context).areYouSureYouWantToDeleteThisConnectionnthisCan),
          actions: [
            TextButton(
              onPressed: () => Navigator.of(context).pop(),
              style: TextButton.styleFrom(
                foregroundColor: Theme.of(context).colorScheme.error,
              ),
              child: Text(S.of(context).cancel),
            ),
            TextButton(
              onPressed: () {
                _performDeleteConnection(context, connection);
                Navigator.of(context).pop();
              },
              style: TextButton.styleFrom(
                foregroundColor: Theme.of(context).colorScheme.error,
              ),
              child: Text(S.of(context).delete),
            ),
          ],
        );
      },
    );
  }

  Future<void> _performDeleteConnection(
      BuildContext context, Connection connection) async {
    final connectionsCubit = ConnectionsCubit.of(context);
    try {
      await connectionsCubit.deleteConnection(connection);
    } catch (e) {
      if (context.mounted) {
        toastification.show(
          context: context,
          title: Text(S.of(context).unableToDeleteConnection),
          description: Text(e.toString()),
          autoCloseDuration: const Duration(seconds: 5),
          type: ToastificationType.error,
          style: ToastificationStyle.fillColored,
          showProgressBar: false,
        );
      }
    }
  }

  Widget _statusToBadge(BuildContext context, String status) {
    Color color = Colors.grey;
    String text = S.of(context).unknown;

    switch (status) {
      case "SYNCHRONIZED":
        color = Colors.green;
        text = S.of(context).synced;
        break;
      case "RATE_LIMITED":
        color = Colors.orange;
        text = S.of(context).rateLimited;
        break;
      case "SYNC_IN_PROGRESS":
        color = Colors.orange;
        text = S.of(context).syncInProgress;
        break;
      case "SUSPENDED":
        color = Colors.red;
        text = S.of(context).suspended;
        break;
    }

    return Chip(
      padding: const EdgeInsets.all(0),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(100),
      ),
      visualDensity: const VisualDensity(horizontal: -4, vertical: -4),
      label: Text(text),
      backgroundColor: color.withAlpha(100),
    );
  }
}
