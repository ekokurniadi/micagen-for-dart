import "package:dartz/dartz.dart";
import "package:injectable/injectable.dart";
import "package:new_micagen/core/error/failures.dart";
import 'package:new_micagen/core/usecases/usecases.dart';
import "surah_remote_datasource.dart";
import "package:new_micagen/features/surah/data/models/surah_model.codegen.dart";

@LazySingleton(as:SurahRemoteDataSource) 
class SurahRemoteDataSourceImpl implements SurahRemoteDataSource {
	@override
	Future<Either<Failures,List<SurahModel>>> getAllSurah(NoParams params) async{
		// TODO: implement execute 
		throw UnimplementedError();
	}
}
