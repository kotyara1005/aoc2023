package main

import (
	"aoc2023"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction [2]int

var (
	Up    Direction = [2]int{-1, 0}
	Down            = [2]int{1, 0}
	Left            = [2]int{0, -1}
	Right           = [2]int{0, 1}
)

var DirectionFromString = map[string]Direction{
	"U": Up,
	"D": Down,
	"L": Left,
	"R": Right,
}

type Point [2]int

func (p Point) Step(d Direction) Point {
	return Point{p[0] + d[0], p[1] + d[1]}
}

func (p Point) BigStep(d Direction, num int) Point {
	return Point{p[0] + d[0]*num, p[1] + d[1]*num}
}

type Command struct {
	Direction     Direction
	NumberOfSteps int
	Colour        string
}

func NewCommandV2(val string) *Command {
	arr := strings.Split(val, " ")

	//log.Println(arr[2], arr[2][2:7], arr[2][7:8])
	num, err := strconv.ParseInt(arr[2][2:7], 16, 32)
	if err != nil {
		log.Fatal(err)
	}
	return &Command{
		[]Direction{Right, Down, Left, Up}[aoc2023.Atoi(arr[2][7:8])],
		int(num),
		arr[2],
	}
}

func NewCommand(val string) *Command {
	arr := strings.Split(val, " ")
	return &Command{
		DirectionFromString[arr[0]],
		aoc2023.Atoi(arr[1]),
		arr[2],
	}
}

func ParseInput(input []string) (result []*Command) {
	for _, line := range input {
		result = append(result, NewCommand(line))
	}
	return result
}

func DrawLine(commands []*Command) map[Point]int {
	pos := Point{0, 0}
	visited := make(map[Point]int)
	visited[pos] = 0

	for _, command := range commands {
		if command.Direction == Up {
			visited[pos] = 1
		} else if command.Direction == Down {
			visited[pos] = -1
		}
		for i := 0; i < command.NumberOfSteps; i += 1 {
			pos = pos.Step(command.Direction)
			if _, ok := visited[pos]; ok {
				log.Println("Found collision", pos)
			}
			if command.Direction == Up {
				visited[pos] = 1
			} else if command.Direction == Down {
				visited[pos] = -1
			} else {
				visited[pos] = 0
			}
		}
	}
	return visited
}

func FindDimensions(line map[Point]int) (xMin int, xMax int, yMin int, yMax int) {
	for point, _ := range line {
		xMin = aoc2023.Min(xMin, point[0])
		xMax = aoc2023.Max(xMax, point[0])
		yMin = aoc2023.Min(yMin, point[1])
		yMax = aoc2023.Max(yMax, point[1])
	}
	return
}

func printTrench(line map[Point]int, insides map[Point]int) {
	xMin, xMax, yMin, yMax := FindDimensions(line)

	for x := xMin; x <= xMax; x += 1 {
		str := []string{}
		for y := yMin; y <= yMax+5; y += 1 {
			v, ok := line[Point{x, y}]
			if ok {
				ch := "#"
				if v == 1 {
					ch = "^"
				}
				if v == -1 {
					ch = "|"
				}
				str = append(str, ch)
			} else {
				v, ok := insides[Point{x, y}]
				if ok {
					str = append(str, strconv.Itoa(v))
				} else {

					str = append(str, ".")
				}
			}
		}
		fmt.Println(strings.Join(str, ""))
	}
}

func PartOne(commands []*Command) int {
	line := DrawLine(commands)
	xMin, xMax, yMin, yMax := FindDimensions(line)
	insides := make(map[Point]int)
	result := 0
	log.Println(xMin, xMax, yMin, yMax)
	for x := xMin; x <= xMax; x += 1 {
		if x%1000 == 0 {
			log.Println(x)
		}
		prevDir := 0
		isInside := 0
		for y := yMin; y <= yMax+5; y += 1 {
			v, ok := line[Point{x, y}]
			if ok {
				//result += 1
				if v != 0 && v != prevDir {
					isInside += v
					prevDir = v
				}
			} else if isInside != 0 {
				insides[Point{x, y}] = isInside
				result += 1
			}
		}
	}

	//PrintTrench(line, insides)
	return result + len(line)
}

func ParseInputV2(input []string) (result []*Command) {
	for _, line := range input {
		result = append(result, NewCommandV2(line))
	}
	return result
}

type Border struct {
	Pos         int
	Direction   int
	IsConnected bool
}

type Borders []Border

func (b Borders) Len() int {
	return len(b)
}

func (b Borders) Less(i, j int) bool {
	return b[i].Pos < b[j].Pos
}

func (b Borders) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func FindIntervals(commands []*Command) map[int]Borders {
	pos := Point{0, 0}
	intervals := make(map[int]Borders)
	isConnected := false
	isStart := true

	for _, command := range commands {
		//log.Println(command)
		if command.Direction == Left || command.Direction == Right {
			if !isStart && command.Direction == Right {
				intervals[pos[0]][len(intervals[pos[0]])-1].IsConnected = true
			}
			if command.Direction == Left {
				isConnected = true
			}

			pos = pos.BigStep(command.Direction, command.NumberOfSteps)
			continue
		}
		isStart = false

		if _, ok := intervals[pos[0]]; !ok {
			intervals[pos[0]] = Borders{}
		}

		if command.Direction == Up {
			intervals[pos[0]] = append(intervals[pos[0]], Border{pos[1], 1, isConnected})
		} else if command.Direction == Down {
			intervals[pos[0]] = append(intervals[pos[0]], Border{pos[1], -1, isConnected})
		} else {
			log.Fatal("?????")
		}
		isConnected = false
		for i := 0; i < command.NumberOfSteps; i += 1 {
			pos = pos.Step(command.Direction)
			//if _, ok := intervals[pos[0]]; ok {
			//	log.Println("Found collision", pos)
			//}
			if command.Direction == Up {
				intervals[pos[0]] = append(intervals[pos[0]], Border{pos[1], 1, false})
			} else if command.Direction == Down {
				intervals[pos[0]] = append(intervals[pos[0]], Border{pos[1], -1, false})
			} else {
				log.Fatal("unreachable")
			}
		}
	}
	return intervals
}

func PrintIntervals(intervals map[int]Borders) {
	start := 0
	stop := 0
	for x, _ := range intervals {
		start = min(start, x)
		stop = max(stop, x)
	}

	for x := start; x <= stop; x += 1 {
		log.Println(intervals[x])
	}
}

func FindPoints(commands []*Command) (result []Point) {
	pos := Point{0, 0}

	for _, command := range commands {
		pos = pos.BigStep(command.Direction, command.NumberOfSteps)
		result = append(result, pos)
	}
	return
}

func PolygonArea(points []Point) int {
	area := 0
	prev := len(points) - 1

	for i, point := range points {
		//area += (points[prev][0] * point[1]) - (points[prev][1] * point[0])
		area += (point[0] + points[prev][0]) * (point[1] - points[prev][1])
		prev = i
	}
	return aoc2023.Abs(area / 2)
}

func PartTwo(commands []*Command) int {
	borderLen := 0
	for _, command := range commands {
		borderLen += command.NumberOfSteps
	}

	points := FindPoints(commands)

	log.Println(points, len(points), borderLen)

	return PolygonArea(points) + borderLen/2 + 1
}

func main() {
	//  952408144115
	//  952408144115
	// 1275575224335

	//97874103728605
	commands := ParseInputV2(aoc2023.ReadLines("day18/input"))
	log.Println(PartTwo(commands))
}
