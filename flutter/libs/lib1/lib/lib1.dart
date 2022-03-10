import 'package:flutter/material.dart';

class StaticText extends StatelessWidget {
  const StaticText({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Text(
      'You have pushed the button this many times:',
      style: TextStyle(
        fontSize: 18,
        fontWeight: FontWeight.bold,
      ),
    );
  }
}
