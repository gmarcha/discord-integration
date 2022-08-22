package env

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
