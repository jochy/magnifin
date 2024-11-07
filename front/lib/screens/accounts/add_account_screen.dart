import 'package:flutter/material.dart';
import 'package:flutter_material_design_icons/flutter_material_design_icons.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:toastification/toastification.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:url_launcher/url_launcher_string.dart';

class AddAccountScreen extends StatefulWidget {
  const AddAccountScreen({super.key});

  @override
  State<AddAccountScreen> createState() => _AddAccountScreenState();
}

class _AddAccountScreenState extends State<AddAccountScreen> {
  var search = "";
  var _previousData = <Connector>[];

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        SizedBox(
          width: 400,
          child: Padding(
            padding: const EdgeInsets.all(4.0),
            child: TextField(
              onChanged: (value) {
                setState(() {
                  search = value;
                });
              },
              decoration: InputDecoration(
                labelText: S.of(context).bankName,
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(100)),
                prefixIcon: const Icon(Icons.search),
              ),
            ),
          ),
        ),
        const SizedBox(height: 20),
        Expanded(
          child: SizedBox(
            width: 400,
            child: FutureBuilder(
              future: ConnectionsCubit.of(context).searchConnectors(search),
              builder: (context, snapshot) {
                if (snapshot.hasError) {
                  return Text("Error: ${snapshot.error}");
                }

                var connectors = <Connector>[];

                if (snapshot.connectionState == ConnectionState.waiting &&
                    _previousData.isEmpty) {
                  return const SizedBox(
                    width: 40,
                    height: 40,
                    child: CircularProgressIndicator(),
                  );
                } else if (snapshot.connectionState ==
                    ConnectionState.waiting) {
                  connectors = _previousData;
                } else {
                  connectors = snapshot.data as List<Connector>;
                }

                if (snapshot.hasData) {
                  _previousData = connectors;
                  return ListView.builder(
                    itemBuilder: (context, index) {
                      var connector = connectors[index];
                      return ListTile(
                        leading: ClipOval(
                          child: CircleAvatar(
                            child: connector.logoUrl != ""
                                ? Image.network(
                                    connector.logoUrl,
                                    loadingBuilder:
                                        (context, child, loadingProgress) {
                                      if (loadingProgress == null) {
                                        return child;
                                      } else {
                                        return Text(connector.name[0]);
                                      }
                                    },
                                  )
                                : Text(connector.name[0]),
                          ),
                        ),
                        title: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Padding(
                              padding: const EdgeInsets.only(left: 8.0),
                              child: Text(connector.name),
                            ),
                            TextButton.icon(
                              onPressed: () {
                                _connect(context, connector);
                              },
                              style: TextButton.styleFrom(
                                visualDensity: const VisualDensity(
                                  horizontal: -4,
                                  vertical: -4,
                                ),
                              ),
                              label: Text(S.of(context).connect),
                              icon: const Icon(MdiIcons.openInNew, size: 15),
                            ),
                          ],
                        ),
                      );
                    },
                    itemCount: connectors.length,
                  );
                }

                return const SizedBox();
              },
            ),
          ),
        ),
      ],
    );
  }

  Future<void> _connect(BuildContext context, Connector connector) async {
    var connectionsCubit = ConnectionsCubit.of(context);

    try {
      var redirect = await connectionsCubit.connect(connector);
      var success = await launchUrlString(redirect);
      if (!success) {
        throw Exception("Failed to launch URL in browser");
      }
    } catch (e) {
      if (context.mounted) {
        toastification.show(
          context: context,
          title: Text(
              S.of(context).unableToConnectToConnectorname(connector.name)),
          description: Text(e.toString()),
          autoCloseDuration: const Duration(seconds: 5),
          type: ToastificationType.error,
          style: ToastificationStyle.fillColored,
          showProgressBar: false,
        );
      }
    }
  }
}
