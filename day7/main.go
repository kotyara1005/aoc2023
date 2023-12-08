package main

import (
	"aoc2023"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Combination int

const (
	Five Combination = 6 - iota
	Four
	FullHouse
	Three
	TwoPairs
	Pair
	HighCard
)

func isFive(counter aoc2023.Counter) bool {
	return counter.HasValue(5)
}

func isFour(counter aoc2023.Counter) bool {
	return counter.HasValue(4)
}

func isFullHouse(c aoc2023.Counter) bool {
	return c.HasValue(2) && c.HasValue(3)
}

func isThree(counter aoc2023.Counter) bool {
	return counter.HasValue(3)
}

func isTwoPairs(c aoc2023.Counter) bool {
	v := 0

	for _, val := range c {
		if val == 2 {
			v += 1
		}
	}
	return v >= 2
}

func isPair(counter aoc2023.Counter) bool {
	return counter.HasValue(2)
}

func FindCombination(counter aoc2023.Counter) Combination {
	if isFive(counter) {
		return Five
	} else if isFour(counter) {
		return Four
	} else if isFullHouse(counter) {
		return FullHouse
	} else if isThree(counter) {
		return Three
	} else if isTwoPairs(counter) {
		return TwoPairs
	} else if isPair(counter) {
		return Pair
	} else {
		return HighCard
	}
}

func FindCombinationWihJokers(counter aoc2023.Counter) Combination {
	jokers, _ := counter['J']
	delete(counter, 'J')

	key, _ := counter.Max()
	counter[key] += jokers

	return FindCombination(counter)

}

var CardPowers = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type Hand struct {
	Cards       []rune
	Bet         int
	Combination Combination
}

func (h *Hand) String() string {
	return fmt.Sprintf("Hand(%s, %d, %d)", string(h.Cards), h.Bet, h.Combination)
}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Less(i, j int) bool {
	l := h[i]
	r := h[j]

	if l.Combination != r.Combination {
		return l.Combination < r.Combination
	}

	for idx, rn := range l.Cards {
		if rn != r.Cards[idx] {
			//log.Println(l.String(), r.String(), rn, r.Cards[idx])
			return CardPowers[rn] < CardPowers[r.Cards[idx]]
		}
	}

	return false
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func parseInput(input []string) Hands {
	result := make(Hands, 0, 0)

	for _, line := range input {
		l := strings.Split(line, " ")
		num, err := strconv.Atoi(l[1])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Hand{[]rune(l[0]), num, FindCombination(aoc2023.NewCounter(l[0]))})
	}

	return result
}

func parseInputV2(input []string) Hands {
	result := make(Hands, 0, 0)

	for _, line := range input {
		l := strings.Split(line, " ")
		num, err := strconv.Atoi(l[1])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Hand{[]rune(l[0]), num, FindCombinationWihJokers(aoc2023.NewCounter(l[0]))})
	}

	return result
}

func PartOne(hands Hands) int {
	sort.Sort(hands)
	result := 0
	for i, h := range hands {
		//log.Println(i+1, h.String())
		result += (i + 1) * h.Bet
	}
	return result
}

func PartTwo(hands Hands) int {
	// 12345
	// 1234J
	// 123JJ
	// 12JJJ
	// 254494947

	CardPowers['J'] = -1
	sort.Sort(hands)
	result := 0
	for i, h := range hands {
		//log.Println(i+1, h.String())
		result += (i + 1) * h.Bet
	}
	return result
}

func main() {
	input := aoc2023.ReadLines("day7/input")
	log.Println("PartOne", PartOne(parseInput(input)))
	log.Println("PartTwo", PartTwo(parseInputV2(input)))
}
