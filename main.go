package main

import (
	"crud-person/internal/config/database"
	"crud-person/internal/router"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	database.InitMysql()
	router.InitRouter().Run(":" + viper.GetString("port"))
}
