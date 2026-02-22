package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/4

	//pathStr := "example.txt"
	pathStr := "chall.txt"
	lines := getData(pathStr)
	m := makeMatrix(lines)
	fmt.Println("--- Part 1 ---")
	nm, c := partOne(m, false)
	fmt.Printf("Count: %d", c)
	fmt.Println("\n--- Part 2 ---")
	partTwo(nm, c)
}

func makeMatrix(lines []string) [][]byte {
	matrix := make([][]byte, len(lines))
	for i, l := range lines {
		matrix[i] = []byte(l)
	}
	return matrix
}

func partOne(matrix [][]byte, ignoreX bool) ([][]byte, int) {
	count := 0
	newMatrix := matrix
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if ignoreX {
				if matrix[i][j] == 64 {
					newMatrix = markAdjacent(newMatrix, i, j, ignoreX)
					if newMatrix[i][j] == 120 {
						count++
					}
				}
			} else {
				if matrix[i][j] == 64 || matrix[i][j] == 120 {
					newMatrix = markAdjacent(newMatrix, i, j, ignoreX)
					if newMatrix[i][j] == 120 {
						count++
					}
				}
			}
		}
	}
	return newMatrix, count
}

func markAdjacent(m [][]byte, i, j int, ignoreX bool) [][]byte {
	newM := m
	countRolls := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			if val, err := cell(newM, i+di, j+dj); err != nil {
				// do nothing
			} else {
				if ignoreX {
					if val == 64 {
						countRolls++
					}
				} else {
					if val == 64 || val == 120 {
						countRolls++
					}
				}
			}
		}
	}
	if countRolls < 4 {
		newM[i][j] = 0x78
	} else {
	}
	return newM
}

func cell(m [][]byte, i, j int) (byte, error) {
	if i < 0 || i >= len(m) {
		return 0, fmt.Errorf("i %d out of range", i)
	}
	if j < 0 || j >= len(m[i]) {
		return 0, fmt.Errorf("j %d out of range", j)
	}
	return m[i][j], nil
}

func partTwo(nm [][]byte, c int) {
	count := c
	prevCount := 0
	newM := nm
	diff := count - prevCount
	for diff != 0 {
		newMa, ca := partOne(newM, true)
		newM = newMa
		prevCount = count
		count += ca
		diff = count - prevCount
	}
	fmt.Printf("Count: %d\n", count)
}

func getData(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("No file %s found", path)
	}
	cleaned := strings.TrimSpace(string(data))
	lines := strings.Split(string(cleaned), "\n")
	return lines
}
