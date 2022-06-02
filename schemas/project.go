package schemas

type Project struct {
	OutputPath  string         `json:"output_path"`
	Config      Config         `json:"config"`
	FeatureName string         `json:"feature_name"`
	Entity      EntityConfig   `json:"entity"`
	BlocOptions BlocOptions    `json:"state_management_options"`
	State       []StateMapping `json:"state"`
	VoidMapping []VoidMapping  `json:"methods"`
}
