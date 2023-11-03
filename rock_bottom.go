package main

import (
	"fmt"
	"strings"
)

// Co-ordinate system:
//
//	o-------> y
//	|
//	|
//	v
//	x

const InputCave string = `################################
~                              #
#         ####                 #
###       ####                ##
###       ####              ####
#######   #######         ######
#######   ###########     ######
################################`

const InputUnits int = 100

type Cave [][]rune

func (c Cave) Print() {
	for _, line := range c {
		fmt.Println(string(line))
	}
}

func (c Cave) At(x, y int) rune {
	return c[x][y]
}
func (c Cave) Above(x, y int) rune {
	return c[x-1][y]
}
func (c Cave) Below(x, y int) rune {
	return c[x+1][y]
}
func (c Cave) Left(x, y int) rune {
	return c[x][y-1]
}
func (c Cave) Right(x, y int) rune {
	return c[x][y+1]
}

func main() {
	lines := strings.Split(InputCave, "\n")
	lines = lines[1:] // Discard ceiling.

	cave := make(Cave, len(lines))
	for i, line := range lines {
		cave[i] = []rune(line)
	}

	cave.Print()

	x := 0
	y := 0

	for n := 1; n < InputUnits; n++ {
		if cave.Below(x, y) == ' ' {
			x++
		} else {
			y++
		}

		if cave.At(x, y) == '#' {
			for cave.Above(x, y) != '~' {
				y--
			}
			x--
			y++
		}

		cave[x][y] = '~'

		print("\n\n\n\n\n\n\n")
		cave.Print()
	}
}
