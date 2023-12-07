package aoc2023

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Number = regexp.MustCompile(`\d+`)

func FindAllNumbers(line string) []int {
	result := make([]int, 0, 0)
	for _, val := range Number.FindAllString(line, -1) {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}
	return result
}

func ReadBytes(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
