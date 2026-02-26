package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort string
	DBhost string
	DBport string
	DBuser string
	DBpassword string
	DBname string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	cfg := &Config{
		HTTPPort: os.Getenv("HTTP_PORT"),
		DBhost : os.Getenv("HOST"),
		DBport : os.Getenv("PORT"),
		DBuser : os.Getenv("USER"),
		DBpassword : os.Getenv("PASSWORD"),
		DBname : os.Getenv("DBNAME"),
	}

	return cfg, nil
}



    
	