package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	target := 2020

	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	sum := 0
	left := 0
	right := len(numbers) - 1

	for sum != target {
		sum = numbers[left] + numbers[right]

		if sum > target {
			right--
		}

		if sum < target {
			left++
		}
	}

	product := numbers[left] * numbers[right]

	fmt.Printf("RESULT (1): %d * %d = %d\n",
		numbers[left], numbers[right], product)

	sum = 0
	left = 0
	right = len(numbers) - 1

	var middle int

	for middle = left + 1; middle < right-1; middle++ {
		for left < right {
			sum = numbers[left] + numbers[middle] + numbers[right]

			if sum > target && middle < right-1 {
				right--
			}

			if sum < target && middle > left+1 {
				left++
			}

			if sum == target {
				break
			}
		}

		if sum == target {
			break
		}
	}

	product = numbers[left] * numbers[middle] * numbers[right]

	fmt.Printf("RESULT (2): %d * %d * %d = %d\n",
		numbers[left], numbers[middle], numbers[right], product)
}
