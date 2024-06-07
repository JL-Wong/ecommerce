import 'package:flutter/material.dart';
import 'package:frontend/controller/LoginController.dart';
import 'package:get/get.dart';


class Loginpage extends StatelessWidget {
  Loginpage({super.key});

  final LoginController controller = Get.put(LoginController());
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        // crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          const Text(
            'Login',
            style: TextStyle(
                fontSize: 40, color: Colors.black, fontWeight: FontWeight.bold),
          ),

          Padding(
            padding: const EdgeInsets.symmetric(vertical: 30,horizontal: 30),
            child: Form(child: Column(
              children: [
                TextFormField(
                  controller: controller.emailController,
                  keyboardType: TextInputType.emailAddress,
                  decoration:  const InputDecoration(
                    labelText: "Email",
                    hintText: "Enter email",
                    prefixIcon: Icon(Icons.email),
                    border:  OutlineInputBorder(),
                  ),
                  onChanged: (value) => controller.email.value = value ,
                  validator: (value) => value!.isEmpty? "Please Enter your email": null,
                ),
            
                const SizedBox(height: 30,),
            
                TextFormField(
                  controller: controller.passwordController,
                  keyboardType: TextInputType.visiblePassword,
                  decoration:  const InputDecoration(
                    labelText: "Password",
                    hintText: "Enter Password",
                    prefixIcon: Icon(Icons.password),
                    border:  OutlineInputBorder(),
                  ),
                  onChanged: (value) => controller.password.value = value ,
                  validator: (value) => value!.isEmpty? "Please Enter your password": null,
                ),

                const SizedBox(height: 30,),

                MaterialButton(
                  onPressed: (){
                    print(controller.email);
                    // if(_formKey.currentState != null && _formKey.currentState!.validate()){
                      controller.login();
                    // }
                  },
                  color: Colors.blue,
                  textColor: Colors.white,
                  child: const Text('Login'),
                )
              ],
            )),
          )
        ],
      ),
    );
  }
}
