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

func GenerateUseCase(project schemas.Project) {
	for i := 0; i < len(project.UseCase); i++ {
		createNewFileUseCase(project.UseCase[i], project)
		writeNewFileUseCase(project.UseCase[i], project)
	}
}

func writeNewFileUseCase(usecase schemas.UseCase, p schemas.Project) (string, error) {
	usecases := camelcase.Split(usecase.MethodName)
	var u []string
	for i := range usecases {

		u = append(u, strings.Title(usecases[i]))

	}
	var x []string
	for y := range usecases {

		x = append(x, strings.ToLower(usecases[y]))

	}

	usesCaseName := strings.Join(u, "")
	usesCaseNameFile := strings.Join(x, "_")

	var realFeatureName []string
	splitText := strings.Split(p.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/domain/usecases/" + usesCaseNameFile + "_usecase" + ".dart")
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
	_, err = file.WriteString("import \"" + p.OutputPath + "/core/error/failures.dart\";\n")
	if isError(err) {
		return "", err
	}
	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"" + p.OutputPath + "/core/usecases/usecases.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"" + p.OutputPath + "/domain/repositories/" + realName + "_repository" + ".dart\";\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("import \"" + p.OutputPath + "/data/models/" + p.Entity.EntityName + "_model.codegen.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	model := strings.Split(p.Entity.EntityName, "_")
	var models []string
	for i := range model {
		models = append(models, strings.Title(model[i]))
	}

	//Write some text line-by-line to file.
	_, err = file.WriteString("class " + usesCaseName + "UseCase implements UseCase<" + usecase.ReturnType + "," + usecase.Parameter + ">{\n")
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

	var s []string
	for j := range usecases {
		if j == 0 {
			s = append(s, strings.ToLower(usecases[j]))
		} else {
			s = append(s, strings.Title(usecases[j]))
		}

	}
	usesCaseMethodName := strings.Join(s, "")

	_, err = file.WriteString("\tfinal " + strings.Join(abstractName, "") + "Repository " + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository;\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t" + usesCaseName + "UseCase({\n")
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

	_, err = file.WriteString("\tFuture<Either<Failures," + usecase.ReturnType + ">> call(" + usecase.Parameter + " params) async{\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\treturn await " + "_" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "Repository." + usesCaseMethodName + "(params);\n")
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

func createNewFileUseCase(usecase schemas.UseCase, p schemas.Project) (string, error) {
	usecases := camelcase.Split(usecase.MethodName)
	var x []string
	for y := range usecases {
		x = append(x, strings.ToLower(usecases[y]))

	}

	usesCaseNameFile := strings.Join(x, "_")
	var realFeatureName []string
	splitText := strings.Split(p.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/domain/usecases/" + usesCaseNameFile + "_usecase" + ".dart")
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
