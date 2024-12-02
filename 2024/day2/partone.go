package main

import "log"

func solvePartOne(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		dir := 0
		safe := true
		for i := 0; i < len(report); i++ {
			if i+1 == len(report) {
				break
			}
			currDir := dir
			level := report[i]
			nextLevel := report[i+1]
			if nextLevel > level {
				currDir = 1
			} else if nextLevel < level {
				currDir = -1
			} else {
				safe = false
				break
			}
			if dir == 0 {
				dir = currDir
			}
			if currDir != dir {
				safe = false
				break
			}
			remainder := abs(nextLevel - level)
			if remainder < 1 || remainder > 3 {
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}
	log.Println(safeCount)
}
