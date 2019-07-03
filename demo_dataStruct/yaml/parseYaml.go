package main

import (
	"fmt"
	"io/ioutil"
	
	"gopkg.in/yaml.v2"
)


type Info struct {
	Config struct {
		Host	string `yaml:"host"`
		User	string	`yaml:"user"`
		Port	int		`yaml:"port"`
		Password	string	`yaml:"password"`
	}
}

func main() {
	yamlFile, _ := ioutil.ReadFile("config.yaml")
	i := Info{}
	err := yaml.Unmarshal(yamlFile, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i.Config.Host)
}