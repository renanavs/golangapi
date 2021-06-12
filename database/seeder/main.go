package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"projectapi/database"
	"projectapi/database/migration"
	"projectapi/model"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		println("Run this program from seeder/ directory")
	}
	file_byte, err := ioutil.ReadFile(pwd + "/MOCK_DATA.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var users []model.User

	json.Unmarshal(file_byte, &users)

	database.StartDB()
	db, err := database.GetDatabase()
	migration.RunMigrations(db)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.Create(&users)
}
