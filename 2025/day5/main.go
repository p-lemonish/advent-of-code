package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// source: https://adventofcode.com/2025/day/5

	pathStr := "example.txt"
	//pathStr := "chall.txt"
	rStr, idStr := getData(pathStr)
	ranges, err := rangesToInts(rStr)
	if err != nil {
		panic("rangesToInts err")
	}
	ids, err := idsToInts(idStr)
	if err != nil {
		panic("idsToInts err")
	}
	fmt.Println("--- Part 1 ---")
	partOne(ranges, ids)
	fmt.Println("\n--- Part 2 ---")
	c := partTwo(ranges)
	fmt.Printf("Count: %d\n", c)
}

func partOne(ranges [][]int, ids []int) {
	count := 0
	for _, i := range ids {
		for _, r := range ranges {
			if i <= r[1] && i >= r[0] {
				count++
				break
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func partTwo(ranges [][]int) int {
	count := 0
	prevMax := 0
	prevMin := 0
	for i, r := range ranges {
		//fmt.Printf("range: %d\n", r)
		if i == 0 {
			prevMax = r[1]
			prevMin = r[0]
			continue
		}
		if r[0] > prevMax {
			//fmt.Printf("start of new independent range (%d, %d)\n", r[0], r[1])
			count += prevMax - prevMin + 1
			//fmt.Printf("counting up: (%d, %d) -> count: %d\n", prevMin, prevMax, count)
			prevMax = r[1]
			prevMin = r[0]
			continue
		}
		if r[0] <= prevMax && prevMax < r[1] {
			//fmt.Printf("in an overlapping range (%d, %d)\nSetting %d as prevMax\n", prevMin, prevMax, r[1])
			prevMax = r[1]
			continue
		}
		//fmt.Printf("The current range (%d, %d) is completely within the prevMin, prevMax range\n", r[0], r[1])
	}
	count += prevMax - prevMin + 1
	return count
}

func getData(path string) ([]string, []string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("No file %s found", path)
	}
	cleaned := strings.TrimSpace(string(data))
	lines := strings.SplitN(string(cleaned), "\n\n", 2)
	ranges := strings.Fields(lines[0])
	ids := strings.Fields(lines[1])
	return ranges, ids
}

func idsToInts(ids []string) ([]int, error) {
	ints := make([]int, len(ids))
	for i, id := range ids {
		val, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		ints[i] = val
	}
	sort.Ints(ints)
	return ints, nil
}

func rangesToInts(ranges []string) ([][]int, error) {
	ints := make([][]int, len(ranges))
	for i, r := range ranges {
		rVals := strings.Split(r, "-")
		iVals := make([]int, len(rVals))
		for j, rv := range rVals {
			val, err := strconv.Atoi(rv)
			if err != nil {
				return nil, err
			}
			iVals[j] = val
		}
		ints[i] = iVals
	}
	sort.Slice(ints, func(i, j int) bool { return ints[i][0] < ints[j][0] })
	return ints, nil
}
