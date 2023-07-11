package schemas

type GeneratorOption struct {
	GenerateModel            bool `json:"generate_model,omitempty"`
	GenerateEntity           bool `json:"generate_entity,omitempty"`
	GenerateUseCase          bool `json:"generate_usecase,omitempty"`
	GenerateLocalDataSource  bool `json:"generate_local_data_source,omitempty"`
	GenerateRemoteDataSource bool `json:"generate_remote_data_source,omitempty"`
	GenerateRepository       bool `json:"generate_repository,omitempty"`
	GenerateModule           bool `json:"generate_module,omitempty"`
	GenerateBloc             bool `json:"generate_bloc,omitempty"`
	GenerateAll              bool `json:"generate_all,omitempty"`
}
