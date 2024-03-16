package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	file = iota
	directory
)

type entry struct {
	parent   *entry
	kind     int
	name     string
	size     int
	children []*entry
}

func (e *entry) totalSize() int {
	if e.kind == file {
		return e.size
	} else {
		totalSize := 0
		for _, c := range e.children {
			totalSize += c.totalSize()
		}
		return totalSize
	}
}

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	root := &entry{
		parent:   nil,
		kind:     directory,
		name:     "/",
		size:     0,
		children: make([]*entry, 0),
	}
	dir := root
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "$ ") {
			dir = executeCommand(text[2:], dir)
		} else {
			parseOutput(text, dir)
		}
	}
	var totalSize int
	sumTotalSizes(root, &totalSize)
	fmt.Printf("total size = %d\n", totalSize)

}

func executeCommand(cmd string, dir *entry) *entry {
	if cmd == "ls" {
		return dir
	}
	if strings.HasPrefix(cmd, "cd ") {
		name := cmd[3:]
		switch name {
		case "/":
			for dir.parent != nil {
				dir = dir.parent
			}
			return dir
		case "..":
			return dir.parent
		default:
			for _, d := range dir.children {
				if d.name == name {
					return d
				}
			}
			log.Fatal(fmt.Errorf("directory '%s' unknown", name))
		}
	}
	return dir
}

func parseOutput(output string, dir *entry) {
	if strings.HasPrefix(output, "dir ") {
		name := output[4:]
		newDir := &entry{
			parent: dir,
			kind:   directory,
			name:   name,
			size:   0,
		}
		dir.children = append(dir.children, newDir)
	} else {
		fields := strings.Fields(output)
		size, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		name := fields[1]
		newFile := &entry{
			parent:   dir,
			kind:     file,
			name:     name,
			size:     size,
			children: nil,
		}
		dir.children = append(dir.children, newFile)
	}
}

func sumTotalSizes(e *entry, totalSize *int) {
	if e.totalSize() <= 100000 {
		*totalSize += e.totalSize()
	}
	for _, c := range e.children {
		if c.kind == directory {
			sumTotalSizes(c, totalSize)
		}
	}
}
