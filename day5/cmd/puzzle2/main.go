package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type move struct {
	count  int
	source int
	target int
}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	temp := make([]string, 0)
	var stacks [][]rune
	var moves []*move
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			slices.Reverse(temp)
			stacks = parseStacks(temp)
			temp = make([]string, 0)
			continue
		}
		temp = append(temp, text)
	}
	moves = parseMoves(temp)
	message := makeMoves(stacks, moves)
	fmt.Printf("message = %s\n", string(message))
}

func parseStacks(input []string) [][]rune {
	fields := strings.Fields(input[0])
	stacks := make([][]rune, len(fields))
	for _, s := range input[1:] {
		layer := []rune(s)
		for i := 0; i < len(stacks); i++ {
			if stacks[i] == nil {
				stacks[i] = make([]rune, 0, 32)
			}
			r := layer[1+4*i]
			if r != ' ' {
				stacks[i] = append(stacks[i], r)
			}
		}
	}
	return stacks
}

func parseMoves(input []string) []*move {
	moves := make([]*move, 0)
	values := [3]int{}
	for _, s := range input {
		fields := strings.Fields(s)
		for i := 0; i < 3; i++ {
			value, err := strconv.Atoi(fields[2*i+1])
			if err != nil {
				log.Fatal(err)
			}
			values[i] = value
		}
		moves = append(moves, &move{values[0], values[1], values[2]})
	}
	return moves
}

func makeMoves(stacks [][]rune, moves []*move) string {
	for _, move := range moves {
		source := move.source - 1
		target := move.target - 1
		stacks[target] = append(stacks[target], stacks[source][len(stacks[source])-move.count:]...)
		stacks[source] = stacks[source][:len(stacks[source])-move.count]
	}
	message := make([]rune, len(stacks))
	for i := 0; i < len(stacks); i++ {
		message[i] = stacks[i][len(stacks[i])-1]
	}
	return string(message)
}
