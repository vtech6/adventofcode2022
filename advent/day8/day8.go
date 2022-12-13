package day8

import (
	"adventofcode/advent"
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func Solve() {
	input := advent.ReadInput("day8")
	splitInput := input.Split("\n")
	// testInput := strings.Split("30373\n25512\n65332\n33549\n35390", "\n")
	// testInput2 := strings.Split("1111111\n1543451\n1443441\n1543451\n1111111", "\n")
	// testInput3 := strings.Split("123\n456\n789", "\n")
	formattedInput := formatInput(splitInput)
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
	treesPivot := make([][]int, len(trees[0]))
	for index, _ := range treesPivot {
		treesPivot[index] = make([]int, len(trees))
	}

	for column := 0; column < len(trees[0]); column++ {
		for row := 0; row < len(trees); row++ {
			// treesPivot[0][0] = trees[len(trees)-1][0]
			treesPivot[column][row] = trees[len(trees)-row-1][column]
		}
	}
	for rowIndex, row := range trees {
		for itemIndex, item := range row {

			leftSlice := row[:itemIndex]
			rightSlice := row[itemIndex+1:]
			upSlice := treesPivot[itemIndex][:len(trees)-rowIndex-1]
			downSlice := treesPivot[itemIndex][len(trees)-rowIndex:]

			allSlices := [][]int{leftSlice, rightSlice, upSlice, downSlice}
			visibleInSlice := []bool{true, true, true, true}
			for sliceIndex, slice := range allSlices {
				for _, sliceItem := range slice {
					if sliceItem >= item {
						visibleInSlice[sliceIndex] = false
					}
				}
			}
			if itemIndex == 1 && rowIndex == 3 {
				fmt.Printf("Item: %v,Left: %v, Right: %v, Up: %v, Down: %v\n %v\n", item, leftSlice, rightSlice, upSlice, downSlice, visibleInSlice)
			}
			if lo.Contains(visibleInSlice, true) {
				visibleTrees[rowIndex][itemIndex] = true
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

// [1,2,3]
// [4,5,6]
// [7,8,9]
// [10,11,12]

// [0][0] = [2][0]
// [0][1] = [1][0]
// [0][2] = [0][0]
// [1][0] = [2][1]
// [1][1] = [1][1]
// [1][2] = [0][1]

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

// [[true true true true true]
//	[true true true false true]
// [true true false true true]
// [true true true false true]
// [true true true true true]]

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
