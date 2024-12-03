package main

import (
	"log"
)

func solvePartTwo(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		unsafeLevelsIndexes := getUnsafeLevelsIndex(report)
		if unsafeLevelsIndexes == -1 {
			safeCount++
			continue
		}
		for i := 0; i < len(report); i++ {
			isSafe := getUnsafeLevelsIndex(copyNewArray(report, i)) == -1
			if isSafe {
				safeCount++
				break
			}
		}

	}
	log.Println("Safe count", safeCount)
}

func getUnsafeLevelsIndex(report []int) int {
	if len(report) <= 1 {
		return -1
	}
	flagIncrease, flagDecrease := false, false
	for i := 1; i < len(report); i++ {
		level := report[i]
		previousLevel := report[i-1]
		remainder := level - previousLevel
		diffValid := abs(remainder) >= 1 && abs(remainder) <= 3

		if remainder > 0 {
			flagIncrease = true
		} else if remainder < 0 {
			flagDecrease = true
		} else {
			return i
		}

		if flagDecrease && flagIncrease {
			return i
		}

		if !diffValid {
			return i
		}
	}
	return -1
}

func copyNewArray(report []int, unsafeLevelIndex int) []int {
	deleted := make([]int, len(report)-1)
	copy(deleted, report[:unsafeLevelIndex])
	copy(deleted[unsafeLevelIndex:], report[unsafeLevelIndex+1:])
	return deleted

}
