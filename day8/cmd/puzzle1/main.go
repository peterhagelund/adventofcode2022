package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	forrest := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		forrest = append(forrest, []rune(scanner.Text()))
	}
	height := len(forrest)
	width := len(forrest[0])
	visible := make(map[string]bool)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if isVisible(forrest, y, x) {
				tree := fmt.Sprintf("%d,%d", y, x)
				visible[tree] = true
			}
		}
	}
	fmt.Printf("count = %d\n", len(visible))
}

func isVisible(forrest [][]rune, y, x int) bool {
	height := len(forrest)
	width := len(forrest[0])
	if y == 0 || y+1 == height || x == 0 || x+1 == width {
		return true
	}
	visible := true
	for i := y - 1; visible && i >= 0; i-- {
		if forrest[i][x] >= forrest[y][x] {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for i := y + 1; i < height; i++ {
		if forrest[i][x] >= forrest[y][x] {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for i := x - 1; i >= 0; i-- {
		if forrest[y][i] >= forrest[y][x] {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for i := x + 1; i < width; i++ {
		if forrest[y][i] >= forrest[y][x] {
			visible = false
		}
	}
	return visible
}
