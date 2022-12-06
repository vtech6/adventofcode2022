package day4

import (
	"adventofcode/advent"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	input := advent.ReadInput("day4")
	pairsOfElves := input.Split("\n")
	containedPairCount, overlappingPairCount := checkElvenPairs(pairsOfElves)
	fmt.Printf("Contained pair count: %v,\nOverlapping pair count: %v\n", containedPairCount, overlappingPairCount)
}

func checkElvenPairs(pairs []string) (int, int) {
	containedPairs := 0
	overlappingPairs := 0
	for _, pair := range pairs {
		pair1, pair2 := splitPairs(pair)
		containedCondition := (pair1[0] <= pair2[0] && pair1[1] >= pair2[1]) || (pair1[0] >= pair2[0] && pair1[1] <= pair2[1])
		if containedCondition {
			containedPairs += 1
		}
		overlappingCondition := (pair1[0] <= pair2[0] && pair1[1] >= pair2[0]) || (pair2[0] <= pair1[0] && pair2[1] >= pair1[0])
		if overlappingCondition {
			overlappingPairs += 1
		}

	}
	return containedPairs, overlappingPairs
}

func getPairRange(pair string) []int64 {
	splitRange := strings.Split(pair, "-")
	lowRange, _ := strconv.ParseInt(splitRange[0], 32, 0)
	highRange, _ := strconv.ParseInt(splitRange[1], 32, 0)
	return []int64{lowRange, highRange}

}

func splitPairs(pair string) ([]int64, []int64) {
	splitPair := strings.Split(pair, ",")
	pair1 := getPairRange(splitPair[0])
	pair2 := getPairRange(splitPair[1])
	return pair1, pair2
}
