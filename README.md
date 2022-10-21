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

### 10. Create new request to endpoint http://localhost:8080/generate with POST Method

### 11. Copy template json and modify

### 12. Klik Send

### Template JSON

```json
{
  "config": {
    "null_safety": true,
    "use_dartz_package": true,
    "use_dependency_injection": true,
    "enable_unit_testing": true
  },
  "feature_name": "users",
  "entity": {
    "entity_name": "users",
    "entity_field": [
      {
        "data_type": "int",
        "field_name": "id"
      },
      {
        "data_type": "String",
        "field_name": "userName"
      },
      {
        "data_type": "String",
        "field_name": "password"
      },
      {
        "data_type": "String",
        "field_name": "token"
      }
    ]
  },
  "state_management_options": {
    "name": "users",
    "state_management": "cubit"
  },
  "state": [
    {
      "state_name": "UsersSuccess",
      "params_or_models": [
        { "type": "model", "use_list": true, "name": "usersData" },
        { "type": "String", "use_list": false, "name": "message" }
      ]
    }
  ],
  "methods": [{ "method_name": "getUsers" }]
}
```
