package generator_v2

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ekokurniadi/micagen-for-dart/schemas"
)

func GenerateModule(project schemas.Project) {
	createFileModuleNetwork(project)
	createFileModuleStorage(project)
}

func createFileModuleNetwork(project schemas.Project) (string, error) {

	url := "https://raw.githubusercontent.com/ekokurniadi/micagen-for-dart/main/generator_v2/network_module.txt"

	filenames := "./core/module/network_module.dart"

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	file, err := os.Create(filenames)
	if err != nil {
		log.Fatal(err.Error())
		return file.Name(), err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create %s is successfully \n", file.Name())
	return file.Name(), nil
}
func createFileModuleStorage(project schemas.Project) (string, error) {

	url := "https://raw.githubusercontent.com/ekokurniadi/micagen-for-dart/main/generator_v2/storage_module.txt"

	filenames := "./core/module/storage_module.dart"

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	file, err := os.Create(filenames)
	if err != nil {
		log.Fatal(err.Error())
		return file.Name(), err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create %s is successfully \n", file.Name())
	return file.Name(), nil
}
