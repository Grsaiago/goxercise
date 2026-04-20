package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Grsaiago/goxercise/internal/entities"
	"github.com/pelletier/go-toml/v2"
)

func main() {
	rawToml := `
	[[exercises]]
	filepath = "filepath"
	testpath = "filepath"
	solutionpath = "optional"
	hint = """
	To finish this exercise, you need to …
	These links might help you …
	"""
`

	list, err := entities.NewExerciseListFromReader(strings.NewReader(rawToml))
	if err != nil {
		newErr, _ := errors.AsType[*toml.StrictMissingError](err)
		panic(fmt.Sprintf("%+v\n", newErr.Errors[0].String()))
	}
	fmt.Printf("%+v", list)
}
