package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetConnectionString() string {
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	dbUser := GetEnv("DB_USER")
	dbPass := GetEnv("DB_PASS")
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, host, port)

}
