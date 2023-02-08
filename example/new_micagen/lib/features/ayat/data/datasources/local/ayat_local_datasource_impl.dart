import "package:dartz/dartz.dart";
import "package:injectable/injectable.dart";
import "package:new_micagen/core/error/failures.dart";
import 'package:new_micagen/core/usecases/usecases.dart';
import "ayat_local_datasource.dart";
import "package:new_micagen/features/ayat/data/models/ayat_model.codegen.dart";

@LazySingleton(as:AyatLocalDataSource) 
class AyatLocalDataSourceImpl implements AyatLocalDataSource {
	@override
	Future<Either<Failures,List<AyatModel>>> getAllSurah(NoParams params) async{
		// TODO: implement execute 
		throw UnimplementedError();
	}
}
