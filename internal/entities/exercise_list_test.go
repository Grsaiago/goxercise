package entities

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshallExerciseList(t *testing.T) {
	testCases := []struct {
		RawToml        string
		ExpectedErr    error
		ExpectedParsed []ExerciseDefinition
		MustErr        bool
	}{
		{
			RawToml: `
					[[exercises]]
					name = "intro1"
					filepath = "filepath"
					testpath = "filepath"
					solutionpath = "optional"
					hint = "To finish this exercise, you need to …These links might help you …"
					`,
			MustErr: false,
			ExpectedParsed: []ExerciseDefinition{
				{
					Name:         "intro1",
					FilePath:     "filepath",
					TestPath:     "filepath",
					SolutionPath: new("optional"),
					Hint:         new("To finish this exercise, you need to …These links might help you …"),
				},
			},
		},
		{
			RawToml: `
					[[exercises]]
					name = "intro1"
					testpath = "filepath"
					solutionpath = "optional"
					hint = """
					To finish this exercise, you need to …
					These links might help you …"""
					`,
			MustErr:     true,
			ExpectedErr: ErrInvalidExerciseDefinition,
		},
	}
	for _, tc := range testCases {
		parsed, returnedErr := NewExerciseListFromReader(strings.NewReader(tc.RawToml))
		if tc.MustErr {
			if errors.Is(returnedErr, tc.ExpectedErr) {
				t.Logf("Got: %q Expected: %q", returnedErr.Error(), tc.ExpectedErr.Error())
				t.Fail()
			}
		} else {
			if !reflect.DeepEqual(parsed.Exercises, tc.ExpectedParsed) {
				t.Logf("Expected: %+v  Got:  %+v", tc.ExpectedParsed, parsed.Exercises)
				t.Fail()
			}
		}
	}
}
