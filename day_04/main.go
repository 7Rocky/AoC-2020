package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidPassport(passport string, level int) bool {
	var (
		rebyr *regexp.Regexp
		reiyr *regexp.Regexp
		reeyr *regexp.Regexp
		rehgt *regexp.Regexp
		rehcl *regexp.Regexp
		reecl *regexp.Regexp
		repid *regexp.Regexp
	)

	if level == 1 {
		rebyr = regexp.MustCompile("byr:")
		reiyr = regexp.MustCompile("iyr:")
		reeyr = regexp.MustCompile("eyr:")
		rehgt = regexp.MustCompile("hgt:")
		rehcl = regexp.MustCompile("hcl:")
		reecl = regexp.MustCompile("ecl:")
		repid = regexp.MustCompile("pid:")
	} else if level == 2 {
		rebyr = regexp.MustCompile("(byr:19[2-9][0-9])|(byr:200[0-2])")
		reiyr = regexp.MustCompile("(iyr:201[0-9])|(iyr:2020)")
		reeyr = regexp.MustCompile("(eyr:202[0-9])|(eyr:2030)")
		rehgt = regexp.MustCompile("(hgt:1[5-8][0-9]cm)|(hgt:19[0-3]cm)|(hgt:59in)|(hgt:6[0-9]in)|(hgt:7[0-6]in)")
		rehcl = regexp.MustCompile("(hcl:#[0-9a-f]{6})")
		reecl = regexp.MustCompile("(ecl:amb)|(ecl:blu)|(ecl:brn)|(ecl:gry)|(ecl:grn)|(ecl:hzl)|(ecl:oth)")
		repid = regexp.MustCompile("pid:[0-9]{9}")
	}

	return rebyr.MatchString(passport) &&
		reiyr.MatchString(passport) &&
		reeyr.MatchString(passport) &&
		rehgt.MatchString(passport) &&
		rehcl.MatchString(passport) &&
		reecl.MatchString(passport) &&
		repid.MatchString(passport)
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	passport := ""

	var passports []string

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			passport += line + " "
		} else {
			passports = append(passports, passport)
			passport = ""
		}
	}

	for _, level := range []int{1, 2} {
		numValid := 0

		for _, p := range passports {
			if isValidPassport(strings.Trim(p, " "), level) {
				numValid++
			}
		}

		fmt.Printf("Valid passports (%d): %d\n", level, numValid)
	}
}
