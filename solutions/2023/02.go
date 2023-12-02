package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	aoc "github.com/oliverflecke/advent-of-code-go/client"
)

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func main() {
	// 	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	// 		`
	input := aoc.GetInput(aoc.Y2023, aoc.Day02)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp2.MustCompile(`Game (?<id>\d+):`, 0)
	reCubes := regexp2.MustCompile(`((?<count>\d+) (?<color>blue|green|red))`, 0)

	var games []Game
	for _, line := range lines {
		m, _ := re.FindStringMatch(line)
		group := m.GroupByName("id")
		id, _ := strconv.Atoi(group.Capture.String())
		rest := line[group.Index+group.Length+2:]
		sets := strings.Split(rest, ";")

		game := Game{
			id:   id,
			sets: []Set{},
		}
		for _, set := range sets {
			s := Set{}
			m, _ := reCubes.FindStringMatch(set)
			for m != nil {
				count, _ := strconv.Atoi(m.Groups()[2].String())
				color := m.Groups()[3].String()

				// fmt.Printf("\nFound %d Count %d color '%s'", id, count, color)
				switch color {
				case "red":
					s.red = count
				case "blue":
					s.blue = count
				case "green":
					s.green = count
				}

				m, _ = reCubes.FindNextMatch(m)
			}

			game.sets = append(game.sets, s)
		}
		games = append(games, game)
	}

	answerA := 0
	for _, game := range games {
		if checkLimits(game) {
			answerA += game.id
		}
	}

	fmt.Printf("Answer A: %d\n", answerA)
	// resultA := aoc.SubmitAnswer(aoc.Y2023, aoc.Day02, aoc.A, string(answerA))
	// fmt.Printf("Anwer A result: %s", resultA)

	// Part Two
	answerB := 0
	for _, game := range games {
		set := game.MinSet()
		// fmt.Printf("%d: min set %s. Power %d\n", game.id, set, set.Power())

		answerB += set.Power()
	}

	fmt.Printf("Answer B: %d\n", answerB)
}

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int
	blue  int
	green int
}

func (s Set) String() string {
	return fmt.Sprintf("%d red, %d blue, %d green", s.red, s.blue, s.green)
}

func (s Set) Power() int {
	return s.red * s.blue * s.green
}

func checkLimits(game Game) bool {
	for _, set := range game.sets {
		if set.red > maxRed || set.blue > maxBlue || set.green > maxGreen {
			return false
		}
	}
	return true
}

func (g Game) MinSet() Set {
	result := Set{}
	for _, set := range g.sets {
		if set.blue > result.blue {
			result.blue = set.blue
		}
		if set.red > result.red {
			result.red = set.red
		}
		if set.green > result.green {
			result.green = set.green
		}
	}

	return result
}
