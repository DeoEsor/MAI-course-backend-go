package create

import (
	"context"
	"github.com/google/uuid"
)

type Command struct {
	Name string
}

type CommandResponse struct {
	PetId uuid.UUID
}

type CommandHandler interface {
	Create(ctx context.Context, command *Command) (*CommandResponse, error)
}
