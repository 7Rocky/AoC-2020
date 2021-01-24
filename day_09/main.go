package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func isSumOfPreviousNumbers(number int, previousNumbers []int) bool {
	sortSlice(previousNumbers)

	left, right := 0, len(previousNumbers)-1

	for left < right {
		if previousNumbers[left]+previousNumbers[right] < number {
			left++
		} else if previousNumbers[left]+previousNumbers[right] > number {
			right--
		} else {
			return true
		}
	}

	return false
}

func sum(arr []int) int {
	sum := 0

	for _, s := range arr {
		sum += s
	}

	return sum
}

func sortSlice(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func findEncryptionWeakness(invalidNumber int, numbers []int) int {
	left, right := 0, 2

	for left < right-1 && right < len(numbers) {
		set := make([]int, right-left)
		copy(set, numbers[left:right])

		if sum(set) < invalidNumber {
			right++
		} else if sum(set) > invalidNumber {
			left++
		} else {
			sortSlice(set)

			return set[0] + set[len(set)-1]
		}
	}

	return -1
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	var invalidNumber int

	for i, n := range numbers {
		if i >= 25 {
			previousNumbers := make([]int, 25)
			copy(previousNumbers, numbers[i-25:i])

			if !isSumOfPreviousNumbers(n, previousNumbers) {
				invalidNumber = n
				break
			}
		}
	}

	fmt.Printf("Invalid number (1): %d\n", invalidNumber)
	fmt.Printf("Encryption weakness (2): %d\n", findEncryptionWeakness(invalidNumber, numbers))
}
