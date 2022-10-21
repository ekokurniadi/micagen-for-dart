package generator_v2

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func GenerateBuildYaml() {
	createFileBuildYaml()
	// writeFileBuildYaml()
}

func createFileBuildYaml() (string, error) {
	copiedFilepath, _ := filepath.Abs("./generator_v2/build.yaml")

	filepath, err := filepath.Abs("./././build.yaml")

	if err != nil {
		log.Fatal("error")
		return filepath, err
	}
	_, err = os.Stat(filepath)
	if os.IsExist(err) {
		fmt.Println("your directory is already exist but it's ok")
		return "", err
	}

	sFile, err := os.Open(copiedFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	filename, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err.Error())
		return filename.Name(), err
	}

	_, err = io.Copy(filename, sFile) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = filename.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}
