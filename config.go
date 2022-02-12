package main

import (
	"encoding/json"
	"errors"
	"os"
)

var ErrConfig = errors.New("wrong config file, should be like:" + configExample)

var configExample = `
[
    {
        "host": "host1.example.com",
        "ports": [ 22 ]
    },
    {
        "host": "host2.example.com",
        "ports": [ 22, 80, 443 ]
    }
]`

type Config []struct {
	Host  string `json:"host"`
	Ports []int  `json:"ports"`
}

type Ports map[string][]int

func readConfigFile(file string) (Ports, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf Config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return nil, ErrConfig
	}

	if err := validateConfig(conf); err != nil {
		return nil, err
	}

	ports := Ports{}
	for _, c := range conf {
		ports[c.Host] = c.Ports
	}

	return ports, nil
}

func validateConfig(conf Config) error {
	if len(conf) == 0 {
		return ErrConfig
	}
	for _, c := range conf {
		if c.Host == "" {
			return ErrConfig
		}
		if c.Ports == nil {
			return ErrConfig
		}
	}
	return nil
}
