package main

import (
	"encoding/json"
	"errors"
	"os"
)

var errConfig = errors.New("wrong config file, should be like:" + configExample)

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

type config []struct {
	Hostname string   `json:"hostname"`
	Ports    []string `json:"ports"`
}

func readConfigFile(file string) (map[string][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return nil, errConfig
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

func validateConfig(conf config) error {
	if len(conf) == 0 {
		return errConfig
	}
	for _, c := range conf {
		if c.Hostname == "" {
			return errConfig
		}
		if c.Ports == nil {
			return errConfig
		}
	}
	return nil
}
