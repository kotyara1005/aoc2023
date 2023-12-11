package main

import (
	"aoc2023"
	"log"
)

func buildPrefSums(universe []string) ([]int, []int) {
	rowsSums := make([]int, len(universe))
	for i, line := range universe {
		isEmpty := true
		for _, val := range line {
			if val != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			rowsSums[i] = 1
		}
	}

	for i := 1; i < len(universe); i += 1 {
		rowsSums[i] += rowsSums[i-1]
	}
	log.Println("rowsSums", rowsSums)

	M := len(universe[0])
	colsSums := make([]int, M)
	for j := 0; j < M; j += 1 {
		isEmpty := true
		for i := 0; i < len(universe); i += 1 {
			if universe[i][j] != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			colsSums[j] = 1
		}
	}
	for i := 1; i < M; i += 1 {
		colsSums[i] += colsSums[i-1]
	}
	log.Println("colsSums", colsSums)
	return rowsSums, colsSums
}

func findPoints(universe []string) (result []aoc2023.Point) {
	for x, line := range universe {
		for y, val := range line {
			if val == '#' {
				result = append(result, aoc2023.Point{x, y})
			}
		}
	}
	return
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func PartOne(universe []string) int {
	rowsSums, colsSums := buildPrefSums(universe)
	galaxies := findPoints(universe)
	log.Println(galaxies)

	result := 0

	for i := 0; i < len(galaxies); i += 1 {
		for j := i + 1; j < len(galaxies); j += 1 {

			a := galaxies[i]
			b := galaxies[j]
			dist := abs(a[0]-b[0]) + abs(a[1]-b[1]) + abs(rowsSums[a[0]]-rowsSums[b[0]])*999999 + abs(colsSums[a[1]]-colsSums[b[1]])*999999
			//log.Println(i, j, dist)
			result += dist
		}
	}

	return result
}

func main() {
	input := aoc2023.ReadLines("day11/input")
	log.Println(PartOne(input))
}
