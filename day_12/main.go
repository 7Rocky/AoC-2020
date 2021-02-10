package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	north   = 'N'
	south   = 'S'
	east    = 'E'
	west    = 'W'
	left    = 'L'
	right   = 'R'
	forward = 'F'
)

type instruction struct {
	action byte
	value  int
}

type waypoint struct {
	position []int
}

type shipObject struct {
	position  []int
	direction byte
	waypoint  waypoint
}

var directions []byte = []byte{north, east, south, west}

func move(instr instruction, ship shipObject, level int) shipObject {
	if level == 1 {
		switch instr.action {
		case north:
			ship.position[1] += instr.value
		case south:
			ship.position[1] -= instr.value
		case east:
			ship.position[0] += instr.value
		case west:
			ship.position[0] -= instr.value
		case left:
			index := (bytes.IndexByte(directions, ship.direction) - instr.value/90 + 4) % 4
			ship.direction = directions[index]
		case right:
			index := (bytes.IndexByte(directions, ship.direction) + instr.value/90 + 4) % 4
			ship.direction = directions[index]
		case forward:
			switch ship.direction {
			case north:
				ship.position[1] += instr.value
			case east:
				ship.position[0] += instr.value
			case south:
				ship.position[1] -= instr.value
			case west:
				ship.position[0] -= instr.value
			}
		}
	} else if level == 2 {
		switch instr.action {
		case north:
			ship.waypoint.position[1] += instr.value
		case south:
			ship.waypoint.position[1] -= instr.value
		case east:
			ship.waypoint.position[0] += instr.value
		case west:
			ship.waypoint.position[0] -= instr.value
		case forward:
			ship.position[0], ship.position[1] =
				ship.position[0]+instr.value*ship.waypoint.position[0],
				ship.position[1]+instr.value*ship.waypoint.position[1]
		}

		if instr.action == right && instr.value == 90 || instr.action == left && instr.value == 270 {
			ship.waypoint.position[0], ship.waypoint.position[1] =
				ship.waypoint.position[1],
				-ship.waypoint.position[0]
		}

		if instr.value == 180 {
			ship.waypoint.position[0], ship.waypoint.position[1] =
				-ship.waypoint.position[0],
				-ship.waypoint.position[1]
		}

		if instr.action == right && instr.value == 270 || instr.action == left && instr.value == 90 {
			ship.waypoint.position[0], ship.waypoint.position[1] =
				-ship.waypoint.position[1],
				ship.waypoint.position[0]
		}
	}

	return ship
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []instruction

	for scanner.Scan() {
		instr := scanner.Text()
		action := instr[0]
		value, _ := strconv.Atoi(instr[1:])
		instructions = append(instructions, instruction{action: action, value: value})
	}

	ship := shipObject{position: []int{0, 0}, direction: east, waypoint: waypoint{position: []int{10, 1}}}

	for _, instr := range instructions {
		ship = move(instr, ship, 1)
	}

	manhattanDistance := int(math.Abs(float64(ship.position[0])) + math.Abs(float64(ship.position[1])))

	fmt.Printf("Final Position: (%d, %d); Manhattan distance to origin (1): %d\n",
		ship.position[0], ship.position[1], manhattanDistance)

	ship = shipObject{position: []int{0, 0}, direction: east, waypoint: waypoint{position: []int{10, 1}}}

	for _, instr := range instructions {
		ship = move(instr, ship, 2)
	}

	manhattanDistance = int(math.Abs(float64(ship.position[0])) + math.Abs(float64(ship.position[1])))

	fmt.Printf("Final Position: (%d, %d); Manhattan distance to origin (2): %d\n",
		ship.position[0], ship.position[1], manhattanDistance)
}
