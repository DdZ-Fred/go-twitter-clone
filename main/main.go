package main

import (
	"log"

	gotwitterclone "github.com/DdZ-Fred/go-twitter-clone"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	gotwitterclone.Run()
}
