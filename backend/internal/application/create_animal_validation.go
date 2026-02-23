package application

import (
	"errors"
	"time"
)

const (
	CodeNameRequired     BusinessCode = "name_required"
	CodeSpeciesInvalid   BusinessCode = "species_invalid"
	CodeBirthdateInvalid BusinessCode = "birthdate_invalid"
)

func validateCreateAnimalInput(in CreateAnimalInput) error {
	if in.Name == "" {
		return BusinessError{Code: CodeNameRequired, Err: errors.New("name is required")}
	}
	switch in.Species {
	case "goat", "pig", "dog", "cat":
	default:
		return BusinessError{Code: CodeSpeciesInvalid, Err: errors.New("species is invalid")}
	}
	if in.Birthdate != "" {
		if _, err := time.Parse(time.DateOnly, in.Birthdate); err != nil {
			return BusinessError{Code: CodeBirthdateInvalid, Err: errors.New("birthdate is invalid")}
		}
	}
	return nil
}
