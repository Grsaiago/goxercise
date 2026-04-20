package entities

import (
	"errors"
	"io"

	"github.com/Grsaiago/goxercise/internal/platform"
	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
)

var (
	ErrFailedToParseToml = errors.New("failed to parse exercises")
)

type ExerciseList struct {
	Exercises []ExerciseDefinition `toml:"exercises"`
}

func NewExerciseListFromReader(reader io.Reader) (*ExerciseList, error) {
	var exerciseList ExerciseList

	if err := toml.NewDecoder(reader).DisallowUnknownFields().Decode(&exerciseList); err != nil {
		return nil, errors.Join(err, ErrFailedToParseToml)
	}

	if err := platform.Validator.Struct(&exerciseList); err != nil {
		castedErr, isCastedErr := errors.AsType[*validator.ValidationErrors](err)
		if isCastedErr {
			return nil, errors.Join(NewExerciseDefinitionValidationError(*castedErr))
		}
		return nil, errors.Join(err, ErrFailedToValidateExercise)
	}
	return &exerciseList, nil
}
