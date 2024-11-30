package pet

import (
	"context"
	"errors"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter"
	"time"
)

func (repository *repository) Update(ctx context.Context, pet *pet.Pet, version time.Time) error {
	const query = `
		update pet
		set
		    status = :status,
		    name = :name,
		    passport = :passport,
		    updated_at = :updated_at
		where id = :id  			
		  and updated_at = :version 
	`

	petDb, err := converter.ToDB(pet)
	if err != nil {
		return err
	}
	petDb.Version = version

	rows, err := repository.db.NamedExecContext(ctx, query, petDb)
	if err != nil {
		return err
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("optimistic error: сущность менялось одновременно с другой операцией")
	}

	return nil
}
