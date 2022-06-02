package schemas

type StateMapping struct {
	StateName      string              `json:"state_name"`
	ParamsOrModels []ParamsOrModelsMap `json:"params_or_models"`
}

type ParamsOrModelsMap struct {
	Type    string `json:"type"`
	UseList bool   `json:"use_list"`
	Name    string `json:"name"`
}
