package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/7

	pathStr := os.Args[1]
	m := getData(pathStr)
	fmt.Println("--- Part 1 ---")
	partOne(m)
	fmt.Println("\n--- Part 2 ---")
	partTwo(m)
}

func partOne(m [][]string) {
	var sRow int
	var sCol int
	maxRows := len(m)
	splitCount := 0
	for rIdx, r := range m {
		for cIdx, c := range r {
			if rIdx == 0 {
				if c == "S" {
					sRow = rIdx
					sCol = cIdx
					m[sRow+1][sCol] = "|"
				}
			} else if rIdx < maxRows-1 {
				if c == "|" {
					if m[rIdx+1][cIdx] == "^" {
						splitCount++
						m[rIdx+1][cIdx-1] = "|"
						m[rIdx+1][cIdx+1] = "|"
					} else {
						m[rIdx+1][cIdx] = "|"
					}
				}
			}
			fmt.Printf("%s", c)
		}
		fmt.Println()
	}
	fmt.Printf("total splits %d\n", splitCount)
}

func partTwo(m [][]string) {
	var sRow int
	var sCol int
	maxRows := len(m)
	pathsN := make([][]int, len(m))
	for i := range pathsN {
		pathsN[i] = make([]int, len(m[i]))
	}
	for rIdx, r := range m {
		for cIdx, c := range r {
			if rIdx == 0 {
				if c == "S" {
					sRow = rIdx
					sCol = cIdx
					m[sRow+1][sCol] = "|"
					pathsN[sRow+1][sCol]++
				}
			} else if rIdx < maxRows-1 {
				if c == "|" {
					if m[rIdx+1][cIdx] == "^" {
						m[rIdx+1][cIdx-1] = "|"
						m[rIdx+1][cIdx+1] = "|"
						pathsN[rIdx+1][cIdx-1] += pathsN[rIdx][cIdx]
						pathsN[rIdx+1][cIdx+1] += pathsN[rIdx][cIdx]
					} else {
						m[rIdx+1][cIdx] = "|"
						pathsN[rIdx+1][cIdx] += pathsN[rIdx][cIdx]
					}
				}
			}
		}
	}
	for _, l := range m {
		fmt.Println(l)
	}
	for _, l := range pathsN {
		fmt.Println(l)
	}
	sum := 0
	for _, n := range pathsN[len(pathsN)-1] {
		sum += n
	}
	fmt.Printf("sum: %d\n", sum)
}

func getData(path string) [][]string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("No file %s found", path)
	}
	lines := strings.Split(string(data), "\n")
	m := [][]string{}
	for _, l := range lines {
		if len(l) != 0 {
			row := strings.Split(l, "")
			m = append(m, row)
		}
	}
	return m
}
