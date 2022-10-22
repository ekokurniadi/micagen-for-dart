Micagen For Dart - Flutter Clean Architecture Generator

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
```

### Template JSON

```json
{
  "config": {
    "use_freezed": true
  },
  "feature_name": "ocr",
  "entity": {
    "entity_name": "ocr",
    "entity_field": [
      {
        "data_type": "String",
        "field_name": "result"
      },
      {
        "data_type": "String",
        "field_name": "version"
      }
    ]
  },
  "usecases": [
    {
      "usecase_name": "InsertImage",
      "parameter": "OcrModel",
      "return_type": "bool"
    },
    {
      "usecase_name": "UpdateImage",
      "parameter": "OcrModel",
      "return_type": "OcrEntity"
    },
    {
      "usecase_name": "GetTextFromImage",
      "parameter": "OcrModel",
      "return_type": "OcrEntity"
    },
    {
      "usecase_name": "DeleteImage",
      "parameter": "int",
      "return_type": "bool"
    }
  ]
}
```
