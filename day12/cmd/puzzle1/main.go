package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	y int
	x int
}

type step struct {
	pos   pos
	depth int
}

var dy = [4]int{1, -1, 0, 0}
var dx = [4]int{0, 0, 1, -1}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	heightMap := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		heightMap = append(heightMap, []rune(scanner.Text()))
	}
	var start pos
	var end pos
	for y, row := range heightMap {
		for x, height := range row {
			if height == 'S' {
				start = pos{y, x}
				heightMap[y][x] = 'a'
			} else if height == 'E' {
				end = pos{y, x}
				heightMap[y][x] = 'z'
			}
		}
	}
	seen := map[pos]bool{start: true}
	queue := []step{{start, 1}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			next := pos{s.pos.y + dy[i], s.pos.x + dx[i]}
			if next.y < 0 || next.y >= len(heightMap) || next.x < 0 || next.x >= len(heightMap[0]) {
				continue
			}
			if _, ok := seen[next]; ok {
				continue
			}
			if heightMap[next.y][next.x]-heightMap[s.pos.y][s.pos.x] > 1 {
				continue
			}
			if next == end {
				fmt.Printf("found end in %d steps\n", s.depth)
				return
			}
			seen[next] = true
			queue = append(queue, step{next, s.depth + 1})
		}
	}
}
