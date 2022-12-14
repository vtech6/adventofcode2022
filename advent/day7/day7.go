package day7

import (
	"adventofcode/advent"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	testInput := advent.ReadInput("day7")
	formattedInput := formatInput(testInput.Split("\n"))
	makeDirs(formattedInput)
}

type Node struct {
	contains []string
	size     int
}

func makeDirs(input [][]string) {
	fileSystem := make(map[string]Node)
	currentDirectory := ""
	totalSizePart2 := 0
	for rowIndex, row := range input {
		if row[0] == "cd" {
			if row[1] == ".." {
				splitDirs := strings.Split(currentDirectory, "/")
				currentDirectory = strings.Join(splitDirs[:len(splitDirs)-1], "/")
			} else {
				if rowIndex != 0 {
					currentDirectory += fmt.Sprintf("/%s", row[1])
				}
			}
		} else if row[0] == "dir" {

			if value, ok := fileSystem[currentDirectory]; ok {
				newContains := append(value.contains, row[1])
				fileSystem[currentDirectory] = Node{contains: newContains, size: value.size}
			} else {
				fileSystem[currentDirectory] = Node{contains: []string{row[1]}}
			}

		} else if row[0] != "ls" {
			conv, _ := strconv.Atoi(row[0])
			totalSizePart2 += conv
			if value, ok := fileSystem[currentDirectory]; ok {
				newSize := value.size + conv
				fileSystem[currentDirectory] = Node{contains: value.contains, size: newSize}
			} else {
				fileSystem[currentDirectory] = Node{size: conv, contains: value.contains}
			}
			splitDirs := strings.Split(currentDirectory, "/")
			for itemIndex, _ := range splitDirs {
				joined := strings.Join(splitDirs[:itemIndex], "/")
				if value, ok := fileSystem[joined]; ok {
					newSize := value.size + conv
					fileSystem[joined] = Node{contains: value.contains, size: newSize}
				}
			}
		}
	}
	freeSpace := 70000000 - totalSizePart2
	smallestDir := 30000000
	totalSizePart1 := 0
	for _, value := range fileSystem {
		if value.size < 100000 {
			totalSizePart1 += value.size
		}
		if freeSpace+value.size > 30000000 {
			if value.size < smallestDir {
				smallestDir = value.size
			}
		}
	}
	fmt.Printf("Smallest dir size: %v\n", smallestDir)
}
func formatInput(input []string) [][]string {
	newInput := [][]string{}
	for _, row := range input {
		newRow := []string{}
		splitRow := strings.Split(row, " ")
		for _, element := range splitRow {
			if !(string(element) == "$") {
				newRow = append(newRow, string(element))
			}
		}
		newInput = append(newInput, newRow)
	}
	return newInput
}
