import 'dart:convert';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:test1/model/usermodel.dart';

class Editusercontroller extends GetxController{

  var user = <User>[].obs;
  var isLoading = false.obs;
  final _accessToken = ''.obs;

  String get accessToken => _accessToken.value;
  void setAccessToken(String value) => _accessToken.value = value;

  Future<bool> getUser() async{
    isLoading(true);
    final response = await http.get(
      Uri.parse('http://127.0.0.1:9080/get-user'),
      headers: {
          'Content-Type': 'application/json',
          'Authorization' : 'Bearer $accessToken'
        },
    );

    if (response.statusCode == 200) {
      final List<dynamic> data = jsonDecode(response.body);
      print(data);
      user.assignAll(data.map((e) => User.fromJson(e)).toList());
      isLoading(false);
      return true;
    }else{
      return false;
    }
  }

  Future<bool> updateUser(User user) async{
    
    final response = await http.put(
      Uri.parse('http://127.0.0.1:9080/update-user'),
      body: jsonEncode({
        'id':user.id,
        'username':user.username,
        'email':user.email
      }),
      headers: {
          'Content-Type': 'application/json',
          'Authorization' : 'Bearer $accessToken'
        },
    );

    print(response.statusCode);
    if (response.statusCode == 200) {
      
      return true;
    }else{
      return false;
    }
  }

  Future<bool> deleteUser(String id) async{
    
    final response = await http.delete(
      Uri.parse('http://127.0.0.1:9080/delete-user'),
      body: jsonEncode({
        'id':id
      }),
      headers: {
          'Content-Type': 'application/json',
          'Authorization' : 'Bearer $accessToken'
        },
    );

    // print(response.statusCode);
    if (response.statusCode == 200) {
      user.removeWhere((u)=> u.id == id);
      return true;
    }else{
      return false;
    }
  }
}