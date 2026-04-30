package interpreter

import (
	"strings"
	"unicode"
)

type syntaxTokenKind int

const (
	syntaxNormal syntaxTokenKind = iota
	syntaxKeyword
	syntaxString
	syntaxComment
	syntaxNumber
	syntaxFunction
)

type syntaxToken struct {
	Text string
	Kind syntaxTokenKind
}

func isSyntaxIdentifierStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isSyntaxIdentifierPart(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
}

func tokenizeSyntax(line string, keywords map[string]bool) []syntaxToken {
	runes := []rune(line)
	tokens := make([]syntaxToken, 0, len(runes))
	indentEnd := 0

	for indentEnd < len(runes) && unicode.IsSpace(runes[indentEnd]) {
		indentEnd++
	}

	trimmed := string(runes[indentEnd:])
	if len(trimmed) >= 3 && strings.ToUpper(trimmed[:3]) == "REM" {
		if len(trimmed) == 3 || unicode.IsSpace([]rune(trimmed)[3]) {
			if indentEnd > 0 {
				tokens = append(tokens, syntaxToken{
					Text: string(runes[:indentEnd]),
					Kind: syntaxNormal,
				})
			}
			tokens = append(tokens, syntaxToken{
				Text: trimmed,
				Kind: syntaxComment,
			})
			return tokens
		}
	}

	for index := 0; index < len(runes); {
		switch {
		case index+1 < len(runes) && runes[index] == '/' && runes[index+1] == '/':
			tokens = append(tokens, syntaxToken{
				Text: string(runes[index:]),
				Kind: syntaxComment,
			})
			return tokens
		case runes[index] == '"':
			end := index + 1
			for end < len(runes) && runes[end] != '"' {
				end++
			}
			if end < len(runes) {
				end++
			}
			tokens = append(tokens, syntaxToken{
				Text: string(runes[index:end]),
				Kind: syntaxString,
			})
			index = end
		case isSyntaxIdentifierStart(runes[index]):
			start := index
			for index < len(runes) && isSyntaxIdentifierPart(runes[index]) {
				index++
			}
			word := string(runes[start:index])
			kind := syntaxNormal
			if keywords[strings.ToUpper(word)] {
				kind = syntaxKeyword
			} else if index < len(runes) && runes[index] == '(' {
				kind = syntaxFunction
			}
			tokens = append(tokens, syntaxToken{
				Text: word,
				Kind: kind,
			})
		case unicode.IsDigit(runes[index]):
			start := index
			for index < len(runes) && (unicode.IsDigit(runes[index]) || runes[index] == '.') {
				index++
			}
			tokens = append(tokens, syntaxToken{
				Text: string(runes[start:index]),
				Kind: syntaxNumber,
			})
		default:
			tokens = append(tokens, syntaxToken{
				Text: string(runes[index]),
				Kind: syntaxNormal,
			})
			index++
		}
	}

	return tokens
}
