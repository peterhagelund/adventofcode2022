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
	bestScenicScore := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			scenicScore := calcScenicScore(forrest, y, x)
			bestScenicScore = max(bestScenicScore, scenicScore)
		}
	}
	fmt.Printf("best scenic score = %d\n", bestScenicScore)
}

func calcScenicScore(forrest [][]rune, y, x int) int {
	height := len(forrest)
	width := len(forrest[0])
	if y == 0 || y+1 == height || x == 0 || x+1 == width {
		return 0
	}
	up, left, down, right := 0, 0, 0, 0
	if y > 0 {
		for i := y - 1; i >= 0; i-- {
			up++
			if forrest[i][x] >= forrest[y][x] {
				break
			}
		}
	}
	if x > 0 {
		for i := x - 1; i >= 0; i-- {
			left++
			if forrest[y][i] >= forrest[y][x] {
				break
			}
		}
	}
	if y+1 < height {
		for i := y + 1; i < height; i++ {
			down++
			if forrest[i][x] >= forrest[y][x] {
				break
			}
		}
	}
	if x+1 < width {
		for i := x + 1; i < width; i++ {
			right++
			if forrest[y][i] >= forrest[y][x] {
				break
			}
		}
	}
	return up * left * down * right
}
