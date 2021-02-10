package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
	visited   bool
}

func executeCode(code []instruction) (int, bool) {
	acc, i := 0, 0

	for i < len(code) {
		if code[i].visited {
			return acc, false
		}

		code[i].visited = true

		if code[i].operation == "nop" {
			i++
		} else if code[i].operation == "acc" {
			acc += code[i].argument
			i++
		} else if code[i].operation == "jmp" {
			i += code[i].argument
		}
	}

	return acc, true
}

func changeJmpForNop(modifiedCode []instruction, index int) []instruction {
	if modifiedCode[index].operation == "jmp" {
		modifiedCode[index].operation = "nop"
	} else if modifiedCode[index].operation == "nop" {
		modifiedCode[index].operation = "jmp"
	}

	return modifiedCode
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var code []instruction

	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")[0]
		argument, _ := strconv.Atoi(strings.Split(line, " ")[1])

		code = append(code, instruction{operation: operation, argument: argument, visited: false})
	}

	modifiedCode := make([]instruction, len(code))

	copy(modifiedCode, code)

	acc, isFinite := executeCode(modifiedCode)

	fmt.Printf("Result before infinite loop (1): %d\n", acc)

	for i, c := range code {
		if c.operation == "acc" {
			continue
		}

		copy(modifiedCode, code)
		modifiedCode = changeJmpForNop(modifiedCode, i)

		acc, isFinite = executeCode(modifiedCode)

		if isFinite {
			break
		}
	}

	fmt.Printf("Result of the clean program (2): %d\n", acc)
}
