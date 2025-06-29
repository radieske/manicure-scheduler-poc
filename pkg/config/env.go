package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*manicure-scheduler-poc)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
