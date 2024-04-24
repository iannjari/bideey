package config

import (
	"bideey/model"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQL interface {
	GetDB(bid *model.Bid) *gorm.DB
}

type PostgresDB struct {
	Database *gorm.DB
}

func NewPostgresDB() *PostgresDB {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	err = ConfigureDb(os.Getenv("DB_HOST"), port,
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	if err != nil {
		panic(err)
	}
	return &PostgresDB{database}
}

var database *gorm.DB

func ConfigureDb(host string, port int, user string, password string, dbname string) (err error) {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, fmt.Sprint(port), user, password,
		dbname)

	db, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
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

	err = db.AutoMigrate(&model.Bid{})
	if err != nil {
		panic("failed to migrate model.Bid after connecting to database" + err.Error())
	}
	err = db.AutoMigrate(&model.Biddable{})
	if err != nil {
		panic("failed to migrate model.Biddable after connecting to database" + err.Error())
	}
	database = db
	return
}

func (*PostgresDB) GetDB() *gorm.DB {
	return database
}
