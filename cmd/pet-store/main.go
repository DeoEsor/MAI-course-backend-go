package main

import (
	"fmt"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
)

func main() {
	var config config.Config
	err := config.Load()
	if err != nil {
		fmt.Printf("error while parsing config: %v", err)
		return
	}

	pet := pet.New("Sugrob")
	fmt.Printf("Created pet with data: %+v\n", pet)

	err = pet.UpdateStatus(status.NONE)
	if err != nil {
		fmt.Printf("error while updating status: %v\n", err)
	}

	err = pet.UpdateStatus(status.Adopted)
	if err != nil {
		fmt.Printf("error while updating status: %v\n", err)
	}

	err = pet.UpdateStatus(status.OnRecover)
	if err != nil {
		panic(err)
	}

	fmt.Println(pet)
}
