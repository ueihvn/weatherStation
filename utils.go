package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvByKey(key string) string {
	return os.Getenv(key)
}
