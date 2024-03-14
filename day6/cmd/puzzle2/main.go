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
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := []rune(scanner.Text())
	for i := 0; i < len(input)-13; i++ {
		unique := true
		for j := 0; unique && j < 14; j++ {
			for k := j + 1; unique && k < 14; k++ {
				unique = input[i+j] != input[i+k]
			}
		}
		if unique {
			fmt.Printf("count = %d\n", i+14)
			break
		}
	}
}
