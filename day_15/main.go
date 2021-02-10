package main

import "fmt"

var lastIndeces = map[int][]int{}

func getLastPositions(arr []int) int {
	length := len(arr)
	return arr[length-1] - arr[length-2]
}

func next(numbers []int) int {
	length := len(numbers)
	lastNumber := numbers[length-1]

	lastNumberIndeces := lastIndeces[lastNumber]

	if len(lastNumberIndeces) < 2 {
		lastIndeces[0] = append(lastIndeces[0], length)
		return 0
	}

	result := getLastPositions(lastNumberIndeces)
	lastIndeces[result] = append(lastIndeces[result], length)

	return result
}

func getIndex(numbers []int, finalLength int) int {
	initialLength := len(numbers)

	lastIndeces = map[int][]int{}

	for i, n := range numbers {
		lastIndeces[n] = append(lastIndeces[n], i)
	}

	for i := initialLength; i < finalLength; i++ {
		numbers = append(numbers, next(numbers))
	}

	return numbers[finalLength-1]
}

func main() {
	input := []int{9, 6, 0, 10, 18, 2, 1}

	fmt.Printf("2020th number spoken (1): %d\n", getIndex(input, 2020))
	fmt.Printf("30000000th number spoken (2): %d\n", getIndex(input, 30000000))
}
