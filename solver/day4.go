package solver

import (
	"fmt"
	"strings"
)

var Day4 = Solver{
	Run: runDay4,
}

func runDay4(input []string) (string, error) {
	count := 0
	for _, line := range input {
		ok := isValidPassphrase(line)
		if ok {
			count++
		}
	}

	output := fmt.Sprintf("%d", count)

	return output, nil
}

func isValidPassphrase(input string) bool {
	memory := make(map[string]bool)

	phrases := strings.Split(input, " ")
	for _, phrase := range phrases {
		if _, ok := memory[phrase]; ok {
			return false
		}

		memory[phrase] = true
	}

	return true
}
