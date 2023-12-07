package main

import (
	"aoc2023"
	"log"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func parseInput(input []string) []*Race {
	ints := [2][]int{
		aoc2023.FindAllNumbers(input[0]),
		aoc2023.FindAllNumbers(input[1]),
	}

	result := make([]*Race, 0, 0)

	for i, time := range ints[0] {
		dist := ints[1][i]
		result = append(result, &Race{time, dist})
	}

	return result
}

func PartOne(races []*Race) int {
	result := 1
	for n, race := range races {
		loc := 0

		for t := 1; t < race.Time; t += 1 {
			dist := t * (race.Time - t)
			if dist > race.Distance {
				//log.Println("Win", t, dist)
				loc += 1
			}
		}
		log.Println("Race ", n, "result", loc)
		result *= loc
	}

	return result
}

func PartTwo(input []string) int {
	time, err := strconv.Atoi(strings.Join(aoc2023.Number.FindAllString(input[0], -1), ""))
	if err != nil {
		log.Fatal(err)
	}
	dist, err := strconv.Atoi(strings.Join(aoc2023.Number.FindAllString(input[1], -1), ""))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(time, dist)
	result := 0

	for t := 1; t < time; t += 1 {
		if t*(time-t) > dist {
			//log.Println("Win", t, dist)
			result += 1
		}
	}

	return result
}

func main() {
	//log.Println(PartOne(parseInput(aoc2023.ReadLines("day6/input"))))
	log.Println(PartTwo(aoc2023.ReadLines("day6/input")))
}
