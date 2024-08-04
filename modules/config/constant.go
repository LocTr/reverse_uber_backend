package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Cluster  string
	AppName  string
	User     string
	Password string
	// Add more fields for other .env values
}

func NewEnvConfig() (*EnvConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	cluster := os.Getenv("CLUSTER_NAME")
	appName := os.Getenv("APP_NAME")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	// Add more fields for other .env values

	return &EnvConfig{
		Cluster:  cluster,
		AppName:  appName,
		User:     user,
		Password: password,
	}, nil
}
