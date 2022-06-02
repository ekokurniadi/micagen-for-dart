package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekokurniadi/micagen-for-dart/schemas"
)

func GenerateDataSourceImpl(project schemas.Project) {
	createFileDataSourceImpl(project)
	writeFileDataSourceImpl(project)
}

func writeFileDataSourceImpl(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/data/datasource/" + realName + "_datasource_impl" + ".dart")
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
	_, err = file.WriteString("import \"package:dartz/dartz.dart\";\n")
	if isError(err) {
		return "", err
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"../../../../core/error/failures.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"" + realName + "_datasource" + ".dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"../../data/models/" + project.Entity.EntityName + "_model.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "DataSourceImpl extends " + strings.Join(abstractName, "") + "DataSource {\n")
	if isError(err) {
		return "", err
	}

	model := strings.Split(project.Entity.EntityName, "_")
	var models []string
	for i := range model {
		models = append(models, strings.Title(model[i]))
	}

	_, err = file.WriteString("\t@override\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\tFuture<Either<Failures," + strings.Join(models, "") + "Model>> execute() async{\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\t// TODO: implement execute \n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\tthrow UnimplementedError();\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	return "", nil

}

func createFileDataSourceImpl(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/data/datasource/" + realName + "_datasource_impl" + ".dart")
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
