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
}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	paths := make([][]pos, 0)
	minX, maxX := 2<<31, 0
	minY, maxY := 2<<31, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := make([]pos, 0)
		text := scanner.Text()
		parts := strings.Split(text, "->")
		for _, part := range parts {
			part = strings.Trim(part, " ")
			xy := strings.Split(part, ",")
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				log.Fatal(err)
			}
			minX = min(minX, x)
			maxX = max(maxX, x)
			minY = min(minY, y)
			maxY = max(maxY, y)
			path = append(path, pos{x, y})
		}
		paths = append(paths, path)
	}
	cave := drawCave(paths, minX, maxX, maxY)
	units := 0
	for {
		units++
		if dropSand(cave) {
			break
		}
	}
	fmt.Printf("units of sand = %d\n", units)
}

func drawCave(paths [][]pos, minX, maxX, maxY int) [][]rune {
	cave := make([][]rune, maxY+3)
	for y := 0; y <= maxY+2; y++ {
		cave[y] = make([]rune, minX+maxX+1)
		for x := 0; x <= maxX+minX; x++ {
			cave[y][x] = '.'
		}
	}
	for x := 0; x <= maxX+minX; x++ {
		cave[maxY+2][x] = '#'
	}
	for _, path := range paths {
		cur := path[0]
		for i := 1; i < len(path); i++ {
			pos := path[i]
			if pos.x == cur.x {
				if pos.y > cur.y {
					for y := cur.y; y <= pos.y; y++ {
						cave[y][cur.x] = '#'
					}
				} else {
					for y := pos.y; y <= cur.y; y++ {
						cave[y][cur.x] = '#'
					}
				}
			} else {
				if pos.x > cur.x {
					for x := cur.x; x <= pos.x; x++ {
						cave[cur.y][x] = '#'
					}
				} else {
					for x := pos.x; x <= cur.x; x++ {
						cave[cur.y][x] = '#'
					}
				}
			}
			cur = pos
		}
	}
	return cave
}

func dropSand(cave [][]rune) bool {
	// height := len(cave)
	width := len(cave[0])
	x, y := 500, 0
	for {
		if cave[y+1][x] == '.' {
			y++
		} else if x > 0 && cave[y+1][x-1] == '.' {
			y++
			x--
		} else if x+1 < width && cave[y+1][x+1] == '.' {
			y++
			x++
		} else {
			cave[y][x] = 'o'
			if x == 500 && y == 0 {
				return true
			} else {
				return false
			}
		}
	}
}
