package main

import (
	"bufio"
	"fmt"
	"os"
)

var loopStart []int
var maxIndex int = 0
var tempLoop int

func Parse(lexTokens []ParseToken) {
	for i, t := range lexTokens {
		fmt.Print(t.Lit)
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
				lb := 0
				for j, tok := range lexTokens[i+1 : len(lexTokens)-1] {
					if tok.Tok == LEFT_BRACKET {
						lb += 1
					}
					if tok.Tok == RIGHT_BRACKET {
						if lb == 0 {
							defer Parse(lexTokens[j+1 : len(lexTokens)-1])
							return
						} else {
							lb -= 1
						}
					}
				}
			} else {
				loopStart = append(loopStart, i+1)
			}
		case RIGHT_BRACKET:
			if memory[index] != 0 {
				tempLoop = loopStart[len(loopStart)-1]
				defer Parse(lexTokens[tempLoop : len(lexTokens)-1])
				return
			} else {
				if len(loopStart) > 0 {
					loopStart = loopStart[:len(loopStart)-1]
				}
			}
		case COMMA:
			fmt.Print("\nEnter a character: ")
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
