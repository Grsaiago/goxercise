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
		ExpectedParsed []ExerciseDefinition
		MustErr        bool
		ErrorMatcher   func(int, *testing.T, error) bool
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
			MustErr: true,
			ErrorMatcher: func(i int, t *testing.T, err error) bool {
				got, returnValue := errors.AsType[ExerciseDefinitionValidationError](err)
				if !returnValue {
					t.Logf("tc[%d] Got: %T Expected: %T", i, got, ExerciseDefinitionValidationError{})
				}
				return returnValue
			},
		},
	}
	for i, tc := range testCases {
		parsed, returnedErr := NewExerciseListFromReader(strings.NewReader(tc.RawToml))
		if tc.MustErr {
			if returnedErr == nil {
				t.Logf("tc[%d] Expected error but got nil", i)
				t.FailNow()
			}
			if !tc.ErrorMatcher(i, t, returnedErr) {
				t.FailNow()
			}
		} else {
			if !reflect.DeepEqual(parsed.Exercises, tc.ExpectedParsed) {
				t.Logf("tc[%d] Expected: %+v  Got:  %+v", i, tc.ExpectedParsed, parsed.Exercises)
				t.FailNow()
			}
		}
	}
}
