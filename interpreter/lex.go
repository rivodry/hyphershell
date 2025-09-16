package interpreter

import (
	"unicode"
)

type Token struct {
	Typ   string
	Value string
}

func Lex(Set string) []Token {
	var Tokens []Token
	for i := 0; i < len(Set); {
		if unicode.IsSpace(rune(Set[i])) {
			i++
			continue
		} else if unicode.IsLetter(rune(Set[i])) || Set[i] == '_' {
			var str string
			for i < len(Set) && (unicode.IsLetter(rune(Set[i])) || unicode.IsDigit(rune(Set[i])) || Set[i] == '_') {
				str = str + string(Set[i])
				i++
			}
			Tokens = append(Tokens, Token{"IDENT", str})

		} else if unicode.IsDigit(rune(Set[i])) {
			dots := 0
			var num string
			for i < len(Set) && dots < 2 && (unicode.IsDigit(rune(Set[i])) || rune(Set[i]) == '.') {
				if rune(Set[i]) == '.' {
					dots++
				}

				num = num + string(Set[i])
				i++
			}
			if num[len(num)-1] == '.' {
				num = num[:len(num)-1]

			}
			Tokens = append(Tokens, Token{"NUMBER", num})

		} else {
			i++
		}
	}
	Tokens = append(Tokens, Token{"EOF", ""})
	return Tokens
}
