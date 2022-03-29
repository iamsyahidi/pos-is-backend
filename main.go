package main

import (
	"fmt"
	"pos-is-backend/api"
	"pos-is-backend/pkg/config"
	"pos-is-backend/pkg/database"
	"runtime"

	"github.com/spf13/viper"
)

func init() {
	config.GetConfig()
}

func main() {
	runtime.GOMAXPROCS(2)

	db, err := database.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := fmt.Sprintf(":%d", viper.GetInt("APP_PORT"))
	app := api.SetupRouter(db)
	go app.Run(port)
	select {}
}
