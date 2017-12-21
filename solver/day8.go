package solver

import (
	"fmt"
	"strconv"
	"strings"
)

var Day8 = Solver{
	Run: runDay8,
}

func runDay8(input []string) (string, error) {
	registers := make(map[string]int)
	for _, instruction := range input {
		tokens := strings.Fields(instruction)

		register1 := tokens[0]
		op1 := tokens[1]
		value1, err := strconv.Atoi(tokens[2])
		if err != nil {
			return "", err
		}
		register2 := tokens[4]
		op2 := tokens[5]
		value2, err := strconv.Atoi(tokens[6])
		if err != nil {
			return "", err
		}

		if evalConditional(register2, op2, value2, registers) {
			value3 := evalExpression(register1, op1, value1, registers)

			registers[register1] = value3
		}
	}

	max := 0
	for _, value := range registers {
		if value > max {
			max = value
		}
	}

	output := fmt.Sprintf("%d", max)

	return output, nil
}

func get(register string, registers map[string]int) int {
	if value, ok := registers[register]; ok {
		return value
	}

	return 0
}

func evalConditional(register string, op string, value int, registers map[string]int) bool {
	rvalue := get(register, registers)

	switch op {
	case "==":
		return rvalue == value
	case "!=":
		return rvalue != value
	case "<":
		return rvalue < value
	case "<=":
		return rvalue <= value
	case ">":
		return rvalue > value
	case ">=":
		return rvalue >= value
	default:
		panic("Uh oh")
	}
}

func evalExpression(register string, op string, value int, registers map[string]int) int {
	rvalue := get(register, registers)

	switch op {
	case "inc":
		return rvalue + value
	case "dec":
		return rvalue - value
	default:
		panic("Uh oh")
	}
}
