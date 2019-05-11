package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenAddress      string  `envconfig:"LISTEN_ADDRESS" default:":8080"`
	JaegerAgentHost    string  `envconfig:"JAEGER_AGENT_HOST" default:"localhost"`
	JaegerAgentPort    string  `envconfig:"JAEGER_AGENT_PORT" default:"6831"`
	JaegerSamplerType  string  `envconfig:"JAEGER_SAMPLER_TYPE" default:"const"`
	JaegerSamplerParam float64 `envconfig:"JAEGER_SAMPLER_PARAM" default:"1"`
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
