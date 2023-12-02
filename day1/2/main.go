package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var validTokens = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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
		token := ""
		lastKey := ""
		for i, c := range v {
			if unicode.IsDigit(c) {
				if first == -1 {
					first = c
				} else {
					last = c
				}
			} else {
				if token == "" {
					if i < len(v) - 3 {
						futureToken := string(c) + string(v[i+1]) + string(v[i+2])
						key, found := isTokenValid(futureToken)
						if found {
							token = string(c)
							lastKey = key
						} else if i != 0 {
							pastToken := string(v[i-1]) + string(c)
							key, found := isTokenValid(pastToken)
							if found {
								token = string(v[i-1]) + string(c)
								lastKey = key
							} else {
								token = ""
								lastKey = ""
							}
						} else {
							token = ""
							lastKey = ""
						}
					}
				} else {
					token += string(c)
					if lastKey != "" {
						remaining, found := strings.CutPrefix(lastKey, token)
						if len(remaining) == 0 && found {
							if first == -1 {
								first = validTokens[lastKey]
							} else {
								last = validTokens[lastKey]
							}
							token = ""
							lastKey = ""
						} else if !found {
							token = ""
							lastKey = ""
						}
					} else {
						key, found := isTokenValid(token)
						if found {
							lastKey = key
						} else {
							token = ""
							lastKey = ""
							continue
						}
					}
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

func isTokenValid(token string) (key string, found bool) {
	for k := range validTokens {
		_, found := strings.CutPrefix(k, token)
		if found {
			return k, true
		}
	}

	return "", false
}
