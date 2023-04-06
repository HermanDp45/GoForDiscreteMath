package main

import (
	"fmt"
	"unicode"
)

type Tag int

type Lexem struct {
	Tag
	Image string
}

const (
	ERROR  Tag = 1 << iota // wrong lexem
	NUMBER                 // number (int)
	VAR                    // var name
	PLUS                   // sign +
	MINUS                  // sign -
	MUL                    // sign *
	DIV                    // sign /
	LPAREN                 // left (
	RPAREN                 // right )
)

func lexer(expr string, lexems chan Lexem) {
	i := 0 // текущая позиция в выражении
	for i < len(expr) {
		c := expr[i] // текущий символ

		// Пропускаем пробелы
		if unicode.IsSpace(rune(c)) {
			i++
			continue
		}

		// Распознаем число
		if unicode.IsDigit(rune(c)) {
			start := i
			for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
				i++
			}
			lexems <- Lexem{Tag: NUMBER, Image: expr[start:i]}
			continue
		}

		//recognize var
		if unicode.IsLetter(rune(c)) {
			start := i
			for i < len(expr) && (unicode.IsLetter(rune(expr[i])) || unicode.IsDigit(rune(expr[i]))) {
				i++
			}
			lexems <- Lexem{Tag: VAR, Image: expr[start:i]}
			continue
		}
		// recognize sign operation
		if c == '+' {
			lexems <- Lexem{Tag: PLUS, Image: "+"}
			i++
			continue
		}
		if c == '-' {
			lexems <- Lexem{Tag: MINUS, Image: "-"}
			i++
			continue
		}
		if c == '*' {
			lexems <- Lexem{Tag: MUL, Image: "*"}
			i++
			continue
		}
		if c == '/' {
			lexems <- Lexem{Tag: DIV, Image: "/"}
			i++
			continue
		}

		// recognize "(" ")"
		if c == '(' {
			lexems <- Lexem{Tag: LPAREN, Image: "("}
			i++
			continue
		}
		if c == ')' {
			lexems <- Lexem{Tag: RPAREN, Image: ")"}
			i++
			continue
		}

		// can't recognise symbol => send panic
		panic(fmt.Sprintf("Unexpected character at position %d: %c", i, c))
	}

	close(lexems)
}


func maink() {

}