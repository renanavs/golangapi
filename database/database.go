package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	database, err := gorm.Open(mysql.Open(""))

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxLifetime(time.Minute * 30)
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)

}

//Singleton
func GetDatabase() *gorm.DB {
	if db == nil {
		StartDB()
	}
	return db
}
