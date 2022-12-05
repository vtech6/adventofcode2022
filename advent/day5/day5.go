package day5

import (
	"adventofcode/advent"
	"fmt"
	"strconv"
	"strings"
)

const (
	bigletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Day5() {
	input := advent.ReadInput("day5")
	formatInput(input.Split("\n"))
}

func formatInput(input []string) {
	crates, instructions := separateInstructions(input)
	fmt.Printf("Part 1: %v, Part 2: %v", rearrangeBoxes(crates, instructions, false), rearrangeBoxes(crates, instructions, true))
}

func shiftArray(crates []string) [][]string {
	// fmt.Printf("Before shifting: %v\n", crates)
	newCrates := [][]string{}
	for crateIndex, crate := range crates {
		for charIndex, character := range crate {
			if crateIndex == 0 {
				newCrates = append(newCrates, []string{string(character)})
			} else {
				newCrates[charIndex] = append(newCrates[charIndex], string(character))
			}
		}
	}
	// fmt.Printf("After shifting %v\n", newCrates)
	for arrayIndex, array := range newCrates {
		newArray := []string{}
		for i := len(array); i > 0; i-- {
			char := array[i-1]
			if char != "-" {
				newArray = append(newArray, array[i-1])
			}
		}
		newCrates[arrayIndex] = newArray
	}
	return newCrates
}

const (
	numbers = "123456789"
)

func formatInstructions(instructions []string) [][]int {
	intArray := [][]int{}
	for _, element := range instructions {
		formattedInstruction := []int{}
		for _, instruction := range strings.Split(element, " ") {
			instructionStr := string(instruction)
			if instructionStr != " " {
				errorCaught := false
				integer, err := strconv.Atoi(instruction)
				if err != nil {
					errorCaught = true
				}
				if !errorCaught {
					formattedInstruction = append(formattedInstruction, int(integer))
				}
			}
		}
		intArray = append(intArray, formattedInstruction)
	}
	return intArray
}

func formatCrates(crates []string) []string {
	formattedCrates := []string{}
	validIndices := []int{}
	for charI, char := range crates[len(crates)-1] {
		if strings.Contains(bigletters, string(char)) {
			validIndices = append(validIndices, charI)
		}
	}
	for _, crate := range crates {
		formattedCrate := ""
		for _, validInt := range validIndices {
			formattedCrate += string(crate[validInt])
		}
		formattedCrate = strings.Replace(formattedCrate, " ", "-", -1)
		formattedCrates = append(formattedCrates, formattedCrate)
	}
	return formattedCrates
}
func separateInstructions(input []string) ([]string, [][]int) {
	crates := []string{}
	instructions := []string{}
	for index, element := range input {
		if element == "" {
			instructions = input[index+1:]
			crates = input[:index-1]
		}
	}
	return formatCrates(crates), formatInstructions(instructions)
}

func rearrangeBoxes(boxes []string, instructions [][]int, stacked bool) string {

	rearrangedBoxes := shiftArray(boxes)
	for _, instruction := range instructions {
		numberOfItems := instruction[0]
		sourceColumn := instruction[1] - 1
		destColumn := instruction[2] - 1
		if !stacked {
			for i := numberOfItems + 1; i > 1; i-- {
				sourceIndex := len(rearrangedBoxes[sourceColumn]) - 1
				rearrangedDest := append(rearrangedBoxes[destColumn], rearrangedBoxes[sourceColumn][sourceIndex])
				rearrangedSource := append(rearrangedBoxes[sourceColumn][:sourceIndex], rearrangedBoxes[sourceColumn][sourceIndex+1:]...)
				rearrangedBoxes[destColumn] = rearrangedDest
				rearrangedBoxes[sourceColumn] = rearrangedSource
			}
		} else {
			rearrangedDest := rearrangedBoxes[destColumn]
			rearrangedSource := rearrangedBoxes[sourceColumn]
			slice := len(rearrangedBoxes[sourceColumn]) - numberOfItems
			// fmt.Printf("%v: Taking %v items from %v to %v\n", instructionIndex, numberOfItems, rearrangedDest, rearrangedSource)
			rearrangedDest = append(rearrangedDest, rearrangedSource[slice:]...)
			rearrangedSource = rearrangedSource[:slice]
			rearrangedBoxes[destColumn] = rearrangedDest
			rearrangedBoxes[sourceColumn] = rearrangedSource
		}
		// fmt.Printf("Rearranged: %v,\n", rearrangedBoxes)
	}
	answer := ""
	for _, arr := range rearrangedBoxes {
		answer += arr[len(arr)-1]
	}
	return answer
}
