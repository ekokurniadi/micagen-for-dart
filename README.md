Micagen For Dart - Flutter Clean Architecture Generator

## Example Project 
https://github.com/ekokurniadi/flutter_mahasiswa

## How to Used ?

### 1. Open your existing or new project flutter

### 2. Open lib folder

### 3. open your terminal and enter to the lib directory

### 4. Running command

```sh
go mod init generated
```

### 5. Create new file main.go

```go
package main

func main(){

}
```

### 6. Running

```sh
go get github.com/ekokurniadi/micagen-for-dart
```

### on your terminal on lib directory

### 7. Modify main.go

```go
package main

import mika "github.com/ekokurniadi/micagen-for-dart"

func main() {
	gen := mika.Initialized()
	gen.RunGenerator()
}
```

### 8. Running command

```sh
go run main.go
```

### on terminal on lib directory

### 9. Open your Postman or something as same

### 10. Create new request to endpoint http://localhost:8080/v2/generate with POST Method

### 11. Copy template json and modify

### 12. Klik Send

### 13. You must add some package on your pubspec.yaml
```yaml
  dartz: ^<version>
  equatable: ^<version>
  freezed: ^<version>
  freezed_annotation: ^<version>
  json_serializable: ^<version>
  dio: ^<version>
  build_runner: ^<version>
  get_it: ^<version>
  injectable: ^<version>
  injectable_generator: ^<version>
  shared_preferences: ^<version>
```

### Template JSON

```json
{
    "config": {
        "use_freezed": true,
        "use_injectable": true
    },
    "feature_name": "ayat",
    "entity": {
        "entity_name": "ayat",
        "entity_field": [
            {
                "data_type": "int",
                "field_name": "nomor"
            },
            {
                "data_type": "String",
                "field_name": "nama"
            },
            {
                "data_type": "String",
                "field_name": "namaLatin"
            },
            {
                "data_type": "int",
                "field_name": "jumlahAyat"
            },
            {
                "data_type": "String",
                "field_name": "tempatTurun"
            },
            {
                "data_type": "String",
                "field_name": "arti"
            },
            {
                "data_type": "String",
                "field_name": "deskripsi"
            },
            {
                "data_type": "String",
                "field_name": "audio"
            }
        ]
    },
    "usecases": [
        {
            "usecase_name": "GetAllSurah",
            "parameter": "NoParams",
            "return_type": "List<AyatModel>"
        }
    ],
    "generator_options": {
        "generate_model": true,
        "generate_usecase": true,
        "generate_local_data_source": true,
        "generate_remote_data_source": true,
        "generate_repository": true,
        "generate_module": true
    }
}
```
