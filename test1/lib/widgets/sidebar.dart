import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:test1/controllers/logoutcontroller.dart';
import 'package:test1/controllers/screencontroller.dart';

class Sidebar extends StatelessWidget {
  Sidebar({super.key});

  final Logoutcontroller _logoutcontroller = Get.find<Logoutcontroller>();
  final Screencontroller _screencontroller = Get.find<Screencontroller>();

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        children: [
          ListTile(
            title: const Text('Home'),
            onTap: (){
              _screencontroller.selectedpage('home');
            },
          ),
          ExpansionTile(
            initiallyExpanded: true,
            title: const Text('Orders'),
            children: [
              ListTile(
                title: const Text('New Orders'),
                onTap: (){
                  _screencontroller.selectedpage('new_order');
                },
              ),
              ListTile(
                title: const Text('Completed Orders'),
                onTap: (){
                  _screencontroller.selectedpage('complete_order');
                },
              ),
            ],
          ),
          ExpansionTile(
            initiallyExpanded: true,
            title: const Text('Products'),
            children: [
              ListTile(
                title: const Text('Add Product'),
                onTap: (){
                  _screencontroller.selectedpage('add_product');
                },
              ),
              ListTile(
                title: const Text('Remove Products'),
                onTap: (){
                  _screencontroller.selectedpage('remove_product');
                },
              ),
            ],
          ),
          ExpansionTile(
            initiallyExpanded: true,
            title: const Text('Users'),
            children: [
              ListTile(
                title: const Text('New User'),
                onTap: (){
                  _screencontroller.selectedpage('new_user');
                },
              ),
              ListTile(
                title: const Text('Edit User'),
                onTap: (){
                  _screencontroller.selectedpage('edit_user');
                },
              ),
              ListTile(
                title: const Text('Delete User'),
                onTap: (){
                  _screencontroller.selectedpage('delete_user');
                },
              ),
            ],
          ),
          const SizedBox(height: 50,),
          MaterialButton(
            onPressed: () async {
              bool success = await _logoutcontroller.logout();
              if (success) {
                Get.offAllNamed('/login');
              } else {
                Get.snackbar("Logout Failed", "I dnt knw why fail");
              }
            },
            color: Colors.blue,
            textColor: Colors.white,
            child: const Text('Logout'),
          )
        ],
      ),
    );
  }
}
