package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora/v3"
)

var memory = make([]byte, 30000)
var index int = 0

var loopStart []int
var maxIndex int = 0

func Parse(lexTokens []ParseToken, startIndex int) {
	loop := lexTokens[startIndex : len(lexTokens)-1]
	for i, t := range loop {
		switch t.Tok {
		case EOF:
			return
		case GREATER_THAN:
			if index < len(memory)-2 {
				index += 1
			} else if index == len(memory)-1 {
				index = 0
			}
			if maxIndex < index {
				maxIndex = index
			}
		case LESS_THAN:
			if index > 0 {
				index -= 1
			} else {
				index = len(memory) - 1
			}
		case PLUS:
			memory[index]++
		case MINUS:
			memory[index]--
		case LEFT_BRACKET:
			if memory[index] == 0 {
				skipTo, err := FindRightMatch(lexTokens, startIndex+i)
				if err != nil {
					HandleError(err, lexTokens, startIndex+i)
				}
				Parse(lexTokens, skipTo+1)
				return
			}
		case RIGHT_BRACKET:
			if memory[index] != 0 {
				loopStart, err := FindLeftMatch(lexTokens, startIndex+i)
				if err != nil {
					HandleError(err, lexTokens, startIndex+i)
				}
				Parse(lexTokens, loopStart+1)
				return
			}
		case COMMA:
			fmt.Print("\n> ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()[0]
			memory[index] = byte(text)
			fmt.Print("\n")
		case PERIOD:
			fmt.Printf("%s", string(memory[index]))
		case SEMI:
			fmt.Print("\nCell no : ")
			for a := range memory[0 : maxIndex+1] {
				fmt.Printf("%d\t\t", a)
			}
			fmt.Print("\nContents: ")
			for _, b := range memory[0 : maxIndex+1] {
				fmt.Printf("%d\t\t", b)
			}
			fmt.Print("\nPointer : ")
			for c := 0; c < index; c++ {
				fmt.Print("\t\t")
			}
			fmt.Print("^\n")
		}
	}
	return
}

func HandleError(err error, lexTokens []ParseToken, currentIndex int) {
	if err != nil {
		fmt.Printf("\nERROR: %s\n", err.Error())
		for i, t := range lexTokens {
			if i < currentIndex {
				fmt.Print(Green(t.Lit))
			} else if i == currentIndex {
				fmt.Print(Red(fmt.Sprintf("%s", t.Lit)))
			} else {
				fmt.Print(t.Lit)
			}
		}
		fmt.Print("\n")
	}
	os.Exit(1)
}

func FindRightMatch(lexTokens []ParseToken, currentIndex int) (int, error) {
	lb := 0
	for i, tok := range lexTokens[currentIndex+1 : len(lexTokens)-1] {
		if tok.Tok == LEFT_BRACKET {
			lb += 1
		}
		if tok.Tok == RIGHT_BRACKET {
			if lb == 0 {
				return i, nil
			}
			lb -= 1
		}
	}
	return 0, errors.New(fmt.Sprintf("No matching right bracket\nlb = %d", lb))
}

func FindLeftMatch(lexTokens []ParseToken, currentIndex int) (int, error) {
	rb := 0
	for i := currentIndex - 1; i >= 0; i-- {
		if lexTokens[i].Lit == "]" {
			rb += 1
		}
		if lexTokens[i].Lit == "[" {
			if rb == 0 {
				return i, nil
			}
			rb -= 1
		}
	}
	return 0, errors.New(fmt.Sprintf("No matching left bracket\nrb = %d", rb))
}

func CheckBrackets(lexTokens []ParseToken) bool {
	var queue []Token
	for _, t := range lexTokens {
		switch t.Tok {
		case LEFT_BRACKET:
			queue = append(queue, RIGHT_BRACKET)
		case RIGHT_BRACKET:
			if 0 < len(queue) && queue[len(queue)-1] == RIGHT_BRACKET {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		}
	}
	return len(queue) == 0
}
