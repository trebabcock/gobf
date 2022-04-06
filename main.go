package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify a file")
	}
	if len(os.Args) > 2 {
		log.Fatal("Too many arguments\nUsage: gobf <file_name>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexTokens := startLex(file)
	if !CheckBrackets(lexTokens) {
		log.Fatal("Mismatched brackets")
	}
	Parse(lexTokens, 0)
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
