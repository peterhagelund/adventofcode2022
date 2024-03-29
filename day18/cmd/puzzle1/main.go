package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
	z int
}

var deltas = [][]int{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	positions := make([]*pos, 0)
	maxX, maxY, maxZ := 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		values := [3]int{}
		for i := 0; i < 3; i++ {
			values[i], err = strconv.Atoi(parts[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		x, y, z := values[0], values[1], values[2]
		maxX = max(maxX, x)
		maxY = max(maxY, y)
		maxZ = max(maxZ, z)
		positions = append(positions, &pos{x, y, z})
	}
	grid := make([][][]bool, maxZ+1)
	for z := 0; z < len(grid); z++ {
		grid[z] = make([][]bool, maxY+1)
		for y := 0; y < len(grid[z]); y++ {
			grid[z][y] = make([]bool, maxX+1)
		}
	}
	for _, p := range positions {
		grid[p.z][p.y][p.x] = true
	}
	surfaceArea := 0
	for _, position := range positions {
		for _, delta := range deltas {
			x, y, z := position.x+delta[0], position.y+delta[1], position.z+delta[2]
			if z < 0 || z == len(grid) || y < 0 || y == len(grid[0]) || x < 0 || x == len(grid[0][0]) {
				surfaceArea++
				continue
			}
			if grid[z][y][x] {
				continue
			}
			surfaceArea++
		}
	}
	fmt.Printf("surface area = %d\n", surfaceArea)
}
