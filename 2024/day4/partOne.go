package main

import "log"

func solvePartOne(grid [][]rune) {
	total := 0
	for x, row := range grid {
		for y, cell := range row {
			if cell == 'X' {
				total += checkValidity(x, y, grid)
			}
		}
	}

	log.Printf("Total: %d", total)
}

func checkValidity(x int, y int, grid [][]rune) int {
	directions := [][]int{
		{0, 1},   // horizontal
		{0, -1},  // horizontal
		{1, 0},   // vertical
		{-1, 0},  // vertical
		{1, 1},   // diagonal
		{-1, 1},  // diagonal
		{1, -1},  // diagonal
		{-1, -1}, // diagonal
	}

	found := 0

	for _, direction := range directions {
		dx, dy := direction[0], direction[1]
		var lastRune rune
	dirMove:
		for i := 0; i < 4; i++ {
			newX := x + i*dx
			newY := y + i*dy

			if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) {
				break
			}

			current := grid[newX][newY]

			switch current {
			case 'X':
				lastRune = current
				continue
			case 'M':
				if lastRune != 'X' {
					break dirMove
				}
				lastRune = current
				continue
			case 'A':
				if lastRune != 'M' {
					break dirMove
				}
				lastRune = current
				continue
			case 'S':
				if lastRune != 'A' {
					break dirMove
				}
				found++
				break dirMove
			default:
				break dirMove
			}

		}
	}
	return found
}
