import "package:dartz/dartz.dart";
import "package:new_micagen/core/error/failures.dart";
import 'package:new_micagen/core/usecases/usecases.dart';
import "package:new_micagen/features/surah/data/models/surah_model.codegen.dart";

abstract class SurahLocalDataSource {
	Future<Either<Failures,List<SurahModel>>> getAllSurah(NoParams params);
}
