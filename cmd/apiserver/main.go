package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"http-rest-api/internal/app/apiserver"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	passwordDB, exists := os.LookupEnv("DATABASE_PASSWORD")
	if exists {
		config.DatabaseURL += fmt.Sprintf(" password=%s", passwordDB)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
