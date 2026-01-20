package interpreter

import (
	"slices"
	"strings"
)

func tokenizer(text string) []string {
	// Trim spaces but don't change case for the entire text
	text = strings.TrimSpace(text)

	breakers := []string{"+", "-", "/", "*", "^", "!", "%", "=", ">", "<", " "}
	parts := make([]string, 0)

	// Split text by breakers, but keep [] and () attached to preceding token
	accumulator := ""
	appendRaw := false
	inBracket := false
	for _, c := range text {
		if c == '\'' || c == '"' {
			accumulator += string(c)
			appendRaw = !appendRaw
			continue
		}
		// Track brackets to keep them with the variable name
		if c == '[' || c == '(' {
			inBracket = true
			accumulator += string(c)
			continue
		}
		if c == ']' || c == ')' {
			inBracket = false
			accumulator += string(c)
			continue
		}
		if slices.Contains(breakers, string(c)) && !appendRaw && !inBracket {
			if len(accumulator) > 0 {
				parts = append(parts, accumulator)
			}
			if c != ' ' {
				parts = append(parts, string(c))
			}
			accumulator = ""
			continue
		}
		accumulator += string(c)
	}

	if len(accumulator) > 0 {
		parts = append(parts, accumulator)
	}

	return parts
}
