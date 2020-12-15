package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rightDir        = false
	leftDir         = true
	upDir           = false
	downDir         = true
	mainDiagUpDir   = true
	mainDiagDownDir = false
	antiDiagUpDir   = false
	antiDiagDownDir = true
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

func visibleSeatsXDirection(seats [][]byte, i, j int, left bool) byte {
	columns := len(seats[0])
	p := 1

	if left {
		for j-p >= 0 && seats[i][j-p] == '.' {
			p++
		}

		if j-p >= 0 {
			return seats[i][j-p]
		}
	} else {
		for j+p < columns && seats[i][j+p] == '.' {
			p++
		}

		if j+p < columns {
			return seats[i][j+p]
		}
	}

	return '.'
}

func visibleSeatsYDirection(seats [][]byte, i, j int, down bool) byte {
	rows := len(seats)
	q := 1

	if down {
		for i+q < rows && seats[i+q][j] == '.' {
			q++
		}

		if i+q < rows {
			return seats[i+q][j]
		}
	} else {
		for i-q >= 0 && seats[i-q][j] == '.' {
			q++
		}

		if i-q >= 0 {
			return seats[i-q][j]
		}
	}

	return '.'
}

func visibleSeatsAntiDiagDirection(seats [][]byte, i, j int, antiDiagDown bool) byte {
	rows, columns := len(seats), len(seats[0])
	r := 1

	if antiDiagDown {
		for i+r < rows && j-r >= 0 && seats[i+r][j-r] == '.' {
			r++
		}

		if i+r < rows && j-r >= 0 {
			return seats[i+r][j-r]
		}
	} else {
		for i-r >= 0 && j+r < columns && seats[i-r][j+r] == '.' {
			r++
		}

		if i-r >= 0 && j+r < columns {
			return seats[i-r][j+r]
		}
	}

	return '.'
}

func visibleSeatsMainDiagDirection(seats [][]byte, i, j int, mainDiagUp bool) byte {
	rows, columns := len(seats), len(seats[0])
	s := 1

	if mainDiagUp {
		for i-s >= 0 && j-s >= 0 && seats[i-s][j-s] == '.' {
			s++
		}

		if i-s >= 0 && j-s >= 0 {
			return seats[i-s][j-s]
		}
	} else {
		for i+s < rows && j+s < columns && seats[i+s][j+s] == '.' {
			s++
		}

		if i+s < rows && j+s < columns {
			return seats[i+s][j+s]
		}
	}

	return '.'
}

func visibleSeats(seats [][]byte, i, j int) []byte {
	var visible []byte
	rows, columns := len(seats), len(seats[0])

	if i == 0 && j == 0 {
		return append(visible,
			visibleSeatsXDirection(seats, 0, 0, rightDir),
			visibleSeatsYDirection(seats, 0, 0, downDir),
			visibleSeatsMainDiagDirection(seats, 0, 0, mainDiagDownDir))
	}

	if i == 0 && j == columns-1 {
		return append(visible,
			visibleSeatsXDirection(seats, 0, columns-1, leftDir),
			visibleSeatsYDirection(seats, 0, columns-1, downDir),
			visibleSeatsAntiDiagDirection(seats, 0, columns-1, antiDiagDownDir))
	}

	if i == 0 {
		return append(visible,
			visibleSeatsXDirection(seats, 0, j, leftDir),
			visibleSeatsXDirection(seats, 0, j, rightDir),
			visibleSeatsYDirection(seats, 0, j, downDir),
			visibleSeatsMainDiagDirection(seats, 0, j, mainDiagDownDir),
			visibleSeatsAntiDiagDirection(seats, 0, j, antiDiagDownDir))
	}

	if i == rows-1 && j == 0 {
		return append(visible,
			visibleSeatsXDirection(seats, rows-1, 0, rightDir),
			visibleSeatsYDirection(seats, rows-1, 0, upDir),
			visibleSeatsAntiDiagDirection(seats, rows-1, 0, antiDiagUpDir))
	}

	if i == rows-1 && j == columns-1 {
		return append(visible,
			visibleSeatsXDirection(seats, rows-1, columns-1, leftDir),
			visibleSeatsYDirection(seats, rows-1, columns-1, upDir),
			visibleSeatsMainDiagDirection(seats, rows-1, columns-1, mainDiagUpDir))
	}

	if i == rows-1 {
		return append(visible,
			visibleSeatsXDirection(seats, rows-1, j, leftDir),
			visibleSeatsXDirection(seats, rows-1, j, rightDir),
			visibleSeatsYDirection(seats, rows-1, j, upDir),
			visibleSeatsMainDiagDirection(seats, rows-1, j, mainDiagUpDir),
			visibleSeatsAntiDiagDirection(seats, rows-1, j, antiDiagUpDir))
	}

	if j == 0 {
		return append(visible,
			visibleSeatsXDirection(seats, i, 0, rightDir),
			visibleSeatsYDirection(seats, i, 0, downDir),
			visibleSeatsYDirection(seats, i, 0, upDir),
			visibleSeatsMainDiagDirection(seats, i, 0, mainDiagDownDir),
			visibleSeatsAntiDiagDirection(seats, i, 0, antiDiagUpDir))
	}

	if j == columns-1 {
		return append(visible,
			visibleSeatsXDirection(seats, i, columns-1, leftDir),
			visibleSeatsYDirection(seats, i, columns-1, downDir),
			visibleSeatsYDirection(seats, i, columns-1, upDir),
			visibleSeatsMainDiagDirection(seats, i, columns-1, mainDiagUpDir),
			visibleSeatsAntiDiagDirection(seats, i, columns-1, antiDiagDownDir))
	}

	return append(visible,
		visibleSeatsXDirection(seats, i, j, rightDir),
		visibleSeatsXDirection(seats, i, j, leftDir),
		visibleSeatsYDirection(seats, i, j, downDir),
		visibleSeatsYDirection(seats, i, j, upDir),
		visibleSeatsMainDiagDirection(seats, i, j, mainDiagUpDir),
		visibleSeatsMainDiagDirection(seats, i, j, mainDiagDownDir),
		visibleSeatsAntiDiagDirection(seats, i, j, antiDiagUpDir),
		visibleSeatsAntiDiagDirection(seats, i, j, antiDiagDownDir))
}

func countOcurrences(str []byte, char byte) int {
	count := 0

	for _, c := range str {
		if c == char {
			count++
		}
	}

	return count
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
				ocurrences := countOcurrences(adjacent, '#')

				if row[j] == 'L' && ocurrences == 0 {
					rowCopy[j] = '#'
				}

				if row[j] == '#' && ocurrences >= 4 {
					rowCopy[j] = 'L'
				}
			} else if level == 2 {
				visible := visibleSeats(seats, i, j)
				ocurrences := countOcurrences(visible, '#')

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
	if len(mat1) != len(mat2) || len(mat1[0]) != len(mat2[0]) {
		return false
	}

	for i, r1 := range mat1 {
		for j, c1 := range r1 {
			if c1 != mat2[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	file, _ := os.Open("./input.txt")

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
		occupiedSeats += countOcurrences(r, '#')
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
		occupiedSeats += countOcurrences(r, '#')
	}

	fmt.Printf("Number of occupied seats when stabilized (2): %d\n", occupiedSeats)
}
