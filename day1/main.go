package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)

	lines := strings.Split(input, "\n")
	sums := 0
	for _, v := range lines {
		var first, last rune = -1, -1
		for i, c := range v {
			if unicode.IsDigit(c) {
				if first == -1 {
					first = c
				} else {
					last = c
				}
			}
			if i == len(v)-1 {
				if last == -1 {
					last = first
				}
				total, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
				sums += total
				first = -1
			}
		}
	}

	fmt.Println(sums)
}