package generator_v2

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GenerateBuildYaml() {
	createFileBuildYaml()
	// writeFileBuildYaml()
}

func createFileBuildYaml() (string, error) {

	filenames := "../build.yaml"
	url := "https://raw.githubusercontent.com/ekokurniadi/micagen-for-dart/main/generator_v2/build.yaml"

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
