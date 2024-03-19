package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	packets := make([][]any, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		packet := parseList([]rune(text[1 : len(text)-1]))
		packets = append(packets, packet)
	}
	dividerPacket1 := []any{[]any{2}}
	dividerPacket2 := []any{[]any{6}}
	packets = append(packets, dividerPacket1)
	packets = append(packets, dividerPacket2)
	slices.SortFunc(packets, sortPacket)
	index1 := slices.IndexFunc(packets, func(packet []any) bool {
		return isPacket(packet, 2)
	})
	index2 := slices.IndexFunc(packets, func(packet []any) bool {
		return isPacket(packet, 6)
	})
	decoderKey := (index1 + 1) * (index2 + 1)
	fmt.Printf("decoder key = %d\n", decoderKey)
}

func parseList(text []rune) []any {
	list := make([]any, 0)
	index := 0
	for index < len(text) {
		if text[index] == '[' {
			level := 0
			i := 0
			for i = index; i < len(text); i++ {
				if text[i] == '[' {
					level++
				}
				if text[i] == ']' {
					level--
					if level == 0 {
						break
					}
				}
			}
			l := parseList(text[index+1 : i])
			list = append(list, l)
			index = i + 1
		} else if unicode.IsDigit(text[index]) {
			i := 0
			for i = index; i < len(text); i++ {
				if !unicode.IsDigit(text[i]) {
					break
				}
			}
			n, err := strconv.Atoi(string(text[index:i]))
			if err != nil {
				log.Fatal(err)
			}
			list = append(list, n)
			index = i
		} else if text[index] == ',' {
			index++
		} else {
			log.Fatal(fmt.Errorf("unexpexted run '%c' at index %d", text[index], index))
		}
	}
	return list
}

func sortPacket(a, b []any) int {
	return determineOrder(a, b)
}

func determineOrder(left, right any) int {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	if leftType.Kind() == reflect.Int && rightType.Kind() == reflect.Int {
		l, r := left.(int), right.(int)
		if l < r {
			return -1
		} else if l == r {
			return 0
		} else {
			return 1
		}
	}
	var l []any
	var r []any
	if leftType.Kind() == reflect.Int {
		l = []any{left.(int)}
	} else {
		l = left.([]any)
	}
	if rightType.Kind() == reflect.Int {
		r = []any{right.(int)}
	} else {
		r = right.([]any)
	}
	for i := 0; i < min(len(l), len(r)); i++ {
		switch determineOrder(l[i], r[i]) {
		case -1:
			return -1
		case 1:
			return 1
		default:
			continue
		}
	}
	if len(l) < len(r) {
		return -1
	} else if len(l) == len(r) {
		return 0
	}
	return 1
}

func isPacket(packet []any, value int) bool {
	if len(packet) != 1 {
		return false
	}
	list, ok := packet[0].([]any)
	if !ok {
		return false
	}
	if len(list) != 1 {
		return false
	}
	v, ok := list[0].(int)
	if !ok {
		return false
	}
	return v == value
}
