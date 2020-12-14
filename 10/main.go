package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func getDifferences(arr []int) []int {
	diffs := make([]int, len(arr)-1)

	for i := 1; i < len(arr); i++ {
		diffs[i-1] = arr[i] - arr[i-1]
	}

	return diffs
}

func countOcurrences(arr, ocurrences []int) []int {
	count := make([]int, len(ocurrences))

	for _, a := range arr {
		for n, i := range ocurrences {
			if a == i {
				count[n]++
			}
		}
	}

	return count
}

func getLengthOfBlocks(arr []int, number int) []int {
	lengths := []int{1}

	for i := 1; i < len(arr); i++ {
		if arr[i] == number && arr[i-1] == number {
			lengths[len(lengths)-1]++
		} else {
			lengths = append(lengths, 1)
		}
	}

	return lengths
}

func divmod(dividend, divisor int) (int, int) {
	return dividend / divisor, dividend % divisor
}

func decimalToTernary(number, length int) []int {
	var remainder int
	remainders := make([]int, length)

	for i := 0; i < length; i++ {
		number, remainder = divmod(number, 3)
		// The set for this approach is not {0, 1, 2}, but {1, 2, 3}
		remainders[i] = remainder + 1
	}

	return remainders
}

func getTernaryTruthTable(length int) [][]int {
	truthTable := make([][]int, int(math.Pow(3, float64(length))))

	for i := 0; i < int(math.Pow(3, float64(length))); i++ {
		truthTable[i] = decimalToTernary(i, length)
	}

	return truthTable
}

func sum(arr []int) int {
	sum := 0

	for _, s := range arr {
		sum += s
	}

	return sum
}

var knownPartitions = map[int]int{}

func getPartitionLength(length int) int {
	if knownPartitions[length] != 0 {
		return knownPartitions[length]
	}

	partitions := 0

	for i := 1; i <= length; i++ {
		possibilities := getTernaryTruthTable(i)

		for _, p := range possibilities {
			try := make([]int, length)
			copy(try, p)

			if sum(try) == length {
				partitions++
			}
		}
	}

	knownPartitions[length] = partitions

	return partitions
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []int{0}

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	numbers = append(numbers, numbers[len(numbers)-1]+3)

	diffs := getDifferences(numbers)
	countDiffs := countOcurrences(diffs, []int{1, 3})

	fmt.Printf("Number of 1-jolt differences multiplied by the number of 3-jolt differences (1): %d\n",
		countDiffs[0]*countDiffs[1])

	numWays := 1

	for _, l := range getLengthOfBlocks(diffs, 1) {
		numWays *= getPartitionLength(l)
	}

	fmt.Printf("Number of distinct ways to arrange the adapters (2): %d\n", numWays)
}
