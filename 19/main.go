package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func matchSimpleRules() map[int][]string {
	matchedRules := map[int][]string{}

	for _, r := range rules {
		if strings.Contains(r, "\"") {
			splitted := strings.Split(r, ":")
			numRule, _ := strconv.Atoi(splitted[0])
			matchedRules[numRule] = append(matchedRules[numRule], strings.Split(splitted[1], "\"")[1])
		}
	}

	return matchedRules
}

func getMatchingCharacters(index int) (int, []string) {
	splitted := strings.Split(rules[index], ":")
	numRule, _ := strconv.Atoi(splitted[0])

	if len(matchingCharacters[numRule]) > 0 {
		return numRule, matchingCharacters[numRule]
	}

	subrules := strings.Split(strings.Trim(splitted[1], " "), " ")
	matchedStrings := []string{""}

	fragment := 0

	for _, s := range subrules {
		if s != "|" {
			numSubrule, _ := strconv.Atoi(s)

			if len(matchingCharacters[numSubrule]) == 0 {
				index := -1

				for i, r := range rules {
					if strings.HasPrefix(r, s+":") {
						index = i
						break
					}
				}

				_, matchingCharacters[numSubrule] = getMatchingCharacters(index)
			}

			newMatchedStrings := make([]string, len(matchedStrings))
			copy(newMatchedStrings, matchedStrings)

			for i, m := range matchedStrings[fragment:] {
				for j, n := range matchingCharacters[numSubrule] {
					if j == 0 {
						newMatchedStrings[i+fragment] = m + n
					} else {
						newMatchedStrings = append(newMatchedStrings, m+n)
					}
				}
			}

			matchedStrings = newMatchedStrings
		} else {
			fragment = len(matchedStrings)
			matchedStrings = append(matchedStrings, "")
		}
	}

	return numRule, matchedStrings
}

var matchingCharacters = map[int][]string{}
var rules []string

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		rules = append(rules, line)
	}

	matchingCharacters = matchSimpleRules()

	for i := range rules {
		num, matchs := getMatchingCharacters(i)
		matchingCharacters[num] = matchs
	}

	var messages []string

	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	numMatches := 0

	for _, m := range messages {
		for _, rule := range matchingCharacters[0] {
			if m == rule {
				numMatches++
			}
		}
	}

	fmt.Printf("Number of matches for rule 0 (1): %d\n", numMatches)

	numMatches = 0

	for _, m := range messages {
		found := true
		num31 := 0
		num42 := 0

		for found {
			for _, rule := range matchingCharacters[31] {
				if strings.HasSuffix(m, rule) {
					m = m[:len(m)-len(rule)]
					num31++
					found = true
					break
				} else {
					found = false
				}
			}
		}

		found = true

		if num31 > 0 {
			for found {
				for _, rule := range matchingCharacters[42] {
					if strings.HasSuffix(m, rule) {
						m = m[:len(m)-len(rule)]
						num42++
						found = true
						break
					} else {
						found = false
					}
				}
			}

			if m == "" && num42 > num31 {
				numMatches++
			}
		}
	}

	fmt.Printf("Number of matches for rule 0 (2): %d\n", numMatches)
}
