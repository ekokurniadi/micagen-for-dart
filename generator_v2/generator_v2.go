package generator_v2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekokurniadi/micagen-for-dart/generator"
	"github.com/ekokurniadi/micagen-for-dart/schemas"
)

type generatorv2 struct {
	project schemas.Project
}

func NewGeneratorHandlerV2(project schemas.Project) generatorv2 {
	return generatorv2{project}
}

func (h *generatorv2) GenerateFeature() {
	current, _ := os.Getwd()
	fmt.Println(current)

	fileName, _ := filepath.Abs("../pubspec.yaml")
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	line := 0
	packageName := ""
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line++
		if line == 1 {
			splitted := strings.Split(sc.Text(), ":")
			packageName = strings.TrimSpace(splitted[1])
			fmt.Println(packageName)
			break
		}
	}
	h.project.OutputPath = "package:" + packageName

	GenerateFeatureName(h.project.FeatureName)
	GenerateEntity(h.project)
	generator.GenerateCoreFailures(h.project)
	generator.GenerateCoreUsecases(h.project)

	if h.project.Config.UseFreezed {
		GenerateBuildYaml()
	}
	GenerateModels(h.project)
	GenerateRepository(h.project)
	GenerateRepositoryImpl(h.project)
	GenerateLocalDataSource(h.project)
	GenerateLocalDataSourceImpl(h.project)
	GenerateRemoteDataSource(h.project)
	GenerateRemoteDataSourceImpl(h.project)
	GenerateUseCase(h.project)
	GenerateDioHelper()

}
