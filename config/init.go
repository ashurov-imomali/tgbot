package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Configs struct {
	Token   string `yaml:"token"`
	Address string `yaml:"address"`
}

func GetConfigs() (conf Configs, err error) {
	file, err := os.ReadFile("./config/configs.yaml")
	if err != nil {
		return Configs{}, err
	}
	err = yaml.Unmarshal(file, &conf)
	return
}
