import "package:freezed_annotation/freezed_annotation.dart";

part "ayat_model.codegen.freezed.dart";
part "ayat_model.codegen.g.dart";

@freezed
class AyatModel with _$AyatModel{

const factory AyatModel({ 
	required int nomor,
	required String nama,
	required String namaLatin,
	required int jumlahAyat,
	required String tempatTurun,
	required String arti,
	required String deskripsi,
	required String audio,
})=_AyatModel; 


factory AyatModel.fromJson(Map<String,dynamic>json) => _$AyatModelFromJson(json);

}
