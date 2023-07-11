package generator_v2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekokurniadi/micagen-for-dart/schemas"
	"github.com/fatih/camelcase"
)

func GenerateRemoteDataSource(project schemas.Project) {
	createFileRemoteDataSource(project)
	writeFileRemoteDataSource(project)
}

func writeFileRemoteDataSource(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/data/datasources/remote/" + realName + "_remote_datasource" + ".dart")
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
	_, err = file.WriteString("import \"" + project.OutputPath + "/core/error/failures.dart\";\n")
	if isError(err) {
		return "", err
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"" + project.OutputPath + "/core/usecases/usecase.dart\";\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("import \"" + project.OutputPath + "/features/" + realName + "/data/models/" + project.Entity.EntityName + "_model.codegen.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("abstract class " + strings.Join(abstractName, "") + "RemoteDataSource {\n")
	if isError(err) {
		return "", err
	}

	model := strings.Split(project.Entity.EntityName, "_")
	var models []string
	for i := range model {
		models = append(models, strings.Title(model[i]))
	}

	for i := 0; i < len(project.UseCase); i++ {
		usecases := camelcase.Split(project.UseCase[i].MethodName)
		var u []string
		for i := range usecases {
			if i == 0 {
				u = append(u, strings.ToLower(usecases[i]))
			} else {
				u = append(u, strings.Title(usecases[i]))
			}

		}
		usesCaseName := strings.Join(u, "")

		_, err = file.WriteString("\tFuture<Either<Failures," + project.UseCase[i].ReturnType + ">> " + usesCaseName + "(" + project.UseCase[i].Parameter + " params" + ");\n")
		if isError(err) {
			return "", err
		}

	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	return "", nil

}

func createFileRemoteDataSource(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/data/datasources/remote/" + realName + "_remote_datasource" + ".dart")
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
