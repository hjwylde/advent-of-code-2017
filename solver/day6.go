package solver

import (
	"fmt"
	"strconv"
	"strings"
)

var Day6 = Solver{
	Run: runDay6,
}

func runDay6(input []string) (string, error) {
	numbers := strings.Split(input[0], "\t")

	banks := make([]int, len(numbers))
	for i, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil {
			return "", err
		}

		banks[i] = value
	}

	fmt.Printf("%v\n", banks)

	memory := make(map[string]bool)

	count := 0
	for {
		key := generateBanksKey(banks)
		if _, ok := memory[key]; ok {
			break
		}

		memory[key] = true

		largestBankIndex := findLargestBankIndex(banks)

		blocks := banks[largestBankIndex]
		banks[largestBankIndex] = 0

		for i := largestBankIndex + 1; i <= largestBankIndex+blocks; i++ {
			banks[i%len(banks)]++
		}

		count++
	}

	output := fmt.Sprintf("%d", count)

	return output, nil
}

func generateBanksKey(banks []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(banks)), ",")
}

func findLargestBankIndex(banks []int) int {
	largestBankIndex := 0
	for i, blocks := range banks {
		if blocks <= banks[largestBankIndex] {
			continue
		}

		largestBankIndex = i
	}

	return largestBankIndex
}
