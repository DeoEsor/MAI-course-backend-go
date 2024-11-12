package main

import (
	"fmt"

	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
)

func main() {
	var config config.Config
	err := config.Load()
	if err != nil {
		fmt.Printf("error while parsing config: %v", err)
		return
	}

}
