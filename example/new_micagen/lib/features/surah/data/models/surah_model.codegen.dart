import "package:freezed_annotation/freezed_annotation.dart";

part "surah_model.codegen.freezed.dart";
part "surah_model.codegen.g.dart";

@freezed
class SurahModel with _$SurahModel{

const factory SurahModel({ 
	required int nomor,
	required String nama,
	required String namaLatin,
	required int jumlahAyat,
	required String tempatTurun,
	required String arti,
	required String deskripsi,
	required String audio,
})=_SurahModel; 


factory SurahModel.fromJson(Map<String,dynamic>json) => _$SurahModelFromJson(json);

}
