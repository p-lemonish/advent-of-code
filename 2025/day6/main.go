package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/6

	//pathStr := "example.txt"
	pathStr := "chall.txt"
	m := getData(pathStr)
	fmt.Println("--- Part 1 ---")
	partOne(m)
	fmt.Println("\n--- Part 2 ---")
	partTwo(m)
}

func result(m [][]string) {
	sum := 0
	for _, l := range m {
		operator := l[len(l)-1]
		if operator == "*" {
			res := 0
			for i, vStr := range l {
				v, err := strconv.Atoi(vStr)
				if err != nil {
					continue
				}
				if i == 0 {
					res += v
					continue
				}
				res *= v
			}
			sum += res
		} else if operator == "+" {
			res := 0
			for _, vStr := range l {
				v, err := strconv.Atoi(vStr)
				if err != nil {
					continue
				}
				res += v
			}
			sum += res
		}
	}
	fmt.Printf("sum: %d\n", sum)
}

func partOne(m [][]string) {
	maxCols := 0
	for _, row := range m {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	mNew := make([][]string, maxCols)
	for i := range maxCols {
		lNew := []string{}
		for j := range m {
			if i < len(m[j]) {
				lNew = append(lNew, strings.TrimSpace(m[j][i]))
			}
		}
		mNew[i] = lNew
	}
	result(mNew)
}

func partTwo(m [][]string) {
	rows := len(m)
	rowStrs := make([]string, rows)
	for i := range m {
		rowStrs[i] = strings.Join(m[i], "")
	}
	cols := len(rowStrs[0])

	mNew := [][]string{}
	curRow := []string{}
	curOp := ""

	for col := cols - 1; col >= 0; col-- {
		newNumStr := ""
		for row := range rows {
			if col >= len(rowStrs[row]) {
				continue
			}
			char := rowStrs[row][col]
			if char == '*' || char == '+' || char == ' ' {
				continue
			}
			newNumStr += string(char)
		}
		if col < len(rowStrs[rows-1]) {
			opChar := rowStrs[rows-1][col]
			if opChar == '*' || opChar == '+' {
				curOp = string(opChar)
			}
		}
		if strings.TrimSpace(newNumStr) != "" {
			curRow = append(curRow, newNumStr)
		}
		isSeparator := strings.TrimSpace(newNumStr) == ""
		isLastCol := col == 0

		if (isSeparator || isLastCol) && len(curRow) > 0 {
			if curOp != "" {
				curRow = append(curRow, curOp)
			}
			mNew = append(mNew, curRow)
			curRow = []string{}
			curOp = ""
		}
	}
	result(mNew)
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
			re := regexp.MustCompile(`\s*(\d+|[*+])\s*`)
			row := re.FindAllString(l, -1)
			m = append(m, row)
		}
	}
	return m
}
