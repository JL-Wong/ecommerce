import 'package:get/get.dart';
import 'package:test1/widgets/completeorderpage.dart';
import 'package:test1/widgets/edituserpage.dart';
import 'package:test1/widgets/homepage.dart';
import 'package:test1/widgets/loginpage.dart';
import 'package:test1/widgets/neworderpage.dart';
import 'package:test1/widgets/newuserpage.dart';
// import 'package:test1/widgets/productspage.dart';

class AppRoutes {
  static const String login = '/login';
  static const String home = '/home';
  static const String new_order = '/neworders';
  static const String complete_order = '/completeorders';
  static const String add_products = '/addproducts';
  static const String remove_products = '/removeproducts';
  static const String new_users = '/newusers';
  static const String edit_users = '/editusers';
  static const String delete_users = '/deleteusers';
  // static const String logout = '/logout';

  static final routes = [
    GetPage(name: login, page: () => Loginpage()),
    GetPage(name: home, page: () => Homepage()),
    GetPage(name: new_order, page: () => const Neworderpage()),
    GetPage(name: complete_order, page: () => const Completeorderpage()),
    // GetPage(name: add_products, page: () => ()),
    // GetPage(name: remove_products, page: () => ()),
    GetPage(name: new_users, page: () => Newuserpage()),
    GetPage(name: edit_users, page: () => Edituserpage()),
    // GetPage(name: delete_users, page: () => ()),
    
  ];
}