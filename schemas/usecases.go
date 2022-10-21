package schemas

type UseCase struct {
	MethodName string `json:"usecase_name"`
	Parameter  string `json:"parameter"`
	ReturnType string `json:"return_type"`
}
