package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	_, b, _, _ = runtime.Caller(0)
)

type Config struct {
	Environment    string `yaml:"Environment"`
	DataServerUrl  string `yaml:"DataServerUrl"`
	MtaServicePort int    `yaml:"MtaServicePort"`
}

func LoadConfig(env string) *Config {
	var config Config
	fileName := "/mta-" + env + ".yml"
	configPath := filepath.Join(b, "..")
	configFilepath := configPath + fileName
	yamlFile, err := ioutil.ReadFile(configFilepath)
	if err != nil || yamlFile == nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &config

}
