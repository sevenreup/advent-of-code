package main

import "fmt"

func solvePartOne(leftNumbers []int, rightNumbers []int) {
	leftNumbers = sortAsc(leftNumbers)
	rightNumbers = sortAsc(rightNumbers)

	totalDistance := 0

	for i := 0; i < len(leftNumbers); i++ {
		distance := abs(rightNumbers[i] - leftNumbers[i])
		totalDistance += distance
	}

	fmt.Println(totalDistance)
}

func sortAsc(numbers []int) []int {
	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)
	for i := 0; i < len(sortedNumbers); i++ {
		for j := i + 1; j < len(sortedNumbers); j++ {
			if sortedNumbers[i] > sortedNumbers[j] {
				sortedNumbers[i], sortedNumbers[j] = sortedNumbers[j], sortedNumbers[i]
			}
		}
	}
	return sortedNumbers
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
