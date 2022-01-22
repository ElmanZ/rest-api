package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ElmanZ/restapi/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
	fmt.Println("Go")

	s := db.Service{}
	s.Init()
	s.Start(os.Getenv("HTTP_PORT"))
}
