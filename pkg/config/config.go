package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config error : ", err.Error())
	}
}
