package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidPassport(passport []byte, level int) bool {
	var (
		patternbyr string
		patterniyr string
		patterneyr string
		patternhgt string
		patternhcl string
		patternecl string
		patternpid string
	)

	if level == 1 {
		patternbyr = `byr:`
		patterniyr = `iyr:`
		patterneyr = `eyr:`
		patternhgt = `hgt:`
		patternhcl = `hcl:`
		patternecl = `ecl:`
		patternpid = `pid:`
	} else if level == 2 {
		patternbyr = `(byr:19[2-9][0-9])|(byr:200[0-2])`
		patterniyr = `(iyr:201[0-9])|(iyr:2020)`
		patterneyr = `(eyr:202[0-9])|(eyr:2030)`
		patternhgt = `(hgt:1[5-8][0-9]cm)|(hgt:19[0-3]cm)|(hgt:59in)|(hgt:6[0-9]in)|(hgt:7[0-6]in)`
		patternhcl = `(hcl:#[0-9a-f]{6})`
		patternecl = `(ecl:amb)|(ecl:blu)|(ecl:brn)|(ecl:gry)|(ecl:grn)|(ecl:hzl)|(ecl:oth)`
		patternpid = `pid:[0-9]{9}`
	}

	rebyr := regexp.MustCompile(patternbyr)
	reiyr := regexp.MustCompile(patterniyr)
	reeyr := regexp.MustCompile(patterneyr)
	rehgt := regexp.MustCompile(patternhgt)
	rehcl := regexp.MustCompile(patternhcl)
	reecl := regexp.MustCompile(patternecl)
	repid := regexp.MustCompile(patternpid)

	return rebyr.Match(passport) &&
		reiyr.Match(passport) &&
		reeyr.Match(passport) &&
		rehgt.Match(passport) &&
		rehcl.Match(passport) &&
		reecl.Match(passport) &&
		repid.Match(passport)
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	passport := ""

	var passports []string

	for scanner.Scan() {
		line := scanner.Text()

		if line != string("") {
			passport += line + " "
		} else {
			passports = append(passports, passport)
			passport = ""
		}
	}

	numValid := []int{0, 0}

	for _, p := range passports {
		if isValidPassport([]byte(strings.Trim(p, " ")), 1) {
			numValid[0]++
		}

		if isValidPassport([]byte(strings.Trim(p, " ")), 2) {
			numValid[1]++
		}
	}

	fmt.Printf("Valid passports (1): %d / %d\n", numValid[0], len(passports))

	// Curiously, it is in fact numValid[1] - 1
	numValid[1]--

	fmt.Printf("Valid passports (2): %d / %d\n", numValid[1], len(passports))
}
