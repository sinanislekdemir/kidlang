package interpreter

// Common utility functions shared across platform implementations

// splitPreservingSpaces splits a string into words while preserving spaces
// Used for syntax highlighting to keep exact spacing
func splitPreservingSpaces(s string) []string {
	var result []string
	var current string
	inQuote := false

	for _, ch := range s {
		if ch == '"' {
			inQuote = !inQuote
			current += string(ch)
		} else if ch == ' ' && !inQuote {
			if current != "" {
				result = append(result, current)
				current = ""
			}
			result = append(result, " ")
		} else {
			current += string(ch)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}

// wrapText wraps text to specified width
func wrapText(text string, width int) []string {
	words := []rune(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	var currentLine []rune
	var currentWord []rune

	for _, ch := range words {
		if ch == ' ' || ch == '\n' {
			// End of word
			if len(currentLine)+len(currentWord)+1 > width && len(currentLine) > 0 {
				// Word doesn't fit, start new line
				lines = append(lines, string(currentLine))
				currentLine = currentWord
			} else {
				// Word fits
				if len(currentLine) > 0 {
					currentLine = append(currentLine, ' ')
				}
				currentLine = append(currentLine, currentWord...)
			}
			currentWord = nil

			if ch == '\n' {
				lines = append(lines, string(currentLine))
				currentLine = nil
			}
		} else {
			currentWord = append(currentWord, ch)
		}
	}

	// Flush remaining word
	if len(currentWord) > 0 {
		if len(currentLine)+len(currentWord)+1 > width && len(currentLine) > 0 {
			lines = append(lines, string(currentLine))
			currentLine = currentWord
		} else {
			if len(currentLine) > 0 {
				currentLine = append(currentLine, ' ')
			}
			currentLine = append(currentLine, currentWord...)
		}
	}

	if len(currentLine) > 0 {
		lines = append(lines, string(currentLine))
	}

	if len(lines) == 0 {
		return []string{""}
	}

	return lines
}

// getSubmenuItemCount returns the number of items in a submenu
// This is the same across platforms
func getSubmenuItemCount(menuSelected int) int {
	switch menuSelected {
	case 0: // File
		return 6
	case 1: // Edit
		return 5
	case 2: // Run
		return 3
	case 3: // Examples
		return 1
	case 4: // Help
		return 3
	case 5: // Language
		return 4
	}
	return 0
}
