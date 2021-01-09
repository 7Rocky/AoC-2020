package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type groupResponse struct {
	answers []string
	people  int
}

func unique(str []byte) string {
	keys := make(map[byte]bool)
	list := ""

	for _, entry := range str {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list += string(entry)
		}
	}

	return list
}

func common(answers []string) string {
	commonAnswers := ""
	minlengthAnswer := ""
	minLength := math.MaxInt32

	for _, answer := range answers {
		if length := len(answer); length < minLength {
			minlengthAnswer = answer
		}
	}

	for _, c := range []byte(minlengthAnswer) {
		common := true

		for _, answer := range answers {
			if !strings.Contains(answer, string(c)) {
				common = false
			}
		}

		if common {
			commonAnswers += string(c)
		}
	}

	return commonAnswers
}

func sum(arr []int) int {
	sum := 0

	for _, a := range arr {
		sum += a
	}

	return sum
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var groups []groupResponse

	answers := ""
	people := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			answers += line + " "
			people++
		} else {
			answers = strings.Trim(answers, " ")
			group := groupResponse{strings.Split(answers, " "), people}
			groups = append(groups, group)
			answers = ""
			people = 0
		}
	}

	var numUniqueAnswers []int

	for _, g := range groups {
		jointAnswers := strings.Join(g.answers, "")
		uniqueAnswers := unique([]byte(jointAnswers))
		numUniqueAnswers = append(numUniqueAnswers, len(uniqueAnswers))
	}

	fmt.Printf("Number of unique YES answers (1): %d\n", sum(numUniqueAnswers))

	var numCommonAnswers []int

	for _, g := range groups {
		commonAnswers := common(g.answers)
		numCommonAnswers = append(numCommonAnswers, len(commonAnswers))
	}

	fmt.Printf("Number of common YES answers (2): %d\n", sum(numCommonAnswers))
}
