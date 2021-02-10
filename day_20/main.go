package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func reverse(str string) string {
	var rev []rune

	for _, b := range str {
		rev = append([]rune{b}, rev...)
	}

	return string(rev)
}

func flip(piece []string) []string {
	var newPiece []string

	for _, p := range piece {
		newPiece = append(newPiece, reverse(p))
	}

	return newPiece
}

func rotate(piece []string) []string {
	var newPiece []string

	for i := 0; i < len(piece); i++ {
		var newRow []byte

		for j := len(piece[0]) - 1; j >= 0; j-- {
			newRow = append(newRow, piece[j][i])
		}

		newPiece = append(newPiece, string(newRow))
	}

	return newPiece
}

func orientate(piece []string, position int) []string {
	var newPiece []string

	switch (position + 8) % 8 {
	case 0:
		newPiece = piece
	case 1:
		newPiece = flip(piece)
	case 2:
		newPiece = rotate(piece)
	case 3:
		newPiece = flip(rotate(piece))
	case 4:
		newPiece = rotate(rotate(piece))
	case 5:
		newPiece = flip(rotate(rotate(piece)))
	case 6:
		newPiece = rotate(rotate(rotate(piece)))
	case 7:
		newPiece = flip(rotate(rotate(rotate(piece))))
	}

	return newPiece
}

func getBorders(piece []string) []string {
	var borders []string

	var leftBorder, rightBorder []byte

	for _, row := range piece {
		leftBorder = append(leftBorder, row[0])
		rightBorder = append(rightBorder, row[len(row)-1])
	}

	return append(borders,
		piece[0], reverse(piece[0]),
		reverse(string(rightBorder)), string(rightBorder),
		piece[len(piece)-1], reverse(piece[len(piece)-1]), string(leftBorder), reverse(string(leftBorder)))
}

func matchBorders(border1, border2 []string) (int, int) {
	for i, b1 := range border1 {
		for j, b2 := range border2 {
			if b1 == b2 {
				if i%2 == 0 {
					return i, j
				}
			}
		}
	}

	return -1, -1
}

func findCorners(pieces map[int][]string) []int {
	var corners []int

	for tile, piece := range pieces {
		borders := getBorders(piece)
		match := 0

		for t, p := range pieces {
			i, j := matchBorders(borders, getBorders(p))
			if t != tile && i != -1 && j != -1 {
				match++
			}
		}

		if match == 2 {
			corners = append(corners, tile)
		}
	}

	return corners
}

func prod(arr []int) int {
	prod := 1

	for _, a := range arr {
		prod *= a
	}

	return prod
}

func match(piece1, piece2 []string, orientation int) bool {
	side1, side2 := "1", "2"

	if orientation == 0 {
		side1 = piece1[0]
		side2 = orientate(piece2, 5)[0]
	}

	if orientation == 2 {
		side1 = orientate(piece1, 6)[0]
		side2 = orientate(piece2, 3)[0]
	}

	if orientation == 4 {
		side1 = orientate(piece1, 5)[0]
		side2 = piece2[0]
	}

	if orientation == 6 {
		side1 = orientate(piece1, 3)[0]
		side2 = orientate(piece2, 6)[0]
	}

	return side1 == side2
}

func indexOf(a int, arr []int) int {
	for i := range arr {
		if a == arr[i] {
			return i
		}
	}

	return -1
}

func countOcurrences(piece []string, substr string) int {
	count := 0

	for _, r := range piece {
		count += strings.Count(r, substr)
	}

	return count
}

func cutImageBorders(image map[int][]string, length int) []string {
	var cuttedImage []string

	for n := 0; n < length; n++ {
		for i := 1; i < len(image[0])-1; i++ {
			row := ""

			for j := n * length; j < (n+1)*length; j++ {
				if image[j] != nil {
					row += image[j][i][1 : len(image[0])-1]
				}
			}

			cuttedImage = append(cuttedImage, row)
		}
	}

	return cuttedImage
}

func findMonsters(image []string) []string {
	for n := 0; n < 8; n++ {
		rimg := orientate(image, n)

		for i := 1; i < len(image)-1; i++ {
			for j := 0; j < len(image)-19; j++ {
				if rimg[i][j] == '#' &&
					string(rimg[i][j+5:j+7]) == "##" &&
					string(rimg[i][j+11:j+13]) == "##" &&
					string(rimg[i][j+17:j+20]) == "###" &&
					rimg[i-1][j+18] == '#' &&
					rimg[i+1][j+1] == '#' &&
					rimg[i+1][j+4] == '#' &&
					rimg[i+1][j+7] == '#' &&
					rimg[i+1][j+10] == '#' &&
					rimg[i+1][j+13] == '#' &&
					rimg[i+1][j+16] == '#' {

					row := []byte(rimg[i])
					row[j] = 'O'
					row[j+5] = 'O'
					row[j+6] = 'O'
					row[j+11] = 'O'
					row[j+12] = 'O'
					row[j+17] = 'O'
					row[j+18] = 'O'
					row[j+19] = 'O'
					rimg[i] = string(row)

					row = []byte(rimg[i-1])
					row[j+18] = 'O'
					rimg[i-1] = string(row)

					row = []byte(rimg[i+1])
					row[j+1] = 'O'
					row[j+4] = 'O'
					row[j+7] = 'O'
					row[j+10] = 'O'
					row[j+13] = 'O'
					row[j+16] = 'O'
					rimg[i+1] = string(row)
				}
			}
		}

		if countOcurrences(rimg, "O") != 0 {
			return rimg
		}
	}

	return image
}

func matchPieces(pieces map[int][]string, corners []int, numRows int) map[int][]string {
	var image = map[int][]string{}
	var upLeftCornerTile int

	for _, c := range corners {
		matches := 0

		for _, p := range pieces {
			i, _ := matchBorders(getBorders(pieces[c]), getBorders(p))

			if i == 2 || i == 4 {
				matches++
			}
		}

		if matches == 2 {
			upLeftCornerTile = c
		}
	}

	image[0] = pieces[upLeftCornerTile]

	pivot := 0
	numImages := 1
	direction := 2
	goLeft := true

	insertedTiles := []int{upLeftCornerTile}

	for numImages < numRows*numRows {
		tryNext := false

		for t, p := range pieces {
			for or := 0; or < 8; or++ {
				if indexOf(t, insertedTiles) == -1 && match(image[pivot], orientate(p, or), direction) {
					if direction == 2 {
						image[pivot+1] = orientate(p, or)
						pivot++
					} else if direction == 4 {
						image[pivot+numRows] = orientate(p, or)
						pivot += numRows
					} else if direction == 6 {
						image[pivot-1] = orientate(p, or)
						pivot--
					}

					insertedTiles = append(insertedTiles, t)
					numImages++
					tryNext = true
					break
				}
			}

			if tryNext {
				break
			}
		}

		if !tryNext && (direction == 2 || direction == 6) {
			direction = 4
			continue
		}

		if direction == 4 {
			if goLeft {
				direction = 6
			} else {
				direction = 2
			}

			goLeft = !goLeft
			continue
		}
	}

	return image
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pieces = map[int][]string{}
	numRows := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Tile") {
			num, _ := strconv.Atoi(line[5:9])
			pieces[num] = []string{}

			for scanner.Scan() {
				row := scanner.Text()

				if len(row) == 0 {
					break
				}

				pieces[num] = append(pieces[num], row)
			}

			numRows++
		}
	}

	numRows = int(math.Sqrt(float64(numRows)))

	corners := findCorners(pieces)

	fmt.Printf("Product of corner tiles (1): %d\n", prod(corners))

	image := matchPieces(pieces, corners, numRows)

	cuttedImage := cutImageBorders(image, numRows)

	monstersImage := findMonsters(cuttedImage)

	fmt.Printf("Number of # in the monsters image (2): %d\n", countOcurrences(monstersImage, "#"))
}
