import "package:freezed_annotation/freezed_annotation.dart";

part "ayat_entity.codegen.freezed.dart";
part "ayat_entity.codegen.g.dart";

@freezed
class AyatEntity with _$AyatEntity{

const factory AyatEntity({ 
	required int nomor,
	required String nama,
	required String namaLatin,
	required int jumlahAyat,
	required String tempatTurun,
	required String arti,
	required String deskripsi,
	required String audio,
})=_AyatEntity; 


factory AyatEntity.fromJson(Map<String,dynamic>json) => _$AyatEntityFromJson(json);

}
