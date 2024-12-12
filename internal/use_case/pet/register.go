package pet

import (
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case/pet/create"
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case/pet/get"
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case/pet/search"
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case/pet/update"
	"github.com/samber/do/v2"
)

func Register(injector do.Injector) error {
	do.Provide(injector, create.New)
	do.Provide(injector, get.New)
	do.Provide(injector, search.New)
	do.Provide(injector, update.New)

	return nil
}
