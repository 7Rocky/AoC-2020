package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func findBestBusID(timestamp int, ids []int) (int, int) {
	var bestTime int = math.MaxInt32
	var bestBusID int

	for _, id := range ids {
		time := (timestamp/id + 1) * id

		if time < bestTime {
			bestTime = time
			bestBusID = id
		}
	}

	return bestBusID, bestTime - timestamp
}

func extendedGcd(a, b int) (int, int, int) {
	if a%b > 0 {
		u, v, d := extendedGcd(b, a%b)
		u = v
		v = (d - a*u) / b

		return u, v, d
	}

	return 0, 1, b
}

func getRemainders(idsList []string) []int {
	var rems []int

	for i, id := range idsList {
		if id != "x" {
			idNumber, _ := strconv.Atoi(id)
			rem := (idNumber - i) % idNumber

			for rem < 0 {
				rem += idNumber
			}

			rems = append(rems, rem)
		}
	}

	return rems
}

func prod(arr []int) int {
	prod := 1

	for _, a := range arr {
		prod *= a
	}

	return prod
}

func inv(n, m int) int {
	a, _, _ := extendedGcd(n, m)

	for a < 0 {
		a += m
	}

	return a
}

func next(rems, mods, partials []int) int {
	last := len(rems) - 1
	result := rems[last]

	for i, p := range partials {
		result -= prod(mods[:i]) * p
	}

	next := (result * inv(prod(mods[:last]), mods[last])) % mods[last]

	for next < 0 {
		next += mods[last]
	}

	return next
}

func solveCrt(rems, mods []int) int {
	var partials []int
	result := 0

	for i := 0; i < len(mods); i++ {
		partial := next(rems[:i+1], mods[:i+1], partials)
		partials = append(partials, partial)
		result += prod(mods[:i]) * partial
	}

	return result
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var timestamp int
	var ids []int

	scanner.Scan()
	timestamp, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	idsList := scanner.Text()

	for _, id := range strings.Split(idsList, ",") {
		if id != "x" {
			idNumber, _ := strconv.Atoi(id)
			ids = append(ids, idNumber)
		}
	}

	busID, timeLapse := findBestBusID(timestamp, ids)

	fmt.Printf("Best Bus ID: %d; Time lapse: %d; Result (1): %d\n", busID, timeLapse, busID*timeLapse)

	rems := getRemainders(strings.Split(idsList, ","))

	fmt.Printf("Result (2): %d\n", solveCrt(rems, ids))
}
