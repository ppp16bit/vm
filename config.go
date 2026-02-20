package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	APIPort string
}

func getEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func LoadConf() Config {
	_ = godotenv.Load()
	return Config{
		DBHost:  getEnv("DB_HOST"),
		DBPort:  getEnv("DB_PORT"),
		DBUser:  getEnv("DB_USER"),
		DBPass:  getEnv("DB_PASS"),
		DBName:  getEnv("DB_NAME"),
		APIPort: getEnv("API_PORT"),
	}
}
