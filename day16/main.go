package main

import (
	"aoc2023"
	"log"
)

type Direction [2]int

var (
	Up    Direction = [2]int{-1, 0}
	Down            = [2]int{1, 0}
	Left            = [2]int{0, -1}
	Right           = [2]int{0, 1}
)

type Node struct {
	X         int
	Y         int
	Direction Direction
}

func isValid(x int, y int, N int, M int) bool {
	return x >= 0 && x < N && y >= 0 && y < M
}

var NextDirection = map[byte]map[Direction][]Direction{
	'.': {
		Up:    {Up},
		Down:  {Down},
		Left:  {Left},
		Right: {Right},
	},
	'/': {
		Up:    {Right},
		Down:  {Left},
		Left:  {Down},
		Right: {Up},
	},
	'\\': {
		Up:    {Left},
		Down:  {Right},
		Left:  {Up},
		Right: {Down},
	},
	'-': {
		Up:    {Left, Right},
		Down:  {Left, Right},
		Left:  {Left},
		Right: {Right},
	},
	'|': {
		Up:    {Up},
		Down:  {Down},
		Left:  {Up, Down},
		Right: {Up, Down},
	},
}

func Next(node Node, val byte, N int, M int) (result []Node) {
	dirs := NextDirection[val][node.Direction]

	for _, nd := range dirs {
		dx := nd[0]
		dy := nd[1]
		nx := node.X + dx
		ny := node.Y + dy
		if isValid(nx, ny, N, M) {
			result = append(result, Node{nx, ny, nd})
		}
	}
	return
}

func PrintState(input []string, visited map[Node]struct{}) {
	log.Println("")
	for x, line := range input {
		buf := []rune{}
		for y, val := range []rune(line) {
			loc := val

			for _, dir := range []Direction{Up, Down, Left, Right} {
				_, ok := visited[Node{x, y, dir}]
				if ok {
					loc = '#'
					break
				}
			}

			buf = append(buf, loc)
		}
		log.Println(string(buf))
	}

	log.Println("==============")
}

func PartOne(start Node, input []string) int {
	N := len(input)
	M := len(input[0])
	visited := make(map[Node]struct{})
	visitedCells := make(map[[2]int]struct{})
	q := []Node{start}

	for len(q) > 0 {
		var nxt []Node
		for _, node := range q {
			_, ok := visited[node]
			if ok {
				continue
			}
			visited[node] = struct{}{}
			visitedCells[[2]int{node.X, node.Y}] = struct{}{}

			nxt = append(nxt, Next(node, input[node.X][node.Y], N, M)...)
		}
		q = nxt
	}

	//PrintState(input, visited)

	return len(visitedCells)
}

func PartTwo(input []string) int {
	N := len(input)
	M := len(input[0])
	result := 0
	for x := 0; x < N; x += 1 {
		loc := PartOne(Node{0, x, Down}, input)
		result = aoc2023.Max(result, loc)
	}

	for x := 0; x < N; x += 1 {
		loc := PartOne(Node{N - 1, x, Up}, input)
		result = aoc2023.Max(result, loc)
	}

	for x := 0; x < M; x += 1 {
		loc := PartOne(Node{x, 0, Right}, input)
		result = aoc2023.Max(result, loc)
	}
	for x := 0; x < M; x += 1 {
		loc := PartOne(Node{x, M - 1, Left}, input)
		result = aoc2023.Max(result, loc)
	}
	return result
}

func main() {
	input := aoc2023.ReadLines("day16/input")
	//log.Println(PartOne(input))
	log.Println(PartTwo(input))
}
