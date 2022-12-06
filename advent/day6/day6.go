package day6

import (
	"adventofcode/advent"
	"fmt"
	"strings"
)

func Solve() {
	input := advent.ReadInput("day6")
	findMarker(input.Value)
}

func findMarker(input string) {
	markerCharacters := ""
	for index, character := range input {
		characterString := string(character)
		characterIndex := strings.Index(markerCharacters, characterString)
		if !strings.Contains(markerCharacters, characterString) {
			markerCharacters += characterString
			if len(markerCharacters) == 14 {
				fmt.Printf("Amount of characters encountered: %v, string: %v\n", index+1, markerCharacters)
				break
			}
		} else {
			markerCharacters = markerCharacters[characterIndex+1:]
			markerCharacters += characterString
		}
	}
}
