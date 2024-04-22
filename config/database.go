package config

import (
	"bideey/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConfigureDb(host string, port int, user string, password string, dbname string) (db *gorm.DB, err error) {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, fmt.Sprint(port), user, password,
		dbname)

	db, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	dbClient, _ := db.DB()
	err = dbClient.Ping()
	if err != nil {
		log.Fatal("error occured while acquiring database connection: ", err)
		return
	}
	fmt.Println("✅ Successfully configured DB.")

	Database = db
	db.AutoMigrate(&model.Bid{})
	db.AutoMigrate(&model.Biddable{})
	return
}
