import 'package:get/get.dart';

class Screencontroller extends GetxController{

  var currentPage = 'home'.obs;

  void selectedpage(String page){
    currentPage.value = page;
  }
}