package main

import (
	"fmt"
	"os"
	"pos-is-backend/api"
	"pos-is-backend/pkg/config"
	"pos-is-backend/pkg/database"
	"runtime"
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

	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))
	app := api.SetupRouter(db)
	go app.Run(port)
	select {}
}
