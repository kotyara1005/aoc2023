package main

import (
	"aoc2023"
	"fmt"
	"log"
	"strings"
)

type Point [2]int

func (p Point) String() string {
	return fmt.Sprintf("%d-%d", p[0], p[1])
}

//	| is a vertical pipe connecting north and south.
//	- is a horizontal pipe connecting east and west.
//	L is a 90-degree bend connecting north and east.
//	J is a 90-degree bend connecting north and west.
//	7 is a 90-degree bend connecting south and west.
//	F is a 90-degree bend connecting south and east.
//	. is ground; there is no pipe in this tile.
//	S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

var connects = map[[2]int]string{
	{-1, 0}: "|7FS",
	{1, 0}:  "|LJS",
	{0, -1}: "-LFS",
	{0, 1}:  "-J7S",
}

//var isConnnected = map[rune]string {
//	'S': "-J7",
//	'|': "",
//	'-': "J7S",
//	'L': "S-J7",
//	'J': "",
//	'7': "",
//	'F': "",
//}

var diffs = map[rune][]Point{
	'S': {{0, -1}, {0, 1}, {-1, 0}, {1, 0}},
	'|': {{-1, 0}, {1, 0}},
	'-': {{0, -1}, {0, 1}},
	'L': {{0, 1}, {-1, 0}},
	'J': {{0, -1}, {-1, 0}},
	'7': {{0, -1}, {1, 0}},
	'F': {{0, 1}, {1, 0}},
}

func findNeighbors(point Point, graph [][]rune) []Point {
	result := make([]Point, 0, 0)
	N := len(graph)
	M := len(graph[0])

	for _, diff := range diffs[graph[point[0]][point[1]]] {
		dx := diff[0]
		dy := diff[1]
		nx := point[0] + dx
		ny := point[1] + dy

		if nx < 0 || nx >= N {
			continue
		}

		if ny < 0 || ny >= M {
			continue
		}

		if !strings.ContainsRune(connects[diff], graph[nx][ny]) {
			continue
		}

		result = append(result, Point{nx, ny})
	}

	return result
}

//func filterVisited(nxt []Point, visited aoc2023.StringSet) []Point {
//	result := make([]Point, 0, 0)
//
//	for _, point := range nxt {
//		if !visited.Has(poi)
//	}
//
//	return result
//}

func BFS(start Point, graph [][]rune) (int, aoc2023.StringSet) {
	q := []Point{start}
	step := 0
	visited := aoc2023.NewStringSet([]string{start.String()})

	for len(q) > 0 {
		//log.Println(q)
		step += 1
		var nxt []Point

		for _, node := range q {
			for _, point := range findNeighbors(node, graph) {
				if visited.Has(point.String()) {
					continue
				}
				visited.Set(point.String())
				nxt = append(nxt, point)
			}
		}
		q = nxt
	}

	return step - 1, visited
}

func printGraph(graph [][]rune) {
	for _, line := range graph {
		log.Println(string(line))
	}
	log.Println("")
	log.Println("==============================")
}

func reverseGraphHorizontally(graph [][]rune) {
	for i := 0; i < len(graph); i += 1 {
		aoc2023.ReverseRunes(graph[i])
	}
}

func reverseGraphVertically(graph [][]rune) {
	aoc2023.ReverseRuneSlices(graph)
}

func checkHorizontally(loop aoc2023.StringSet, graph [][]rune, prevCond rune, cond rune, left string, right string) {
	for x, line := range graph {
		inLoop := false
		prev := '.'
		for y, val := range line {
			//if inLoop {
			//	graph[x][y] = 'I'
			//}
			areConnected := strings.ContainsRune(left, prev) && strings.ContainsRune(right, val)
			if loop.Has(Point{x, y}.String()) {
				if !areConnected {
					inLoop = !inLoop
				}
			} else if inLoop {
				//result += 1
				if graph[x][y] == prevCond || prevCond == '*' {
					graph[x][y] = cond
				}
			}

			prev = val
		}
		//log.Println(x, result)
	}
}

func checkVertically(loop aoc2023.StringSet, graph [][]rune, prevCond rune, cond rune, top string, bottom string) {
	N := len(graph)
	M := len(graph[0])
	for y := 0; y < M; y += 1 {
		inLoop := false
		prev := '.'
		for x := 0; x < N; x += 1 {
			val := graph[x][y]

			areConnected := strings.ContainsRune(top, prev) && strings.ContainsRune(bottom, val)
			if loop.Has(Point{x, y}.String()) {
				if !areConnected {
					inLoop = !inLoop
				}
			} else if inLoop {
				if graph[x][y] == prevCond || prevCond == '*' {
					graph[x][y] = cond
				}
			}
			prev = val
		}
	}
}

func PartTwoW(loop aoc2023.StringSet, graph [][]rune) int {
	left := "S-LF"
	right := "S-J7"
	top := "|7FS"
	bottom := "|LJS"
	checkHorizontally(loop, graph, '*', 'I', left, right)
	printGraph(graph)

	reverseGraphHorizontally(graph)
	checkHorizontally(loop, graph, 'I', 'D', right, left)
	reverseGraphHorizontally(graph)
	printGraph(graph)

	checkVertically(loop, graph, 'D', 'V', top, bottom)
	printGraph(graph)

	//reverseGraphVertically(graph)
	//checkVertically(loop, graph, 'V', 'A', bottom, top)
	//reverseGraphVertically(graph)
	//printGraph(graph)

	result := 0
	return result
}

func HasPathOutside(point Point, loop aoc2023.StringSet, graph [][]rune, cache map[Point]bool, visited map[Point]struct{}) bool {
	N := len(graph)
	M := len(graph[0])
	if point[0] < 0 || point[0] >= N {
		return true
	}

	if point[1] < 0 || point[1] >= M {
		return true
	}

	result, ok := cache[point]
	if ok {
		return result
	}
	if loop.Has(point.String()) {
		return false
	}

	_, ok = visited[point]
	if ok {
		return false
	}
	visited[point] = struct{}{}

	for _, diff := range diffs['S'] {
		dx := diff[0]
		dy := diff[1]
		nx := point[0] + dx
		ny := point[1] + dy

		if HasPathOutside(Point{nx, ny}, loop, graph, cache, visited) {
			return true
		}
	}
	return false

}

func blowGraphUp(graph [][]rune) [][]rune {
	left := "S-LF"
	right := "S-J7"
	top := "|7FS"
	bottom := "|LJS"

	_, loop := findStartAndLoop(graph)

	N := len(graph)
	M := len(graph[0])
	for i, _ := range graph {
		newLine := make([]rune, 0, 0)
		for j := 0; j < M-1; j += 1 {
			prev := graph[i][j]
			val := graph[i][j+1]
			areConnected := loop.Has(Point{i, j}.String()) && strings.ContainsRune(left, prev) && strings.ContainsRune(right, val)
			filler := '*'
			if areConnected {
				filler = '-'
			}

			newLine = append(newLine, prev, filler)
		}
		newLine = append(newLine, graph[i][M-1])
		graph[i] = newLine
	}

	//printGraph(graph)

	_, loop = findStartAndLoop(graph)

	result := make([][]rune, 0, 0)
	N = len(graph)
	M = len(graph[0])

	for i := 0; i < N-1; i += 1 {
		line := make([]rune, 0, 0)
		for j := 0; j < M; j += 1 {
			prev := graph[i][j]
			val := graph[i+1][j]
			areConnected := loop.Has(Point{i, j}.String()) && strings.ContainsRune(top, prev) && strings.ContainsRune(bottom, val)
			filler := '*'
			if areConnected {
				filler = '|'
			}
			line = append(line, filler)
		}
		result = append(result, graph[i], line)
	}

	result = append(result, graph[N-1])
	return result
}

func PartTwo(_ aoc2023.StringSet, graph [][]rune) int {
	//printGraph(graph)
	graph = blowGraphUp(graph)
	printGraph(graph)

	_, loop := findStartAndLoop(graph)

	N := len(graph)
	M := len(graph[0])
	cache := make(map[Point]bool)

	for x := 0; x < N; x += 1 {
		for y := 0; y < M; y += 1 {
			visited := make(map[Point]struct{})
			rv := HasPathOutside(Point{x, y}, loop, graph, cache, visited)
			//log.Println(x, y, rv, len(visited))
			for key, _ := range visited {
				cache[key] = rv
			}
			//if x == 3 && y == 3 {
			//	log.Println(visited)
			//}
		}
	}

	result := 0
	for key, val := range cache {
		if !val && !loop.Has(key.String()) && graph[key[0]][key[1]] != '*' {
			result += 1
			graph[key[0]][key[1]] = 'I'
		}
	}

	printGraph(graph)
	//log.Println(cache)
	//log.Println(loop)

	return result
}

func findStartAndLoop(graph [][]rune) (Point, aoc2023.StringSet) {
	start := Point{}

	for x, line := range graph {
		y := strings.IndexRune(string(line), 'S')
		if y != -1 {
			start = Point{x, y}
		}
	}

	_, loop := BFS(start, graph)
	return start, loop
}

func main() {
	input := aoc2023.ReadLines("day10/input")
	graph := make([][]rune, 0, 0)
	start := Point{}

	for x, line := range input {
		y := strings.IndexRune(line, 'S')
		if y != -1 {
			start = Point{x, y}
		}
		graph = append(graph, []rune(line))
	}

	result, loop := BFS(start, graph)
	log.Println(result)

	log.Println(PartTwo(loop, graph))

}
