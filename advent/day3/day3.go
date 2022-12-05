package day3

import (
	"adventofcode/advent"
	"fmt"
	"math"
	"strings"
)

const (
	smallLetters = "abcdefghijklmnopqrstuvwxyz"
)

func Day3() {
	input := advent.ReadInput("day3")
	elvenRucksacks := input.Split("\n")
	examineRucksacks(elvenRucksacks)
	examineBadges(elvenRucksacks)
}

func examineBadges(rucksacks []string) {
	groupedRucksacks := make(map[float64][]string)
	for rucksackIndex, rucksack := range rucksacks {
		groupIndex := math.Floor(float64(rucksackIndex) / 3)
		group := groupedRucksacks[groupIndex]
		groupedRucksacks[groupIndex] = append(group, rucksack)
	}
	letterValues := getLetterValues()

	totalBadgeValue := 0
	for _, rucksack := range groupedRucksacks {
		repeated := ""
		for _, element := range rucksack[0] {
			stringifiedElement := string(element)
			if strings.Contains(rucksack[1], stringifiedElement) && strings.Contains(rucksack[2], stringifiedElement) {
				if !strings.Contains(repeated, stringifiedElement) {
					repeated += stringifiedElement
					totalBadgeValue += letterValues[stringifiedElement]
				}
			}

		}
	}
	fmt.Printf("Total Badge Value (Part 2): %v\n", totalBadgeValue)
}

func getLetterValues() map[string]int {
	letterValues := make(map[string]int)
	for letterIndex, letter := range smallLetters {
		letterValues[string(letter)] = letterIndex + 1
		letterValues[strings.ToUpper(string(letter))] = letterIndex + 27
	}
	return letterValues
}

func examineRucksacks(rucksacks []string) {
	letterValues := getLetterValues()
	totalScore := 0
	for _, rucksack := range rucksacks {
		rucksackScore := 0
		compartments := []string{rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]}
		repeated := ""
		for _, element := range compartments[0] {
			stringifiedElement := string(element)
			if strings.Contains(compartments[1], stringifiedElement) {
				if !strings.Contains(repeated, stringifiedElement) {
					repeated += stringifiedElement
					rucksackScore += letterValues[stringifiedElement]
				}
			}

		}
		totalScore += rucksackScore
	}
	fmt.Printf("Total rucksack value (Part 1): %v\n", totalScore)
}
