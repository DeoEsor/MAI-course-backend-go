package pet

import (
	"context"
	"errors"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter"
	"github.com/google/uuid"
)

func (repository *repository) GetRequired(ctx context.Context, petId uuid.UUID) (*pet.Pet, error) {
	pet, err := repository.Get(ctx, petId)
	if err != nil {
		return nil, err
	}
	if pet == nil {
		return nil, errors.New("pet not found")
	}
	return pet, nil
}

func (repository *repository) Get(ctx context.Context, petId uuid.UUID) (*pet.Pet, error) {

	const query = `
		SELECT 
		    id,
		    status,
		    name,
		    passport,
		    created_at,
		    updated_at
		FROM pet
		WHERE id = :id
		LIMIT 1
	`

	rows, err := repository.db.NamedQueryContext(ctx, query, map[string]any{
		"id": petId,
	})
	if err != nil {
		return nil, err
	}

	result, err := converter.FromDbRows(rows)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}

	return result[0], nil
}
