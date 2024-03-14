package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		start := [2]int{}
		end := [2]int{}
		for i, pair := range pairs {
			sections := strings.Split(pair, "-")
			start[i], err = strconv.Atoi(sections[0])
			if err != nil {
				log.Fatal(err)
			}
			end[i], err = strconv.Atoi(sections[1])
			if err != nil {
				log.Fatal(err)
			}
			if ((start[0] >= start[1] && start[0] <= end[1]) || (end[0] >= start[1] && end[0] <= end[1])) ||
				((start[1] >= start[0] && start[1] <= end[0]) || (end[1] >= start[0] && end[1] <= end[0])) {
				sum++
			}
		}
	}
	fmt.Printf("sum = %d\n", sum)
}
