import "package:dartz/dartz.dart";
import "package:new_micagen/core/error/failures.dart";
import 'package:new_micagen/core/usecases/usecases.dart';
import "package:new_micagen/features/ayat/data/models/ayat_model.codegen.dart";

abstract class AyatRepository {
	Future<Either<Failures,List<AyatModel>>> getAllSurah(NoParams params);
}
