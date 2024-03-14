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
	scanner := bufio.NewScanner(file)
	maxCalories := 0
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
			maxCalories = max(maxCalories, elfCalories)
			elfCalories = 0
		}
	}
	if elfCalories > 0 {
		maxCalories = max(maxCalories, elfCalories)
	}
	fmt.Printf("max calories = %d\n", maxCalories)
}
