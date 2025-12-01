package domain

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

const YAML_NOT_FOUND_ERROR = "ERROR: config.yaml not readed"
const UNMARSHAL_YAML_ERROR = "ERROR: config.yaml unmarshal problem"

type SecurityConfig struct {
    JWTKey           string `yaml:"jwt-key"`
    TokenLiveInHours int    `yaml:"token-live-hours"`
}

type AppConfiguration struct {
    Security    SecurityConfig `yaml:"security"`
    LoggerDebug bool           `yaml:"logger-debug"`
}

var (
	configInstance *AppConfiguration
	once           sync.Once
)

func GetConfig() *AppConfiguration {
	once.Do(func() {
		f, err := os.ReadFile("config/config.yaml")
		if err != nil {
			panic(YAML_NOT_FOUND_ERROR)
		}

		var cfg AppConfiguration
		if err := yaml.Unmarshal(f, &cfg); err != nil {
			panic(UNMARSHAL_YAML_ERROR)
		}

		configInstance = &cfg
	})
	return configInstance
}
