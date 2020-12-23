package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	player1 int = iota
	player2
)

func pop(arr []int) (int, []int) {
	return arr[0], arr[1:]
}

func shift(arr []int) []int {
	n, arr := pop(arr)
	return append(arr, n)
}

func game(decks [2][]int) []int {
	round := 0
	number := 0

	for len(decks[player1]) != 0 && len(decks[player2]) != 0 {
		round++

		if decks[player1][0] < decks[player2][0] {
			decks[player2] = shift(decks[player2])
			number, decks[player1] = pop(decks[player1])
			decks[player2] = append(decks[player2], number)
			continue
		}

		if decks[player1][0] > decks[player2][0] {
			decks[player1] = shift(decks[player1])
			number, decks[player2] = pop(decks[player2])
			decks[player1] = append(decks[player1], number)
			continue
		}
	}

	if len(decks[player2]) == 0 {
		return decks[player1]
	}

	return decks[player2]
}

func score(deck []int) int {
	score := 0

	for i, n := range deck {
		score += (len(deck) - i) * n
	}

	return score
}

func areEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	result := true

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			result = false
		}
	}

	return result
}

func indexOf(decks [2][]int, game int) int {
	for i, d := range games[game] {
		if areEqual(decks[player1], d[player1]) && areEqual(decks[player2], d[player2]) {
			return i
		}
	}

	return -1
}

var games = map[int][][2][]int{}

func gameRecursive(decks [2][]int, game int) (int, []int) {
	number, round := 0, 0

	// fmt.Printf("== Game %d ==\n", game)
	games[game] = [][2][]int{}

	for len(decks[player1]) != 0 && len(decks[player2]) != 0 {
		round++

		if indexOf([2][]int{decks[player1], decks[player2]}, game) != -1 {
			// fmt.Printf("The winner of game %d is player 1!\n\n", game)
			games[game] = [][2][]int{}

			return player1, decks[player1]
		}

		games[game] = append(games[game], [2][]int{decks[player1], decks[player2]})

		// fmt.Printf("\n-- Round %d (Game %d) --\n", round, game)
		// fmt.Printf("Player 1's deck: ")
		// fmt.Println(decks[player1])
		// fmt.Printf("Player 2's deck: ")
		// fmt.Println(decks[player2])
		// fmt.Printf("Player 1 plays: %d\n", decks[player1][0])
		// fmt.Printf("Player 2 plays: %d\n", decks[player2][0])

		var winner int

		if decks[player1][0] > decks[player2][0] {
			winner = player1
		} else {
			winner = player2
		}

		if decks[player1][0] < len(decks[player1]) && decks[player2][0] < len(decks[player2]) {
			// fmt.Println("Playing a sub-game to determine the winner...\n")
			copyDeck1, copyDeck2 := make([]int, decks[player1][0]), make([]int, decks[player2][0])
			copy(copyDeck1, decks[player1][1:decks[player1][0]+1])
			copy(copyDeck2, decks[player2][1:decks[player2][0]+1])
			winner, _ = gameRecursive([2][]int{copyDeck1, copyDeck2}, game+1)
			// fmt.Printf("...anyway, back to game %d.\n", game)
		}

		if winner == player1 {
			// fmt.Printf("Player 1 wins round %d of game %d!\n", round, game)
			decks[player1] = shift(decks[player1])
			number, decks[player2] = pop(decks[player2])
			decks[player1] = append(decks[player1], number)
			continue
		}

		if winner == player2 {
			// fmt.Printf("Player 2 wins round %d of game %d!\n", round, game)
			decks[player2] = shift(decks[player2])
			number, decks[player1] = pop(decks[player1])
			decks[player2] = append(decks[player2], number)
			continue
		}
	}

	games[game] = [][2][]int{}

	if len(decks[player2]) == 0 {
		// fmt.Printf("The winner of game %d is player 1!\n\n", game)
		return player1, decks[player1]
	}

	// fmt.Printf("The winner of game %d is player 2!\n\n", game)
	return player2, decks[player2]
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var decks1, decks2 [2][]int

	player := player1

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Player") {
			continue
		}

		if line == "" {
			player = player2
			continue
		}

		num, _ := strconv.Atoi(line)
		decks1[player] = append(decks1[player], num)
		decks2[player] = append(decks2[player], num)
	}

	winnerDeck := game(decks1)

	fmt.Printf("Score for Combat winner (1): %d\n", score(winnerDeck))

	_, winnerDeck = gameRecursive(decks2, 1)

	fmt.Printf("Score for Recursive Combat winner (2): %d\n", score(winnerDeck))
}
