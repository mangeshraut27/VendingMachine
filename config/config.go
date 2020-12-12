package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config is a package variable, which is populated during init() execution and shared to whole application
var Config Configuration

//Configuration : application level config
type Configuration struct {
	Products []ProductList `json:"products"`
}

//ProductList : Application products
type ProductList struct {
	Name            string `json:"name"`
	Price           int    `json:"price"`
	DefaultQuantity int    `json:"defaultQuantity"`
}

var (
	ioutilReadFile = ioutil.ReadFile
)

//LoadConfig load configs from config.json
var LoadConfig = func(configFilePath string) error {
	log.Printf("Looking for JSON config file (%s)", configFilePath)

	contents, err := ioutilReadFile(configFilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
	} else {
		log.Printf("Configuration has been read from JSON (%s) successfully\n", configFilePath)
	}

	return err
}
