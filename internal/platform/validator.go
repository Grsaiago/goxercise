package platform

import "github.com/go-playground/validator/v10"

var (
	Validator *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
)
