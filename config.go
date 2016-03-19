package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Host     string
	Port     int
	Path     string
	Protocol string
}

func configExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		// Expecting NotExist error, if error is something else let the user know
		if !os.IsNotExist(err) {
			fmt.Println(err)
		}

		return false
	}

	return true
}

func defaultConfig() Config {
	return Config{
		Host:     "",
		Port:     -1,
		Path:     "",
		Protocol: "http",
	}
}

func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	return &config, nil
}

func (c Config) saveConfig(filename string) error {
	data, err := json.Marshal(&c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0)
	if err != nil {
		return err
	}
	return nil
}
