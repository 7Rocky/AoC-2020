package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func productMod(a, b, m int) int {
	return (a * b) % m
}

func getLoopSize(subjectNumber, divisor, publicKey int) int {
	num := 1
	loopSize := 0

	for num != publicKey {
		loopSize++
		num = productMod(num, subjectNumber, divisor)
	}

	return loopSize
}

func getEncryptionKey(subjectNumber, divisor, loopSize int) int {
	num := 1

	for i := 0; i < loopSize; i++ {
		num = productMod(num, subjectNumber, divisor)
	}

	return num
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cardPublicKey, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	doorPublicKey, _ := strconv.Atoi(scanner.Text())

	subjectNumber := 7
	divisor := 20201227

	cardLoopSize := getLoopSize(subjectNumber, divisor, cardPublicKey)
	doorLoopSize := getLoopSize(subjectNumber, divisor, doorPublicKey)

	fmt.Println("Encryption key (1):", getEncryptionKey(cardPublicKey, divisor, doorLoopSize))
	fmt.Println("Encryption key (1):", getEncryptionKey(doorPublicKey, divisor, cardLoopSize))
}
