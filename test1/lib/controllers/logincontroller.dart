import 'dart:convert';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:url_launcher/url_launcher_string.dart';

class Logincontroller extends GetxController{
  final _url = ''.obs;
  var isLoading = false.obs;
  

  String get url => _url.value;

  Future<bool> login() async {
    isLoading(true);
    final response = await http.post(
      Uri.parse('http://127.0.0.1:9080/google-login'),
      headers: {
        'Content-Type': 'application/json',
      },
    );

    if (response.statusCode == 200) {
      final responseBody = jsonDecode(response.body);
      // print(responseBody);
      
      _url.value = responseBody['authUrl'];
      // print(url);
      if (await canLaunchUrlString(url)) {
        await launchUrlString(url,webOnlyWindowName: '_self');
      } else {
        throw 'Could not launch the url';
      }  

      isLoading(false);
      return true;   
    } else {
      isLoading(false);
      return false;
    }
  }
}