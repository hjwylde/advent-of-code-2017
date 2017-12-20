package day01

import (
	"fmt"
)

func Run(input string) (string, error) {
	sum := 0

	chars := []rune(input)
	for i, c := range chars {
		next := (i + 1) % len(chars)

		if c == chars[next] {
			sum += int(c - '0')
		}
	}

	output := fmt.Sprintf("%d", sum)

	return output, nil
}
