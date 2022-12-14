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

// func navigate(row string, commands [][]string, fileSystem map[string]interface{}) map[string]interface{} {
// 	currentDir := make(map[string]interface{})
// 	currentDir[row] = fileSystem

// 	return currentDir
// }

func navigate(path string, currentDir map[string]interface{}) map[string]interface{} {

	if path == ".." {
		fmt.Printf("navigate to .. (%v)\n", currentDir["_parent"].(map[string]interface{})["_name"])
		return currentDir["_parent"].(map[string]interface{})
	}

	if currentDir[path] == nil {
		child := make(map[string]interface{})
		child["_name"] = path
		child["_parent"] = currentDir
		child["_size"] = "0"
		currentDir[path] = child
	}

	fmt.Printf("navigate to %v\n", path)
	return currentDir[path].(map[string]interface{})

}

func addChild(row []string, currentDir map[string]interface{}) map[string]interface{} {

	child := make(map[string]interface{})
	child["_name"] = row[1]
	child["_parent"] = currentDir

	if row[0] != "dir" {
		child["_size"] = row[0]
	}

	currentDir[row[1]] = child
	fmt.Println("addedChild")
	return currentDir

}

func makeDirs(input [][]string) {
	fileSystem := make(map[string]interface{})
	currentDir := make(map[string]interface{})
	currentDir["_name"] = "root"
	currentDir["_size"] = "0"

	fileSystem["_name"] = "filesystem"
	fileSystem["_size"] = "0"
	fileSystem["root"] = currentDir

	for _, row := range input {
		fmt.Printf(">> %v\n", row)
		if row[0] == "cd" {
			currentDir = navigate(row[1], currentDir)
		} else if row[0] == "ls" {
			//We don't give a fuck
		} else {
			currentDir = addChild(row, currentDir)
		}
	}

	fmt.Println("------------------")
	debug(fileSystem)
	//fmt.Printf("size %v", size(fileSystem))
}

func debug(item map[string]interface{}) {
	path := pathRecursive(item, "")
	fmt.Printf("%v (%v)\n", path, size(item))
	for key, value := range item {
		if !strings.HasPrefix(key, "_") {
			debug(value.(map[string]interface{}))
		}
	}

}

func pathRecursive(item map[string]interface{}, path string) string {
	parent := item["_parent"]
	name := item["_name"]
	result := fmt.Sprintf("%v/%v", name, path)
	if parent != nil {
		return pathRecursive(parent.(map[string]interface{}), result)
	}

	return result
}

func isLeaf(item map[string]interface{}) bool {
	for key := range item {
		if !strings.HasPrefix(key, "_") {
			return false
		}
	}
	// fmt.Printf("Leaf %v\n", item["_name"])
	return true
}

func size(item map[string]interface{}) int {

	currentSize := 0
	if item["_size"] != nil {
		currentSize, _ = strconv.Atoi(item["_size"].(string))
	}

	if isLeaf(item) {
		return currentSize
	}
	childrenSize := 0
	for key, value := range item {
		if !strings.HasPrefix(key, "_") {
			childrenSize += size(value.(map[string]interface{}))
		}
	}

	return currentSize + childrenSize
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
