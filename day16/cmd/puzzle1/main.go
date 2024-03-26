package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	name      string
	flowRate  int
	neighbors []string
	index     int
}

type travel struct {
	name     string
	distance int
}

type entry struct {
	name    string
	minutes int
	bitMask int
}

func main() {
	valves := parseInput("puzzle_input.txt")
	distances := calcDistances(valves)
	cache := make(map[entry]int)
	maxValue := depthFirstSearch(entry{"AA", 30, 0}, valves, distances, cache)
	fmt.Printf("max value = %d\n", maxValue)
}

func parseInput(name string) map[string]*valve {
	re := regexp.MustCompile(`Valve (?P<valve>[A-Z][A-Z]) has flow rate=(?P<rate>[0-9]+); tunnel(s?) lead(s?) to valve(s?) (?P<valves>[A-Z,\s]+)`)
	valveIndex := re.SubexpIndex("valve")
	rateIndex := re.SubexpIndex("rate")
	valvesIndex := re.SubexpIndex("valves")
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	valves := make(map[string]*valve)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		name := matches[valveIndex]
		flowRate, err := strconv.Atoi(matches[rateIndex])
		if err != nil {
			log.Fatal(err)
		}
		parts := strings.Split(matches[valvesIndex], ",")
		neighbors := make([]string, len(parts))
		for i, part := range parts {
			neighbors[i] = strings.Trim(part, " ")
		}
		valves[name] = &valve{
			name:      name,
			flowRate:  flowRate,
			neighbors: neighbors,
			index:     0,
		}
	}
	return valves
}

func calcDistances(valves map[string]*valve) map[string]map[string]int {
	distances := make(map[string]map[string]int)
	functioningValves := make([]string, 0)
	for name, valve := range valves {
		if name != "AA" && valve.flowRate == 0 {
			continue
		}
		distances[name] = make(map[string]int)
		if name != "AA" {
			functioningValves = append(functioningValves, name)
		}
		distances[name][name] = 0
		distances[name]["AA"] = 0
		visited := make(map[string]bool)
		visited[name] = true
		queue := make([]*travel, 0)
		queue = append(queue, &travel{name, 0})
		for len(queue) > 0 {
			t := queue[0]
			queue = queue[1:]
			currentValve := valves[t.name]
			for _, neighbor := range currentValve.neighbors {
				if _, ok := visited[neighbor]; ok {
					continue
				}
				visited[neighbor] = true
				neighborValve := valves[neighbor]
				if neighborValve.flowRate > 0 {
					distances[name][neighbor] = t.distance + 1
				}
				queue = append(queue, &travel{neighbor, t.distance + 1})
			}
		}
		delete(distances[name], name)
		if name != "AA" {
			delete(distances[name], "AA")
		}
	}
	for index, name := range functioningValves {
		valves[name].index = index
	}
	return distances
}

func depthFirstSearch(e entry, valves map[string]*valve, distances map[string]map[string]int, cache map[entry]int) int {
	if maxValue, ok := cache[e]; ok {
		return maxValue
	}
	maxValue := 0
	for name, distance := range distances[e.name] {
		valve := valves[name]
		bit := 1 << valve.index
		if e.bitMask&bit != 0 {
			continue
		}
		remainingMinutes := e.minutes - distance - 1
		if remainingMinutes <= 0 {
			continue
		}
		maxValue = max(maxValue, depthFirstSearch(entry{name, remainingMinutes, e.bitMask | bit}, valves, distances, cache)+valve.flowRate*remainingMinutes)
	}
	return maxValue
}
