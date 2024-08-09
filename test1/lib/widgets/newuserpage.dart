import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:test1/controllers/addusercontroller.dart';

class Newuserpage extends StatelessWidget {
  Newuserpage({super.key});

  final Addusercontroller _addusercontroller = Get.put(Addusercontroller());
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('User Page'),),
      body: Center(
        child: Column(
          
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            const Text("Create User", style: TextStyle(color: Colors.black, fontSize: 50),),
            TextFormField(
              controller: _usernameController,
              onChanged: (value) => _addusercontroller.setUsername(value),
              decoration: const InputDecoration(
                labelText: "Username / Email",
                border: OutlineInputBorder(),
                
              ),
            ),
            const SizedBox(height: 20,),
            TextFormField(
              controller: _emailController,
              onChanged: (value) => _addusercontroller.setEmail(value),
              keyboardType: TextInputType.visiblePassword,
              // obscureText: true,
              decoration: const InputDecoration(
                labelText: "Email",
                border: OutlineInputBorder(),
                
              ),
            ),
            const SizedBox(height: 50,),
            Obx(() =>  SegmentedButton<String>(
                segments: const [
                  ButtonSegment(
                    value: "Admin",
                    icon: Icon(Icons.people),
                    label: Text('Admin'),
                  ),
                  ButtonSegment(
                    value: "Packer",
                    icon: Icon(Icons.work),
                    label: Text('Packer'),
                  )
                ], 
                selected: {_addusercontroller.role},
                onSelectionChanged: (newSelection){
                  _addusercontroller.setRole(newSelection as String);
                },
              ),
            ),
            const SizedBox(height: 50,),
            MaterialButton(
              onPressed: () async{
                // handle button press
                // Handle button press
                bool success = await _addusercontroller.add();
                
                if (success) {
                  // Get.to(Homepage(token: _controller.accessToken,));
                  _usernameController.clear();
                  _emailController.clear();
                  Get.snackbar("Created Success", "User created");
                  
                } else {
                  Get.snackbar("Created Failed", "Failed");
                }
              },
              color: Colors.blue,
              textColor: Colors.white,
              child: const Text('Create'),
            )
        ],
      ),
      )
    );
  }
}