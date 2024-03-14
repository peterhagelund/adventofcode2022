package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock = iota
	paper
	scissors
)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hands := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}
	points := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		opponent := hands[fields[0]]
		me := hands[fields[1]]
		points += determineOutcome(opponent, me)
	}
	fmt.Printf("points = %d\n", points)
}

func determineOutcome(opponent, me int) int {
	switch opponent {
	case rock:
		if me == rock {
			return (1 + 3)
		} else if me == paper {
			return (2 + 6)
		} else if me == scissors {
			return (3 + 0)
		}
	case paper:
		if me == rock {
			return (1 + 0)
		} else if me == paper {
			return (2 + 3)
		} else if me == scissors {
			return (3 + 6)
		}
	case scissors:
		if me == rock {
			return (1 + 6)
		} else if me == paper {
			return (2 + 0)
		} else if me == scissors {
			return (3 + 3)
		}
	}
	return 0
}
