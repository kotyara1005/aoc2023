package main

import (
	"aoc2023"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	Num    int
	Rounds []Round
}

func parseGame(line string) *Game {
	var rexp = regexp.MustCompile("(\\d+) (red|blue|green)")

	l := strings.Split(strings.TrimPrefix(line, "Game "), ":")
	num, err := strconv.Atoi(l[0])
	if err != nil {
		log.Fatal(err)
	}

	result := new(Game)
	result.Num = num
	result.Rounds = make([]Round, 0, 3)
	rounds := strings.Split(l[1], ";")
	for _, round := range rounds {
		r := Round{0, 0, 0}
		buckets := rexp.FindAllStringSubmatch(round, -1)
		for _, b := range buckets {
			num, err := strconv.Atoi(b[1])
			if err != nil {
				log.Fatal(err)
			}
			if b[2] == "blue" {
				r.Blue += num
			} else if b[2] == "green" {
				r.Green += num
			} else if b[2] == "red" {
				r.Red += num
			} else {
				log.Fatal(b)
			}
		}
		result.Rounds = append(result.Rounds, r)
		//log.Println(num, buckets, r)
	}

	//log.Println(result)

	return result
}

const MAX_BLUE = 14
const MAX_RED = 12
const MAX_GREEN = 13

func PartOne(input []string) int {
	result := 0
	for _, line := range input {
		game := parseGame(line)
		log.Println("game", game)

		flag := true

		for _, round := range game.Rounds {
			if round.Blue > MAX_BLUE || round.Green > MAX_GREEN || round.Red > MAX_RED {
				flag = false
				break
			}
		}
		if flag {
			result += game.Num
		}
	}
	return result
}

func PartTwo(input []string) int {
	result := 0
	for _, line := range input {
		game := parseGame(line)
		log.Println("game", game)

		var (
			maxRed   = 0
			maxBlue  = 0
			maxGreen = 0
		)
		for _, round := range game.Rounds {
			maxRed = aoc2023.Max(round.Red, maxRed)
			maxBlue = aoc2023.Max(round.Blue, maxBlue)
			maxGreen = aoc2023.Max(round.Green, maxGreen)
		}
		result += maxRed * maxGreen * maxBlue
	}
	return result
}

func main() {
	var input = aoc2023.ReadLines("day2/input")
	//log.Println("PartOne", PartOne(input))
	log.Println("PartTwo", PartTwo(input))
}
