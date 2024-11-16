package create

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
)

type Command struct {
	Name string
	//registrationPlace string, uuid uuid.UUID
}

// CQRS
// Command Query Responsebility Separation
type CommandResponse struct {
	Pet *pet.Pet
}

type Handler interface {
	Handle(ctx context.Context, command *Command) (*CommandResponse, error)
}
