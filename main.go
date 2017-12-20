package main

import (
	"fmt"
	"github.com/hjwylde/advent-of-code-2017/day01"
	"io/ioutil"
	"os"
	"strings"
)

const usageTemplate = "Usage: advent-of-code-2017 DAY\n"
const unknownArgumentTemplate = "Unknown argument: %s\n"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		usage()
		os.Exit(2)
	}

	arg := args[0]
	switch arg {
	case "1":
		data, err := ioutil.ReadFile("./input/day01.txt")
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}

		input := strings.TrimSuffix(string(data), "\n")

		output, err := day01.Run(input)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s\n", output)
	default:
		fmt.Printf(unknownArgumentTemplate, arg)
		os.Exit(2)
	}
}

func usage() {
	fmt.Printf(usageTemplate)
}
