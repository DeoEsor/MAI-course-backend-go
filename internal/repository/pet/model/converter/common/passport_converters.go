package common

import (
	"encoding/json"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/passport"
)

func PassportToJson(entity *passport.Passport) (string, error) {
	jsonBytes, err := json.Marshal(entity)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func JsonToPassport(jsonValue string) (*passport.Passport, error) {
	var entity passport.Passport
	err := json.Unmarshal([]byte(jsonValue), &entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
