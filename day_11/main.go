package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	left = iota
	right
	up
	down
	mainDiagUp
	mainDiagDown
	antiDiagUp
	antiDiagDown
)

func adjacentSeats(seats [][]byte, i, j int) []byte {
	var adjacent []byte
	rows, columns := len(seats), len(seats[0])

	if i == 0 && j == 0 {
		return append(adjacent,
			seats[1][1],
			seats[0][1],
			seats[1][0])
	}

	if i == 0 && j == columns-1 {
		return append(adjacent,
			seats[0][columns-2],
			seats[1][columns-2],
			seats[1][columns-1])
	}

	if i == 0 {
		return append(adjacent,
			seats[0][j-1],
			seats[0][j+1],
			seats[1][j-1],
			seats[1][j+1],
			seats[1][j+0])
	}

	if i == rows-1 && j == 0 {
		return append(adjacent,
			seats[rows-1][1],
			seats[rows-2][1],
			seats[rows-2][0])
	}

	if i == rows-1 && j == columns-1 {
		return append(adjacent,
			seats[rows-1][columns-2],
			seats[rows-2][columns-2],
			seats[rows-2][columns-1])
	}

	if i == rows-1 {
		return append(adjacent,
			seats[rows-1][j-1],
			seats[rows-1][j+1],
			seats[rows-2][j-1],
			seats[rows-2][j+1],
			seats[rows-2][j+0])
	}

	if j == 0 {
		return append(adjacent,
			seats[i-1][0],
			seats[i-1][1],
			seats[i+0][1],
			seats[i+1][1],
			seats[i+1][0])
	}

	if j == columns-1 {
		return append(adjacent,
			seats[i-1][columns-1],
			seats[i-1][columns-2],
			seats[i+0][columns-2],
			seats[i+1][columns-2],
			seats[i+1][columns-1])
	}

	return append(adjacent,
		seats[i-1][j+1],
		seats[i-1][j+0],
		seats[i-1][j-1],
		seats[i+0][j+1],
		seats[i+0][j-1],
		seats[i+1][j+1],
		seats[i+1][j+0],
		seats[i+1][j-1])
}

func visibleSeatsDirection(seats [][]byte, i, j, dir int) byte {
	rows, columns := len(seats), len(seats[0])
	p := 1

	switch dir {
	case left:
		for j-p >= 0 && seats[i][j-p] == '.' {
			p++
		}

		if j-p >= 0 {
			return seats[i][j-p]
		}
	case right:
		for j+p < columns && seats[i][j+p] == '.' {
			p++
		}

		if j+p < columns {
			return seats[i][j+p]
		}
	case up:
		for i-p >= 0 && seats[i-p][j] == '.' {
			p++
		}

		if i-p >= 0 {
			return seats[i-p][j]
		}
	case down:
		for i+p < rows && seats[i+p][j] == '.' {
			p++
		}

		if i+p < rows {
			return seats[i+p][j]
		}
	case mainDiagUp:
		for i-p >= 0 && j-p >= 0 && seats[i-p][j-p] == '.' {
			p++
		}

		if i-p >= 0 && j-p >= 0 {
			return seats[i-p][j-p]
		}
	case mainDiagDown:
		for i+p < rows && j+p < columns && seats[i+p][j+p] == '.' {
			p++
		}

		if i+p < rows && j+p < columns {
			return seats[i+p][j+p]
		}
	case antiDiagUp:
		for i-p >= 0 && j+p < columns && seats[i-p][j+p] == '.' {
			p++
		}

		if i-p >= 0 && j+p < columns {
			return seats[i-p][j+p]
		}
	case antiDiagDown:
		for i+p < rows && j-p >= 0 && seats[i+p][j-p] == '.' {
			p++
		}

		if i+p < rows && j-p >= 0 {
			return seats[i+p][j-p]
		}
	}

	return '.'
}

func visibleSeats(seats [][]byte, i, j int) []byte {
	var visible []byte
	rows, columns := len(seats), len(seats[0])

	if i == 0 && j == 0 {
		return append(visible,
			visibleSeatsDirection(seats, 0, 0, right),
			visibleSeatsDirection(seats, 0, 0, down),
			visibleSeatsDirection(seats, 0, 0, mainDiagDown))
	}

	if i == 0 && j == columns-1 {
		return append(visible,
			visibleSeatsDirection(seats, 0, columns-1, left),
			visibleSeatsDirection(seats, 0, columns-1, down),
			visibleSeatsDirection(seats, 0, columns-1, antiDiagDown))
	}

	if i == 0 {
		return append(visible,
			visibleSeatsDirection(seats, 0, j, left),
			visibleSeatsDirection(seats, 0, j, right),
			visibleSeatsDirection(seats, 0, j, down),
			visibleSeatsDirection(seats, 0, j, mainDiagDown),
			visibleSeatsDirection(seats, 0, j, antiDiagDown))
	}

	if i == rows-1 && j == 0 {
		return append(visible,
			visibleSeatsDirection(seats, rows-1, 0, right),
			visibleSeatsDirection(seats, rows-1, 0, up),
			visibleSeatsDirection(seats, rows-1, 0, antiDiagUp))
	}

	if i == rows-1 && j == columns-1 {
		return append(visible,
			visibleSeatsDirection(seats, rows-1, columns-1, left),
			visibleSeatsDirection(seats, rows-1, columns-1, up),
			visibleSeatsDirection(seats, rows-1, columns-1, mainDiagUp))
	}

	if i == rows-1 {
		return append(visible,
			visibleSeatsDirection(seats, rows-1, j, left),
			visibleSeatsDirection(seats, rows-1, j, right),
			visibleSeatsDirection(seats, rows-1, j, up),
			visibleSeatsDirection(seats, rows-1, j, mainDiagUp),
			visibleSeatsDirection(seats, rows-1, j, antiDiagUp))
	}

	if j == 0 {
		return append(visible,
			visibleSeatsDirection(seats, i, 0, right),
			visibleSeatsDirection(seats, i, 0, down),
			visibleSeatsDirection(seats, i, 0, up),
			visibleSeatsDirection(seats, i, 0, mainDiagDown),
			visibleSeatsDirection(seats, i, 0, antiDiagUp))
	}

	if j == columns-1 {
		return append(visible,
			visibleSeatsDirection(seats, i, columns-1, left),
			visibleSeatsDirection(seats, i, columns-1, down),
			visibleSeatsDirection(seats, i, columns-1, up),
			visibleSeatsDirection(seats, i, columns-1, mainDiagUp),
			visibleSeatsDirection(seats, i, columns-1, antiDiagDown))
	}

	return append(visible,
		visibleSeatsDirection(seats, i, j, right),
		visibleSeatsDirection(seats, i, j, left),
		visibleSeatsDirection(seats, i, j, down),
		visibleSeatsDirection(seats, i, j, up),
		visibleSeatsDirection(seats, i, j, mainDiagUp),
		visibleSeatsDirection(seats, i, j, mainDiagDown),
		visibleSeatsDirection(seats, i, j, antiDiagUp),
		visibleSeatsDirection(seats, i, j, antiDiagDown))
}

func round(seats [][]byte, level int) [][]byte {
	seatsCopy := make([][]byte, len(seats))
	copy(seatsCopy, seats)
	rows, columns := len(seatsCopy), len(seatsCopy[0])

	for i := 0; i < rows; i++ {
		row := seats[i]
		rowCopy := make([]byte, len(row))
		copy(rowCopy, row)
		seatsCopy[i] = rowCopy

		for j := 0; j < columns; j++ {
			if row[j] == '.' {
				continue
			}

			if level == 1 {
				adjacent := adjacentSeats(seats, i, j)
				ocurrences := bytes.Count(adjacent, []byte{'#'})

				if row[j] == 'L' && ocurrences == 0 {
					rowCopy[j] = '#'
				}

				if row[j] == '#' && ocurrences >= 4 {
					rowCopy[j] = 'L'
				}
			} else if level == 2 {
				visible := visibleSeats(seats, i, j)
				ocurrences := bytes.Count(visible, []byte{'#'})

				if row[j] == 'L' && ocurrences == 0 {
					rowCopy[j] = '#'
				}

				if row[j] == '#' && ocurrences >= 5 {
					rowCopy[j] = 'L'
				}
			}
		}
	}

	return seatsCopy
}

func areEqual(mat1, mat2 [][]byte) bool {
	if len(mat1) != len(mat2) {
		return false
	}

	for i := range mat1 {
		if !bytes.Equal(mat1[i], mat2[i]) {
			return false
		}
	}

	return true
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var seats [][]byte

	for scanner.Scan() {
		seats = append(seats, []byte(scanner.Text()))
	}

	seatsCopy1 := make([][]byte, len(seats))
	seatsCopy2 := make([][]byte, len(seats))

	copy(seatsCopy1, seats)
	copy(seatsCopy2, seats)

	seats = round(seats, 1)

	for !areEqual(seatsCopy1, seats) {
		copy(seatsCopy1, seats)
		seats = round(seats, 1)
	}

	occupiedSeats := 0

	for _, r := range seats {
		occupiedSeats += bytes.Count(r, []byte{'#'})
	}

	fmt.Printf("Number of occupied seats when stabilized (1): %d\n", occupiedSeats)

	copy(seats, seatsCopy2)
	seats = round(seats, 2)

	for !areEqual(seatsCopy2, seats) {
		copy(seatsCopy2, seats)
		seats = round(seats, 2)
	}

	occupiedSeats = 0

	for _, r := range seats {
		occupiedSeats += bytes.Count(r, []byte{'#'})
	}

	fmt.Printf("Number of occupied seats when stabilized (2): %d\n", occupiedSeats)
}
