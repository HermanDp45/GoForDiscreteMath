package main

import (
	"fmt"
	"unicode"
)

type Tag int;

type Lexem struct {
	Tag;
	Image string;
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
	i := 0;
	for i < len(expr) {
		char := expr[i]; //get current symbol

		// skip space symbols
		if unicode.IsSpace(rune(char)) {
			i++;
			continue;
		}
		
		// recognize a word
		if unicode.IsLetter(rune(char)) {
			q := i + 1;

			for (q < len(expr)) && (unicode.IsLetter(rune(expr[q])) || unicode.IsDigit(rune(expr[q]))) {
				q++;
			}

			lexems <- Lexem{
				Tag: VAR,
				Image: expr[i:q],
			}

			i = q;
			continue;
		}

		// recognize a number
		if unicode.IsDigit(rune(char)) {
			q := i + 1;

			for q < len(expr) && unicode.IsDigit(rune(expr[q])) {
				q++;
			}

			lexems <- Lexem{
				Tag: NUMBER,
				Image: expr[i:q],
			}

			i = q;
			continue;
		}

		// recognize sign operation and "(" ")" symbols
		switch char {
		case '+':
			lexems <- Lexem{
				Tag: PLUS,
				Image: "+",
			}
			i++;
			continue;
		case '-':
			lexems <- Lexem{
				Tag: MINUS,
				Image: "-",
			}
			i++;
			continue;
		case '*':
			lexems <- Lexem{
				Tag: MUL,
				Image: "*",
			}
			i++;
			continue;
		case '/':
			lexems <- Lexem{
				Tag: DIV,
				Image: "/",
			}
			i++;
			continue;
		case '(':
			lexems <- Lexem{
				Tag: LPAREN,
				Image: "(",
			}
			i++;
			continue;
		case ')':
			lexems <- Lexem{
				Tag: RPAREN,
				Image: ")",
			}
			i++;
			continue;
		default:
			panic(fmt.Sprintf("Unexpected character at position %d: %c", i, char));
		}
	
	}
}