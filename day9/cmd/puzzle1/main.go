package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	y int
	x int
}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	head := pos{0, 0}
	tail := pos{0, 0}
	positions := map[pos]bool{tail: true}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		direction := []rune(fields[0])[0]
		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		for s := 0; s < steps; s++ {
			switch direction {
			case 'U':
				head.y++
			case 'L':
				head.x--
			case 'D':
				head.y--
			case 'R':
				head.x++
			default:
				log.Fatal(fmt.Errorf("unknown direction %c", direction))
			}
			dy := head.y - tail.y
			dx := head.x - tail.x
			if int(math.Abs(float64(dy))) > 1 || int(math.Abs(float64(dx))) > 1 {
				if dx == 0 {
					tail.y += dy / 2
				} else if dy == 0 {
					tail.x += dx / 2
				} else {
					if dy > 0 {
						tail.y++
					} else {
						tail.y--
					}
					if dx > 0 {
						tail.x++
					} else {
						tail.x--
					}
				}
				positions[tail] = true
			}
		}
	}
	fmt.Printf("number of positions = %d\n", len(positions))
}
