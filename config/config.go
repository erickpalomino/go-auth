package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// GlobalConfig
var GlobalConfig Config

type Config struct {
	Database struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBname   string `yaml:"dbname"`
	} `yaml:"database"`

	Security struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"security"`
}

func InitGlobalConfiguration(configPath string) error {
	conf, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer conf.Close()
	decoder := yaml.NewDecoder(conf)
	err = decoder.Decode(&GlobalConfig)
	return err
}
