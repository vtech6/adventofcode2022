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

	// testInput := strings.Split("R 1,U 1,L 2,D 2,R 2,U 2,L 1,D 1", ",")
	// drawMap(testInput)

	// testInput := strings.Split("R 2,U 2,D 4,L 1", ",")
	// drawMap(testInput)

	// testInput3 := strings.Split("R 9,U 2,D 4", ",")
	// drawMap(testInput3)

	// testInput2 := strings.Split("R 5,U 8,L 8,D 3,R 17,D 10,L 25,R 20", ",")
	// drawMap(testInput2)

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
	steps     int
	direction string
	xAxis     int
	yAxis     int
	snek      []positions
	length    int
}

func (snakeMap *SnakeMap) addPosition() {
	for step := 0; step < snakeMap.steps; step++ {
		var position positions
		switch snakeMap.direction {
		case "D":
			snakeMap.yAxis -= 1
		case "U":
			snakeMap.yAxis += 1
		case "R":
			snakeMap.xAxis += 1
		case "L":
			snakeMap.xAxis -= 1
		default:
			return
		}
		position = positions{snakeMap.xAxis, snakeMap.yAxis}
		snakeMap.positions = append(snakeMap.positions, position)
		snakeMap.drawTrail()
	}
}

func (snakeMap *SnakeMap) drawTrail(head ...positions) {
	var leadPositions []positions
	if len(head) != 0 {
		leadPositions = head
	} else {
		leadPositions = snakeMap.positions
	}
	snakeLen := len(leadPositions)
	lastIndex := snakeLen - 1
	if snakeLen > snakeMap.length {
		lastPosition := leadPositions[lastIndex]
		previousPosition := leadPositions[lastIndex-1]
		xAx := math.Abs(float64(lastPosition.x - snakeMap.snek[lastIndex-1].x))
		yAx := math.Abs(float64(lastPosition.y - snakeMap.snek[lastIndex-1].y))
		// lastSnakeIndex := len(snakeMap.snek) - 1
		// lastSnakePosition := snakeMap.snek[lastSnakeIndex]
		if xAx+yAx >= 2 {

			bothAx := lastPosition.x != snakeMap.snek[lastIndex-1].x && lastPosition.y != snakeMap.snek[lastIndex-1].y
			if xAx+yAx >= 3 {
				snakeMap.snek = append(snakeMap.snek, previousPosition)
			} else if bothAx {
				snakeMap.snek = append(snakeMap.snek, snakeMap.snek[lastIndex-1])
			} else {
				snakeMap.snek = append(snakeMap.snek, previousPosition)
			}
		} else {
			snakeMap.snek = append(snakeMap.snek, snakeMap.snek[lastIndex-1])
		}
	}
}

func (snakeMap *SnakeMap) setLength(length int) {
	for i := 0; i <= length-1; i++ {
		snakeMap.snek = append(snakeMap.snek, positions{0, 0})
	}
	snakeMap.length = length
}

func drawMap(input []string) {
	headMap := SnakeMap{positions: []positions{{0, 0}}}
	headMap.setLength(2)
	for _, row := range input {
		splitRow := strings.Split(row, " ")
		conv, _ := strconv.Atoi(splitRow[1])
		direction := splitRow[0]
		headMap.direction = direction
		headMap.steps = conv
		headMap.addPosition()
	}
	fmt.Printf("SNEK: %v\n", len(lo.Uniq(headMap.snek)))
}
