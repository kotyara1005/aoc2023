package main

import (
	"aoc2023"
	"log"
	"math/bits"
	"strings"
)

func buildInt(s string) uint {
	result := 0

	for _, val := range s {
		result = result << 1
		if val == '#' {
			result += 1
		}
	}

	return uint(result)
}

type Matrix struct {
	Rows []uint
	Cols []uint
}

func ParseMatrix(input []string) Matrix {
	var rows []uint
	var cols []uint
	for _, line := range input {
		rows = append(rows, buildInt(line))
	}

	N := len(input)
	M := len(input[0])
	for i := 0; i < M; i += 1 {
		col := make([]rune, 0)
		for j := 0; j < N; j += 1 {
			col = append(col, rune(input[j][i]))
		}
		cols = append(cols, buildInt(string(col)))
	}

	return Matrix{rows, cols}
}

func FindLongestPalindrom(arr []uint) int {
	result := 0
	for i := 0; i < len(arr)-1; i += 1 {
		left := i
		right := i + 1
		for left >= 0 && right < len(arr) && arr[left] == arr[right] {
			left -= 1
			right += 1
		}
		if left == -1 || right == len(arr) {
			return i + 1
			//result = aoc2023.Max(result, i+1)
		}
	}
	return result
}

func PartOne(input []string) int {
	matrixes := []Matrix{}
	for _, mat := range input {
		matrixes = append(matrixes, ParseMatrix(strings.Split(mat, "\n")))
	}
	result := 0
	for _, mat := range matrixes {
		a := FindLongestPalindrom(mat.Cols)
		b := FindLongestPalindrom(mat.Rows)

		result += a + 100*b
		//log.Println(i, FindLongestPalindrom(mat.Cols), FindLongestPalindrom(mat.Rows))
	}
	return result
}

func FindLongestPalindromV2(arr []uint) int {
	result := 0
	for i := 0; i < len(arr)-1; i += 1 {
		left := i
		right := i + 1
		corr := 0
		for left >= 0 && right < len(arr) && (arr[left] == arr[right] || (bits.OnesCount(arr[left]^arr[right]) <= 1 && corr == 0)) {
			corr += bits.OnesCount(arr[left] ^ arr[right])
			left -= 1
			right += 1
		}
		if corr == 1 && (left == -1 || right == len(arr)) {
			//log.Println("corr", corr, i)
			return i + 1
			//result = aoc2023.Max(result, i+1)
		}
	}
	return result
}

func PartTwo(input []string) int {
	matrixes := []Matrix{}
	for _, mat := range input {
		matrixes = append(matrixes, ParseMatrix(strings.Split(mat, "\n")))
	}
	result := 0
	for i, mat := range matrixes {
		//log.Println(strconv.FormatInt(int64(mat.Rows[0]), 2), strconv.FormatInt(int64(mat.Rows[5]), 2))
		//log.Println(mat.Rows[0], mat.Rows[5])
		//aPrev := FindLongestPalindrom(mat.Cols)
		//bPrev := FindLongestPalindrom(mat.Rows)
		a := FindLongestPalindromV2(mat.Cols)
		b := FindLongestPalindromV2(mat.Rows)

		//if a == aPrev {
		//	a = 0
		//}
		//if b == bPrev {
		//	b = 0
		//}

		result += a + 100*b
		log.Println(i, "Vertical", a, "Horizon", b)
	}
	return result
}

func main() {
	log.Println(1 ^ 3)
	input := strings.Split(string(aoc2023.ReadBytes("day13/input")), "\n\n")
	//log.Println(strconv.FormatInt(int64(buildInt("###....")), 2))
	//log.Println(PartOne(input))
	log.Println(PartTwo(input))
}
