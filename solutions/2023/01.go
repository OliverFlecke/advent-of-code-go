package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	aoc "github.com/oliverflecke/advent-of-code-go/client"
)

func main() {
	input := aoc.GetInput(aoc.Y2023, aoc.Day01)
	// 	input := `1abc2
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet`
	lines := strings.Split(input, "\n")

	var sum = 0
	for _, line := range lines {
		var first int
		var last int
		for _, c := range line {
			if digit, err := strconv.Atoi(string(c)); err == nil {
				first = digit
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if digit, err := strconv.Atoi(string(line[i])); err == nil {
				last = digit
				break
			}
		}

		number := first*10 + last
		// fmt.Printf("Line %s: %d\n", line, number)
		sum += number
	}

	fmt.Printf("Answer: %d\n", sum)

	// aoc.SubmitAnswer(aoc.Y2023, aoc.Day01, aoc.A, fmt.Sprintf("%d", sum))

	// 	testInput := `two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen
	// eighthree`
	// 	lines2 := strings.Split(testInput, "\n")

	re2 := regexp2.MustCompile(`(?=(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine))`, 0)

	var answerB = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		other := regexp2FindAllString(re2, line)
		first := toDigit(other[0])
		last := toDigit(other[len(other)-1])

		answerB += first*10 + last
	}

	fmt.Printf("Answer B: %d\n", answerB)
	// isBCorrect := aoc.SubmitAnswer(aoc.Y2023, aoc.Day01, aoc.B, fmt.Sprintf("%d", answerB))
	// fmt.Printf("Answer for B was: %s\n", isBCorrect)

}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.Groups()[1].String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func toDigit(value string) int {
	if digit, err := strconv.Atoi(value); err == nil {
		return digit
	}

	switch value {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}

	// panic(fmt.Sprintf("Unable to convert %s to a digit", value))
	return -1
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
