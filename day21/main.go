package main

import (
	"aoc2023"
	"log"
)

func findStart(graph []string) [2]int {

	for x, line := range graph {
		for y, val := range line {
			if val == 'S' {
				return [2]int{x, y}
			}
		}
	}

	rv := [2]int{-1, -1}
	log.Fatal("no start")
	return rv
}

var Diffs = [][2]int{
	{1, 0}, {0, 1}, {0, -1}, {-1, 0},
}

func findNeighbors(graph []string, point [2]int) [][2]int {
	result := [][2]int{}
	N := len(graph)
	M := len(graph[0])

	for _, diff := range Diffs {
		x := point[0] + diff[0]
		y := point[1] + diff[1]

		if x >= 0 && x < N && y >= 0 && y < M && graph[x][y] != '#' {
			result = append(result, [2]int{x, y})
		}
	}

	return result
}

func PartOne(graph []string, maxSteps int) int {
	start := findStart(graph)
	q := map[[2]int]struct{}{
		start: {},
	}

	for i := 0; i < maxSteps; i += 1 {
		log.Println(i, len(q))
		nxt := make(map[[2]int]struct{})

		for point, _ := range q {
			for _, p := range findNeighbors(graph, point) {
				nxt[p] = struct{}{}
			}
		}

		q = nxt
	}

	return len(q)
}

type Point struct {
	X    int
	Y    int
	MapX int
	MapY int
}

func findNeighborsV2(graph []string, point Point) []Point {
	result := []Point{}
	N := len(graph)
	M := len(graph[0])

	for _, diff := range Diffs {
		x := point.X + diff[0]
		y := point.Y + diff[1]
		mapX := point.MapX
		mapY := point.MapY

		if x < 0 {
			mapX -= 1
			x = N - 1
		}
		if x == N {
			mapX += 1
			x = 0
		}

		if y < 0 {
			mapY -= 1
			y = M - 1
		}
		if y == N {
			mapY += 1
			y = 0
		}

		if graph[x][y] != '#' {
			result = append(result, Point{x, y, mapX, mapY})
		}
	}

	return result
}

func PartTwo(graph []string, maxSteps int) int {
	start := findStart(graph)
	q := map[Point]struct{}{
		Point{start[0], start[1], 0, 0}: {},
	}

	for i := 0; i < maxSteps; i += 1 {
		if i < 100 || i%100 == 0 {
			log.Println(i, len(q))
		}
		nxt := make(map[Point]struct{})

		for point, _ := range q {
			for _, p := range findNeighborsV2(graph, point) {
				nxt[p] = struct{}{}
			}
		}

		q = nxt
	}

	return len(q)
}

func main() {
	input := aoc2023.ReadLines("day21/test_input")
	//log.Println(PartOne(input, 64))
	log.Println(PartTwo(input, 5000))
	//log.Println(PartTwo(input, 26501365))
}
