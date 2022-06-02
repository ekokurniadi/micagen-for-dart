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

func GenerateRepositoryImpl(project schemas.Project) {
	createFileRepositoryImpl(project)
	writeFileRepositoryImpl(project)
}

func writeFileRepositoryImpl(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/data/repository/" + realName + "_repository_impl" + ".dart")
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
	_, err = file.WriteString("import \"../../../core/error/failures.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"../../domain/repository/" + realName + "_repository" + ".dart\";\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("import \"../../data/datasource/" + realName + "_datasource" + ".dart\";\n")
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
	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "RepositoryImpl extends " + strings.Join(abstractName, "") + "Repository {\n")
	if isError(err) {
		return "", err
	}

	datasource := camelcase.Split(realName)
	var datasources []string
	for i := range datasource {
		if i > 0 {
			datasources = append(datasources, strings.Title(datasource[i]))
		} else {
			datasources = append(datasources, datasource[i])
		}

	}

	_, err = file.WriteString("\tfinal " + strings.Join(abstractName, "") + "DataSource " + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "DataSource;\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t" + strings.Join(abstractName, "") + "RepositoryImpl({\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t\trequired " + strings.Join(abstractName, "") + "DataSource " + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "DataSource,\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}):" + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "DataSource =" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "DataSource;\n\n")
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

	_, err = file.WriteString("\t\treturn await " + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "DataSource.execute();\n")
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

func createFileRepositoryImpl(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/data/repository/" + realName + "_repository_impl" + ".dart")
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
