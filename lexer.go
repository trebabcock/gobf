package main

import (
	"bufio"
	"io"
)

type Token int

type ParseToken struct {
	Pos Position
	Lit string
	Tok Token
}

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	Pos    Position
	Reader *bufio.Reader
}

const (
	EOF = iota
	GREATER_THAN
	LESS_THAN
	PLUS
	MINUS
	LEFT_BRACKET
	RIGHT_BRACKET
	COMMA
	PERIOD
	SEMI
)

var tokens = []string{
	EOF:           "EOF",
	GREATER_THAN:  ">",
	LESS_THAN:     "<",
	PLUS:          "+",
	MINUS:         "-",
	LEFT_BRACKET:  "[",
	RIGHT_BRACKET: "]",
	COMMA:         ",",
	PERIOD:        ".",
	SEMI:          ";",
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		Pos:    Position{Line: 1, Column: 0},
		Reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Reset() {
	l = NewLexer(l.Reader)
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.Reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.Pos, EOF, "EOF"
			}
			panic(err)
		}
		l.Pos.Column++

		switch r {
		case '\n':
			l.resetPosition()
		case '>':
			return l.Pos, GREATER_THAN, tokens[GREATER_THAN]
		case '<':
			return l.Pos, LESS_THAN, tokens[LESS_THAN]
		case '+':
			return l.Pos, PLUS, tokens[PLUS]
		case '-':
			return l.Pos, MINUS, tokens[MINUS]
		case '[':
			return l.Pos, LEFT_BRACKET, tokens[LEFT_BRACKET]
		case ']':
			return l.Pos, RIGHT_BRACKET, tokens[RIGHT_BRACKET]
		case ',':
			return l.Pos, COMMA, tokens[COMMA]
		case '.':
			return l.Pos, PERIOD, tokens[PERIOD]
		case ';':
			return l.Pos, SEMI, tokens[SEMI]
		default:
			continue
		}
	}
}

func (l *Lexer) resetPosition() {
	l.Pos.Line++
	l.Pos.Column = 0
}
