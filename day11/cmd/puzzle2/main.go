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

type monkey struct {
	id          int
	items       []int
	operator    rune
	operand     string
	value       int
	divisor     int
	trueMonkey  int
	falseMonkey int
	inspections int
}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	monkeys := parseMonkeys(file)
	modulo := 1
	for _, monkey := range monkeys {
		modulo *= monkey.divisor
	}
	for round := 1; round <= 10000; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				var worryLevel int
				var value int
				if monkey.operand == "old" {
					value = item
				} else {
					value = monkey.value
				}
				if monkey.operator == '*' {
					worryLevel = item * value
				} else {
					worryLevel = item + value
				}
				worryLevel %= modulo
				var index int
				if worryLevel%monkey.divisor == 0 {
					index = monkey.trueMonkey
				} else {
					index = monkey.falseMonkey
				}
				monkeys[index].items = append(monkeys[index].items, worryLevel)
				monkey.inspections++
			}
			monkey.items = make([]int, 0)
		}
	}
	inspections := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspections[i] = monkey.inspections
	}
	slices.Sort(inspections)
	slices.Reverse(inspections)
	monkeyBusiness := inspections[0] * inspections[1]
	fmt.Printf("monkey business = %d\n", monkeyBusiness)
}

func parseMonkeys(file *os.File) []*monkey {
	monkeys := make([]*monkey, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "Monkey ") {
			m := parseMonkey(scanner)
			monkeys = append(monkeys, m)
		}
	}
	return monkeys
}

func parseMonkey(scanner *bufio.Scanner) *monkey {
	m := &monkey{}
	var text string
	for i := 0; i < 6; i++ {
		text = scanner.Text()
		switch i {
		case 0:
			id, err := strconv.Atoi(text[7:8])
			if err != nil {
				log.Fatal(err)
			}
			m.id = id
		case 1:
			parts := strings.Split(text[18:], ",")
			m.items = make([]int, len(parts))
			for j, part := range parts {
				item, err := strconv.Atoi(strings.TrimSpace(part))
				if err != nil {
					log.Fatal(err)
				}
				m.items[j] = item
			}
		case 2:
			fields := strings.Fields(text)
			m.operator = []rune(fields[4])[0]
			m.operand = fields[5]
			if m.operand != "old" {
				value, err := strconv.Atoi(fields[5])
				if err != nil {
					log.Fatal(err)
				}
				m.value = value
			}
		case 3:
			fields := strings.Fields(text)
			divisor, err := strconv.Atoi(fields[3])
			if err != nil {
				log.Fatal(err)
			}
			m.divisor = divisor
		case 4:
			fields := strings.Fields(text)
			trueMonkey, err := strconv.Atoi(fields[5])
			if err != nil {
				log.Fatal(err)
			}
			m.trueMonkey = trueMonkey
		case 5:
			fields := strings.Fields(text)
			falseMonkey, err := strconv.Atoi(fields[5])
			if err != nil {
				log.Fatal(err)
			}
			m.falseMonkey = falseMonkey
		}
		scanner.Scan()
	}
	return m
}
