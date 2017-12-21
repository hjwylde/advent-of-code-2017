package solver

import (
	"fmt"
	"regexp"
	"strings"
)

var Day7 = Solver{
	Run: runDay7,
}

func runDay7(input []string) (string, error) {
	memory := make(map[string]bool)
	for _, line := range input {
		matcher, err := regexp.Compile("[^a-z ]+")
		if err != nil {
			return "", err
		}

		normalisedLine := matcher.ReplaceAllString(line, "")
		nodes := strings.Fields(normalisedLine)

		program := nodes[0]
		_, ok := memory[program]
		memory[program] = ok

		subprograms := nodes[1:]
		for _, subprogram := range subprograms {
			memory[subprogram] = true
		}
	}

	var bottomProgram string
	for program, parents := range memory {
		if !parents {
			bottomProgram = program
			break
		}
	}

	output := fmt.Sprintf("%s", bottomProgram)

	return output, nil
}
