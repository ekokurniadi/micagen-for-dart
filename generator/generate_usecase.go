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

func GenerateUseCase(project schemas.Project) {
	createFileUseCase(project)
	writeFileUseCase(project)
}

func writeFileUseCase(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/domain/usecases/" + realName + "_usecase" + ".dart")
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
	_, err = file.WriteString("import \"package:equatable/equatable.dart\";\n")
	if isError(err) {
		return "", err
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"../../../../core/error/failures.dart\";\n")
	if isError(err) {
		return "", err
	}
	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"../../../../core/usecases/usecases.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"../../domain/repository/" + realName + "_repository" + ".dart\";\n")
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

	model := strings.Split(project.Entity.EntityName, "_")
	var models []string
	for i := range model {
		models = append(models, strings.Title(model[i]))
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "UseCase implements UseCase<" + strings.Join(models, "") + "Model" + ",NoParams>{\n")
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

	_, err = file.WriteString("\tfinal " + strings.Join(abstractName, "") + "Repository " + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository;\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t" + strings.Join(abstractName, "") + "UseCase({\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t\trequired " + strings.Join(abstractName, "") + "Repository " + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository,\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}):" + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository =" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository;\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t@override\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\tFuture<Either<Failures," + strings.Join(models, "") + "Model>> call(NoParams params) async{\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\treturn await " + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository.execute();\n")
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

	_, err = file.WriteString("class Params extends Equatable{\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\tconst Params();\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t@override\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\tList<Object?> get props => [];\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	return "", nil

}

func createFileUseCase(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/domain/usecases/" + realName + "_usecase" + ".dart")
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
