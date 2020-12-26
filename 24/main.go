package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	x, y int
}

func (t tile) String() string {
	return strconv.Itoa(t.x) + "," + strconv.Itoa(t.y)
}

const (
	black bool = true
	white bool = false

	e  string = "e"
	ne string = "ne"
	se string = "se"
	w  string = "w"
	sw string = "sw"
	nw string = "nw"
)

var directions = []string{e, ne, se, w, sw, nw}

var dirs = map[string]tile{
	e:  {2, 0},
	ne: {1, 1},
	se: {1, -1},
	w:  {-2, 0},
	sw: {-1, -1},
	nw: {-1, 1},
}

func toPosition(position string) tile {
	coord := strings.Split(position, ",")

	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])

	return tile{x, y}
}

func dailyFlip(tiles map[string]bool) map[string]bool {
	blackNeighbours := make(map[string]int)

	for pos, color := range tiles {
		if color == black {
			pTile := toPosition(pos)

			for _, d := range dirs {
				t := tile{pTile.x + d.x, pTile.y + d.y}
				blackNeighbours[t.String()]++
			}
		}
	}

	newBlackTiles := make(map[string]bool)

	for pos, count := range blackNeighbours {
		if (tiles[pos] == black && !(count == 0 || count > 2)) || (tiles[pos] == white && count == 2) {
			newBlackTiles[pos] = black
		}
	}

	return newBlackTiles
}

func processDays(tiles map[string]bool, days int) {
	result := 0

	for _, color := range tiles {
		if color == black {
			result++
		}
	}

	fmt.Printf("Number of black tiles (1): %d\n", result)

	for day := 0; day < days; day++ {
		result = 0
		tiles = dailyFlip(tiles)

		for _, color := range tiles {
			if color == black {
				result++
			}
		}

	}

	fmt.Printf("Number of black tiles in day %d (2): %d\n", days, result)
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tiles := make(map[string]bool)

	for scanner.Scan() {
		rule := scanner.Text()
		pTile := tile{0, 0}

		for rule != "" {
			for _, d := range directions {
				if strings.HasPrefix(rule, d) {
					pTile.x += dirs[d].x
					pTile.y += dirs[d].y

					if len(rule) == len(d) {
						tiles[pTile.String()] = !tiles[pTile.String()]
					}

					rule = rule[len(d):]
				}
			}

		}
	}

	processDays(tiles, 100)
}
