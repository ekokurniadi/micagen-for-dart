package generator

import "github.com/ekokurniadi/micagen-for-dart/schemas"

type generator struct {
	project schemas.Project
}

func NewGeneratorHandler(project schemas.Project) generator {
	return generator{project}
}

func (h *generator) GenerateFeature() {
	GenerateFeatureName(h.project.FeatureName)
	GenerateEntity(h.project)
	GenerateModels(h.project)
	GenerateDataSource(h.project)
	GenerateDataSourceImpl(h.project)
	GenerateRepository(h.project)
	GenerateRepositoryImpl(h.project)
	GenerateCoreUsecases(h.project)
	GenerateCoreFailures(h.project)
	GenerateUseCase(h.project)
	GenerateCubit(h.project)
}
