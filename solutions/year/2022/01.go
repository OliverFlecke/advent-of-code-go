package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/oliverflecke/advent-of-code-go/client"
)

func main() {
	var input = aoc.GetInput(aoc.Y2022, aoc.Day01)
	var groups = strings.Split(input, "\n\n")

	var max = 0
	for _, g := range groups {
		var lines = strings.Split(g, "\n")
		var sum = 0
		for _, s := range lines {
			var x, _ = strconv.Atoi(s)
			sum += x
		}

		if sum > max {
			max = sum
		}
	}

	fmt.Printf("Answer: %d\n", max)
	aoc.SubmitAnswer(aoc.Y2022, aoc.Day01, aoc.A, fmt.Sprintf("%d", max))
}
