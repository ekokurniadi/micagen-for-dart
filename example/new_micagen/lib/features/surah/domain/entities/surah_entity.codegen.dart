import "package:freezed_annotation/freezed_annotation.dart";

part "surah_entity.codegen.freezed.dart";
part "surah_entity.codegen.g.dart";

@freezed
class SurahEntity with _$SurahEntity{

const factory SurahEntity({ 
	required int nomor,
	required String nama,
	required String namaLatin,
	required int jumlahAyat,
	required String tempatTurun,
	required String arti,
	required String deskripsi,
	required String audio,
})=_SurahEntity; 


factory SurahEntity.fromJson(Map<String,dynamic>json) => _$SurahEntityFromJson(json);

}
