import 'dart:convert';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:web/web.dart' as web;

class Exchangecontroller extends GetxController{

  final _code = ''.obs;
  final _idToken = ''.obs;
  final _accessToken = ''.obs;

  String get code => _code.value;
  String get idToken => _idToken.value;
  String get accessToken => _accessToken.value;

  void setCode(String value) => _code.value = value;

  Future<bool> exchange() async {
    final response = await http.post(
      Uri.parse('http://127.0.0.1:9080/exchange'),
      body: jsonEncode({
        'code':code
      }),
      headers:{
        'Content-Type':'application/json',
      }
    );

    if(response.statusCode == 200){
      final responseBody = json.decode(response.body);
      _idToken.value = responseBody['id_token'];
      // print(idToken);
      web.window.localStorage['token'] = _idToken.value;
      _accessToken.value = responseBody['access_token'];
      
      // print('This is access token : $accessToken');
      return true;
    }else{
      return false;
    }
  }
}