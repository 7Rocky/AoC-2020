package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	values              []int
	satisfiedConditions []int
}

type valueRange struct {
	min, max int
}

func isInvalid(value int, ranges []valueRange) int {
	for _, r := range ranges {
		if r.min <= value && value <= r.max {
			return 0
		}
	}

	return value
}

func intersect(arr1, arr2 []int) []int {
	var intersection []int

	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}

	for _, a1 := range arr1 {
		for _, a2 := range arr2 {
			if a1 == a2 {
				intersection = append(intersection, a1)
			}
		}
	}

	return intersection
}

func matchFields(value, column int, columns map[int][]int, ranges []valueRange) {
	var validColumns []int

	for i, r := range ranges {
		if r.min <= value && value <= r.max {
			validColumns = append(validColumns, i/2)
		}
	}

	columns[column] = intersect(columns[column], validColumns)
}

func indexOf(number int, arr []int) int {
	for i, n := range arr {
		if n == number {
			return i
		}
	}

	return -1
}

func keyOf(number int, columns map[int][]int) int {
	for k, v := range columns {
		if v[0] == number {
			return k
		}
	}

	return -1
}

func deleteColumnFromOthers(column int, columns map[int][]int) {
	for c, m := range columns {
		if len(m) > 1 {
			index := indexOf(column, m)

			if index != -1 {
				newMatch := []int{}
				newMatch = append(newMatch, m[:index]...)
				newMatch = append(newMatch, m[index+1:]...)

				columns[c] = newMatch
			}
		}
	}
}

func allMatchesLengthEquals(number int, columns map[int][]int) bool {
	for _, m := range columns {
		if len(m) != number {
			return false
		}
	}

	return true
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numBlankLines := 0

	var ranges []valueRange
	var yourTicket ticket = ticket{[]int{}, []int{}}
	var nearbyTickets []ticket
	var fields []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			numBlankLines++
			continue
		}

		switch numBlankLines {
		case 0:
			splitted := strings.Split(line, " ")
			fields = append(fields, strings.Split(line, ":")[0])

			var indeces []int

			if strings.Count(line, " ") == 4 {
				indeces = []int{2, 4}
			} else {
				indeces = []int{1, 3}
			}

			for _, i := range indeces {
				rangeStr := strings.Split(splitted[i], "-")

				min, _ := strconv.Atoi(rangeStr[0])
				max, _ := strconv.Atoi(rangeStr[1])

				ranges = append(ranges, valueRange{min, max})
			}

		case 1:
			if line != "your ticket:" {
				for _, s := range strings.Split(line, ",") {
					n, _ := strconv.Atoi(s)
					yourTicket.values = append(yourTicket.values, n)
				}
			}
		case 2:
			if line != "your ticket:" {
				values := []int{}

				for _, s := range strings.Split(line, ",") {
					n, _ := strconv.Atoi(s)
					values = append(values, n)
				}

				nearbyTickets = append(nearbyTickets, ticket{values, []int{}})
			}
		}
	}

	scanningErrorRate := 0

	for _, t := range nearbyTickets {
		for _, v := range t.values {
			scanningErrorRate += isInvalid(v, ranges)
		}
	}

	fmt.Printf("Scanning Error Rate (1): %d\n", scanningErrorRate)

	var correctTickets []ticket

	for _, t := range nearbyTickets {
		for _, v := range t.values {
			if isInvalid(v, ranges) == 0 {
				correctTickets = append(correctTickets, t)
			}
		}
	}

	columns := map[int][]int{}

	for _, t := range nearbyTickets {
		for i, v := range t.values {
			matchFields(v, i, columns, ranges)
		}
	}

	checked := map[int]bool{}

	for !allMatchesLengthEquals(1, columns) {
		for c, m := range columns {
			if len(m) == 1 && !checked[c] {
				checked[c] = true
				deleteColumnFromOthers(m[0], columns)
				break
			}
		}
	}

	var departureIndeces []int

	for i, f := range fields {
		if strings.Contains(f, "departure") {
			departureIndeces = append(departureIndeces, i)
		}
	}

	product := 1

	for _, d := range departureIndeces {
		index := keyOf(d, columns)
		product *= yourTicket.values[index]
	}

	fmt.Printf("Product of departure values of your ticket (2): %d\n", product)
}
