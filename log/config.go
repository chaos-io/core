package log

type Config struct {
	Level        string                 `json:"level" default:"debug"`
	Encode       string                 `json:"encode" default:"console"`
	LevelPort    int                    `json:"levelPort" default:"0"`
	LevelPattern string                 `json:"levelPattern" default:""`
	Output       string                 `json:"output" default:"console"`
	InitFields   map[string]interface{} `json:"initFields"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Level:  "debug",
		Encode: "console",
		Output: "console",
	}
}
