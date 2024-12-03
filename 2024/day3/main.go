package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

type OpState int
type OpType int

const (
	Initial OpState = iota
	MultFunction
	DoFunction
	DontFunction
	Arg1
	Arg2
)

const (
	Multiply OpType = iota
	Do
	Dont
)

type Lexer struct {
	r         *bufio.Reader
	file      *os.File
	lastRune  rune
	currentOp Op
	state     int
	allOps    []Op
}

type Op struct {
	FirstArg  float64
	SecondArg float64
	Type      OpType
}

func newLexer() *Lexer {
	file := "input.txt"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewReader(f)

	return &Lexer{r: scanner, file: f, state: int(Initial), currentOp: Op{}}
}

func main() {
	lexer := newLexer()
	defer lexer.file.Close()
	for {
		r, _, err := lexer.r.ReadRune()
		if err != nil {
			log.Println(err)
			break
		}

		switch r {
		case 'm':
			lexer.lastRune = r
		case 'u':
			if !lexer.lastEqualOrReset('m') {
				continue
			}
			lexer.lastRune = r
		case 'l':
			if !lexer.lastEqualOrReset('u') {
				continue
			}
			lexer.lastRune = r
		case 'd':
			lexer.lastRune = r
		case 'o':
			if !lexer.lastEqualOrReset('d') {
				continue
			}
			lexer.lastRune = r
		case 'n':
			if !lexer.lastEqualOrReset('o') {
				continue
			}
			lexer.lastRune = r
		case '\'':
			if !lexer.lastEqualOrReset('n') {
				continue
			}
			lexer.lastRune = r
		case 't':
			if !lexer.lastEqualOrReset('\'') {
				continue
			}
			lexer.lastRune = r
		case '(':
			if lexer.lastEqual('l') {
				lexer.lastRune = r
				lexer.state = int(Arg1)
				lexer.currentOp.Type = Multiply
				continue
			} else if lexer.lastEqual('t') {
				lexer.lastRune = r
				lexer.state = int(DontFunction)
				lexer.currentOp.Type = Dont
				continue
			} else if lexer.lastEqual('o') {
				lexer.lastRune = r
				lexer.state = int(DoFunction)
				lexer.currentOp.Type = Do
				continue
			}
			lexer.reset()
		case ')':
			if lexer.state == int(Arg2) || lexer.state == int(DoFunction) || lexer.state == int(DontFunction) {
				lexer.finishOp()
			} else {
				lexer.reset()
			}
		case ',':
			if lexer.state == int(Arg1) {
				lexer.state = int(Arg2)
				continue
			} else {
				lexer.reset()
			}
		default:
			if (lexer.state == int(Arg1) || lexer.state == int(Arg2)) && unicode.IsDigit(r) {
				rawNumber := lexer.ReadNumber(r)
				lexer.setArg(rawNumber)
				continue
			}
			lexer.reset()
		}
	}

	log.Println(lexer.allOps)

	total := 0.0
	var lastOpType OpType
	for _, op := range lexer.allOps {
		if lastOpType == Dont {
			if op.Type == Do {
				lastOpType = op.Type
			}
			continue
		}
		if op.Type == Multiply {
			total += op.FirstArg * op.SecondArg
		} else {
			lastOpType = op.Type
		}
	}
	log.Println("Total", total)
}

func (l *Lexer) reset() {
	l.lastRune = ' '
	l.state = int(Initial)
}

func (l *Lexer) lastEqual(expected rune) bool {
	return l.lastRune == expected
}

func (l *Lexer) lastEqualOrReset(expected rune) bool {
	if l.lastRune == expected {
		return true
	}
	l.reset()
	return false
}

func (l *Lexer) ReadNumber(current rune) string {
	number := string(current)
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if unicode.IsDigit(r) {
			number += string(r)
		} else {
			l.r.UnreadRune()
			break
		}
	}
	return number
}

func (l *Lexer) finishOp() {
	l.allOps = append(l.allOps, l.currentOp)
	l.currentOp = Op{}
	l.reset()
}

func (l *Lexer) setArg(rawNumber string) {
	number, err := strconv.Atoi(rawNumber)
	if err != nil {
		log.Fatal(err)
	}
	if l.state == int(Arg1) {
		l.currentOp.FirstArg = float64(number)
	} else {
		l.currentOp.SecondArg = float64(number)
	}
}
