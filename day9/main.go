package main

import (
	"aoc2023"
	"log"
)

func parseInput(input []string) (result [][]int) {
	for _, line := range input {
		result = append(result, aoc2023.FindAllNumbers(line))
	}
	return
}

func findNext(seq []int) (result int) {
	allZeros := false

	cur := seq
	for !allZeros {
		log.Println(cur, result)
		result += cur[len(cur)-1]
		next := make([]int, len(cur)-1)
		allZeros = true

		for i, val := range cur[:len(cur)-1] {
			next[i] = cur[i+1] - val
			if next[i] != 0 {
				allZeros = false
			}
		}

		cur = next
	}
	return
}

func PartOne(input [][]int) int {
	result := 0
	//log.Println(findNext(input[0]))

	for _, line := range input {
		result += findNext(line)
	}

	return result
}

func PartTwo(input [][]int) int {
	result := 0

	//log.Println(findNext(aoc2023.Reverse(input[2])))

	for _, line := range input {
		result += findNext(aoc2023.Reverse(line))
	}

	return result
}

func main() {
	input := parseInput(aoc2023.ReadLines("day9/input"))
	//log.Println("PartOne", PartOne(input))
	log.Println("PartTwo", PartTwo(input))
}
