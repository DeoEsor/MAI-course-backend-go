package main

import (
	"context"
	"fmt"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
	"github.com/DeoEsor/MAI-course-backend-go/internal/use-case/pet/create"
)

func main() {
	ctx := context.Background()
	var config config.Config
	err := config.Load()
	if err != nil {
		fmt.Printf("error while parsing config: %v", err)
		return
	}

	handler, err := create.New()
	if err != nil {
		fmt.Printf("error while creating handler: %v", err)
		return
	}
	response, err := handler.Handle(ctx, &create.Command{
		Name: "SpecificName",
	})
	if err != nil {
		fmt.Printf("error while creating handler: %v", err)
		return
	}

	fmt.Println(response.Pet)
}
