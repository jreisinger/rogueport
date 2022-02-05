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

type config []struct {
	Host  string `json:"host"`
	Ports []int  `json:"ports"`
}

func readConfigFile(file string) ([]*host, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return nil, ErrConfig
	}

	if err := validateConfig(conf); err != nil {
		return nil, err
	}

	var hosts []*host
	for _, c := range conf {
		hosts = append(hosts, &host{name: c.Host, portsExpected: c.Ports})
	}

	return hosts, nil
}

func validateConfig(conf config) error {
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
