package initializer

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func  LoadEnvVariable() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("successfully Load env")
}