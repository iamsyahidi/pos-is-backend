package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("config error : ", err.Error())
	}
}
