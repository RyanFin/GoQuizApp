package main

import (
	"RyanFin/GoQuizApp/pkg/constants"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	File string `yaml:"file"`
}

func main() {
	// load config file from scratch
	filename, err := filepath.Abs(constants.FileName)
	if err != nil {
		panic(err)
	}

	// load file
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config.File)
	if err != nil {
		panic(err)
	}

	fmt.Println(config.File)
}
