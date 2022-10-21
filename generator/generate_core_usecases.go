package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ekokurniadi/micagen-for-dart/schemas"
)

func GenerateCoreUsecases(project schemas.Project) {
	createFileCoreUsecases(project)
	writeFileCoreUsecases(project)
}

func writeFileCoreUsecases(project schemas.Project) (string, error) {
	filepath, err := filepath.Abs("./core/usecases/usecases.dart")
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
	_, err = file.WriteString("import \"" + project.OutputPath + "/core/error/failures.dart\";\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("abstract class UseCase<Type, Params> {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\tFuture<Either<Failures,Type>> call(Params params);\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("class NoParams extends Equatable {\n")
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

func createFileCoreUsecases(project schemas.Project) (string, error) {
	filepath, err := filepath.Abs("./core/usecases/usecases.dart")

	if err != nil {
		log.Fatal("error")
		return filepath, err
	}
	_, err = os.Stat(filepath)
	if os.IsExist(err) {
		fmt.Println("your directory is already exist but it's ok")
		return "", err
	}

	filename, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err.Error())
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}
