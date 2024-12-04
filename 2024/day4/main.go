package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() [][]rune {
	file := "input.txt"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	grid := [][]rune{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		values := []rune(scanner.Text())
		grid = append(grid, values)
	}
	return grid
}

func main() {
	grid := readInput()
	solvePartOne(grid)
	solvePartTwo(grid)
}
