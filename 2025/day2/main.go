package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/2

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
	for _, line := range lines {
		if line != "" {
			parts := strings.Split(string(line), "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			for i := start; i <= end; i++ {
				if isInvalidPartOne(i) {
					sum += i
				}
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func isInvalidPartOne(i int) bool {
	iStr := strconv.Itoa(i)
	iLen := len(iStr)
	if iLen%2 != 0 {
		return false
	}
	iHalf := int(iLen / 2)
	iSplitOne := iStr[:iHalf]
	iSplitTwo := iStr[iHalf:]
	if iSplitOne == iSplitTwo {
		return true
	}
	return false
}

func isInvalidPartTwo(i int) bool {
	iStr := strconv.Itoa(i)
	iLen := len(iStr)
	lenHalf := int(iLen / 2)
	for j := range lenHalf {
		if iLen%(j+1) == 0 {
			parts := splitN(iStr, j+1)
			allMatch := true
			for k := range len(parts) {
				if parts[0] != parts[k] {
					allMatch = false
				}
			}
			if allMatch {
				return true
			}
		}
	}
	return false
}

func splitN(s string, n int) []string {
	size := len(s) / n
	parts := make([]string, 0, size)
	for i := 0; i < len(s); i += n {
		parts = append(parts, s[i:i+n])
	}
	return parts
}

func partTwo(lines []string) {
	sum := 0
	for _, line := range lines {
		if line != "" {
			parts := strings.Split(string(line), "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			for i := start; i <= end; i++ {
				if isInvalidPartTwo(i) {
					sum += i
				}
			}
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
	lines := strings.Split(string(cleaned), ",")
	return lines
}
