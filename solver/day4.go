package solver

import (
	"fmt"
	"sort"
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
		chars := strings.Split(phrase, "")
		sort.Strings(chars)

		key := strings.Join(chars, "")
		if _, ok := memory[key]; ok {
			return false
		}

		memory[key] = true
	}

	return true
}
