// Package validator provides data validation utilities
package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validate is the instance of the validator
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
