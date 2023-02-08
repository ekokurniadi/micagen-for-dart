// GENERATED CODE - DO NOT MODIFY BY HAND

// ignore_for_file: implicit_dynamic_parameter

part of 'ayat_entity.codegen.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_AyatEntity _$$_AyatEntityFromJson(Map<String, dynamic> json) =>
    _$_AyatEntity(
      nomor: json['nomor'] as int,
      nama: json['nama'] as String,
      namaLatin: json['nama_latin'] as String,
      jumlahAyat: json['jumlah_ayat'] as int,
      tempatTurun: json['tempat_turun'] as String,
      arti: json['arti'] as String,
      deskripsi: json['deskripsi'] as String,
      audio: json['audio'] as String,
    );

Map<String, dynamic> _$$_AyatEntityToJson(_$_AyatEntity instance) =>
    <String, dynamic>{
      'nomor': instance.nomor,
      'nama': instance.nama,
      'nama_latin': instance.namaLatin,
      'jumlah_ayat': instance.jumlahAyat,
      'tempat_turun': instance.tempatTurun,
      'arti': instance.arti,
      'deskripsi': instance.deskripsi,
      'audio': instance.audio,
    };
