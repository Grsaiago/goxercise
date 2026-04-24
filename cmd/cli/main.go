package main

import (
	"fmt"
	"strings"

	"github.com/Grsaiago/goxercise/internal/entities"
)

func main() {
	// 	rawToml := `
	// 	filepath = "filepath"
	// 	testpath = "filepath"
	// 	solutionpath = "optional"
	// 	hint = """
	// 	To finish this exercise, you need to …
	// 	These links might help you …
	// 	"""
	// `
	// var example entities.ExerciseDefinition
	// 	if err := toml.NewDecoder(tomlReader).DisallowUnknownFields().Decode(&example); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%+v", example)

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

	tomlReader := strings.NewReader(rawToml)

	list, err := entities.NewExerciseListFromReader(tomlReader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", list)
}
