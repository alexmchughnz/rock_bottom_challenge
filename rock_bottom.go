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

type WaterState int

const (
	Flowing WaterState = iota
	Falling
	Filling
)

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

	// State variables.
	col := 0
	height := maxCapacity
	state := Flowing

	spacesBelow := func(c, h int) int {
		nRocks := maxCapacity - capacities[c]

		nBelow := (h - 1) - nRocks
		return nBelow
	}

	for i := 1; i < InputUnits; i++ {
		println(spacesBelow(col, height))

		if state == Falling {
			height--
		} else {
			col++

			if spacesBelow(col, height) < 0 {
				// Hit a wall! Take a step back.
				col--
				// Return to the column water is falling from.
				currentDepth := depths[col]
				for depths[col] == currentDepth {
					col--
				}
				// Water goes up a level and flows right.
				height++
				col++

				state = Filling
			}
		}

		// Place a water tile.
		depths[col]++

		// Determine state change.
		nBelow := spacesBelow(col, height)

		if nBelow == 0 {
			state = Flowing
		}
		if state == Filling {
			nBelow -= depths[col]
		}
		if nBelow > 0 {
			state = Falling
		}

		fmt.Println(depths)
	}
}
