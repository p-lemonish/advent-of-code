package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/1

	//pathStr := "example.txt"
	pathStr := "chall.txt"
	lines := getData(pathStr)
	fmt.Println("--- Part 1 ---")
	partOne(lines)
	fmt.Println("\n--- Part 2 ---")
	partTwo(lines)
}

func partOne(lines []string) {
	// count++ every time the dial ends at 0
	dial := 50
	count := 0
	for _, line := range lines {
		if line != "" {
			side := line[:1]
			distance, _ := strconv.Atoi(line[1:])
			//fmt.Printf("Turning to %s for distance %d\n", side, distance)
			if side == "L" {
				dial = (dial - distance + 100) % 100
			} else if side == "R" {
				dial = (dial + distance + 100) % 100
			}
			//fmt.Printf("Dial is at %d\n", dial)
			if dial == 0 {
				count++
			}
		}
	}
	fmt.Printf("Password is %d\n", count)
}

func partTwo(lines []string) {
	dial := 50
	count := 0
	for _, line := range lines {
		if line != "" {
			side := line[:1]
			distance, _ := strconv.Atoi(line[1:])
			//fmt.Printf("\nDial is at %d, turning %s by %d\n", dial, side, distance)
			if side == "L" {
				// wow very elegant
				for range distance {
					dial--
					dial = (dial + 100) % 100
					if dial == 0 {
						count++
					}
				}
			} else if side == "R" {
				for range distance {
					dial++
					dial = (dial + 100) % 100
					if dial == 0 {
						count++
					}
				}
			}
		}
	}
	fmt.Printf("Password is %d\n", count)
}

func getData(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("No file %s found", path)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
