package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type shinyGoldBagStatusType int

const (
	notChecked shinyGoldBagStatusType = iota
	notContaining
	containing
)

type bagType struct {
	innerBags          []string
	totalInnerBags     int
	shinyGoldBagStatus shinyGoldBagStatusType
}

var bags = map[string]*bagType{}

var targetBag string = "shiny gold bag"

func getShinyGoldBagStatus(bag string) shinyGoldBagStatusType {
	if len(bags[bag].innerBags) == 0 {
		return notContaining
	}

	if bags[bag].shinyGoldBagStatus == notChecked {
		contains := false

		for _, b := range bags[bag].innerBags {
			if strings.Contains(b, targetBag) {
				contains = true
			}
		}

		if contains {
			return containing
		}

		for _, b := range bags[bag].innerBags {
			if getShinyGoldBagStatus(b[2:]) == containing {
				return containing
			}
		}

		return notContaining
	}

	return bags[bag].shinyGoldBagStatus
}

func getTotalInnerBags(bag string) int {
	if len(bags[bag].innerBags) == 0 {
		return 0
	}

	totalInnerBags := 0

	for _, b := range bags[bag].innerBags {
		numSubbags, _ := strconv.Atoi(b[:1])

		totalInnerBags += numSubbags

		if totalInnerSubbags := getTotalInnerBags(b[2:]); totalInnerSubbags != 0 {
			totalInnerBags += numSubbags * totalInnerSubbags
		}
	}

	return totalInnerBags
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bagData := strings.Trim(scanner.Text(), ".")
		splitted := strings.Split(bagData, " contain ")
		bag := strings.TrimRight(splitted[0], "s")

		innerBags := []string{}

		if strings.Contains(splitted[1], ", ") {
			innerBags = strings.Split(splitted[1], ", ")

			for i, b := range innerBags {
				innerBags[i] = strings.TrimRight(b, "s")
			}
		} else if !strings.Contains(splitted[1], "no other bags") {
			innerBags = []string{strings.TrimRight(splitted[1], "s")}
		}

		shinyGoldStatus := notChecked

		if strings.Contains(splitted[1], targetBag) {
			shinyGoldStatus = containing
		}

		bags[bag] = &bagType{innerBags: innerBags, totalInnerBags: 0, shinyGoldBagStatus: shinyGoldStatus}
	}

	for b := range bags {
		bags[b].shinyGoldBagStatus = getShinyGoldBagStatus(b)
		bags[b].totalInnerBags = getTotalInnerBags(b)
	}

	numBags := 0

	for _, d := range bags {
		if d.shinyGoldBagStatus == containing {
			numBags++
		}
	}

	fmt.Printf("Number of bags that contain a %s (1): %d\n",
		targetBag, numBags)

	fmt.Printf("Number of bags contained inside a %s (2): %d\n",
		targetBag, bags[targetBag].totalInnerBags)
}
