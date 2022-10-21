package schemas

type StateMapping struct {
	StateName string   `json:"state_name"`
	Params    []Params `json:"params"`
}

type Params struct {
	Type    string `json:"data_type"`
	UseList bool   `json:"declare_as_list"`
	Name    string `json:"name"`
}
