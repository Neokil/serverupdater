package main

import (
	"flag"
	"serverupdater/internal/config"
)

func main() {
	configPath := flag.String("config", "./config.json", "path to the config file")
	if !flag.Parsed() {
		flag.Parse()
	}

	err := config.ToFile(*configPath, &config.Config{
		AppName: "Go-Server-Project-Base",
		Host:    "0.0.0.0",
		Port:    80,
	})
	if err != nil {
		panic(err.Error())
	}
}
