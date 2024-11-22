import 'package:flutter/material.dart';
import 'package:front/components/styled_text.dart';
import 'package:front/cubit/auth/auth_cubit.dart';
import 'package:front/cubit/connections/connections_cubit.dart';
import 'package:front/generated/l10n.dart';
import 'package:go_router/go_router.dart';
import 'package:toastification/toastification.dart';

class SigninScreen extends StatefulWidget {
  const SigninScreen({super.key});

  @override
  State<SigninScreen> createState() => _SignInScreenState();
}

class _SignInScreenState extends State<SigninScreen> {
  var awaiting = false;
  var failure = false;
  var _canSubmit = false;
  var usernameController = TextEditingController();
  var passwordController = TextEditingController();

  @override
  void initState() {
    usernameController.addListener((){
      var newSubmit = usernameController.text.isNotEmpty && passwordController.text.isNotEmpty;
      if (newSubmit != _canSubmit) {
        setState(() {
          _canSubmit = newSubmit;
        });
      }
    });
    passwordController.addListener((){
      var newSubmit = usernameController.text.isNotEmpty && passwordController.text.isNotEmpty;
      if (newSubmit != _canSubmit) {
        setState(() {
          _canSubmit = newSubmit;
        });
      }
    });
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: SizedBox(
          width: 300,
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Card.outlined(
              child: Padding(
                padding: const EdgeInsets.all(20.0),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    StyledText.headlineMedium(S.of(context).newAccount),
                    const SizedBox(height: 20),
                    if (failure) ...[
                      StyledText.errorText(
                          S.of(context).loginFailedPleaseTryAgain),
                      const SizedBox(height: 8)
                    ],
                    TextField(
                      controller: usernameController,
                      decoration:
                          InputDecoration(labelText: S.of(context).username),
                    ),
                    TextField(
                      controller: passwordController,
                      decoration:
                          InputDecoration(labelText: S.of(context).password),
                      obscureText: true,
                      onSubmitted: (_) => _signIn(context),
                    ),
                    const SizedBox(height: 20),
                    FilledButton(
                      onPressed: awaiting || !_canSubmit ? null : () => _signIn(context),
                      child: awaiting
                          ? const SizedBox(
                              width: 14,
                              height: 14,
                              child: CircularProgressIndicator(
                                strokeWidth: 2,
                              ),
                            )
                          : Text(S.of(context).createNewAccount),
                    ),
                    const SizedBox(height: 10),
                    TextButton(
                      onPressed: () => GoRouter.of(context).replace('/login'),
                      child: Text(S.of(context).login),
                    ),
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }

  Future<void> _signIn(BuildContext context) async {
    var cubit = AuthCubit.of(context);
    var router = GoRouter.of(context);
    var connectionsCubit = ConnectionsCubit.of(context);

    setState(() {
      awaiting = true;
    });

    // Simulate a network request
    try {
      var res =
          await cubit.createUser(usernameController.text, passwordController.text);
      if (res) {
        router.replace('/');
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
          title: Text(S.of(context).createUserFailed),
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
