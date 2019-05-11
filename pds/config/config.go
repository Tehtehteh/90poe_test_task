package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenAddress string `envconfig:"LISTEN_ADDRESS" default:":8081"`
}

func CreateConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env, assuming we are in production mode!")
	}
	var env Config
	err = envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}
