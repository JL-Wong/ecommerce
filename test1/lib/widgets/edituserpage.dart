import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:test1/controllers/editusercontroller.dart';
import 'package:test1/controllers/exchangecontroller.dart';
import 'package:test1/model/usermodel.dart';

class Edituserpage extends StatelessWidget {
  Edituserpage({super.key});

  final Editusercontroller _editusercontroller = Get.put(Editusercontroller());
  final Exchangecontroller _exchangecontroller = Get.find<Exchangecontroller>();

  @override
  Widget build(BuildContext context) {
    final accessToken = _exchangecontroller.accessToken;
    _editusercontroller.setAccessToken(accessToken);
    _editusercontroller.getUser();

    return Scaffold(
      body: Center(
        child: Obx(() {
          if (_editusercontroller.isLoading.value) {
            return const CircularProgressIndicator();
          } else {
            return DataTable(
              columns: const [
                DataColumn(label: Text('Username')),
                DataColumn(label: Text('Email')),
                DataColumn(label: Text('Edit')),
                DataColumn(label: Text('Delete')),
              ],
              rows: _editusercontroller.user.map((u) {
                return DataRow(
                  cells: [
                    DataCell(Text(u.username)),
                    DataCell(Text(u.email)),
                    DataCell(
                      IconButton(
                        icon: const Icon(Icons.edit),
                        onPressed: () {
                          // Navigate to edit screen or handle edit action
                          showEditDialog(context, u);
                        },
                      ),
                    ),
                    DataCell(
                      IconButton(
                        icon: const Icon(Icons.delete),
                        onPressed: () async {
                          bool success = await _editusercontroller.deleteUser(u.id);

                          if (success) {
                            _editusercontroller.getUser(); // Refresh user list
                          }
                        },
                      ),
                    ),
                  ],
                );
              }).toList(),
            );
          }
        }),
      ),
    );
  }

  void showEditDialog(BuildContext context, User user) {
    final TextEditingController usernameController = TextEditingController(text: user.username);
    final TextEditingController emailController = TextEditingController(text: user.email);

    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Edit User'),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              TextField(
                enabled: false,
                controller: usernameController,
                decoration: const InputDecoration(labelText: 'Username'),
              ),
              TextField(
                controller: emailController,
                decoration: const InputDecoration(labelText: 'Email'),
              ),
            ],
          ),
          actions: [
            TextButton(
              child: const Text('Cancel'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            TextButton(
              child: const Text('Save'),
              onPressed: () async {
                user.email = emailController.text;

                bool success = await _editusercontroller.updateUser(user);
                // print(success);
                if (success) {
                  bool success = await _editusercontroller.getUser(); // Refresh user list
                  if (success) {
                    Navigator.of(context).pop();
                  }
                  
                } else {
                  // Handle error
                }
              },
            ),
          ],
        );
      },
    );
  }
}
