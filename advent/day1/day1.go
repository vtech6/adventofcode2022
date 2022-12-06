package day1

import (
	"adventofcode/advent"
	"fmt"
	"strconv"
	"strings"
)

type elfDetails struct {
	amount int
}

type highestValues struct {
	first  elfDetails
	second elfDetails
	third  elfDetails
}

func Solve() {
	input := advent.ReadInput("day1")
	var highestValues = highestValues{
		first: elfDetails{
			amount: 0,
		},
		second: elfDetails{
			amount: 0,
		},
		third: elfDetails{
			amount: 0,
		},
	}
	splitElves(input.Split("\n\n"), &highestValues)
	fmt.Println(highestValues)
	fmt.Println(highestValues.Sum())

}

func (values *highestValues) Sum() int {
	return values.first.amount + values.second.amount + values.third.amount
}

func findHighestValues(elementCount int, highestValues *highestValues) {
	if elementCount < highestValues.third.amount {
		return
	}
	if elementCount < highestValues.second.amount {
		highestValues.third.amount = elementCount
		return
	}
	if elementCount < highestValues.first.amount {
		highestValues.third.amount = highestValues.second.amount
		highestValues.second.amount = elementCount
		return
	}
	highestValues.second.amount = highestValues.first.amount
	highestValues.first.amount = elementCount
}

func splitElves(input []string, highestValues *highestValues) {
	for _, element := range input {
		elementCount := 0
		elfElements := strings.Split(element, "\n")
		for _, item := range elfElements {
			value, err := strconv.ParseInt(item, 0, 32)
			if err == nil {
				elementCount += int(value)
			}
		}
		findHighestValues(elementCount, highestValues)
	}
}
