package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/3

	//pathStr := "example.txt"
	pathStr := "chall.txt"
	lines := getData(pathStr)
	fmt.Println("--- Part 1 ---")
	partOne(lines)
	fmt.Println("\n--- Part 2 ---")
	partTwo(lines)
}

func partOne(lines []string) {
	sum := 0
	batteries := 2
	for _, line := range lines {
		if line != "" {
			joltages := make([]int, batteries)
			digits := splitN(line, 1)
			for i := range batteries {
				size := len(digits)
				foundMax, foundIdx := findMax(digits[:size-(batteries-1-i)])
				digits = digits[foundIdx+1:]
				joltages[i] = foundMax
			}
			maxJoltage := concat(joltages)
			sum += maxJoltage
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func concat(values []int) int {
	s := ""
	for _, v := range values {
		a := strconv.Itoa(v)
		s = fmt.Sprintf("%s%s", s, a)
	}
	i, _ := strconv.Atoi(s)
	return i
}

func findMax(arr []int) (int, int) {
	idx := 0
	m := arr[idx]
	for i, v := range arr[1:] {
		if v > m {
			m = v
			idx = i + 1
		}
	}
	return m, idx
}

// s = "123456", n = 3 -> [123, 456]
func splitN(s string, n int) []int {
	parts := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		parts[i] = num
	}
	return parts
}

func partTwo(lines []string) {
	sum := 0
	batteries := 12
	for _, line := range lines {
		if line != "" {
			joltages := make([]int, batteries)
			digits := splitN(line, 1)
			for i := range batteries {
				size := len(digits)
				foundMax, foundIdx := findMax(digits[:size-(batteries-1-i)])
				digits = digits[foundIdx+1:]
				joltages[i] = foundMax
			}
			maxJoltage := concat(joltages)
			sum += maxJoltage
		}
	}
	fmt.Printf("Sum: %d\n", sum)
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
