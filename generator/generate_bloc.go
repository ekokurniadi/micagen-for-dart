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

func GenerateBloc(project schemas.Project) {
	createFileBloc(project)
	createFileEvent(project)
	createFileState(project)
	writeFileBloc(project)
	writeFileEvent(project)
	writeFileState(project)
}

func writeFileEvent(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/presentations/bloc/" + realName + "_event" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString("part of \"" + realName + "_bloc.dart" + "\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	_, err = file.WriteString("@freezed\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "Event with _$" + strings.Join(abstractName, "") + "Event{\n")
	if isError(err) {
		return "", err
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

		_, err = file.WriteString("\tconst factory " + strings.Join(abstractName, "") + "Event." + usesCaseName + "(\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\t" + project.UseCase[i].Parameter + " params,\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString(")=" + strings.Title(usesCaseName) + "Event\n")
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

func writeFileBloc(project schemas.Project) (string, error) {

	var realFeatureName []string
	var splitText []string
	if strings.Contains(project.FeatureName, " ") {
		splitText = strings.Split(project.FeatureName, " ")
	} else if strings.Contains(project.FeatureName, "-") {
		splitText = strings.Split(project.FeatureName, "_")
	}

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/presentations/bloc/" + realName + "_bloc" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString("import \"package:bloc/bloc.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"package:freezed_annotation/freezed_annotation.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"package:injectable/injectable.dart\";\n")
	if isError(err) {
		return "", err
	}

	for i := 0; i < len(project.UseCase); i++ {
		usecases := camelcase.Split(project.UseCase[i].MethodName)
		var u []string
		for i := range usecases {

			u = append(u, strings.ToLower(usecases[i]))

		}
		usesCaseName := strings.Join(u, "_")

		_, err = file.WriteString("import \"" + project.OutputPath + "/features/" + realName + "/domain/usecases/" + usesCaseName + "usecase" + ".dart\";\n")
		if isError(err) {
			return "", err
		}

	}

	_, err = file.WriteString("part \"" + realName + "_event.dart" + "\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("part \"" + realName + "_state.dart" + "\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("part \"" + realName + "_bloc.freezed.dart" + "\";\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("@injectable\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	className := strings.Join(abstractName, "")

	_, err = file.WriteString("class " + className + "Bloc extends Bloc<" + className + "Event," + className + "State>{\n")
	if isError(err) {
		return "", err
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

		_, err = file.WriteString("final " + strings.Title(usesCaseName) + "UseCase _" + usesCaseName + "Usecase;\n")
		if isError(err) {
			return "", err
		}

	}

	_, err = file.WriteString("class " + className + "Bloc(\n")
	if isError(err) {
		return "", err
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

		_, err = file.WriteString("\tthis._" + usesCaseName + "Usecase,\n")
		if isError(err) {
			return "", err
		}

	}
	_, err = file.WriteString("\t):super(" + className + "State.initial()){\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("on<" + className + "Event>((event,emit){});\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	return "", nil
}

func writeFileState(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/presentations/bloc/" + realName + "_state" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString("part of \"" + realName + "_bloc.dart" + "\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	_, err = file.WriteString("enum " + strings.Join(abstractName, "") + "Status {initial,loading,loaded,failure}\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("@freezed\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "State with _$" + strings.Join(abstractName, "") + "State{\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\tconst factory " + strings.Join(abstractName, "") + "State({\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\trequired " + strings.Join(abstractName, "") + "Status status,\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\trequired List<" + strings.Join(abstractName, "") + "Entity> data,")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\trequired String errorMessage,\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t})= _" + strings.Join(abstractName, "") + "State\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\tfactory " + strings.Join(abstractName, "") + "State.initial()=> const " + strings.Join(abstractName, "") + "State(\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\t\tstatus: " + strings.Join(abstractName, "") + "Status.initial,\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t\tdata: []\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t\terrorMessage: \"\"\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t);\n\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	return "", nil
}

func createFileBloc(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/presentations/bloc/" + realName + "_bloc" + ".dart")
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

func createFileEvent(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/presentations/bloc/" + realName + "_event" + ".dart")
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

func createFileState(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./features/" + realName + "/presentations/bloc/" + realName + "_state" + ".dart")
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
