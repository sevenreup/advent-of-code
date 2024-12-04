package main

import (
	"fmt"
)

func solvePartTwo(grid [][]rune) {
	total := 0
	for x, row := range grid {
		for y, cell := range row {
			if cell == 'A' {
				if checkXMasValidity(x, y, grid) {
					total++
				}
			}
		}
	}

	fmt.Printf("Total: %d", total)
}

func checkXMasValidity(x int, y int, grid [][]rune) bool {
	directions := [][]int{
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
	}
	top := []rune{}
	bottom := []rune{}
	for _, direction := range directions {
		dx, dy := direction[0], direction[1]

		newX := x + dx
		newY := y + dy

		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) {
			return false
		}

		current := grid[newX][newY]
		if current != 'M' && current != 'S' {
			return false
		}

		if dy == -1 {
			top = append(top, current)
		} else {
			bottom = append(bottom, current)
		}
	}
	if validCorner(top[0], bottom[1]) && validCorner(top[1], bottom[0]) {
		return true
	}
	return false
}

func validCorner(top rune, bottom rune) bool {
	return top == 'M' && bottom == 'S' || top == 'S' && bottom == 'M'
}
