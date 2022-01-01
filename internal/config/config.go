package config

import (
	"io/ioutil"
	"log"

	"github.com/kemal576/go-pw-manager/models"
	"gopkg.in/yaml.v2"
)

func ReadConfiguration(configName string) models.Config {
	var config models.Config

	yamlFile, err := ioutil.ReadFile(configName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
