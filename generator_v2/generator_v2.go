package generator_v2

import (
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
}
