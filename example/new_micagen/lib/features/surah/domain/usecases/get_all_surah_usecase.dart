import "package:dartz/dartz.dart";
import 'package:injectable/injectable.dart';
import "package:new_micagen/core/error/failures.dart";
import "package:new_micagen/core/usecases/usecases.dart";
import "package:new_micagen/features/surah/domain/repositories/surah_repository.dart";
import "package:new_micagen/features/surah/data/models/surah_model.codegen.dart";

@injectable
class GetAllSurahUseCase implements UseCase<List<SurahModel>,NoParams>{
	final SurahRepository _surahRepository;

	GetAllSurahUseCase({
		required SurahRepository surahRepository,
}):_surahRepository =surahRepository;

	@override
	Future<Either<Failures,List<SurahModel>>> call(NoParams params) async{
		return await _surahRepository.getAllSurah(params);
	}
}
