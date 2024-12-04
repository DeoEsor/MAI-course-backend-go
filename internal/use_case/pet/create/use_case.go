package create

import (
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet"
	"github.com/samber/do/v2"
)

type useCase struct {
	petRepository pet.RepositoryI
}

func New(injector do.Injector) (CommandHandler, error) {
	return &useCase{
		petRepository: do.MustInvoke[pet.RepositoryI](injector),
	}, nil
}
