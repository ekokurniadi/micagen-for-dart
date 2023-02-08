import "package:dartz/dartz.dart";
import "package:injectable/injectable.dart";
import "package:new_micagen/core/error/failures.dart";
import 'package:new_micagen/core/usecases/usecases.dart';
import "package:new_micagen/features/surah/domain/repositories/surah_repository.dart";
import "package:new_micagen/features/surah/data/datasources/remote/surah_remote_datasource.dart";
import "package:new_micagen/features/surah/data/datasources/local/surah_local_datasource.dart";
import "package:new_micagen/features/surah/data/models/surah_model.codegen.dart";

@LazySingleton(as:SurahRepository) 
class SurahRepositoryImpl implements SurahRepository {
	final SurahLocalDataSource _surahLocalDataSource;
	final SurahRemoteDataSource _surahRemoteDataSource;

	const SurahRepositoryImpl({
		required SurahLocalDataSource surahLocalDataSource,
		required SurahRemoteDataSource surahRemoteDataSource,
}):_surahLocalDataSource =surahLocalDataSource,
_surahRemoteDataSource =surahRemoteDataSource;

	@override
	Future<Either<Failures,List<SurahModel>>> getAllSurah(NoParams params) async{
		return await _surahRemoteDataSource.getAllSurah(params);
	}
}
