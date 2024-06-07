import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:io';

class LoginController extends GetxController {
  var email = ''.obs;
  var password = ''.obs;

  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  Future<void> login() async {
    final url = Uri.parse('http://172.16.7.160:9080');
    // final url = Uri.parse('http://172.21.0.8:9080');
    
    try {
      // final response = await http.get(url);
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          // 'Authorization': 'Basic ${base64Encode(utf8.encode('$email:$password'))}',
        },
        body: json.encode({
          'email': email.value,
          'password': password.value,
        }),
      );

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        // Handle successful response
        print('Login successful: $responseData');
      } else {
        // Handle error response
        print('Login failed: ${response.statusCode}');
      }
    } catch (error) {
      // Handle network or other errors
      print('Error: $error');
    }
  }

  @override
  void onClose() {
    emailController.dispose();
    passwordController.dispose();
    super.onClose();
  }
}
