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

	testInput := strings.Split("R 1,U 1,L 2,D 2,R 2,U 2,L 1,D 1", ",")
	drawMap(testInput)
	testInput2 := strings.Split("R 4,U 4,L 3,D 1,R 4,D 1,L 5,R 2", ",")
	drawMap(testInput2)

	input := advent.ReadInput("day9")
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
	drawTail(headMap.positions, &snekMap)
	fmt.Printf("Unique positions for head: %v, tail: %v\n", len(lo.Uniq(headMap.positions)), len(lo.Uniq(snekMap.positions)))
}
func (snakeMap *SnakeMap) addPosition(xValue int, yValue int) {
	position := positions{x: xValue, y: yValue}
	snakeMap.positions = append(snakeMap.positions, position)
}
func drawTail(headPositions []positions, snekMap *SnakeMap) {
	snekMap.positions = []positions{{x: 0, y: 0}, {x: 0, y: 0}}
	for posIndex, position := range headPositions {
		if posIndex > 1 {
			prevPosition := headPositions[posIndex-1]
			xAx := math.Abs(float64(position.x - snekMap.positions[posIndex-1].x))
			yAx := math.Abs(float64(position.y - snekMap.positions[posIndex-1].y))
			// distance := math.Abs(float64(position.x - tailPosition.x + position.y - tailPosition.y))
			if xAx+yAx >= 2 {
				bothAx := position.x != snekMap.positions[posIndex-1].x && position.y != snekMap.positions[posIndex-1].y
				if xAx+yAx >= 3 {
					snekMap.positions = append(snekMap.positions, prevPosition)
				} else if bothAx {
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
