package main

import (
	"io/ioutil"

	yml "gopkg.in/yaml.v2"
)

var (
	//Config - yaml config reader object
	Config Yaml
	//Token - Dropbox Token
	Token string
	//Getlink - Dropbox HTTP API URL for getting file temporary link
	Getlink string
	//Listlink - Dropbox HTTP API URL for listing metadata for all files in dropbox folder
	Listlink string
)

//Yaml struct used to parse the conf.yml file.
type Yaml struct {
	Token    string `yaml:"token"`
	Key      string `yaml:"app_key"`
	Secret   string `yaml:"app_secret"`
	Getlink  string `yaml:"get_link_api_url"`
	Listlink string `yaml:"list_link_api_url"`
}

// Yaml method for reading and Unmarshaling bytes
// of the conf.yml into the Yaml struct
func (c *Yaml) getConf() *Yaml {
	yamlFile, err := ioutil.ReadFile("conf.yml")
	check(err)

	err = yml.Unmarshal(yamlFile, c)
	check(err)

	return c
}

// Getconfigs parses the conf.yml file and retrieves all necessary
// configuration parameters
func Getconfigs() {

	// Call getConf method of the Yaml struct "Config"
	Config.getConf()

	// Assign config variables
	Token = Config.Token
	Getlink = Config.Getlink
	Listlink = Config.Listlink
}
