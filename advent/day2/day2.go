package day2

import (
	"adventofcode/advent"
	"fmt"
	"strings"
)

func Solve() {
	input := advent.ReadInput("day2")
	formattedInput := input.Split("\n")
	_, _, overallScorePart1, overallScorePart2 := splitElementsAndCountScore(formattedInput)
	fmt.Printf("Part 1 score: %v, Part 2 score: %v", overallScorePart1, overallScorePart2)
}

func part1(shapes []string) int {
	shapeValue := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	score := 0
	score += shapeValue[shapes[1]]
	transformedShapes := []int{shapeValue[shapes[0]], shapeValue[shapes[1]]}
	if shapes[0] == shapes[1] {
		score += 3
	}
	if (transformedShapes[1]-transformedShapes[0] == 1) || (transformedShapes[0] == 3 && transformedShapes[1] == 1) {
		score += 6
	}
	return score
}

func part2(shapes []string) []string {
	shapeValue := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	valueToShape := map[int]string{
		0: "C",
		1: "A",
		2: "B",
		3: "C",
		4: "A",
	}
	var newShape string
	if shapes[1] == "X" {
		newShape = valueToShape[shapeValue[shapes[0]]-1]
	}
	if shapes[1] == "Y" {
		newShape = shapes[0]
	}
	if shapes[1] == "Z" {
		newShape = valueToShape[shapeValue[shapes[0]]+1]
	}
	return []string{shapes[0], newShape}
}

func splitElementsAndCountScore(input []string) ([]interface{}, []interface{}, int, int) {
	cipher := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	decipheredArray := []interface{}{}
	arrayWithScores := []interface{}{}
	overallScorePart1 := 0
	overallScorePart2 := 0
	for _, element := range input {
		shapes := strings.Split(element, " ")
		winDrawLose := part2(shapes)
		decipheredShapes := []string{shapes[0], cipher[shapes[1]]}
		decipheredArray = append(decipheredArray, decipheredShapes)
		totalScorePart1 := part1(decipheredShapes)
		totalScorePart2 := part1(winDrawLose)
		arrayWithScores = append(arrayWithScores, totalScorePart1)
		overallScorePart1 += totalScorePart1
		overallScorePart2 += totalScorePart2
	}
	return decipheredArray, arrayWithScores, overallScorePart1, overallScorePart2
}
