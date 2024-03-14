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
	group := [3]string{}
	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		group[index] = scanner.Text()
		index++
		if index == 3 {
			for _, r := range group[0] {
				if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
					if r >= 'a' && r <= 'z' {
						sum += int(r - 'a' + 1)
					} else {
						sum += int(r - 'A' + 27)
					}
					break
				}
			}
			index = 0
		}
	}
	fmt.Printf("sum = %d\n", sum)
}
