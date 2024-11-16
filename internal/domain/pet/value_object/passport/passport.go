package passport

import (
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/gender"
	"github.com/google/uuid"
)

type Passport struct {
	Id       uuid.UUID     `json:"id"`
	Name     string        `json:"name"`
	LastName string        `json:"last_name"`
	BreedID  uuid.UUID     `json:"breed"`
	Gender   gender.Gender `json:"gender"`
}

func New(
	name, lastName string,
	breedID uuid.UUID,
	gender gender.Gender,
) (*Passport, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Passport{
		Id:       id,
		Name:     name,
		LastName: lastName,
		BreedID:  breedID,
		Gender:   gender,
	}, nil
}
