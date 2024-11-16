package create

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
)

func (uc *UseCase) Handle(ctx context.Context, command *Command) (*CommandResponse, error) {
	// TODO проверка что в БД нет такого питомца
	newPet := pet.New(command.Name)
	// TODO более сложные действия
	// TODO сохранение в БД

	return &CommandResponse{
		Pet: newPet,
	}, nil
}
