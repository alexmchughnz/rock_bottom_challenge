package main

import (
	"fmt"
	"strings"
)

const InputMap string = `################################
~                              #
#         ####                 #
###       ####                ##
###       ####              ####
#######   #######         ######
#######   ###########     ######
################################`

const InputUnits int = 100

func transpose(m string) string {
	lines := strings.Split(m, "\n")

	mT := make([]string, len(lines[0]))

	for _, line := range lines {
		for i, c := range line {
			mT[i] += string(c)
		}
	}
	return strings.Join(mT, "\n")
}

// Capacity of a column = how many vertical stacks of water it can hold.
func capacity(col string) int {
	return len(col) - strings.Count(col, "#")
}

func main() {
	mT := transpose(InputMap)
	cols := strings.Split(mT, "\n")
	maxCapacity := len(cols[0]) - 2 // Less floor and ceiling.

	capacities := make([]int, len(cols))
	for i := range capacities {
		capacities[i] = capacity(cols[i])
	}

	depths := make([]int, len(cols))
	depths[0] = 1 // Start with a single water unit in col 0.
	// capacities[0] = 0

	// Within a column, water will:
	// 1. increase depth to deplete "fillable depth", where at any time,
	//    a column's "fillable depth" = spaces between water height and rock height
	//								  = (water height - 1) - rock height
	// 								  = (`height` - 1) - (`maxCapacity` - `capacity`)
	// 2. flow to the next column

	col := 0
	height := maxCapacity
	isFalling := false
	fillableDepth := func(c, h int) int { return (h - 1) - (maxCapacity - capacities[c]) }

	for i := 0; i < InputUnits; i++ {
		println(fillableDepth(col, height))
		if fillableDepth(col, height) > 0 {
			// Same column, water falls down.
			height--
			fmt.Printf("Fall to height %v: ", height)
			isFalling = true
		} else {
			// Column full, water flows right.
			col++
			fmt.Printf("Move right to col %v: ", col)
			isFalling = false
		}

		depths[col]++

		fmt.Println(depths)
		if isFalling {
			println("~")
		}
	}
}
