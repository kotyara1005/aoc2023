package main

import (
	"aoc2023"
	"log"
	"strconv"
	"strings"
)

const DIGITS = "1234567890"
const NON_SYMBOLS = DIGITS + "."

var Diffs = [][2]int{
	{1, 1}, {1, 0}, {1, -1},
	{0, 1}, {0, -1},
	{-1, 1}, {-1, 0}, {-1, -1},
}

func hasSymbol(input []string, x int, y int) bool {
	N := len(input)
	M := len(input[0])
	for _, vals := range Diffs {
		dx, dy := vals[0], vals[1]
		nx := x + dx
		ny := y + dy

		if !(0 <= nx && nx < N && 0 <= ny && ny < M) {
			continue
		}

		if !strings.ContainsRune(NON_SYMBOLS, rune(input[nx][ny])) {
			return true
		}
	}
	return false
}

func PartOne(input []string) (result int) {
	for x, line := range input {
		buf := make([]rune, 0, 0)
		isPart := false
		for y, r := range []rune(line) {
			if strings.ContainsRune(DIGITS, r) {
				buf = append(buf, r)
				isPart = isPart || hasSymbol(input, x, y)
			} else {
				//log.Println(x, y, isPart, buf)
				if isPart {
					num, err := strconv.Atoi(string(buf))
					if err != nil {
						log.Fatalln(err)
					}
					result += num
				}
				buf = make([]rune, 0, 0)
				isPart = false
			}
		}
		if isPart {
			num, err := strconv.Atoi(string(buf))
			if err != nil {
				log.Fatalln(err)
			}
			result += num
		}
	}
	return
}

func findGears(input []string, x int, y int) [][2]int {
	N := len(input)
	M := len(input[0])
	result := make([][2]int, 0, 0)
	for _, vals := range Diffs {
		dx, dy := vals[0], vals[1]
		nx := x + dx
		ny := y + dy

		if !(0 <= nx && nx < N && 0 <= ny && ny < M) {
			continue
		}

		if rune(input[nx][ny]) == '*' {
			result = append(result, [2]int{nx, ny})
		}
	}
	return result
}

type Gear struct {
	Count int
	Ratio int
}

func addGears(gears map[[2]int]*Gear, localGears [][2]int, num int) {
	uniqGears := make(map[[2]int]struct{})
	for _, gear := range localGears {
		uniqGears[gear] = struct{}{}
	}

	for gear, _ := range uniqGears {
		_, prs := gears[gear]
		if !prs {
			gears[gear] = &Gear{0, 1}
		}
		gears[gear].Count += 1
		gears[gear].Ratio *= num
	}
}

func PartTwo(input []string) (result int) {
	gears := make(map[[2]int]*Gear)
	for x, line := range input {
		buf := make([]rune, 0, 0)
		localGears := make([][2]int, 0, 0)
		for y, r := range []rune(line) {
			if strings.ContainsRune(DIGITS, r) {
				buf = append(buf, r)
				localGears = append(localGears, findGears(input, x, y)...)
			} else {
				//log.Println(x, y, isPart, buf)
				if len(buf) > 0 {
					num, err := strconv.Atoi(string(buf))
					if err != nil {
						log.Fatalln(err)
					}
					addGears(gears, localGears, num)
				}
				buf = make([]rune, 0, 0)
				localGears = make([][2]int, 0, 0)
			}
		}
		if len(buf) > 0 {
			num, err := strconv.Atoi(string(buf))
			if err != nil {
				log.Fatalln(err)
			}
			addGears(gears, localGears, num)
		}
	}

	for _, gear := range gears {
		if gear.Count == 2 {
			result += gear.Ratio
		}
	}
	return
}

func main() {
	input := aoc2023.ReadLines("day3/input")

	//log.Println("PartOne", PartOne(input))
	log.Println("PartTwo", PartTwo(input))
}
