package solver

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var Day3 = Solver{
	Run: runDay3,
}

func runDay3(input []string) (string, error) {
	number := strings.TrimSpace(input[0])
	value, err := strconv.ParseFloat(number, 64)
	if err != nil {
		return "", err
	}

	sqrt := math.Sqrt(value)
	lower := math.Trunc(sqrt)
	upper := lower + 1

	output := fmt.Sprintf("%0f - %0f", lower, upper)

	return output, nil
}
