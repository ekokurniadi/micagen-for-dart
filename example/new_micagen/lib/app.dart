import 'package:new_micagen/core/helpers/dio_helper.dart';
import 'package:new_micagen/injector.dart';
import 'package:shared_preferences/shared_preferences.dart';

class App{
  const App._();

  static Future<void>init()async{
    DioHelper.initialDio("");
    DioHelper.setDioLogger("");
    String? token = getIt<SharedPreferences>().getString("token");
    DioHelper.setDioHeader(token);
  }
}