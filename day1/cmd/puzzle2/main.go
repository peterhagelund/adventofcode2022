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

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	allCalories := make([]int, 0)
	elfCalories := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if len(text) > 0 {
			calories, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			elfCalories += calories
		} else {
			allCalories = append(allCalories, elfCalories)
			elfCalories = 0
		}
	}
	if elfCalories > 0 {
		allCalories = append(allCalories, elfCalories)
	}
	slices.Sort(allCalories)
	slices.Reverse(allCalories)
	topThreeCalories := allCalories[0] + allCalories[1] + allCalories[2]
	fmt.Printf("top three calories = %d\n", topThreeCalories)
}
