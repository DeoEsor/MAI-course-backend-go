package main

import (
	"context"
	"fmt"
	"github.com/DeoEsor/MAI-course-backend-go/internal/bootstrap"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
	repositoryPet "github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet"
)

func main() {
	ctx := context.Background()
	var cfg config.Config
	err := cfg.Load()
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

	db, closeConnection, err := bootstrap.ConfigureDb(ctx, &cfg.DatabaseConfig)
	defer closeConnection()
	if err != nil {
		panic(err)
	}

	repository, err := repositoryPet.New(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created repository: %T\n", repository)
}
