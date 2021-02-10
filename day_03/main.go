package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTrees(treeMap []string, slope [2]int) int {
	height := len(treeMap)
	width := len(treeMap[0])

	position := [2]int{0, 0}

	numTrees := 0

	for position[1] < height {
		if treeMap[position[1]][position[0]] == '#' {
			numTrees++
		}

		position[0] = (position[0] + slope[0]) % width
		position[1] += slope[1]
	}

	return numTrees
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var treeMap []string

	for scanner.Scan() {
		treeMap = append(treeMap, scanner.Text())
	}

	result := countTrees(treeMap, [2]int{3, 1})

	fmt.Printf("Number of trees (1): %d\n", result)

	product := countTrees(treeMap, [2]int{1, 1}) *
		result *
		countTrees(treeMap, [2]int{5, 1}) *
		countTrees(treeMap, [2]int{7, 1}) *
		countTrees(treeMap, [2]int{1, 2})

	fmt.Printf("Number of trees (2): %d\n", product)
}
