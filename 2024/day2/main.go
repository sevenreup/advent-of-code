package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]int {
	file := "input.txt"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reports := [][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		levels := []int{}
		for _, v := range values {
			number, _ := strconv.Atoi(v)
			levels = append(levels, number)
		}
		reports = append(reports, levels)
	}
	return reports
}

func main() {
	reports := readInput()
	solvePartOne(reports)
	solvePartTwo(reports)
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
