package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	log.Println("::: LOADING ENV :::")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
