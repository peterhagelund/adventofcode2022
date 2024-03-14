package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		text := scanner.Text()
		l := len(text) / 2
		c1 := text[:l]
		c2 := text[l:]
		for _, r := range c1 {
			if strings.ContainsRune(c2, r) {
				if r >= 'a' && r <= 'z' {
					sum += int(r - 'a' + 1)
				} else {
					sum += int(r - 'A' + 27)
				}
				break
			}
		}
	}
	fmt.Printf("sum = %d\n", sum)
}
