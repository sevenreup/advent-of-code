package main

import "fmt"

func solvePartTwo(leftNumbers []int, rightNumbers []int) {
	totalScore := 0

	for i := 0; i < len(leftNumbers); i++ {
		times := numberOfOccurrences(rightNumbers, leftNumbers[i])
		totalScore += (times * leftNumbers[i])
	}

	fmt.Println(totalScore)
}

func numberOfOccurrences(numbers []int, number int) int {
	count := 0
	for _, n := range numbers {
		if n == number {
			count++
		}
	}
	return count
}
