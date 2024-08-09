import 'package:flutter/material.dart';
import 'package:get/get.dart';
// import 'package:test1/controllers/activetabscontroller.dart';
import 'package:test1/routes/routes.dart';
import 'package:url_strategy/url_strategy.dart';

void main() {
  setPathUrlStrategy();

  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  MyApp({super.key});
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      initialRoute: AppRoutes.login,
      getPages: AppRoutes.routes,
      // home: Homepage(),
      // initialBinding: BindingsBuilder((){
      //   Get.put(Activetabscontroller());
      // }),
    );
  }
}

