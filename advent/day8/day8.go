package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	// input := advent.ReadInput("day8")
	// splitInput := input.Split("\n")
	// testInput := strings.Split("30373\n25512\n65332\n33549\n35390", "\n")
	testInput2 := strings.Split("111111\n1543451\n1443441\n1543451\n111111", "\n")
	formattedInput := formatInput(testInput2)
	findVisibleTrees(formattedInput)
}

func formatInput(input []string) [][]int {
	formattedInput := [][]int{}
	for _, row := range input {
		formattedRow := []int{}
		for _, item := range row {
			convItem, _ := strconv.Atoi(string(item))
			formattedRow = append(formattedRow, convItem)
		}
		formattedInput = append(formattedInput, formattedRow)
	}
	return formattedInput
}

func findVisibleTrees(trees [][]int) {
	fmt.Println(trees)
	visibleTrees := [][]bool{}
	for vIndex, row := range trees {
		visibleTrees = append(visibleTrees, []bool{})
		for i := 0; i < len(row); i++ {
			if vIndex == 0 || vIndex == len(trees)-1 || i == 0 || i == len(row)-1 {
				visibleTrees[vIndex] = append(visibleTrees[vIndex], true)
			} else {
				visibleTrees[vIndex] = append(visibleTrees[vIndex], false)
			}
		}
	}

	for verticalIndex, row := range trees {
		tallestTreesH := []int{0, 0}
		for horizontalIndex, _ := range row {
			if horizontalIndex < len(row)/2+1 {
				leftTree := row[horizontalIndex]
				lastIndex := len(row) - horizontalIndex - 1
				rightTree := row[lastIndex]
				if leftTree > tallestTreesH[0] {
					visibleTrees[verticalIndex][horizontalIndex] = true
					tallestTreesH[0] = leftTree

				}
				if rightTree > tallestTreesH[1] {
					visibleTrees[verticalIndex][lastIndex] = true
					tallestTreesH[1] = rightTree
				}
			}
		}
	}
	for column := 0; column < len(trees[0]); column++ {
		tallestTreesV := []int{0, 0}
		for row := 0; row < len(trees); row++ {
			if column < len(trees)/2+1 {
				firstTree := trees[column][row]
				lastIndex := len(trees) - 1
				lastTree := trees[column][lastIndex]
				if firstTree > tallestTreesV[0] {
					visibleTrees[column][row] = true
					tallestTreesV[0] = trees[column][row]
				}
				if lastTree > tallestTreesV[1] {
					visibleTrees[column][lastIndex] = true
					tallestTreesV[1] = trees[column][lastIndex]
				}
			}
		}
	}

	totalVisible := 0
	for _, row := range visibleTrees {
		for _, item := range row {
			if item {
				totalVisible += 1
			}
		}
	}
	fmt.Printf("%v\n", visibleTrees)
	fmt.Printf("%v\n", totalVisible)
}

func findHighestInRow() {

}

// [3 0 3 7 3]
// [2 5 5 1 2]
// [6 5 3 3 2]
// [3 3 5 4 9]
// [3 5 3 9 0]

// [[true true true true true]
// [true true true false true]
// [true false false true true]
// [true false true true true]
// [true true false true true]]

// [[1 1 1 1 1 1]
// [1 5 4 3 4 5 1]
// [1 4 4 3 4 4 1]
// [1 5 4 3 4 5 1]
// [1 1 1 1 1 1]]

// [true true true true true true]
// [true true false false false true true]
// [true true false false false true true]
// [true true false false false true true]
// [true true true true true true]]
