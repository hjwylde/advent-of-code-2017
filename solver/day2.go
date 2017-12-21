package solver

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var Day2 = Solver{
	Run: runDay2,
}

func runDay2(input []string) (string, error) {
	checksum := 0
	for _, line := range input {
		value, err := calculateLineRange(line)
		if err != nil {
			return "", err
		}

		checksum += value
	}

	output := fmt.Sprintf("%d", checksum)

	return output, nil
}

func calculateLineRange(input string) (int, error) {
	low := math.MaxInt64
	high := math.MinInt64

	numbers := strings.Split(input, "\t")
	for _, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}

		low = min(low, value)
		high = max(value, high)
	}

	return high - low, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
