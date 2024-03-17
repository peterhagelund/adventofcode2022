package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	starts := make([]pos, 0)
	var end pos
	for y, row := range heightMap {
		for x, height := range row {
			if height == 'S' {
				height = 'a'
				heightMap[y][x] = height
			} else if height == 'E' {
				height = 'z'
				heightMap[y][x] = height
				end = pos{y, x}
			}
			if height == 'a' {
				starts = append(starts, pos{y, x})
			}
		}
	}
	minSteps := int(math.Pow(2, 31))
	for _, start := range starts {
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
					minSteps = min(minSteps, s.depth)
					queue = queue[:0]
					break
				}
				seen[next] = true
				queue = append(queue, step{next, s.depth + 1})
			}
		}
	}
	fmt.Printf("fewest steps = %d\n", minSteps)
}
