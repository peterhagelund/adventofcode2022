package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type sensor struct {
	sx       int
	sy       int
	bx       int
	by       int
	distance int
}

var re = regexp.MustCompile(`Sensor at x=(?P<sx>-?[0-9]+), y=(?P<sy>-?[0-9]+): closest beacon is at x=(?P<bx>-?[0-9]+), y=(?P<by>-?[0-9]+)`)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sensors := make([]*sensor, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindStringSubmatch(text)
		sx, sy := parsePos(matches[1], matches[2])
		bx, by := parsePos(matches[3], matches[4])
		distance := calcDistance(sx, sy, bx, by)
		s := &sensor{sx, sy, bx, by, distance}
		sensors = append(sensors, s)
	}
	count := 0
	count += findInvalidBeaconLocations(sensors, 2_000_000, -1)
	count += findInvalidBeaconLocations(sensors, 2_000_000, +1)
	fmt.Printf("count = %d\n", count)
}

func parsePos(s1, s2 string) (int, int) {
	x, err := strconv.Atoi(s1)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(s2)
	if err != nil {
		log.Fatal(err)
	}
	return x, y
}

func calcDistance(sx, sy, bx, by int) int {
	return (max(sx, bx) - min(sx, bx)) + (max(sy, by) - min(sy, by))
}

func findBeaconX(sensors []*sensor, by int) int {
	for _, s := range sensors {
		if s.by == by {
			return s.bx
		}
	}
	return -1
}

func findInvalidBeaconLocations(sensors []*sensor, by int, delta int) int {
	count := 0
	bx := findBeaconX(sensors, by) + delta
	for {
		valid := true
		for _, s := range sensors {
			distance := calcDistance(s.sx, s.sy, bx, by)
			if distance <= s.distance {
				valid = false
				break
			}
		}
		if valid {
			return count
		}
		count++
		bx += delta
	}
}
