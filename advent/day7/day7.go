// package day7

// import (
// 	"adventofcode/advent"
// 	"fmt"
// 	"strings"
// )

// func Solve() {
// 	testInput := advent.ReadInput("day7")
// 	formattedInput := formatInput(testInput.Split("\n")[:10])
// 	makeDirs(formattedInput)
// }

// // func navigate(row string, commands [][]string, fileSystem map[string]interface{}) map[string]interface{} {
// // 	currentDir := make(map[string]interface{})
// // 	currentDir[row] = fileSystem

// // 	return currentDir
// // }

// func navigate(path string, currentDir map[string]interface{}) map[string]interface{} {

// 	if path == ".." {
// 		fmt.Println(path)
// 		return currentDir["_parent"].(map[string]interface{})
// 	}

// 	if currentDir[path] == nil {
// 		child := make(map[string]interface{})
// 		child["_name"] = path
// 		child["_parent"] = currentDir
// 		fmt.Println(child)
// 		currentDir[path] = child
// 	}

// 	fmt.Println(currentDir[path])
// 	return currentDir[path].(map[string]interface{})

// }

// func addChild(row []string, currentDir map[string]interface{}) map[string]interface{} {

// 	child := make(map[string]interface{})
// 	child["_name"] = row[1]
// 	child["_parent"] = currentDir

// 	if row[0] != "dir" {
// 		child["_size"] = row[0]
// 	}

// 	currentDir[row[1]] = child
// 	return currentDir

// }

// func makeDirs(input [][]string) {
// 	fileSystem := make(map[string]interface{})
// 	currentDir := make(map[string]interface{})
// 	currentDir["_name"] = "root"
// 	fileSystem["root"] = currentDir

// 	for _, row := range input {
// 		if row[0] == "cd" {
// 			currentDir = navigate(row[1], currentDir)
// 			fmt.Println(currentDir)
// 			// navigate(row[1], input[index:], fileSystem)
// 		} else if row[0] == "ls" {
// 			//We don't give a fuck
// 		} else {
// 			//currentDir = addChild(row, currentDir)
// 		}

// 	}

// 	//fmt.Println(fileSystem)
// }

// func formatInput(input []string) [][]string {
// 	newInput := [][]string{}
// 	for _, row := range input {
// 		newRow := []string{}
// 		splitRow := strings.Split(row, " ")
// 		for _, element := range splitRow {
// 			if !(string(element) == "$") {
// 				newRow = append(newRow, string(element))
// 			}
// 		}
// 		newInput = append(newInput, newRow)
// 	}
// 	return newInput
// }
