package main

import (
	"log"
	"os"

	"github.com/ElmanZ/restapi/db"
	"github.com/joho/godotenv"
)

//entry point for the application
func main() {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}

	//Database initialization
	s := db.Service{}
	s.Init(
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_NAME"),
		os.Getenv("PG_SSL"),
	)

	//Start the app
	s.Start(os.Getenv("HTTP_PORT"))
}
