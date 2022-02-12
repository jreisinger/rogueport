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
        "hostname": "host1.example.com",
        "ports": [ "22/tcp" ]
    },
    {
        "hostname": "host2.example.com",
        "ports": [ "22/tcp", "80/tcp", "443/tcp" ]
    }
]`

type Config []struct {
	Hostname string   `json:"hostname"`
	Ports    []string `json:"ports"`
}

func readConfigFile(file string) (map[string][]string, error) {
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

	ports := make(map[string][]string)

	for _, c := range conf {
		ports[c.Hostname] = c.Ports
	}

	return ports, nil
}

func validateConfig(conf Config) error {
	if len(conf) == 0 {
		return ErrConfig
	}
	for _, c := range conf {
		if c.Hostname == "" {
			return ErrConfig
		}
		if c.Ports == nil {
			return ErrConfig
		}
	}
	return nil
}
