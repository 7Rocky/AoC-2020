package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getSeatID(seat string) int {
	rows, colums := seat[:7], seat[7:]

	rowsRange := []int{0, 127}

	for _, r := range rows {
		if r == 'F' {
			rowsRange[1] = (rowsRange[1] + rowsRange[0]) / 2
		} else if r == 'B' {
			rowsRange[0] = (rowsRange[1]+rowsRange[0])/2 + 1
		}
	}

	columnsRange := []int{0, 7}

	for _, c := range colums {
		if c == 'L' {
			columnsRange[1] = (columnsRange[1] + columnsRange[0]) / 2
		} else if c == 'R' {
			columnsRange[0] = (columnsRange[1]+columnsRange[0])/2 + 1
		}
	}

	return rowsRange[0]*8 + columnsRange[0]
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ids []int

	for scanner.Scan() {
		seat := scanner.Text()
		ids = append(ids, getSeatID(seat))
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	fmt.Printf("Maximum sear ID (1): %d\n", ids[len(ids)-1])

	for i, id := range ids {
		if i > 0 && id != ids[i-1]+1 {
			fmt.Printf("Missing ID (2): %d\n", id-1)
			break
		}
	}
}
