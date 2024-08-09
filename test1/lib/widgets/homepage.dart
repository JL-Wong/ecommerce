import 'dart:js_interop';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
// import 'package:test1/controllers/activetabscontroller.dart';
import 'package:test1/controllers/exchangecontroller.dart';
import 'package:test1/controllers/logoutcontroller.dart';
import 'package:test1/controllers/screencontroller.dart';
import 'package:test1/widgets/completeorderpage.dart';
import 'package:test1/widgets/edituserpage.dart';
import 'package:test1/widgets/neworderpage.dart';
import 'package:test1/widgets/newuserpage.dart';
// import 'package:test1/widgets/sidebar.dart';
import 'package:web/web.dart' as web;

class Homepage extends StatelessWidget {
  Homepage({super.key});

  final Exchangecontroller _exchangecontroller = Get.put(Exchangecontroller());
  final Logoutcontroller _logoutcontroller = Get.put(Logoutcontroller());
  final Screencontroller _screencontroller = Get.put(Screencontroller());
  // final Activetabscontroller _tabcontroller = Get.find();

  
  @override
  Widget build(BuildContext context) {
    //extract the current url
    final Uri currentUri = Uri.base;

    final String? code = currentUri.queryParameters['code'];
    print(code);

    if (web.window.localStorage['activeTabs'] == null) {
      web.window.localStorage['activeTabs'] = '0';
    }
    // Initialize localStorage for tracking active tabs
    int activeTabs = int.parse(web.window.localStorage['activeTabs']!);
    web.window.localStorage['activeTabs'] = (activeTabs + 1).toString();

    // Handle window unload
    web.window.addEventListener('click',(void event){
      int activeTabs = int.parse(web.window.localStorage['activeTabs']!);
      web.window.localStorage['activeTabs'] = (activeTabs - 1).toString();
      if (int.parse(web.window.localStorage['activeTabs']!) <= 0) {
        _logoutcontroller.logout();
      }
    }.toJS);


    // web.window.addEventListener('storage', callback)

    WidgetsBinding.instance.addPostFrameCallback((_) {
      web.window.history.replaceState(null, 'unused', '/home');
      _exchangecontroller.setCode(code!);
      _exchangecontroller.exchange();
    });

    return Scaffold(
        appBar: AppBar(
          title: const Text('Main page'),
        ),
        // drawer: Sidebar(),
        body: Row(
          children: [
            Container(
              width: 200,
              color: Colors.grey[600],
              child: ListView(
                children: [
                  ListTile(
                    title: const Text('Home'),
                    onTap: () {
                      _screencontroller.selectedpage('home');
                    },
                  ),
                  ExpansionTile(
                    initiallyExpanded: true,
                    title: const Text('Orders'),
                    children: [
                      ListTile(
                        title: const Text('New Orders'),
                        onTap: () {
                          _screencontroller.selectedpage('new_order');
                        },
                      ),
                      ListTile(
                        title: const Text('Completed Orders'),
                        onTap: () {
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
                        onTap: () {
                          _screencontroller.selectedpage('add_product');
                        },
                      ),
                      ListTile(
                        title: const Text('Remove Products'),
                        onTap: () {
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
                        onTap: () {
                          _screencontroller.selectedpage('new_user');
                        },
                      ),
                      ListTile(
                        title: const Text('Edit User'),
                        onTap: () {
                          _screencontroller.selectedpage('edit_user');
                        },
                      ),
                      ListTile(
                        title: const Text('Delete User'),
                        onTap: () {
                          _screencontroller.selectedpage('delete_user');
                        },
                      ),
                    ],
                  ),
                  const SizedBox(
                    height: 50,
                  ),
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
              )
            ),
            Expanded(child: Obx(() {
              switch (_screencontroller.currentPage.value) {
                case 'home':
                  return const Center(child: Text('Home page'),);
                case 'new_order':
                  return const Neworderpage();
                case 'complete_order':
                  return const Completeorderpage();
                // case 'add_product':
                //   return Homepage();
                // case 'remove_product':
                //   return Homepage();
                case 'new_user':
                  return Newuserpage();
                case 'edit_user':
                  return Edituserpage();
                // case 'delete_user':
                //   return Homepage();
                default:
                  return const Center(child: Text('Not found'),);
              }
            }))
          ],
        ));
  }
}
