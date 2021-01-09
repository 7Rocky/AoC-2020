package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pocket1 = map[int][][]byte{}
var pocket2 = map[int]map[int][][]byte{}

func initPocket1(centralSlice []string, iterations int) {
	initialLength := len(centralSlice)
	maxSpan := initialLength + 2*iterations + 2
	maxHeight := 1 + 2*iterations + 2

	for i := -maxHeight / 2; i <= maxHeight/2; i++ {
		span := [][]byte{}

		for j := 0; j < maxSpan; j++ {
			span = append(span, []byte(strings.Repeat(".", maxSpan)))
		}

		pocket1[i] = span
	}

	for i, s := range centralSlice {
		index := (maxSpan - initialLength) / 2
		pocket1[0][index+i] = []byte(strings.Repeat(".", index) + s + strings.Repeat(".", index))
	}
}

func countActiveGrid1(grid [][]byte, x, y int) int {
	length := len(grid)

	if x > 0 && x < length-1 && y > 0 && y < length-1 {
		return strings.Count(string(grid[y-1][x-1:x+2]), "#") +
			strings.Count(string(grid[y+0][x-1:x+2]), "#") +
			strings.Count(string(grid[y+1][x-1:x+2]), "#")
	}
	return 0
}

func countActive1(x, y, z int) int {
	active := 0

	for _, i := range []int{-1, 0, 1} {
		grid := pocket1[z+i]

		if grid != nil {
			active += countActiveGrid1(grid, x, y)
		}
	}

	if pocket1[z][y][x] == '#' {
		active--
	}

	return active
}

func getMapCopy1() map[int][][]byte {
	var nextPocket1 = map[int][][]byte{}

	for z, grid := range pocket1 {
		nextGrid := [][]byte{}

		for _, row := range grid {
			nextRow := []byte{}

			for x := range row {
				nextRow = append(nextRow, row[x])
			}

			nextGrid = append(nextGrid, nextRow)
		}

		nextPocket1[z] = nextGrid
	}

	return nextPocket1
}

func cycle1() map[int][][]byte {
	nextPocket1 := getMapCopy1()

	for z, grid := range pocket1 {
		for y, row := range grid {
			for x := range row {
				count := countActive1(x, y, z)

				if count == 3 && pocket1[z][y][x] == '.' {
					nextPocket1[z][y][x] = '#'
				}

				if (count < 2 || count > 3) && pocket1[z][y][x] == '#' {
					nextPocket1[z][y][x] = '.'
				}
			}
		}
	}

	return nextPocket1
}

func countTotalActive1() int {
	total := 0

	for _, grid := range pocket1 {
		for _, row := range grid {
			total += strings.Count(string(row), "#")
		}
	}

	return total
}

func initPocket2(centralSlice []string, iterations int) {
	initialLength := len(centralSlice)
	maxSpan := initialLength + 2*iterations + 16
	maxHeight := 1 + 2*iterations + 16
	for k := -maxHeight / 2; k <= maxHeight/2; k++ {
		var volume = map[int][][]byte{}

		for i := -maxHeight / 2; i <= maxHeight/2; i++ {
			span := [][]byte{}

			for j := 0; j < maxSpan; j++ {
				span = append(span, []byte(strings.Repeat(".", maxSpan)))
			}

			volume[i] = span
		}

		pocket2[k] = volume
	}

	for i, s := range centralSlice {
		index := (maxSpan - initialLength) / 2
		pocket2[0][0][index+i] = []byte(strings.Repeat(".", index) + s + strings.Repeat(".", index))
	}
}

func countActiveGrid2(volume map[int][][]byte, x, y, z int) int {
	length := len(volume)

	if x > 0 && x < length-1 && y > 0 && y < length-1 && z > -length/2 && z < length/2 {
		return strings.Count(string(volume[z-1][y-1][x-1:x+2]), "#") +
			strings.Count(string(volume[z+0][y-1][x-1:x+2]), "#") +
			strings.Count(string(volume[z+1][y-1][x-1:x+2]), "#") +
			strings.Count(string(volume[z-1][y+0][x-1:x+2]), "#") +
			strings.Count(string(volume[z+0][y+0][x-1:x+2]), "#") +
			strings.Count(string(volume[z+1][y+0][x-1:x+2]), "#") +
			strings.Count(string(volume[z-1][y+1][x-1:x+2]), "#") +
			strings.Count(string(volume[z+0][y+1][x-1:x+2]), "#") +
			strings.Count(string(volume[z+1][y+1][x-1:x+2]), "#")
	}

	return 0
}

func countActive2(x, y, z, w int) int {
	active := 0

	for _, i := range []int{-1, 0, 1} {
		volume := pocket2[w+i]

		if volume != nil {
			active += countActiveGrid2(volume, x, y, z)
		}
	}

	if pocket2[w][z][y][x] == '#' {
		active--
	}

	return active
}

func getMapCopy2() map[int]map[int][][]byte {
	var nextPocket2 = map[int]map[int][][]byte{}

	for w, volume := range pocket2 {
		nextVolume := map[int][][]byte{}

		for i, grid := range volume {
			nextGrid := [][]byte{}

			for _, row := range grid {
				nextRow := []byte{}

				for x := range row {
					nextRow = append(nextRow, row[x])
				}

				nextGrid = append(nextGrid, nextRow)
			}

			nextVolume[i] = nextGrid
		}

		nextPocket2[w] = nextVolume
	}

	return nextPocket2
}

func cycle2() map[int]map[int][][]byte {
	nextPocket2 := getMapCopy2()

	for w, volume := range pocket2 {
		for z, grid := range volume {
			for y, row := range grid {
				for x := range row {
					count := countActive2(x, y, z, w)

					if count == 3 && pocket2[w][z][y][x] == '.' {
						nextPocket2[w][z][y][x] = '#'
					}

					if (count < 2 || count > 3) && pocket2[w][z][y][x] == '#' {
						nextPocket2[w][z][y][x] = '.'
					}
				}
			}
		}
	}

	return nextPocket2
}

func countTotalActive2() int {
	total := 0

	for _, volume := range pocket2 {
		for _, grid := range volume {
			for _, row := range grid {
				total += strings.Count(string(row), "#")
			}
		}
	}

	return total
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var centralSlice []string

	for scanner.Scan() {
		centralSlice = append(centralSlice, scanner.Text())
	}

	numCycles := 6

	initPocket1(centralSlice, numCycles)

	for i := 0; i < numCycles; i++ {
		pocket1 = cycle1()
	}

	fmt.Printf("Total active cubes (1): %d\n", countTotalActive1())

	initPocket2(centralSlice, numCycles)

	for i := 0; i < numCycles; i++ {
		pocket2 = cycle2()
	}

	fmt.Printf("Total active cubes (2): %d\n", countTotalActive2())
}
