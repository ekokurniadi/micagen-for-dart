package schemas

type Config struct {
	NullSafety             bool `json:"null_safety"`
	UseDartzPackage        bool `json:"use_dartz_package"`
	UseDependencyInjection bool `json:"use_dependency_injection"`
	EnableUnitTesting      bool `json:"enable_unit_testing"`
}
