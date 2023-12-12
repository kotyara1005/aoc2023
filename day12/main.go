package main

import (
	"aoc2023"
	"log"
	"strings"
)

type Record struct {
	Springs string
	Groups  []int
}

func parseInput(input []string) (result []Record) {
	for _, line := range input {
		arr := strings.Split(line, " ")
		result = append(
			result,
			Record{
				Springs: arr[0],
				Groups:  aoc2023.FindAllNumbers(arr[1]),
			})
	}
	return
}

func FindAllSolutions(Springs string, Groups []int) int {
	result := 0
	springs := []rune(Springs)
	groups := append([]int{0}, Groups...)

	var Inner func(pos int, group int)
	Inner = func(pos int, group int) {
		if groups[group] < 0 {
			return
		}
		if pos == len(springs) {
			if group == len(groups)-1 && groups[group] == 0 {
				//log.Println(string(springs), groups)
				result += 1
			}
			return
		}

		prev := '.'
		if pos > 0 {
			prev = springs[pos-1]
		}
		state := springs[pos]

		switch state {
		case '.':
			if prev == '#' && groups[group] > 0 {
				return
			}
			Inner(pos+1, group)
		case '#':
			if prev == '.' {
				group += 1
			}
			if group == len(groups) {
				return
			}
			groups[group] -= 1
			Inner(pos+1, group)
			groups[group] += 1
		case '?':
			for _, sym := range ".#" {
				springs[pos] = sym
				Inner(pos, group)
			}
			springs[pos] = '?'
		}

	}

	Inner(0, 0)

	return result
}

func PartOne(input []string) int {
	records := parseInput(input)
	result := 0
	for i, record := range records {
		loc := FindAllSolutionsCached(record.Springs, record.Groups)
		log.Println(input[i], "=>", loc)
		result += loc
	}

	return result
}

func parseInputV2(input []string) (result []string) {
	for _, line := range input {
		arr := strings.Split(line, " ")
		l := strings.Join([]string{
			strings.Join([]string{arr[0], arr[0], arr[0], arr[0], arr[0]}, "?"),
			strings.Join([]string{arr[1], arr[1], arr[1], arr[1], arr[1]}, ","),
		},
			" ",
		)
		result = append(result, l)
	}
	return
}

func FindAllSolutionsCached(Springs string, Groups []int) int {
	result := 0
	springs := []rune(Springs)
	groups := append([]int{0}, Groups...)
	cache := make(map[[5]int]int)

	var Inner func(pos int, group int) int
	Inner = func(pos int, group int) int {
		if groups[group] < 0 {
			return 0
		}
		if pos == len(springs) {
			if group == len(groups)-1 && groups[group] == 0 {
				return 1
			}
			return 0
		}

		prev := '.'
		if pos > 0 {
			prev = springs[pos-1]
		}
		state := springs[pos]
		key := [5]int{pos, group, groups[group], int(prev), int(state)}
		if rv, ok := cache[key]; ok {
			return rv
		}

		loc := 0

		switch state {
		case '.':
			if prev == '#' && groups[group] > 0 {
				return 0
			}
			loc += Inner(pos+1, group)
		case '#':
			if prev == '.' {
				group += 1
			}
			if group == len(groups) {
				return 0
			}
			groups[group] -= 1
			loc += Inner(pos+1, group)
			groups[group] += 1
		case '?':
			for _, sym := range ".#" {
				springs[pos] = sym
				loc += Inner(pos, group)
			}
			springs[pos] = '?'
		}

		cache[key] = loc
		return loc

	}

	result = Inner(0, 0)

	return result
}

func PartTwo(input []string) int {
	return PartOne(parseInputV2(input))
	//log.Println(parseInputV2(input)[0])
	//return 0
}

func main() {
	input := aoc2023.ReadLines("day12/input")
	//log.Println(PartOne(input))
	log.Println(PartTwo(input))
}
