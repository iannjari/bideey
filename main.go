package main

import (
	"bideey/api"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	server := api.NewServer(os.Getenv("SERVER_HOST"))
	server.Run()

}
