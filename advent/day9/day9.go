package day9

import (
	"adventofcode/advent"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Solve() {
	input := advent.ReadInput("day9")
	// testInput := strings.Split("R 1,U 1,L 2,D 2,R 2,U 2,L 1,D 1", ",")
	splitInput := input.Split("\n")
	drawMap(splitInput)
}

type positions struct {
	x int
	y int
}

type SnakeMap struct {
	positions []positions
}

func drawMap(input []string) {
	yAxis := 0
	yAxisMax := 0
	yAxisMin := 0
	xAxis := 0
	xAxisMax := 0
	xAxisMin := 0

	headMap := SnakeMap{positions: []positions{}}
	snekMap := SnakeMap{positions: []positions{}}
	for _, row := range input {
		splitRow := strings.Split(row, " ")
		conv, _ := strconv.Atoi(splitRow[1])
		if splitRow[0] == "D" {
			for step := 0; step < conv; step++ {
				headMap.addPosition(xAxis, yAxis-step)
			}
			yAxis -= conv
			if yAxis < yAxisMin {
				yAxisMin = yAxis
			}

		}
		if splitRow[0] == "U" {
			for step := 0; step < conv; step++ {
				headMap.addPosition(xAxis, yAxis+step)
			}
			yAxis += conv
			if yAxis > yAxisMax {
				yAxisMax = yAxis
			}
		}
		if splitRow[0] == "R" {
			for step := 0; step < conv; step++ {
				headMap.addPosition(xAxis+step, yAxis)
			}
			xAxis += conv
			if xAxis > xAxisMax {
				xAxisMax = xAxis
			}
		}
		if splitRow[0] == "L" {
			for step := 0; step < conv; step++ {
				headMap.addPosition(xAxis-step, yAxis)
			}
			xAxis -= conv
			if xAxis < xAxisMin {
				xAxisMin = xAxis
			}
		}
	}
	fmt.Printf("X: %v, Xmin: %v, Xmax: %v, Y: %v, Ymin: %v, Ymax: %v\n", xAxis, xAxisMin, xAxisMax, yAxis, yAxisMin, yAxisMax)
	fmt.Println(headMap.positions)
	drawTail(headMap.positions, &snekMap)
	fmt.Println(snekMap.positions)
	fmt.Printf("Unique positions for input: %v, snake: %v\n", len(lo.Uniq(headMap.positions)), len(lo.Uniq(snekMap.positions)))

}
func (snakeMap *SnakeMap) addPosition(xValue int, yValue int) {
	position := positions{x: xValue, y: yValue}
	snakeMap.positions = append(snakeMap.positions, position)
}
func drawTail(headPositions []positions, snekMap *SnakeMap) {
	snekMap.positions = []positions{{x: 0, y: 0}, {x: 0, y: 0}}
	for posIndex, position := range headPositions {
		if posIndex > 1 {
			tailPosition := headPositions[posIndex-2]
			prevPosition := headPositions[posIndex-1]
			pointDiff := float64(position.x - tailPosition.x + position.y - tailPosition.y)
			pointDiff2 := float64(position.x - snekMap.positions[posIndex-1].x + position.y - snekMap.positions[posIndex-1].y)

			if posIndex > 4 {
				fmt.Printf("PointDiff: %v, PointDiff2: %v, Pos: %v, Snake: %v\n", pointDiff, pointDiff2, position, snekMap.positions[posIndex-1])
			}
			if math.Abs(pointDiff) > 1 && math.Abs(pointDiff2) > 0 {
				bothAx := position.x != tailPosition.x && position.y != tailPosition.y
				if bothAx {
					snekMap.positions = append(snekMap.positions, snekMap.positions[posIndex-1])
				} else {
					snekMap.positions = append(snekMap.positions, prevPosition)
				}
			} else {
				snekMap.positions = append(snekMap.positions, snekMap.positions[posIndex-1])
			}
		}
	}
}

// [{0 0} {0 -1} {0 -2} {1 -2} {2 -2} {2 -3} {2 -2} {1 -2} {0 -2} {0 -1} {0  0} {0 -1}]
// [{0 0} {0  0} {0 -1} {0 -1} {1 -2} {1 -2} {1 -2} {1 -2} {1 -2} {1 -2} {0 -1} {0 -1}]

// [{0 0} {1 0} {1 1} {0 1} {-1 1} {-1 0} {-1 -1} {0 -1} {1 -1} {1  0} {1 1} {0 1}]
// [{0 0} {0 0} {0 0} {0 0} { 0 0} { 0 0} {-1  0} {-1 0} {0 -1} {0 -1} {1 0} {1 0}]

// {3 4} {2 4} {1 4} {1 3} {2 3} {3 3} {4 3}
// {4 3} {3 4} {2 4} {2 4} {2 4} {2 4} {3 3}
