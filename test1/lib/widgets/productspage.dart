import 'package:flutter/material.dart';

class Productspage extends StatelessWidget {
  const Productspage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Products Page'),),
      body: const Text('This is Products page'),
    );
  }
}