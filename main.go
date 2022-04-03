package main

import (
	"io"
	"log"
	"os"
)

var memory = make([]byte, 30000)
var index int = 0

func main() {
	file, err := os.Open("hw.bf")
	if err != nil {
		panic(err)
	}

	lexTokens := startLex(file)
	if !CheckBrackets(lexTokens) {
		log.Fatal("Mismatched brackets")
	}
	Parse(lexTokens)
}

func startLex(file io.Reader) []ParseToken {
	var curTokens = []ParseToken{}
	lexer := NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}
		curTokens = append(curTokens, ParseToken{pos, lit, tok})
	}
	curTokens = append(curTokens, ParseToken{Position{0, 0}, "EOF", EOF})
	return curTokens
}
