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
	grid := make([][][]rune, maxZ+1)
	for z := 0; z < len(grid); z++ {
		grid[z] = make([][]rune, maxY+1)
		for y := 0; y < len(grid[z]); y++ {
			grid[z][y] = make([]rune, maxX+1)
			for x := 0; x < len(grid[z][y]); x++ {
				grid[z][y][x] = ' '
			}
		}
	}
	for _, p := range positions {
		grid[p.z][p.y][p.x] = '#'
	}
	surfaceArea := 0
	for _, position := range positions {
		for _, delta := range deltas {
			x, y, z := position.x+delta[0], position.y+delta[1], position.z+delta[2]
			if z < 0 || z == len(grid) || y < 0 || y == len(grid[0]) || x < 0 || x == len(grid[0][0]) {
				surfaceArea++
				continue
			}
			if grid[z][y][x] == '#' {
				continue
			}
			surfaceArea++
		}
	}
	for z := 1; z < len(grid)-1; z++ {
		for y := 1; y < len(grid[0])-1; y++ {
			for x := 1; x < len(grid[0][0])-1; x++ {
				if grid[z][y][x] != ' ' {
					continue
				}
				area := determineBubbleArea(grid, x, y, z)
				surfaceArea -= area
			}
		}
	}
	fmt.Printf("surface area = %d\n", surfaceArea)
}

func determineBubbleArea(grid [][][]rune, x, y, z int) int {
	p := &pos{x, y, z}
	bubble := make([]*pos, 0)
	visited := map[pos]bool{*p: true}
	queue := []*pos{p}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		bubble = append(bubble, p)
		for _, delta := range deltas {
			x = p.x + delta[0]
			y = p.y + delta[1]
			z = p.z + delta[2]
			if z < 0 || z == len(grid) || y < 0 || y == len(grid[0]) || x < 0 || x == len(grid[0][0]) {
				return 0
			}
			if grid[z][y][x] != ' ' {
				continue
			}
			a := &pos{x, y, z}
			if _, ok := visited[*a]; ok {
				continue
			}
			visited[*a] = true
			queue = append(queue, a)
		}
	}
	area := 0
	for _, p := range bubble {
		grid[p.z][p.y][p.x] = 'B'
		for _, delta := range deltas {
			x = p.x + delta[0]
			y = p.y + delta[1]
			z = p.z + delta[2]
			if grid[z][y][x] == '#' {
				area++
			}
		}
	}
	return area
}
