package pet

import (
	"context"

	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter"
)

func (repository *repository) Create(ctx context.Context, pet *pet.Pet) error {
	const query = `
		INSERT INTO pet
		(	
		 	id,
		    status,
		    name,
		    passport,
		    created_at,
		    updated_at
		)
		VALUES 
		(
		 	:id,
			:status,
			:name,
			:passport,
			:created_at,
			:updated_at
		)
	`

	petDb, err := converter.ToDB(pet)
	if err != nil {
		return err
	}

	_, err = repository.db.NamedExecContext(ctx, query, petDb)
	if err != nil {
		return err
	}

	return nil
}
