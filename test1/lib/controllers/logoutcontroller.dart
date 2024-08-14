import 'package:get/get.dart';
import 'package:test1/controllers/exchangecontroller.dart';
// import 'package:test1/controllers/logincontroller.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:web/web.dart' as web;

class Logoutcontroller extends GetxController{
  final Exchangecontroller _exchangecontroller = Get.find<Exchangecontroller>();

  String get idToken => _exchangecontroller.idToken;


  Future<bool> logout() async {
    final response = await http.post(
      Uri.parse('http://127.0.0.1:9080/logout'),
      body: json.encode({
        'id_token': idToken,
        // 'id_token': web.window.localStorage.getItem('token'),
      }),
      headers: {
        'Content-Type': 'application/json',
      },
    );

    if (response.statusCode == 200) {
      web.window.localStorage.removeItem('token');
      return true;
    } else {
      print(response.statusCode);
      return false;
    }
  }
}