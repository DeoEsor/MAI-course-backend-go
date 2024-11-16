package breed

import "github.com/google/uuid"

type Breed struct {
	Id          uuid.UUID
	Name        string
	ParentId    *uuid.UUID
	ChildBreeds []uuid.UUID
	Country     string
	Description string
}

func New(name string, parentId *uuid.UUID) (*Breed, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Breed{
		Id:       id,
		Name:     name,
		ParentId: parentId,
	}, nil
}

func (b *Breed) UpdateName(newName string) {
	b.Name = newName
}

func (b *Breed) Update(
	country, description string,
	childBreeds []uuid.UUID,
) {
	b.ChildBreeds = childBreeds
	b.Description = description
	b.Country = country
}
