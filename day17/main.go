package main

import (
	"aoc2023"
	"container/heap"
	"log"
)

type Direction [2]int

var (
	Up    Direction = [2]int{-1, 0}
	Down            = [2]int{1, 0}
	Left            = [2]int{0, -1}
	Right           = [2]int{0, 1}
)

var NextDirections = map[Direction][3]Direction{
	Up:    {Up, Right, Left},
	Down:  {Down, Left, Right},
	Left:  {Left, Up, Down},
	Right: {Right, Up, Down},
}

type Node struct {
	X         int
	Y         int
	Direction Direction
	Steps     int
	HeatLoss  int
}

func (n Node) IsValid(N int) bool {
	return n.X >= 0 && n.X < N && n.Y >= 0 && n.Y < N && n.Steps < 3
}

func (n Node) Next(graph [][]int) (result []Node) {
	for _, nextDirection := range NextDirections[n.Direction] {
		steps := 0
		if n.Direction == nextDirection {
			steps = n.Steps + 1
		}
		node := Node{
			n.X + nextDirection[0],
			n.Y + nextDirection[1],
			nextDirection,
			steps,
			n.HeatLoss,
		}
		if node.IsValid(len(graph)) {
			//log.Println(node.X, node.Y, graph)
			node.HeatLoss += graph[node.X][node.Y]
			result = append(result, node)
		}
	}
	return
}

func (n Node) IsValidV2(N int, M int) bool {
	return n.X >= 0 && n.X < N && n.Y >= 0 && n.Y < M && n.Steps < 10
}

func (n Node) NextV2(graph [][]int) (result []Node) {
	for _, nextDirection := range NextDirections[n.Direction] {
		steps := 0
		if n.Direction == nextDirection {
			steps = n.Steps + 1
		} else if n.Steps < 3 {
			continue
		}
		node := Node{
			n.X + nextDirection[0],
			n.Y + nextDirection[1],
			nextDirection,
			steps,
			n.HeatLoss,
		}
		if node.IsValidV2(len(graph), len(graph[0])) {
			node.HeatLoss += graph[node.X][node.Y]
			result = append(result, node)
		}
	}
	return
}

type Nodes []Node

func (n *Nodes) Len() int {
	return len(*n)
}

func (n *Nodes) Less(i, j int) bool {
	pq := *n
	return pq[i].HeatLoss < pq[j].HeatLoss
}

func (n *Nodes) Swap(i, j int) {
	pq := *n
	pq[i], pq[j] = pq[j], pq[i]
}

func (n *Nodes) Push(x any) {
	*n = append(*n, x.(Node))
}

func (n *Nodes) Pop() any {
	old := *n
	l := len(old)
	item := old[l-1]
	*n = old[0 : l-1]
	return item
}

func PartOne(input [][]int) int {
	N := len(input)
	visited := make(map[[5]int]struct{})
	pq := new(Nodes)
	heap.Push(pq, Node{0, 0, Down, 0, 0})
	//heap.Push(pq, Node{0, 0, Right, 0, 0})

	result := 500
	stepNum := 0
	for pq.Len() > 0 {
		stepNum += 1
		if stepNum%1000 == 0 {
			log.Println("stepNum", stepNum)
		}
		node := heap.Pop(pq).(Node)
		key := [5]int{node.X, node.Y, node.Direction[0], node.Direction[1], node.Steps}
		_, ok := visited[key]
		if ok {
			continue
		}
		visited[key] = struct{}{}

		if node.X == N-1 && node.Y == N-1 {
			result = node.HeatLoss
			break
		}

		for _, nxt := range node.Next(input) {
			heap.Push(pq, nxt)
		}
	}
	log.Println(result, pq.Len(), len(visited), stepNum)

	//log.Println(visited)
	return result
}

func parseInput(input []string) (result [][]int) {
	for _, line := range input {
		loc := make([]int, 0)
		for _, val := range line {
			loc = append(loc, aoc2023.Atoi(string([]rune{val})))
		}
		result = append(result, loc)
	}
	return
}

func PartTwo(input [][]int) int {
	N := len(input)
	M := len(input[0])
	visited := make(map[[5]int]struct{})
	pq := new(Nodes)
	heap.Push(pq, Node{0, 0, Down, 0, 0})
	heap.Push(pq, Node{0, 0, Right, 0, 0})

	result := 5000
	stepNum := 0
	for pq.Len() > 0 {
		stepNum += 1
		if stepNum%1000 == 0 {
			log.Println("stepNum", stepNum)
		}
		node := heap.Pop(pq).(Node)
		key := [5]int{node.X, node.Y, node.Direction[0], node.Direction[1], node.Steps}
		_, ok := visited[key]
		if ok {
			continue
		}
		visited[key] = struct{}{}

		if node.X == N-1 && node.Y == M-1 && node.Steps > 2 {
			result = node.HeatLoss
			break
		}

		nodes := node.NextV2(input)
		for _, nxt := range nodes {
			heap.Push(pq, nxt)
		}
	}
	//log.Println(visited)
	// 839
	log.Println(result, pq.Len(), len(visited), stepNum)

	return result
}

func main() {
	input := parseInput(aoc2023.ReadLines("day17/input"))
	//log.Println(PartOne(input))
	log.Println(PartTwo(input))
}
