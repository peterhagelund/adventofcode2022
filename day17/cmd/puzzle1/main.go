package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var rocks = [][]uint8{
	{
		0b0011110,
	},
	{
		0b0001000,
		0b0011100,
		0b0001000,
	},
	{
		0b0000100,
		0b0000100,
		0b0011100,
	},
	{
		0b0010000,
		0b0010000,
		0b0010000,
		0b0010000,
	},
	{
		0b0011000,
		0b0011000,
	},
}

var filler = [7]uint8{0, 0, 0, 0, 0, 0, 0}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	pattern := []rune(scanner.Text())
	patternIndex := 0
	chamber := make([]uint8, 0)
	var rock []uint8
	rockIndex := 0
	rockCount := 0
	rockY := 0
	towerHeight := 0
	for {
		// New rock or drop it
		if rock == nil {
			rock = make([]uint8, len(rocks[rockIndex]))
			copy(rock, rocks[rockIndex])
			chamberHeight := towerHeight + 3 + len(rock)
			if len(chamber) < chamberHeight {
				chamber = append(chamber, filler[:chamberHeight-len(chamber)]...)
			} else if len(chamber) > chamberHeight {
				chamber = chamber[:chamberHeight]
			}
			rockY = len(chamber) - len(rock)
		} else {
			canDrop := true
			if rockY == 0 {
				canDrop = false
			} else {
				for i := 0; canDrop && i < len(rock); i++ {
					if chamber[rockY+len(rock)-2-i]^rock[i] != chamber[rockY+len(rock)-2-i]|rock[i] {
						canDrop = false
					}
				}
			}
			if !canDrop {
				for i := 0; i < len(rock); i++ {
					chamber[rockY+len(rock)-1-i] |= rock[i]
				}
				towerHeight = max(towerHeight, rockY+len(rock))
				rockCount++
				if rockCount == 2022 {
					break
				}
				rock = nil
				rockIndex = (rockIndex + 1) % len(rocks)
				continue
			} else {
				rockY--
			}
		}
		// Handle jet pattern
		if pattern[patternIndex] == '>' {
			canMove := true
			for i := 0; canMove && i < len(rock); i++ {
				if rock[i]&0b0000001 != 0 {
					canMove = false
				} else {
					if chamber[rockY+len(rock)-1-i]^(rock[i]>>1) != chamber[rockY+len(rock)-1-i]|(rock[i]>>1) {
						canMove = false
					}
				}
			}
			if canMove {
				for i := 0; i < len(rock); i++ {
					rock[i] >>= 1
				}
			}
		} else {
			canMove := true
			for i := 0; canMove && i < len(rock); i++ {
				if rock[i]&0b1000000 != 0 {
					canMove = false
				} else {
					if chamber[rockY+len(rock)-1-i]^(rock[i]<<1) != chamber[rockY+len(rock)-1-i]|(rock[i]<<1) {
						canMove = false
					}
				}
			}
			if canMove {
				for i := 0; i < len(rock); i++ {
					rock[i] <<= 1
				}
			}
		}
		patternIndex = (patternIndex + 1) % len(pattern)
	}
	fmt.Printf("height of tower after %d rocks = %d\n", rockCount, towerHeight)
}
