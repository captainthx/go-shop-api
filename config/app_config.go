package config

import (
	"os"
)

var (
	Mode          string
	ServerPort    string
	UploadPath    string
	ImageBaseUrl  string
	ImageBasePath string
	DbHost        string
	DbPort        string
	DbName        string
	DbUsername    string
	DbPassword    string
)

func Init() {
	Mode = getEnv("GIN_MODE")
	ServerPort = getEnv("SERVER_PORT")
	UploadPath = getEnv("UPLOAD_PATH")
	ImageBaseUrl = getEnv("BASE_URL")
	ImageBasePath = getEnv("IMAGE_BASE_PATH")
	DbHost = getEnv("DB_HOST")
	DbPort = getEnv("DB_PORT")
	DbName = getEnv("DB_NAME")
	DbUsername = getEnv("DB_USER")
	DbPassword = getEnv("DB_PASSWORD")
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}
