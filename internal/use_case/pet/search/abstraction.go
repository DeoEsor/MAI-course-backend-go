package search

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
)

type Command struct {
	Name   *string
	Offset int
	Limit  int
}

type CommandResponse struct {
	Pets []*pet.Pet
}

type CommandHandler interface {
	Search(ctx context.Context, command *Command) (*CommandResponse, error)
}
