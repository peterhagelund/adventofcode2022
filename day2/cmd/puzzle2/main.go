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

const (
	loss = iota
	draw
	win
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
	}
	outcomes := map[string]int{
		"X": loss,
		"Y": draw,
		"Z": win,
	}
	points := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		opponent := hands[fields[0]]
		outcome := outcomes[fields[1]]
		me := determineMyHand(opponent, outcome)
		points += determineOutcome(opponent, me)
	}
	fmt.Printf("points = %d\n", points)
}

func determineMyHand(opponent, outcome int) int {
	switch opponent {
	case rock:
		if outcome == loss {
			return scissors
		} else if outcome == draw {
			return rock
		} else {
			return paper
		}
	case paper:
		if outcome == loss {
			return rock
		} else if outcome == draw {
			return paper
		} else {
			return scissors
		}
	case scissors:
		if outcome == loss {
			return paper
		} else if outcome == draw {
			return scissors
		} else {
			return rock
		}
	}
	return 0
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
