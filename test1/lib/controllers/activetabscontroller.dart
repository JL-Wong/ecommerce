import 'package:get/get.dart';
import 'package:web/web.dart' as web;

class Activetabscontroller extends GetxController{
  final _activeTabs = 0.obs;

  @override
  void onInit() {
    super.onInit();
    _loadActiveTabs();
  }

  void _loadActiveTabs() {
    _activeTabs.value = int.parse(web.window.localStorage['activeTabs'] ?? '0');
  }

  void incrementTabs() {
    _activeTabs.value += 1;
    web.window.localStorage['activeTabs'] = _activeTabs.value.toString();
  }

  void decrementTabs() {
    _activeTabs.value -= 1;
    web.window.localStorage['activeTabs'] = _activeTabs.value.toString();
  }
}