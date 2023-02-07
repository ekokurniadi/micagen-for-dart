package schemas

type Project struct {
	OutputPath      string          `json:"output_path,omitempty"`
	Config          Config          `json:"config"`
	FeatureName     string          `json:"feature_name"`
	Entity          EntityConfig    `json:"entity"`
	State           []StateMapping  `json:"state,omitempty"`
	UseCase         []UseCase       `json:"usecases"`
	GeneratorOption GeneratorOption `json:"generator_options"`
}
