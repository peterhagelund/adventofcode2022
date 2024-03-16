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
	knots := [10]pos{}
	positions := map[pos]bool{knots[9]: true}
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
				knots[0].y++
			case 'L':
				knots[0].x--
			case 'D':
				knots[0].y--
			case 'R':
				knots[0].x++
			default:
				log.Fatal(fmt.Errorf("unknown direction %c", direction))
			}
			for k := 1; k < 10; k++ {
				dy := knots[k-1].y - knots[k].y
				dx := knots[k-1].x - knots[k].x
				if int(math.Abs(float64(dy))) > 1 || int(math.Abs(float64(dx))) > 1 {
					if dx == 0 {
						knots[k].y += dy / 2
					} else if dy == 0 {
						knots[k].x += dx / 2
					} else {
						if dy > 0 {
							knots[k].y++
						} else {
							knots[k].y--
						}
						if dx > 0 {
							knots[k].x++
						} else {
							knots[k].x--
						}
					}
				}
			}
			positions[knots[9]] = true
		}
	}
	fmt.Printf("number of positions = %d\n", len(positions))
}
