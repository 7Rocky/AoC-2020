package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const nBits int = 36

var mask string

var mem1 = map[int]int{}
var mem2 = map[int]int{}

func decToBin(number, length int) []int {
	var remainder int
	remainders := make([]int, length)

	for i := 0; i < length; i++ {
		number, remainder = number/2, number%2
		remainders[length-i-1] = remainder
	}

	return remainders
}

func binToDec(bits []int) int {
	number := 0

	for i := 0; i < len(bits); i++ {
		number += bits[i] * int(math.Pow(2, float64(len(bits)-i-1)))
	}

	return number
}

func applyMask(number, level int) []int {
	bin := decToBin(number, nBits)

	if level == 1 {
		for i := range mask {
			if mask[i] != 'X' {
				bin[i], _ = strconv.Atoi(string(mask[i]))
			}
		}

		return []int{binToDec(bin)}
	}

	binStr := make([]byte, nBits)

	for i := range mask {
		binStr[i] = strconv.Itoa(bin[i])[0]

		if mask[i] != '0' {
			binStr[i] = mask[i]
		}
	}

	return findMaskCombinations(binStr)
}

func getBinaryTruthTable(length int) [][]int {
	truthTable := make([][]int, int(math.Pow(2, float64(length))))

	for i := 0; i < len(truthTable); i++ {
		truthTable[i] = decToBin(i, length)
	}

	return truthTable
}

func findMaskCombinations(stream []byte) []int {
	numFloating := strings.Count(string(stream), "X")

	possibilities := getBinaryTruthTable(numFloating)

	var combinations []int
	bin := make([]int, nBits)

	for _, p := range possibilities {
		j := 0

		for i, b := range stream {
			if stream[i] == 'X' {
				bin[i] = p[j]
				j++
			} else {
				bin[i], _ = strconv.Atoi(string(b))
			}
		}

		combinations = append(combinations, binToDec(bin))
	}

	return combinations
}

func executeInstruction(instruction string) {
	if strings.Contains(instruction, "mem") {
		splitted := strings.Split(instruction, " = ")
		index, _ := strconv.Atoi(splitted[0][4 : len(splitted[0])-1])
		value, _ := strconv.Atoi(splitted[1])

		mem1[index] = applyMask(value, 1)[0]

		indeces := applyMask(index, 2)

		for _, i := range indeces {
			mem2[i] = value
		}
	} else if strings.Contains(instruction, "mask") {
		mask = instruction[7:]
	}
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		executeInstruction(scanner.Text())
	}

	sum := 0

	for _, m := range mem1 {
		sum += m
	}

	fmt.Printf("Sum of memory values (1): %d\n", sum)

	sum = 0

	for _, m := range mem2 {
		sum += m
	}

	fmt.Printf("Sum of memory values (2): %d\n", sum)
}
