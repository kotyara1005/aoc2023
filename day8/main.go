package main

import (
	"aoc2023"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Graph map[string][2]string
type GraphIterator struct {
	Graph    Graph
	Path     []rune
	Idx      int
	Position string
}

func (gi *GraphIterator) Next() {
	gi.Idx = (gi.Idx + 1) % len(gi.Path)
	action := gi.Path[gi.Idx]
	if action == 'L' {
		gi.Position = gi.Graph[gi.Position][0]
	} else {
		gi.Position = gi.Graph[gi.Position][1]
	}
}

func (gi *GraphIterator) IsEnd() bool {
	return gi.Position[2] == 'Z'
}

func Brent(path_ string, graph Graph, start string) (cycleLength int, distanceToFirstEnd int) {
	slow := GraphIterator{graph, []rune(path_), -1, start}
	fast := slow
	fast.Next()
	//log.Println("Hare pos", fast.Position, "tort pos", slow.Position)

	power := 1
	cycleLength = 1

	for !reflect.DeepEqual(slow, fast) {
		if power == cycleLength {
			slow = fast
			power *= 2
			cycleLength = 0
		}
		fast.Next()
		cycleLength += 1
	}

	//log.Println("Found cycle length", cycleLength)
	slow = GraphIterator{graph, []rune(path_), -1, start}
	fast = slow

	for i := 0; i < cycleLength; i += 1 {
		fast.Next()
	}

	//log.Println("Fast pos", fast.Position)

	distanceToFirstEnd = 0
	for !reflect.DeepEqual(slow, fast) {
		slow.Next()
		fast.Next()
		distanceToFirstEnd += 1
	}

	//log.Println(distanceToFirstEnd, slow.Position, fast.Position)

	for !slow.IsEnd() {
		slow.Next()
		distanceToFirstEnd += 1
	}
	//log.Println(distanceToFirstEnd, slow.Position)

	//cycleLength = 0

	//log.Println(cycleLength == distanceToFirstEnd)

	return
}

func parseInput(input []string) (string, Graph) {
	regex := regexp.MustCompile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)
	graph := make(Graph)

	for _, line := range input[2:] {
		rv := regex.FindAllStringSubmatch(line, -1)
		//log.Println(rv)
		graph[rv[0][1]] = [2]string{rv[0][2], rv[0][3]}
	}

	//log.Println(graph)
	return input[0], graph
}

func PartOne(path_ string, graph Graph) int {
	result := 0

	cur := "AAA"
	end := "ZZZ"
	i := 0
	path := []rune(path_)

	for {
		if cur == end {
			break
		}
		result += 1
		action := path[i]
		i = (i + 1) % len(path)
		if action == 'L' {
			cur = graph[cur][0]
		} else {
			cur = graph[cur][1]
		}
	}

	return result
}

func findStartNodes(graph Graph) []string {
	result := make([]string, 0, 0)

	for node, _ := range graph {
		if node[2] == 'A' {
			result = append(result, node)
		}
	}

	return result
}

func makeStep(positions []string, graph Graph, direction rune) {
	for i, _ := range positions {
		if direction == 'L' {
			positions[i] = graph[positions[i]][0]

		} else {
			positions[i] = graph[positions[i]][1]
		}
	}
}

func isEnd(positions []string) bool {
	for _, pos := range positions {
		if pos[2] != 'Z' {
			return false
		}
	}
	return true
}

func PartTwoBruteForce(path_ string, graph Graph) int {
	result := 0
	i := 0
	path := []rune(path_)
	visited := aoc2023.NewStringSet(nil)

	positions := findStartNodes(graph)
	log.Println(positions)

	for {
		if isEnd(positions) {
			break
		}
		result += 1
		action := path[i]
		sort.Strings(positions)
		key := strings.Join(append(append([]string{}, positions...), strconv.Itoa(i)), "-")
		if visited.Has(key) {
			log.Fatal("visited")
		}
		visited.Set(key)

		if result%1000000 == 0 {
			log.Println(result, key, len(visited))
		}

		i = (i + 1) % len(path)
		makeStep(positions, graph, action)
	}

	return result
}

func PartTwo(path_ string, graph Graph) int {
	positions := findStartNodes(graph)

	for _, pos := range positions {
		cl, _ := Brent(path_, graph, pos)
		log.Print(pos, " ", cl)
	}
	return 0
}

func main() {
	path, graph := parseInput(aoc2023.ReadLines("day8/input"))
	//log.Println("PartOne", PartOne(path, graph))
	log.Println("PartTwo", PartTwo(path, graph))
}
