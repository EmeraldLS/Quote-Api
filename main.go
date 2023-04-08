package main

import (
	"fmt"
	"os"

	"github.com/EmeraldLS/quote-generator/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	err := router.Router(os.Getenv(":PORT"))
	if err != nil {
		fmt.Println(err)
	}
}
