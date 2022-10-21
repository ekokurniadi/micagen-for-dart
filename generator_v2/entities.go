package generator_v2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekokurniadi/micagen-for-dart/schemas"
)

func GenerateEntity(entity schemas.Project) {
	createFile(entity)
	writeFileEntity(entity)
}

func writeFileEntity(entity schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(entity.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/domain/entities/" + strings.ToLower(entity.Entity.EntityName) + "_entity.codegen" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if entity.Config.UseFreezed {
		_, err = file.WriteString("import \"package:freezed_annotation/freezed_annotation.dart\";\n\n")
		if isError(err) {
			return "", err
		}
	} else {
		_, err = file.WriteString("import \"package:equatable/equatable.dart\";\n\n")
		if isError(err) {
			return "", err
		}
	}

	className := strings.Split(entity.Entity.EntityName, "_")
	var abstractName []string
	for i := range className {
		abstractName = append(abstractName, strings.Title(className[i]))
	}

	if entity.Config.UseFreezed {
		_, err = file.WriteString("part \"" + strings.ToLower(entity.Entity.EntityName) + "_entity.codegen.freezed" + ".dart\";\n")
		if isError(err) {
			return "", err
		}
		_, err = file.WriteString("part \"" + strings.ToLower(entity.Entity.EntityName) + "_entity.codegen.g" + ".dart\";\n\n")
		if isError(err) {
			return "", err
		}
		_, err = file.WriteString("@freezed\n")
		if isError(err) {
			return "", err
		}

		//Write some text line-by-line to file.
		_, err = file.WriteString("class " + strings.Join(abstractName, "") + "Entity with _$" + strings.Join(abstractName, "") + "Entity{\n")
		if isError(err) {
			return "", err
		}
		_, err = file.WriteString("\nconst factory " + strings.Join(abstractName, "") + "Entity({" + " \n")
		if isError(err) {
			return "", err
		}

		for a := 0; a < len(entity.Entity.EntityField); a++ {
			_, err = file.WriteString("\t" + "required " + entity.Entity.EntityField[a].DataType + " " + entity.Entity.EntityField[a].FieldName + ",\n")
			if isError(err) {
				return "", err
			}
		}

		_, err = file.WriteString("})=" + "_" + strings.Join(abstractName, "") + "Entity;" + " \n\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\nfactory " + strings.Join(abstractName, "") + "Entity.fromJson(Map<String,dynamic>json) => _$" + strings.Join(abstractName, "") + "EntityFromJson(json);\n\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("}\n")
		if isError(err) {
			return "", err
		}

	} else {
		_, err = file.WriteString("\n")
		if isError(err) {
			return "", err
		}

		//Write some text line-by-line to file.
		_, err = file.WriteString("class " + strings.Join(abstractName, "") + "Entity extends Equatable{\n")
		if isError(err) {
			return "", err
		}

		for a := 0; a < len(entity.Entity.EntityField); a++ {
			_, err = file.WriteString("\t" + "final " + entity.Entity.EntityField[a].DataType + "? " + entity.Entity.EntityField[a].FieldName + ";\n")
			if isError(err) {
				return "", err
			}
		}

		_, err = file.WriteString("\nconst " + strings.Join(abstractName, "") + "Entity({" + " \n")
		if isError(err) {
			return "", err
		}

		for b := 0; b < len(entity.Entity.EntityField); b++ {
			_, err = file.WriteString("\t" + "required this." + entity.Entity.EntityField[b].FieldName + ",\n")
			if isError(err) {
				return "", err
			}
		}

		_, err = file.WriteString("\n});" + " \n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\t@override\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\tList<Object?> get props =>[")
		if isError(err) {
			return "", err
		}

		for e := 0; e < len(entity.Entity.EntityField); e++ {
			_, err = file.WriteString("\t" + entity.Entity.EntityField[e].FieldName + ",")
			if isError(err) {
				return "", err
			}
		}

		_, err = file.WriteString("];\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("}\n")
		if isError(err) {
			return "", err
		}

	}

	return "", nil

}

func createFile(entity schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(entity.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/domain/entities/" + strings.ToLower(entity.Entity.EntityName) + "_entity.codegen" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	filename, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err.Error())
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}

func isError(err error) bool {

	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
