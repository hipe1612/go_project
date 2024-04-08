package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func InitMysql() {
	var connectionString = viper.GetString("database.mysql.connection-string")
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
