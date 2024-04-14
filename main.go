package main

import (
	"bideey/api"
	"bideey/config"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		panic(err)
	}

	db, err := config.ConfigureDb(os.Getenv("DB_HOST"), port,
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	if db != nil {
		server := api.NewServer(8000)
		server.Run()

	} else {
		panic(err)
	}

}
