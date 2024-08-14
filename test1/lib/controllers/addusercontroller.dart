import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:get/get.dart';

class Addusercontroller extends GetxController{
  final _username = ''.obs;
  final _email = ''.obs;
  final _role = 'Packer'.obs;

  String get username => _username.value;
  String get email => _email.value;
  String get role => _role.value;

  void setUsername(String value) => _username.value = value;
  void setEmail(String value) => _email.value = value;
  void setRole(String value) => _role.value = value;

  Future<bool> add() async{
    final response = await http.post(
      Uri.parse('http://127.0.0.1:9080/add'),
      body: jsonEncode({
        'username':username,
        'email':email,
        'realmRoles' :role
      }),
      headers:{
        'Content-Type':'application/json',
      },
    );

    if(response.statusCode == 200){
      return true;
    }else{
      return false;
    }
  }
}