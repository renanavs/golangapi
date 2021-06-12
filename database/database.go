package database

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() error {
	dsn := "root@tcp(127.0.0.1:3306)/projectapi"
	database, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err)
		return err
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxLifetime(time.Minute * 30)
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)

	return nil

}

func GetDatabase() (*gorm.DB, error) {
	if db == nil {
		return db, errors.New("FAILED TO RECEIVE DATABASE CONNECTION")
	}
	return db, nil
}
