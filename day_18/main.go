package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(arr []int) int {
	sum := 0

	for _, a := range arr {
		sum += a
	}

	return sum
}

func findSumSign(op string) int {
	for i, t := range op {
		if t == '+' {
			return i
		}
	}

	return -1
}

func getLastNumber(op string) (int, int) {
	terms := strings.Split(op, " ")
	number, _ := strconv.Atoi(terms[len(terms)-1])

	return number, int(math.Log10(float64(number))) + 1
}

func getFirstNumber(op string) (int, int) {
	terms := strings.Split(op, " ")
	number, _ := strconv.Atoi(terms[0])

	return number, int(math.Log10(float64(number))) + 1
}

func evaluateSums(op string) string {
	for strings.Count(op, "+") != 0 {
		sum := findSumSign(op)

		num1, digits1 := getLastNumber(op[:sum-1])
		num2, digits2 := getFirstNumber(op[sum+2:])

		partial := num1 + num2

		op = op[:sum-1-digits1] + strconv.Itoa(partial) + op[sum+2+digits2:]
	}

	return op
}

func evaluateArithmetic(op string, level int) int {
	if level == 1 {
		terms := strings.Split(op, " ")
		result, _ := strconv.Atoi(terms[0])

		for i := 1; i < len(terms); i += 2 {
			if terms[i] == "+" {
				num, _ := strconv.Atoi(terms[i+1])
				result += num
			}

			if terms[i] == "*" {
				num, _ := strconv.Atoi(terms[i+1])
				result *= num
			}
		}

		return result
	}

	op = evaluateSums(op)

	terms := strings.Split(op, " ")
	result, _ := strconv.Atoi(terms[0])

	for i := 1; i < len(terms); i += 2 {
		if terms[i] == "*" {
			num, _ := strconv.Atoi(terms[i+1])
			result *= num
		}
	}

	return result
}

func findMostInnerParentheses(op string) (int, int) {
	opening, closing := -1, -1

	for i, t := range op {
		if t == '(' {
			opening = i
		}

		if t == ')' {
			closing = i
			break
		}
	}

	return opening, closing
}

func evaluate(op string, level int) int {
	for strings.Count(op, "(") > 0 {
		opening, closing := findMostInnerParentheses(op)

		partial := evaluateArithmetic(op[opening+1:closing], level)

		op = op[:opening] + strconv.Itoa(partial) + op[closing+1:]
	}

	return evaluateArithmetic(op, level)
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var operations []string

	for scanner.Scan() {
		operations = append(operations, scanner.Text())
	}

	var results1 []int

	for _, op := range operations {
		results1 = append(results1, evaluate(op, 1))
	}

	fmt.Printf("Sum of all operations (1): %d\n", sum(results1))

	var results2 []int

	for _, op := range operations {
		results2 = append(results2, evaluate(op, 2))
	}

	fmt.Printf("Sum of all operations (2): %d\n", sum(results2))
}
