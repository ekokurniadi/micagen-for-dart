// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// InjectableConfigGenerator
// **************************************************************************

import 'package:dio/dio.dart' as _i9;
import 'package:get_it/get_it.dart' as _i1;
import 'package:injectable/injectable.dart' as _i2;
import 'package:shared_preferences/shared_preferences.dart' as _i12;

import 'core/helpers/dio_helper.dart' as _i10;
import 'core/module/network_module.dart' as _i20;
import 'core/module/storage_module.dart' as _i21;
import 'features/ayat/data/datasources/local/ayat_local_datasource.dart' as _i3;
import 'features/ayat/data/datasources/local/ayat_local_datasource_impl.dart'
    as _i4;
import 'features/ayat/data/datasources/remote/ayat_remote_datasource.dart'
    as _i5;
import 'features/ayat/data/datasources/remote/ayat_remote_datasource_impl.dart'
    as _i6;
import 'features/ayat/data/repositories/ayat_repository_impl.dart' as _i8;
import 'features/ayat/domain/repositories/ayat_repository.dart' as _i7;
import 'features/ayat/domain/usecases/get_all_surah_usecase.dart' as _i11;
import 'features/surah/data/datasources/local/surah_local_datasource.dart'
    as _i13;
import 'features/surah/data/datasources/local/surah_local_datasource_impl.dart'
    as _i14;
import 'features/surah/data/datasources/remote/surah_remote_datasource.dart'
    as _i15;
import 'features/surah/data/datasources/remote/surah_remote_datasource_impl.dart'
    as _i16;
import 'features/surah/data/repositories/surah_repository_impl.dart' as _i18;
import 'features/surah/domain/repositories/surah_repository.dart' as _i17;
import 'features/surah/domain/usecases/get_all_surah_usecase.dart'
    as _i19; // ignore_for_file: unnecessary_lambdas

// ignore_for_file: lines_longer_than_80_chars
/// initializes the registration of provided dependencies inside of [GetIt]
Future<_i1.GetIt> $initGetIt(_i1.GetIt get,
    {String? environment, _i2.EnvironmentFilter? environmentFilter}) async {
  final gh = _i2.GetItHelper(get, environment, environmentFilter);
  final networkModule = _$NetworkModule();
  final storageModule = _$StorageModule();
  gh.lazySingleton<_i3.AyatLocalDataSource>(
      () => _i4.AyatLocalDataSourceImpl());
  gh.lazySingleton<_i5.AyatRemoteDataSource>(
      () => _i6.AyatRemoteDataSourceImpl());
  gh.lazySingleton<_i7.AyatRepository>(() => _i8.AyatRepositoryImpl(
      ayatLocalDataSource: get<_i3.AyatLocalDataSource>(),
      ayatRemoteDataSource: get<_i5.AyatRemoteDataSource>()));
  gh.lazySingleton<_i9.Dio>(() => networkModule.dio);
  gh.factory<_i10.DioHelper>(() => _i10.DioHelper());
  gh.factory<_i11.GetAllSurahUseCase>(
      () => _i11.GetAllSurahUseCase(ayatRepository: get<_i7.AyatRepository>()));
  await gh.lazySingletonAsync<_i12.SharedPreferences>(
      () => storageModule.sharedPreference,
      preResolve: true);
  gh.lazySingleton<_i13.SurahLocalDataSource>(
      () => _i14.SurahLocalDataSourceImpl());
  gh.lazySingleton<_i15.SurahRemoteDataSource>(
      () => _i16.SurahRemoteDataSourceImpl());
  gh.lazySingleton<_i17.SurahRepository>(() => _i18.SurahRepositoryImpl(
      surahLocalDataSource: get<_i13.SurahLocalDataSource>(),
      surahRemoteDataSource: get<_i15.SurahRemoteDataSource>()));
  gh.factory<_i19.GetAllSurahUseCase>(() =>
      _i19.GetAllSurahUseCase(surahRepository: get<_i17.SurahRepository>()));
  return get;
}

class _$NetworkModule extends _i20.NetworkModule {}

class _$StorageModule extends _i21.StorageModule {}
