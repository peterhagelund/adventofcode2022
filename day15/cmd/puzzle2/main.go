package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type sensor struct {
	sx       int
	sy       int
	bx       int
	by       int
	distance int
}

type interval struct {
	lowX  int
	highX int
}

var re = regexp.MustCompile(`Sensor at x=(?P<sx>-?[0-9]+), y=(?P<sy>-?[0-9]+): closest beacon is at x=(?P<bx>-?[0-9]+), y=(?P<by>-?[0-9]+)`)

const count = 4_000_000

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
	for y := 0; y <= count; y++ {
		intervals := make([]*interval, 0)
		for _, s := range sensors {
			offset := s.distance - int(math.Abs(float64(s.sy-y)))
			if offset < 0 {
				continue
			}
			lowX := s.sx - offset
			highX := s.sx + offset
			intervals = append(intervals, &interval{lowX, highX})
		}
		slices.SortFunc(intervals, func(a, b *interval) int {
			if a.lowX < b.lowX {
				return -1
			} else if a.lowX > b.lowX {
				return 1
			} else if a.highX < b.highX {
				return -1
			} else if a.highX > b.highX {
				return 1
			}
			return 0
		})
		queue := make([]*interval, 0)
		for _, interval := range intervals {
			if len(queue) == 0 {
				queue = append(queue, interval)
				continue
			}
			queueInterval := queue[len(queue)-1]
			if interval.lowX > queueInterval.lowX+1 {
				queue = append(queue, interval)
				continue
			}
			queue[len(queue)-1].highX = max(queueInterval.highX, interval.highX)
		}
		x := 0
		for _, q := range queue {
			if x < q.lowX {
				fmt.Printf("tuning frequency at %d, %d = %d\n", x, y, x*4_000_000+y)
				return
			}
			x = max(x, q.highX+1)
			if x > count {
				break
			}
		}
	}
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
