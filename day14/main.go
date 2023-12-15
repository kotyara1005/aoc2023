package main

import (
	"aoc2023"
	"crypto/md5"
	"hash/crc32"
	"log"
	"strings"
)

func PartOne(input [][]rune) int {
	//N := len(input)
	//M := len(input[0])
	////log.Println(N, M)
	//result := 0
	//
	//for j := 0; j < M; j += 1 {
	//	start := 0
	//	num := 0
	//
	//	for i := 0; i < N; i += 1 {
	//		//log.Println(rune(input[i][j]), start, num)
	//		if input[i][j] == '#' {
	//			for t := 0; t < num; t += 1 {
	//				result += N - start - t
	//			}
	//			start = i + 1
	//			num = 0
	//		} else if input[i][j] == 'O' {
	//			num += 1
	//		}
	//
	//	}
	//
	//	for t := 0; t < num; t += 1 {
	//		result += N - start - t
	//	}
	//}
	Rotate(input)
	RollLeft(input)
	RotateRight(input)
	//PrintInput(input)
	return CountResult(input)
}

func CountResult(input [][]rune) int {
	result := 0
	for i, line := range input {
		result += strings.Count(string(line), "O") * (len(input) - i)
	}
	return result
}

func RollLeft(input [][]rune) {
	for _, line := range input {
		empty := make([]int, 0)
		for j, val := range line {
			switch val {
			case '#':
				empty = nil
			case 'O':
				if len(empty) > 0 {
					line[j] = '.'
					e := empty[0]
					empty = append(empty[1:], j)
					line[e] = 'O'
				}
			case '.':
				empty = append(empty, j)
			}
		}
	}
}

func PrintInput(input [][]rune) {
	//log.Println("\n===============================================\n")
	for _, line := range input {
		log.Println(string(line))
	}
	log.Println("\n===============================================\n")
}

func Rotate(input [][]rune) {
	// 10 x 10
	// (0, 0) <= (0, 9) <= (9, 9) <= (9, 0)
	// (0, 1) <= (1, 9) <= (9, 8) <= (9, 0)
	// (i, j) <= (j, N-i) <= (N-i, N-j) <= (N-j, i)
	N := len(input)

	for i := 0; i < N/2; i += 1 {
		for j := i; j < N-1-i; j += 1 {
			buf := input[i][j]
			//db := []rune{
			//	input[i][j],
			//	input[j][N-1-i],
			//	input[N-1-i][N-1-j],
			//	input[N-1-j][i],
			//}
			//log.Println(db)
			input[i][j] = input[j][N-1-i]
			input[j][N-1-i] = input[N-1-i][N-1-j]
			input[N-1-i][N-1-j] = input[N-1-j][i]
			input[N-1-j][i] = buf
			//break
			//db = []rune{
			//	input[i][j],
			//	input[i][N-1-j],
			//	input[N-1-i][N-1-j],
			//	input[N-1-i][j],
			//}
			//log.Println(db)

			//break
		}
		//break
	}
}

func RotateRight(input [][]rune) {
	Rotate(input)
	Rotate(input)
	Rotate(input)
}

func ApplyCycle(input [][]rune) {
	//Rotate(input)
	RollLeft(input)

	RotateRight(input)
	RollLeft(input)

	RotateRight(input)
	RollLeft(input)

	RotateRight(input)
	RollLeft(input)
	RotateRight(input)
}

func HashInput(input [][]rune) uint32 {
	N := len(input)
	buf := make([]byte, N*N)

	for i := 0; i < N; i += 1 {
		copy(buf[i*N:], string(input[i]))
	}

	v := md5.Sum(buf)
	return crc32.ChecksumIEEE(v[:])
}

func PartTwo(input [][]rune) int {
	//PrintInput(input)
	//PrintInput(input)
	//RotateRight(input)
	PrintInput(input)
	Rotate(input)
	cache := map[uint32]int{}
	for i := 0; i < 122+(1000000000-122)%(148-122); i += 1 {
		if i%1000000 == 0 {
			log.Println(i)
		}
		key := HashInput(input)
		v, ok := cache[key]
		if ok {
			log.Println("Cycle found", v, i)
			//break
		}
		ApplyCycle(input)
		cache[key] = i

		RotateRight(input)
		log.Println(i, CountResult(input), key)
		Rotate(input)
	}
	//RotateRight(input)
	RotateRight(input)
	//PrintInput(input)

	//log.Println([]byte([]rune{'q'}))
	return CountResult(input)
}

func main() {
	input := make([][]rune, 0)
	for _, line := range aoc2023.ReadLines("day14/input") {
		input = append(input, []rune(line))
	}

	//start := time.Now()
	//log.Println(PartOne(input))
	//log.Println(time.Now().Sub(start))

	//log.Println(CountResult(input))
	log.Println(PartTwo(input))
}
