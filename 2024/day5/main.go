package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (map[int]ComparisonsRule, [][]int) {
	file := "input.txt"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	compMode := true
	scanner := bufio.NewScanner(f)

	rules := make(map[int]ComparisonsRule)
	pages := make([][]int, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Trim(text, " ") == "" {
			compMode = false
			continue
		}

		if compMode {
			values := strings.Split(text, "|")
			before, _ := strconv.Atoi(strings.Trim(values[0], " "))
			after, _ := strconv.Atoi(strings.Trim(values[1], " "))

			beforeRule, ok := rules[before]
			if !ok {
				beforeRule = ComparisonsRule{}
			}
			beforeRule.Before = append(beforeRule.Before, after)
			rules[before] = beforeRule

			afterRule, ok := rules[after]
			if !ok {
				afterRule = ComparisonsRule{}
			}
			afterRule.After = append(afterRule.After, before)
			rules[after] = afterRule
		} else {
			values := strings.Split(text, ",")
			lines := make([]int, 0)
			for _, v := range values {
				page, _ := strconv.Atoi(strings.Trim(v, " "))
				lines = append(lines, page)
			}
			pages = append(pages, lines)
		}
	}

	return rules, pages
}

type ComparisonsRule struct {
	Before []int
	After  []int
}

func main() {
	rules, updates := readInput()
	totalMiddle := 0
	for batch, update := range updates {
		valid := true
		for i, page := range update {
			rule := rules[page]
			before := update[:i]
			after := update[i+1:]

			isValid := comparePages(before, rule.Before)
			if !isValid {
				log.Println("Invalid before", batch)
				valid = false
				break
			}

			isValid = comparePages(after, rule.After)
			if !isValid {
				log.Println("Invalid after", batch)
				valid = false
				break
			}
		}
		if valid {
			middleNum := len(update) / 2
			totalMiddle += update[middleNum]
		}
	}

	log.Println("Total middle", totalMiddle)
}

func comparePages(pages []int, rules []int) bool {
	for _, page := range pages {
		for _, rule := range rules {
			if page == rule {
				return false
			}
		}
	}
	return true
}
