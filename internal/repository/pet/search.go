package pet

import (
	"context"
	pet2 "github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter"
)

func (repository *repository) Search(ctx context.Context, searchQuery *RepositorySearchQuery) ([]*pet2.Pet, error) {

	const query = `
		SELECT 
		    id,
		    status,
		    name,
		    passport,
		    created_at,
		    updated_at
		FROM pet
		WHERE 
		    (CAST(:name as varchar(50)) is null or name like :name)
		ORDER BY created_at DESC 	
		LIMIT :limit
		OFFSET :offset
	`

	rows, err := repository.db.NamedQueryContext(ctx, query, searchQuery)
	if err != nil {
		return nil, err
	}

	result, err := converter.FromDbRows(rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}
