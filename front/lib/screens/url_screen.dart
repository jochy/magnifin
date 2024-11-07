import 'package:flutter/material.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/config.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:go_router/go_router.dart';
import 'package:toastification/toastification.dart';

class UrlScreen extends StatefulWidget {
  const UrlScreen({super.key});

  @override
  State<UrlScreen> createState() => _UrlScreenState();
}

class _UrlScreenState extends State<UrlScreen> {
  var urlController = TextEditingController();
  var awaiting = false;
  var failure = false;
  var _canSubmit = false;

  @override
  void initState() {
    urlController.addListener(() {
      var newSubmit = urlController.text.isNotEmpty;
      if (newSubmit != _canSubmit) {
        setState(() {
          _canSubmit = newSubmit;
        });
      }
    });
    urlController.text = "http://localhost:8080";
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: SizedBox(
          width: 400,
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Card.outlined(
              child: Padding(
                padding: const EdgeInsets.all(20.0),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    StyledText.headlineMedium(S.of(context).url),
                    const SizedBox(height: 20),
                    if (failure) ...[
                      StyledText.errorText(
                        S
                            .of(context)
                            .anErrorOccurredWhileTryingToContactTheServerNplease,
                        maxLines: 2,
                      ),
                      const SizedBox(height: 8)
                    ],
                    TextField(
                      controller: urlController,
                      decoration:
                          InputDecoration(labelText: S.of(context).backendUrl),
                    ),
                    const SizedBox(height: 20),
                    FilledButton(
                      onPressed: awaiting || !_canSubmit
                          ? null
                          : () => _validateUrl(context),
                      child: awaiting
                          ? const SizedBox(
                              width: 14,
                              height: 14,
                              child: CircularProgressIndicator(
                                strokeWidth: 2,
                              ),
                            )
                          : Text(S.of(context).validate),
                    )
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }

  Future<void> _validateUrl(BuildContext context) async {
    var cubit = AuthCubit.of(context);
    var router = GoRouter.of(context);

    setState(() {
      awaiting = true;
    });

    // Simulate a network request
    try {
      var res = await cubit.ping(urlController.text);
      if (res) {
        Configuration.instance.baseUrl = urlController.text;
        router.go('/');
      } else {
        setState(() {
          failure = true;
        });
      }
    } catch (e) {
      setState(() {
        failure = true;
      });
      if (mounted && context.mounted) {
        toastification.show(
          context: context,
          title: Text(S.of(context).loginError),
          description: Text(e.toString()),
          autoCloseDuration: const Duration(seconds: 5),
          type: ToastificationType.error,
          style: ToastificationStyle.fillColored,
          showProgressBar: false,
        );
      }
    }

    setState(() {
      awaiting = false;
    });
  }
}
