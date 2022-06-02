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

func GenerateCubit(project schemas.Project) {
	createFileCubit(project)
	writeFileCubit(project)
	createFileCubitState(project)
	writeFileCubitState(project)
}

func createFileCubit(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/presentation/bloc/" + realName + "_cubit" + ".dart")
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
func writeFileCubit(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/presentation/bloc/" + realName + "_cubit" + ".dart")
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
	_, err = file.WriteString("import \"package:bloc/bloc.dart\";\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("import \"package:equatable/equatable.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"../../domain/usecases/" + realName + "_usecase.dart\";\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"../../data/models/" + project.Entity.EntityName + "_model.dart\";\n")
	if isError(err) {
		return "", err
	}
	//Write some text line-by-line to file.
	_, err = file.WriteString("import \"../../../../core/usecases/usecases.dart\";\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("part \"" + realName + "_state.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "Cubit " + "extends Cubit<" + strings.Join(abstractName, "") + "State> {\n")
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

	_, err = file.WriteString("\n\tfinal " + strings.Join(abstractName, "") + "UseCase _" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "UseCase;\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("\n\t" + strings.Join(abstractName, "") + "Cubit({required " + strings.Join(abstractName, "") + "UseCase " + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "UseCase,}) " + ": _" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "UseCase = " + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "UseCase, super(" + strings.Join(abstractName, "") + "Initial());\n")
	if isError(err) {
		return "", err
	}

	for _, method := range project.VoidMapping {
		_, err = file.WriteString("\n\t" + "Future<void> " + method.VoidName + "() async{\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\tfinal result = await _" + strings.Replace(strings.Join(datasources, ""), "_", "", -1) + "UseCase.call(NoParams());\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\tresult.fold((l)=>null,(r)=>null);\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\t}\n")
		if isError(err) {
			return "", err
		}
	}

	_, err = file.WriteString("}")
	if isError(err) {
		return "", err
	}

	return "", nil
}

func createFileCubitState(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/presentation/bloc/" + realName + "_state" + ".dart")
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
func writeFileCubitState(project schemas.Project) (string, error) {
	var realFeatureName []string
	splitText := strings.Split(project.FeatureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	filepath, err := filepath.Abs("./" + realName + "/presentation/bloc/" + realName + "_state" + ".dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString("part of \"" + realName + "_cubit.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	abstract := strings.Split(realName, "_")
	var abstractName []string
	for i := range abstract {
		abstractName = append(abstractName, strings.Title(abstract[i]))
	}

	_, err = file.WriteString("abstract class " + strings.Join(abstractName, "") + "State " + "extends Equatable {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\n\tconst " + strings.Join(abstractName, "") + "State();\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\t@override\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\tList<Object> get props => [];\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("class " + strings.Join(abstractName, "") + "Initial extends " + strings.Join(abstractName, "") + "State {}\n\n")
	if isError(err) {
		return "", err
	}

	for x := range project.State {
		stateName := strings.Split(project.State[x].StateName, "_")
		var states []string
		for i := range stateName {
			states = append(states, strings.Title(project.State[i].StateName))
		}
		_, err = file.WriteString("class " + strings.Join(states, "") + " extends " + strings.Join(abstractName, "") + "State {\n")
		if isError(err) {
			return "", err
		}

		model := strings.Split(project.Entity.EntityName, "_")
		var models []string
		for i := range model {
			models = append(models, strings.Title(model[i]))
		}
		for y := range project.State[x].ParamsOrModels {
			if project.State[x].ParamsOrModels[y].Type == "model" && project.State[x].ParamsOrModels[y].UseList {
				_, err = file.WriteString("\tfinal" + " " + "List<" + strings.Join(models, "") + "Model" + ">" + " " + project.State[x].ParamsOrModels[y].Name + ";\n")
				if isError(err) {
					return "", err
				}
			} else if project.State[x].ParamsOrModels[y].Type == "model" && !project.State[x].ParamsOrModels[y].UseList {
				_, err = file.WriteString("\tfinal" + " " + strings.Join(models, "") + "Model" + " " + project.State[x].ParamsOrModels[y].Name + ";\n")
				if isError(err) {
					return "", err
				}
			} else {
				_, err = file.WriteString("\tfinal" + " " + project.State[x].ParamsOrModels[y].Type + " " + project.State[x].ParamsOrModels[y].Name + ";\n")
				if isError(err) {
					return "", err
				}
			}

		}

		_, err = file.WriteString("\nconst " + strings.Join(states, "") + "({\n")
		if isError(err) {
			return "", err
		}
		for xy := range project.State[x].ParamsOrModels {
			_, err = file.WriteString("required this." + project.State[x].ParamsOrModels[xy].Name + ",\n")
			if isError(err) {
				return "", err
			}

		}
		_, err = file.WriteString("});\n\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\t@override\n")
		if isError(err) {
			return "", err
		}
		_, err = file.WriteString("\tList<Object> get props => [")
		if isError(err) {
			return "", err
		}
		for xyz := range project.State[x].ParamsOrModels {
			_, err = file.WriteString("" + project.State[x].ParamsOrModels[xyz].Name + ",")
			if isError(err) {
				return "", err
			}
		}
		_, err = file.WriteString("];")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("\n\n")
		if isError(err) {
			return "", err
		}

		_, err = file.WriteString("}\n\n")
		if isError(err) {
			return "", err
		}
	}

	return "", nil
}
