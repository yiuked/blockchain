package main

import (
	"block-chain/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go cmd.RunMiner()
	cmd.RunWeb()
}
