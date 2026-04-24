package entities

import (
	"errors"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/Grsaiago/goxercise/internal/platform"
	"github.com/go-playground/validator/v10"
)

var (
	ErrFailedToParseToml = errors.New("failed to parse exercises config toml")
)

type ExerciseList struct {
	Exercises []ExerciseDefinition `toml:"exercises" validate:"dive"`
}

// Returns either a sentinel error [ErrFailedToParseToml] or
// a complex [ExerciseDefinitionValidationError], which has an array of all invalid keys
func NewExerciseListFromReader(reader io.Reader) (*ExerciseList, error) {
	var exerciseList ExerciseList

	if _, err := toml.NewDecoder(reader).Decode(&exerciseList); err != nil {
		return nil, errors.Join(err, ErrFailedToParseToml)
	}

	if err := platform.Validator.Struct(&exerciseList); err != nil {
		castedErr, isCastedErr := errors.AsType[validator.ValidationErrors](err)
		switch isCastedErr {
		case true:
			return nil, NewExerciseDefinitionValidationError(castedErr)
		case false:
			panic("unknown error occured when validating the exercise list config")
		}
	}
	return &exerciseList, nil
}
