package solver

import (
	"fmt"
)

var Day1 = Solver{
	Run: runDay1,
}

func runDay1(input []string) (string, error) {
	sum := 0

	chars := []rune(input[0])
	for i, c := range chars {
		next := (i + len(chars)/2) % len(chars)

		if c == chars[next] {
			sum += int(c - '0')
		}
	}

	output := fmt.Sprintf("%d", sum)

	return output, nil
}
