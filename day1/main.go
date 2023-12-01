package main

import (
	"aoc2023"
	"log"
	"strconv"
	"strings"
)

const DIGITS = "1234567890"

var StringDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func PartOne() {
	lines := aoc2023.ReadLines("day1/input")
	sum := 0
	for _, line := range lines {
		first := strings.IndexAny(line, DIGITS)
		last := strings.LastIndexAny(line, DIGITS)

		n, err := strconv.Atoi(string([]rune{rune(line[first]), rune(line[last])}))
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	log.Println("Sum ", sum)
}

func findFirst(line string) rune {
	pos := strings.IndexAny(line, DIGITS)
	value := rune(line[pos])

	for i, num := range StringDigits {
		cur := strings.Index(line, num)
		if cur > -1 && cur < pos {
			pos = cur
			value = rune(strconv.Itoa(i + 1)[0])
		}
	}

	return value
}

func findLast(line string) rune {
	pos := strings.LastIndexAny(line, DIGITS)
	value := rune(line[pos])

	for i, num := range StringDigits {
		cur := strings.LastIndex(line, num)
		if cur > -1 && cur > pos {
			pos = cur
			value = rune(strconv.Itoa(i + 1)[0])
		}
	}

	return value
}

func PartTwo() {
	lines := aoc2023.ReadLines("day1/input")
	sum := 0
	for _, line := range lines {
		n, err := strconv.Atoi(string([]rune{
			findFirst(line), findLast(line),
		}))
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	log.Println("Sum ", sum)
}

func main() {
	//PartOne()
	PartTwo()
}
