package schemas

type GeneratorOption struct {
	GenerateModel            bool `json:"generate_model,omitempty"`
	GenerateEntity           bool `json:"generate_entity,omitempty"`
	GenerateUseCase          bool `json:"generate_usecase,omitempty"`
	GenerateLocalDataSource  bool `json:"generate_local_data_source,omitempty"`
	GenerateRemoteDataSource bool `json:"generate_remote_data_source,omitempty"`
	GenerateRepository       bool `json:"generate_repository,omitempty"`
}
