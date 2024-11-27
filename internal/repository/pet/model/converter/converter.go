package converter

import (
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model"
)

//go:generate goverter gen ./

// goverter:variables
// goverter:output:file ./converter.gen.go
// goverter:output:package github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter
// goverter:matchIgnoreCase
// goverter:ignoreMissing
// goverter:ignoreUnexported
// goverter:ignoreMissing
// goverter:useZeroValueOnPointerInconsistency
// goverter:skipCopySameType
// goverter:matchIgnoreCase
// goverter:extend github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter/common:.*
var (
	ToDB   func(request *pet.Pet) (*model.PetDB, error)
	FromDb func(request *model.PetDB) (*pet.Pet, error)
)
