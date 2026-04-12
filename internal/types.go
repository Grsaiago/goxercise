package types

import (
	"errors"
	"io"

	"github.com/pelletier/go-toml/v2"
)

type ExerciseDefinition struct {
	Name         string  `toml:"name" validate:"required,filepath"`
	FilePath     string  `toml:"filepath" validate:"required,filepath"`
	TestPath     string  `toml:"testpath" validate:"required,filepath"`
	SolutionPath *string `toml:"solutionpath,omitempty" validate:"filepath"`
	Hint         *string `toml:"hint,omitempty" validate:"gt=0"`
}

var (
	ErrInvalidExerciseDefinition = errors.New("Invalid exercise definition")
	ErrInvalidExerciseFile       = errors.New("Invalid exercises file")
	ErrFailedToOpenExerciseFile  = errors.New("failed to open the exercieses file")
)

type ExerciseList struct {
	exercises []ExerciseDefinition `toml:"exercises"`
}

func LoadExerciseListFromReader(reader io.Reader) (*ExerciseList, error) {
	var exerciseList ExerciseList

	toml.NewDecoder(reader).Decode(exerciseList)

	return &exerciseList, nil
}
