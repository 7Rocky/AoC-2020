package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type requirement struct {
	minimum   int
	maximum   int
	character byte
	password  string
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	var reqs = []requirement{}

	scanner := bufio.NewScanner(file)

	reMinimum := regexp.MustCompile(`^\d+`)
	reMaximum := regexp.MustCompile(`\d+ `)
	reCharacter := regexp.MustCompile(`\D:`)
	rePassword := regexp.MustCompile(`: \w+`)

	for scanner.Scan() {
		line := scanner.Text()

		minimum, _ := strconv.Atoi(reMinimum.FindString(line))
		maximum, _ := strconv.Atoi(strings.Trim(reMaximum.FindString(line), " "))
		character := reCharacter.FindString(line)[0]
		password := rePassword.FindString(line)[2:]

		reqs = append(reqs, requirement{minimum: minimum, maximum: maximum, character: character, password: password})
	}

	numValid := 0

	for _, r := range reqs {
		if r.minimum <= strings.Count(r.password, string(r.character)) {
			if strings.Count(r.password, string(r.character)) <= r.maximum {
				numValid++
			}
		}
	}

	fmt.Printf("Number of valid passwords (1): %d\n", numValid)

	numValid = 0

	for _, r := range reqs {
		if (r.password[r.minimum-1] == r.character) != (r.password[r.maximum-1] == r.character) {
			numValid++
		}
	}

	fmt.Printf("Number of valid passwords (2): %d\n", numValid)
}
