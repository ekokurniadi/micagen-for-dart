package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFeatureName(featureName string) {
	writeFolder(featureName)
}

func writeFolder(featureName string) error {
	var realFeatureName []string
	splitText := strings.Split(featureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")

	path := filepath.Join("./", realName)

	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println("your directory is already exist but it's ok")
	}

	var childCoreFolder = []string{"error", "usecases"}
	var baseFolder = []string{"data", "domain", "presentation"}
	var dataChildFolder = []string{"datasource", "models", "repository"}
	var domainChildFolder = []string{"entity", "repository", "usecases"}
	var presentationChildFolder = []string{"bloc", "pages", "widgets"}

	core := filepath.Join("./", "core")
	os.Mkdir(core, 0755)

	for x := 0; x < len(childCoreFolder); x++ {
		pathCore := filepath.Join("./core/", childCoreFolder[x])
		os.Mkdir(pathCore, 0755)
	}

	for a := 0; a < len(baseFolder); a++ {
		basePath := filepath.Join(path, baseFolder[a])
		os.Mkdir(basePath, 0755)
		for b := 0; b < len(dataChildFolder); b++ {
			dataChildFolderPath := filepath.Join(path+"/"+baseFolder[0], dataChildFolder[b])
			os.Mkdir(dataChildFolderPath, 0755)

		}
		for c := 0; c < len(domainChildFolder); c++ {
			domainChildFolderPath := filepath.Join(path+"/"+baseFolder[1], domainChildFolder[c])
			os.Mkdir(domainChildFolderPath, 0755)

		}
		for d := 0; d < len(presentationChildFolder); d++ {
			presentationChildFolderPath := filepath.Join(path+"/"+baseFolder[2], presentationChildFolder[d])
			os.Mkdir(presentationChildFolderPath, 0755)

		}
	}

	return nil
}
