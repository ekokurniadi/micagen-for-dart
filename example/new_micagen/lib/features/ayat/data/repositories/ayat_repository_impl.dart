import "package:dartz/dartz.dart";
import "package:injectable/injectable.dart";
import "package:new_micagen/core/error/failures.dart";
import 'package:new_micagen/core/usecases/usecases.dart';
import "package:new_micagen/features/ayat/domain/repositories/ayat_repository.dart";
import "package:new_micagen/features/ayat/data/datasources/remote/ayat_remote_datasource.dart";
import "package:new_micagen/features/ayat/data/datasources/local/ayat_local_datasource.dart";
import "package:new_micagen/features/ayat/data/models/ayat_model.codegen.dart";

@LazySingleton(as:AyatRepository) 
class AyatRepositoryImpl implements AyatRepository {
	final AyatLocalDataSource _ayatLocalDataSource;
	final AyatRemoteDataSource _ayatRemoteDataSource;

	const AyatRepositoryImpl({
		required AyatLocalDataSource ayatLocalDataSource,
		required AyatRemoteDataSource ayatRemoteDataSource,
}):_ayatLocalDataSource =ayatLocalDataSource,
_ayatRemoteDataSource =ayatRemoteDataSource;

	@override
	Future<Either<Failures,List<AyatModel>>> getAllSurah(NoParams params) async{
		return await _ayatRemoteDataSource.getAllSurah(params);
	}
}
