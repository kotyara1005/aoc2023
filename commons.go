package aoc2023

import (
	"io"
	"log"
	"os"
	"strings"
)

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
