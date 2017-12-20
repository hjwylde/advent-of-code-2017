package main

import (
	"fmt"
	"github.com/hjwylde/advent-of-code-2017/solver"
	"io/ioutil"
	"os"
	"strings"
)

const usageTemplate = "Usage: advent-of-code-2017 DAY\n"
const unknownArgumentTemplate = "Unknown argument: %s\n"

var days = map[string]solver.Solver{
	"1": solver.Day1,
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		usage()
		os.Exit(2)
	}

	arg := args[0]

	day, ok := days[arg]
	if !ok {
		fmt.Printf(unknownArgumentTemplate, arg)
		os.Exit(2)
	}

	filename := fmt.Sprintf("./input/day%s.txt", arg)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	input := strings.TrimSuffix(string(data), "\n")

	output, err := day.Run(input)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", output)
}

func usage() {
	fmt.Printf(usageTemplate)
}
