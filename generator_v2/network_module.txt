import 'package:dio/dio.dart';
import 'package:injectable/injectable.dart';
import '../helpers/dio_helper.dart';

@module
abstract class NetworkModule {
  @lazySingleton
  Dio get dio => DioHelper.dio!;
  
}
