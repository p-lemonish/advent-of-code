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
	for _, l := range m {
		fmt.Println(l)
	}
	mNew := make([][]string, len(m))
	for i := range m {
		lNew := make([]string, len(m[i]))
		for j := range m[i] {
			lNew[j] = strings.TrimSpace(m[j][i])
		}
		mNew[i] = lNew
	}
	result(mNew)
}

func partTwo(m [][]string) {
	for _, l := range m {
		fmt.Println(l)
	}

	rows := len(m)
	rowStrs := make([]string, rows)
	for i := range m {
		rowStrs[i] = strings.Join(m[i], "")
	}
	cols := len(rowStrs[0])

	mNew := [][]string{}
	curRow := []string{}
	curOp := ""

	for col := range cols {
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
		isLastCol := col == cols-1

		if (isSeparator || isLastCol) && len(curRow) > 0 {
			if curOp != "" {
				curRow = append(curRow, curOp)
			}
			mNew = append(mNew, curRow)
			curRow = []string{}
			curOp = ""
		}
	}
	for _, l := range mNew {
		fmt.Println(l)
	}
	result(mNew)
}

func getData(path string) [][]string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("No file %s found", path)
	}
	lines := strings.Split(string(data), "\n")
	fmt.Printf("lines from data: %s\n", lines)
	m := make([][]string, len(lines)-1)
	for i, l := range lines {
		if len(l) != 0 {
			re := regexp.MustCompile(`\s*(\d+|[*+])\s*`)
			m[i] = re.FindAllString(l, -1)
			fmt.Printf("adding to matrix: %s\n", m[i])
		}
	}
	return m
}
