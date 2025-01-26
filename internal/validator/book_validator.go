package validator

import (
	"errors"

	"github.com/Gierdiaz/Book/internal/dto"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}

func ValidateBookDTO(bookDTO *dto.BookDTO) error {
	if err := ValidateStruct(bookDTO); err != nil {
		return err
	}

	if bookDTO.Quantity == 0 && bookDTO.Available {
		return errors.New("quantity cannot be zero if available is true")
	}

	return nil
}
