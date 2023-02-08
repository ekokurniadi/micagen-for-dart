import "package:dartz/dartz.dart";
import 'package:injectable/injectable.dart';
import "package:new_micagen/core/error/failures.dart";
import "package:new_micagen/core/usecases/usecases.dart";
import "package:new_micagen/features/ayat/domain/repositories/ayat_repository.dart";
import "package:new_micagen/features/ayat/data/models/ayat_model.codegen.dart";

@injectable
class GetAllSurahUseCase implements UseCase<List<AyatModel>,NoParams>{
	final AyatRepository _ayatRepository;

	GetAllSurahUseCase({
		required AyatRepository ayatRepository,
}):_ayatRepository =ayatRepository;

	@override
	Future<Either<Failures,List<AyatModel>>> call(NoParams params) async{
		return await _ayatRepository.getAllSurah(params);
	}
}
