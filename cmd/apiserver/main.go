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

	passwordDB, exists := os.LookupEnv("DATABASE_PASSWORD")
	if exists {
		config.Store.DatabaseURL += fmt.Sprintf(" password=%s", passwordDB)
	}

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
