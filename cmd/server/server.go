package main

import (
	"flag"
	"fmt"
	"go-project-base/internal/config"
)

func main() {
	configPath := flag.String("config", "./config.json", "path to the config file")
	if !flag.Parsed() {
		flag.Parse()
	}

	config, err := config.FromFile(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read config: %v", config)
	fmt.Println("ready to go!")
}
