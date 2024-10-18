package main

<<<<<<< HEAD
import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	language := "go"

	fmt.Printf("Hello %v", language)

=======
import "fmt"

func main() {
	language := "go"
	fmt.Printf("Hello %v", language)
>>>>>>> 756f684 (lesson1: added simplest program)
}
