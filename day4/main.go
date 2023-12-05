package main

import (
	"aoc2023"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var NumberRegexp = regexp.MustCompile(`\d+`)

type Card struct {
	Num        int
	Numbers    map[int]struct{}
	WinNumbers map[int]struct{}
}

func (c Card) NumberOfMatches() int {
	result := 0

	for num, _ := range c.Numbers {
		_, ok := c.WinNumbers[num]
		if ok {
			result += 1
		}
	}

	return result
}

func (c Card) Score() int {
	degree := c.NumberOfMatches()

	if degree == 0 {
		return 0
	}
	return 1 << (degree - 1)
}

func parse(input []string) []Card {
	cards := make([]Card, 0, 0)

	for _, line := range input {
		l := strings.Split(strings.TrimPrefix(line, "Card "), ":")
		nums := strings.Split(l[1], "|")
		n, err := strconv.Atoi(strings.TrimSpace(l[0]))
		if err != nil {
			log.Fatal(err)
		}
		card := Card{
			n,
			make(map[int]struct{}),
			make(map[int]struct{}),
		}

		for _, val := range NumberRegexp.FindAllString(nums[0], -1) {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			card.Numbers[num] = struct{}{}
		}

		for _, val := range NumberRegexp.FindAllString(nums[1], -1) {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			card.WinNumbers[num] = struct{}{}
		}

		cards = append(cards, card)
	}

	return cards
}

func PartOne(cards []Card) int {
	result := 0

	for _, card := range cards {
		result += card.Score()
	}

	return result
}

func PartTwo(cards []Card) int {
	nums := make([]int, len(cards))

	for i, _ := range nums {
		nums[i] = 1
	}

	for i, card := range cards {
		n := card.NumberOfMatches()

		for j := 0; j < n; j += 1 {
			nums[i+j+1] += nums[i]
		}
	}

	log.Println(nums)

	result := 0
	for _, num := range nums {
		result += num
	}

	return result
}

func main() {
	input := parse(aoc2023.ReadLines("day4/input"))
	//log.Println(1 << 0)
	//log.Println(PartOne(input))
	log.Println("PartTwo", PartTwo(input))
}
