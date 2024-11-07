import 'package:flutter/material.dart';

class TransactionsScreen extends StatelessWidget {
  const TransactionsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: [
          Card.outlined(child: Padding(
            padding: const EdgeInsets.all(50.0),
            child: Text("Transactions"),
          )),
          Card.filled(child: Padding(
            padding: const EdgeInsets.all(50.0),
            child: Text("Transactions"),
          )),
          Card(child: Padding(
            padding: const EdgeInsets.all(50.0),
            child: Text("Transactions"),
          )),
        ],
      ),
    );
  }
}
