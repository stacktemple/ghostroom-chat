package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

var (
	Cfg  Config
	once sync.Once
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}
}

func loadConfig() {
	loadEnv()

	Cfg = Config{
		Port:        getEnv("PORT", "3000"),
		DatabaseURL: mustGet("DATABASE_URL"),
		JWTSecret:   mustGet("JWT_SECRET"),
	}
}

func Init() {
	once.Do(loadConfig)
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func mustGet(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Fatalf("Environment variable %s not set", key)
	return ""
}
