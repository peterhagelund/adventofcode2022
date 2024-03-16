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
	x := 1
	cycleNumber := 0
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var cycles int
		var value int
		fields := strings.Fields(scanner.Text())
		if fields[0] == "noop" {
			cycles = 1
		} else if fields[0] == "addx" {
			cycles = 2
			value, err = strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}
		}
		for i := 0; i < cycles; i++ {
			cycleNumber++
			if cycleNumber == 20 || (cycleNumber-20)%40 == 0 {
				sum += cycleNumber * x
			}
		}
		if cycles == 2 {
			x += value
		}
	}
	fmt.Printf("sum = %d\n", sum)
}
