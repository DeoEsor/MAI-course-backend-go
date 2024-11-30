package main

import (
	"context"
	"fmt"
	"github.com/DeoEsor/MAI-course-backend-go/internal/bootstrap"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
	repositoryPet "github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet"
	"github.com/brianvoe/gofakeit/v7"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	var cfg config.Config
	err := cfg.Load()
	if err != nil {
		fmt.Printf("error while parsing config: %v", err)
		return
	}

	db, closeConnection, err := bootstrap.ConfigureDb(ctx, &cfg.DatabaseConfig)
	if err != nil {
		panic(err)
	}
	defer closeConnection()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	fmt.Println("ping database success")

	repository, err := repositoryPet.New(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created repository: %T\n", repository)

	pets, err := repository.Search(ctx, &repositoryPet.RepositorySearchQuery{
		Offset: 0,
		Limit:  1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found pets: %+v\n", pets)

	var wg sync.WaitGroup

	for i := range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pet := pets[0]
			version := pet.UpdatedAt

			pet.UpdateName(gofakeit.Name())
			sleepTime := time.Duration(i) * time.Second
			time.Sleep(sleepTime) // long operation

			err := repository.Update(ctx, pet, version)
			if err != nil {
				fmt.Printf("error while updating pet: %v %v", pet.ID, err)
			}
			fmt.Printf("Updated pet: %+v\n", pet)
		}()
	}
	wg.Wait()
}
