package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekokurniadi/micagen-for-dart/schemas"

	"github.com/fatih/camelcase"
)

func GenerateModels(project schemas.Project) {
	createFileModels(project)
	writeFileModels(project)
}

func writeFileModels(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/data/models/" + strings.ToLower(project.Entity.EntityName) + "_model" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"../../domain/entity/" + project.Entity.EntityName + "_entity.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	className := strings.Split(project.Entity.EntityName, "_")
	var abstractName []string
	for i := range className {
		abstractName = append(abstractName, strings.Title(className[i]))
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "Model extends " + strings.Join(abstractName, "") + "Entity{\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\nconst " + strings.Join(abstractName, "") + "Model({" + " \n")
	if isError(err) {
		return "", err
	}

	for b := 0; b < len(project.Entity.EntityField); b++ {
		_, err = file.WriteString("\t\t" + "required " + project.Entity.EntityField[b].DataType + "? " + project.Entity.EntityField[b].FieldName + ",\n")
		if isError(err) {
			return "", err
		}
	}

	_, err = file.WriteString("\t}):super(" + " \n")
	if isError(err) {
		return "", err
	}

	for b := 0; b < len(project.Entity.EntityField); b++ {
		_, err = file.WriteString("\t" + project.Entity.EntityField[b].FieldName + ":" + project.Entity.EntityField[b].FieldName + ",\n")
		if isError(err) {
			return "", err
		}
	}

	_, err = file.WriteString("\t);" + " \n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\n\tfactory " + strings.Join(abstractName, "") + "Model" + ".fromJson(Map<String, dynamic> json) => " + strings.Join(abstractName, "") + "Model" + "(\n")
	if isError(err) {
		return "", err
	}

	for c := 0; c < len(project.Entity.EntityField); c++ {
		var camelCase_formatter []string

		splitted := camelcase.Split(project.Entity.EntityField[c].FieldName)
		for a := 0; a < len(splitted); a++ {
			camelCase_formatter = append(camelCase_formatter, strings.ToLower(splitted[a]))
		}
		replaceCamelCase := strings.Join(camelCase_formatter, "_")
		_, err = file.WriteString("\t\t" + project.Entity.EntityField[c].FieldName + ": json['" + replaceCamelCase + "'],\n")
		if isError(err) {
			return "", err
		}
	}

	_, err = file.WriteString("\t);\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\tMap<String, dynamic> toJson() => {\n")
	if isError(err) {
		return "", err
	}

	for d := 0; d < len(project.Entity.EntityField); d++ {
		var camelCase_formatter []string

		splitted := camelcase.Split(project.Entity.EntityField[d].FieldName)
		for a := 0; a < len(splitted); a++ {
			camelCase_formatter = append(camelCase_formatter, strings.ToLower(splitted[a]))
		}
		replaceCamelCase := strings.Join(camelCase_formatter, "_")
		_, err = file.WriteString("\t\t" + "\"" + replaceCamelCase + "\": " + project.Entity.EntityField[d].FieldName + ",\n")
		if isError(err) {
			return "", err
		}
	}
	_, err = file.WriteString("};\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}" + " \n")
	if isError(err) {
		return "", err
	}

	return "", nil
}

func createFileModels(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/data/models/" + strings.ToLower(project.Entity.EntityName) + "_model" + ".dart")
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
