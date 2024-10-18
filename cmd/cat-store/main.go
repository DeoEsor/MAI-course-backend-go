package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	language := "go"

	fmt.Printf("Hello %v", language)

}
