package types

import (
	"errors"
	"strings"
	"testing"
)

func TestUnmarshallExerciseList(t *testing.T) {

	testCases := []struct {
		RawToml     string
		ExpectedErr error
		MustErr     bool
	}{
		{
			RawToml: `
					[[exercises]]
					name = "intro1"
					filepath = "filepath"
					testpath = "filepath"
					solutionpath = "optional"
					hint = """
					To finish this exercise, you need to …
					These links might help you …"""
					`,
			MustErr: false,
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
		_, got := LoadExerciseListFromReader(strings.NewReader(tc.RawToml))
		if tc.MustErr {
			if errors.Is(got, tc.ExpectedErr) {
				t.Logf("Got: %q Expected: %q", got.Error(), tc.ExpectedErr.Error())
				t.Fail()
			}
		}
	}
}
