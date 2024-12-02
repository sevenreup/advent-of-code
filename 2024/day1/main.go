package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([]int, []int) {
	file := "input.txt"

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer f.Close()

	leftNumbers := []int{}
	rightNumbers := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(v[0])
		rightNumber, _ := strconv.Atoi(v[1])
		leftNumbers = append(leftNumbers, leftNumber)
		rightNumbers = append(rightNumbers, rightNumber)
	}

	return leftNumbers, rightNumbers
}

func main() {
	leftNumbers, rightNumbers := parseInput()

	solvePartOne(leftNumbers, rightNumbers)
	solvePartTwo(leftNumbers, rightNumbers)
}