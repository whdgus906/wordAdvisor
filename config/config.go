package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DictAPIKey string
}

func (cfg *Config) ReadCfg() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(os.Getenv("GOPATH") + "/src/github.com/whdgus906/wordAdvisor/config/")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("Fatal error config file: %s", err)
	}

	return nil
}
