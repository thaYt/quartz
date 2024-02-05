package main

import (
	"os"
	"quartz/api"

	"github.com/go-yaml/yaml"
)

type YamlConfig struct {
	Method api.Method `yaml:"method"`
	LogDir string     `yaml:"logdir"`
}

var Config YamlConfig

func WriteConfig() {
	a, err := yaml.Marshal(Config)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(a)
}

func ReadConfig() {
	f, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &Config)
	if err != nil {
		panic(err)
	}
}
