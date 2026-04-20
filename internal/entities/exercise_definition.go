package entities

import (
	"errors"
	"fmt"

	"github.com/Grsaiago/goxercise/internal/platform"
	"github.com/go-playground/validator/v10"
)

var (
	ErrFailedToValidateExercise  = errors.New("failed to validate exercise definition")
	ErrInvalidExerciseDefinition = errors.New("Invalid exercise definition")
)

type ExerciseDefinition struct {
	Name         string  `toml:"name" validate:"required,filepath"`
	FilePath     string  `toml:"filepath" validate:"required,filepath"`
	TestPath     string  `toml:"testpath" validate:"required,filepath"`
	SolutionPath *string `toml:"solutionpath,omitempty" validate:"filepath"`
	Hint         *string `toml:"hint,omitempty" validate:"gt=0"`
}

type ExerciseDefinitionValidationError struct {
	fieldErrors []string
}

func NewExerciseDefinitionValidationError(validationErrors validator.ValidationErrors) *ExerciseDefinitionValidationError {
	var fieldErrors []string
	for _, err := range validationErrors {
		fieldErrors = append(fieldErrors, fmt.Sprintf("[%s]: [%s]", err.Field(), err.Error()))
	}
	return &ExerciseDefinitionValidationError{
		fieldErrors,
	}
}

func (err ExerciseDefinitionValidationError) Error() string {
	return fmt.Sprintf("%v", err.fieldErrors)
}

// Returns a type erased [ExerciseDefinitionValidationError]
func (ed ExerciseDefinition) Validate() error {
	err := platform.Validator.Struct(&ed)
	castedErr, isCastedErr := errors.AsType[*validator.ValidationErrors](err)
	if !isCastedErr {
		panic("ExerciseDefinition failed to validate with unknown error")
	}
	return NewExerciseDefinitionValidationError(*castedErr)
}
