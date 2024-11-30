// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package converter

import (
	pet "github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	status "github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
	model "github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model"
	common "github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet/model/converter/common"
	"time"
)

func init() {
	FromDb = func(source *model.PetDB) (*pet.Pet, error) {
		var pPetPet *pet.Pet
		if source != nil {
			var petPet pet.Pet
			petPet.ID = (*source).ID
			petPet.Status = status.Status((*source).Status)
			petPet.Name = (*source).Name
			pPassportPassport, err := common.JsonToPassport((*source).Passport)
			if err != nil {
				return nil, err
			}
			petPet.Passport = pPassportPassport
			petPet.CreatedAt = timeTimeToTimeTime((*source).CreatedAt)
			petPet.UpdatedAt = timeTimeToTimeTime((*source).UpdatedAt)
			pPetPet = &petPet
		}
		return pPetPet, nil
	}
	ToDB = func(source *pet.Pet) (*model.PetDB, error) {
		var pModelPetDB *model.PetDB
		if source != nil {
			var modelPetDB model.PetDB
			modelPetDB.ID = (*source).ID
			modelPetDB.Status = string((*source).Status)
			modelPetDB.Name = (*source).Name
			xstring, err := common.PassportToJson((*source).Passport)
			if err != nil {
				return nil, err
			}
			modelPetDB.Passport = xstring
			modelPetDB.CreatedAt = timeTimeToTimeTime((*source).CreatedAt)
			modelPetDB.UpdatedAt = timeTimeToTimeTime((*source).UpdatedAt)
			pModelPetDB = &modelPetDB
		}
		return pModelPetDB, nil
	}
}
func timeTimeToTimeTime(source time.Time) time.Time {
	return source
}
