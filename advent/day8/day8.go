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
			treesPivot[column][row] = trees[len(trees)-row-1][column]
		}
	}
	highestScore := 1
	for rowIndex, row := range trees {
		for itemIndex, item := range row {
			leftSlice := row[:itemIndex]
			rightSlice := row[itemIndex+1:]
			upSlice := treesPivot[itemIndex][:len(trees)-rowIndex-1]
			downSlice := treesPivot[itemIndex][len(trees)-rowIndex:]
			allSlices := [][]int{leftSlice, rightSlice, upSlice, downSlice}
			visibleInSlice := []bool{true, true, true, true}
			normalSlices := [][]int{rightSlice, downSlice}
			inverseSlices := [][]int{leftSlice, upSlice}
			scenicScore := 1
			for _, nSlice := range normalSlices {
				containsHigher := false
				for nSliceItemIndex, nSliceItem := range nSlice {
					if !containsHigher {
						if nSliceItem >= item {
							containsHigher = true
							scenicScore *= (nSliceItemIndex + 1)
						}
					}
				}
				if !containsHigher {
					scenicScore *= len(nSlice)
				}
			}
			for _, inSlice := range inverseSlices {
				containsHigher := false
				for inIndex := len(inSlice) - 1; inIndex > 0; inIndex-- {
					if !containsHigher {
						if inSlice[inIndex] >= item {
							containsHigher = true
							scenicScore *= (len(inSlice) - inIndex)
						}
					}
				}
				if !containsHigher {
					scenicScore *= len(inSlice)
				}
			}
			for sliceIndex, slice := range allSlices {
				for _, sliceItem := range slice {
					if sliceItem >= item && visibleInSlice[sliceIndex] {
						visibleInSlice[sliceIndex] = false
					}
				}
			}
			if lo.Contains(visibleInSlice, true) {
				visibleTrees[rowIndex][itemIndex] = true
			}
			if scenicScore > highestScore {
				highestScore = scenicScore
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
	fmt.Printf("Total visible: %v\n", totalVisible)
	fmt.Printf("Highest score: %v\n", highestScore)
}
