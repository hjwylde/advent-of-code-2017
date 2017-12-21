package solver

import (
	"fmt"
	"strconv"
)

var Day5 = Solver{
	Run: runDay5,
}

func runDay5(input []string) (string, error) {
	instructions := make([]int, len(input))

	for i, instruction := range input {
		value, err := strconv.Atoi(instruction)
		if err != nil {
			return "", err
		}

		instructions[i] = value
	}

	count := 0
	for i, instruction := 0, 0; i >= 0 && i < len(instructions); i += instruction {
		count++

		instruction = instructions[i]
		if instruction >= 3 {
			instructions[i]--
		} else {
			instructions[i]++
		}
	}

	output := fmt.Sprintf("%d", count)

	return output, nil
}
