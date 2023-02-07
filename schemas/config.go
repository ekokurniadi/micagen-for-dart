package schemas

type Config struct {
	UseFreezed    bool `json:"use_freezed,omitempty"`
	UseInjectable bool `json:"use_injectable,omitempty"`
}
