import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:test1/controllers/logincontroller.dart';

class Loginpage extends StatelessWidget {
  Loginpage({super.key});

  final Logincontroller _logincontroller = Get.put(Logincontroller());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text("Login", style: TextStyle(fontSize: 50),),
            Obx((){
              if(_logincontroller.isLoading.value){
                return const CircularProgressIndicator();
              }else{
                return MaterialButton(
                  onPressed: () async{
                    bool success = await _logincontroller.login();
                    // print(success);
                    if(success){
                      // Get.offAllNamed('/home');
                    }else{
                      Get.snackbar('Login Failed','Please try again');
                    }
                  },
                  color: Colors.blue,
                  textColor: Colors.white,
                  child: const Text('Login'),
                
                );
              }
            })
            
          ],
        ),
      ),
    );
  }
}