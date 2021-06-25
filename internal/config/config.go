package config

import (
	"encoding/json"
	"io/ioutil"
)

type ShellCommand struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type Handler struct {
	Uri      string         `json:"uri"`
	Commands []ShellCommand `json:"commands"`
}

type Config struct {
	AppName  string    `json:"app-name"`
	Host     string    `json:"host"`
	Port     int       `json:"port"`
	Secret   string    `json:"secret"`
	Handlers []Handler `json:"handlers"`
}

func FromFile(path string) (c *Config, err error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c = &Config{}
	err = json.Unmarshal(f, c)
	if err != nil {
		return nil, err
	}
	return c, err
}

func ToFile(path string, c *Config) error {
	b, err := json.MarshalIndent(&c, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0o666)
	if err != nil {
		return err
	}
	return nil
}
