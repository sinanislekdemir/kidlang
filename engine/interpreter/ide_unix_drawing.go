//go:build !windows
// +build !windows

package interpreter

import (
	"fmt"
	"strings"
)

// draw renders the IDE interface
func (ide *UnixIDE) draw() {
	ide.screen.Clear()

	t := getIDETranslation()

	// Draw menu bar (Batman style - black on yellow)
	ide.screen.ColorOn(4)
	ide.screen.MovePrint(0, 0, strings.Repeat(" ", ide.maxX))

	// Menu items with proper spacing
	menuItems := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}
	menuX := 1
	for i, item := range menuItems {
		if ide.menuActive && i == ide.menuSelected {
			// Highlight selected menu item
			ide.screen.ColorOff(4)
			ide.screen.ColorOn(5) // Black on white for selection
			ide.screen.MovePrint(0, menuX, " "+item+" ")
			ide.screen.ColorOff(5)
			ide.screen.ColorOn(4)
		} else {
			ide.screen.MovePrint(0, menuX, " "+item+" ")
		}
		menuX += len([]rune(item)) + 3 // Use rune length for UTF-8
	}
	ide.screen.ColorOff(4)

	// Draw editor area with black background
	ide.screen.ColorOn(1)
	editorHeight := ide.maxY - 2 // Reserve space for menu and status bar

	// Show welcome text if enabled and no content
	if ide.showWelcome && len(ide.lines) == 1 && ide.lines[0] == "" {
		t := getIDETranslation()
		for row := 0; row < editorHeight; row++ {
			ide.screen.Move(row+1, 0)
			ide.screen.Print(strings.Repeat(" ", ide.maxX))

			if row < len(t.WelcomeText) {
				line := t.WelcomeText[row]
				startX := (ide.maxX - len([]rune(line))) / 2
				if startX < 0 {
					startX = 0
				}
				ide.drawSyntaxHighlightedLine(row+1, startX, line)
			}
		}
	} else {
		for row := 0; row < editorHeight; row++ {
			lineNum := row + ide.scrollY
			ide.screen.Move(row+1, 0)

			// Clear line with black background
			ide.screen.Print(strings.Repeat(" ", ide.maxX))

			if lineNum < len(ide.lines) {
				// Draw line number in yellow
				ide.screen.ColorOff(1)
				ide.screen.ColorOn(7)
				ide.screen.MovePrint(row+1, 0, fmt.Sprintf("%4d ", lineNum+1))
				ide.screen.ColorOff(7)
				ide.screen.ColorOn(1)

				// Draw line content with syntax highlighting
				line := ide.lines[lineNum]
				if len(line) > ide.maxX-6 {
					line = line[:ide.maxX-6]
				}
				ide.drawSyntaxHighlightedLine(row+1, 5, line)
			}
		}
	}
	ide.screen.ColorOff(1)

	// Draw dropdown menu AFTER editor so it appears on top
	if ide.menuActive && ide.submenuActive {
		ide.drawMenuDropdown()
	}

	// Draw status bar at bottom
	statusY := ide.maxY - 1
	ide.screen.ColorOn(5)
	ide.screen.MovePrint(statusY, 0, strings.Repeat(" ", ide.maxX))

	state := editorState{
		Filename: ide.filename,
		Modified: ide.modified,
	}
	status := fmt.Sprintf(" %s %s | %s:%d %s:%d ",
		state.getFilenameDisplay(),
		state.getModifiedIndicator(),
		t.StatusLine,
		ide.cursorY+1,
		t.StatusCol,
		ide.cursorX+1)

	var helpText string
	if ide.menuActive && ide.submenuActive {
		helpText = t.HelpMenuOpen
	} else {
		helpText = t.HelpRunning
	}

	ide.screen.MovePrint(statusY, 0, status)
	ide.screen.MovePrint(statusY, ide.maxX-len(helpText)-1, helpText)
	ide.screen.ColorOff(5)

	// Position cursor
	displayY := ide.cursorY - ide.scrollY + 1
	displayX := ide.cursorX + 5 // Offset for line numbers
	ide.screen.Move(displayY, displayX)

	ide.screen.Refresh()
}

// drawMenuDropdown draws the dropdown menu for the selected menu item
func (ide *UnixIDE) drawMenuDropdown() {
	t := getIDETranslation()

	// Calculate menu position dynamically based on menu items
	menuItems := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}
	menuX := 1
	for i := 0; i < ide.menuSelected && i < len(menuItems); i++ {
		menuX += len([]rune(menuItems[i])) + 3
	}
	menuY := 1

	var items []string
	switch ide.menuSelected {
	case 0: // File
		items = []string{t.FileNew, t.FileOpen, t.FileSave, t.FileSaveAs, "---", t.FileExit}
	case 1: // Run
		items = []string{t.RunRun, "---", t.RunStop}
	case 2: // Examples
		items = []string{t.ExamplesBrowse}
	case 3: // Help
		items = []string{t.HelpAbout, t.HelpDocs, t.HelpShortcuts}
	case 4: // Language
		items = []string{t.LangEnglish, t.LangTurkish, t.LangFinnish, t.LangGerman}
	}

	// Find max width using rune length
	maxWidth := 0
	for _, item := range items {
		runeLen := len([]rune(item))
		if runeLen > maxWidth {
			maxWidth = runeLen
		}
	}
	maxWidth += 4 // Padding

	// Draw dropdown box
	for i, item := range items {
		if ide.submenuActive && i == ide.submenuSelected {
			ide.screen.ColorOn(4) // Yellow background for selection
		} else {
			ide.screen.ColorOn(5) // Black on white
		}
		line := fmt.Sprintf(" %-*s ", maxWidth-2, item)
		ide.screen.MovePrint(menuY+i, menuX, line)
		if ide.submenuActive && i == ide.submenuSelected {
			ide.screen.ColorOff(4)
		} else {
			ide.screen.ColorOff(5)
		}
	}
}

// calculateLanguageMenuX calculates the X position for the Language menu dropdown
func calculateLanguageMenuX(maxX int, t IDEStrings) int {
	// Calculate total width of menu items before Language
	totalWidth := 1 // Start position
	menuItems := []string{t.MenuFile, t.MenuEdit, t.MenuRun, t.MenuExamples, t.MenuHelp}
	for _, item := range menuItems {
		totalWidth += len(item) + 3
	}
	return totalWidth
}

// isLetter checks if a rune is a letter (including UTF-8 letters)
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'à' && ch <= 'ö') || (ch >= 'ø' && ch <= 'ÿ') || // Latin extended
		(ch >= 'Ğ' && ch <= 'ğ') || (ch >= 'İ' && ch <= 'ı') || // Turkish
		(ch >= 'Ş' && ch <= 'ş') || (ch >= 'Ç' && ch <= 'ç') || // Turkish
		(ch >= 'Ö' && ch <= 'ö') || (ch >= 'Ü' && ch <= 'ü') || // German/Turkish
		(ch >= 'Ä' && ch <= 'ä') // German
}

// isDigit checks if a rune is a digit
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// drawSyntaxHighlightedLine draws a line with syntax highlighting
func (ide *UnixIDE) drawSyntaxHighlightedLine(y, x int, line string) {
	// Get keywords from translation
	t := getIDETranslation()
	keywords := t.SyntaxKeywords

	col := x
	i := 0
	inString := false

	// Convert line to runes for proper UTF-8 handling
	runes := []rune(line)

	for i < len(runes) {
		ch := runes[i]

		// Handle comments
		if !inString && i+1 < len(runes) && runes[i] == '/' && runes[i+1] == '/' {
			ide.screen.ColorOn(6) // Green for comments
			ide.screen.MovePrint(y, col, string(runes[i:]))
			ide.screen.ColorOff(6)
			return
		}

		// Handle strings
		if ch == '"' {
			inString = !inString
			if inString {
				// Find end of string
				end := i + 1
				for end < len(runes) && runes[end] != '"' {
					end++
				}
				if end < len(runes) {
					end++ // Include closing quote
				}
				ide.screen.ColorOn(3) // Cyan for strings
				ide.screen.MovePrint(y, col, string(runes[i:end]))
				ide.screen.ColorOff(3)
				col += end - i
				i = end
				inString = false
				continue
			}
		}

		// Handle keywords and identifiers
		if isLetter(ch) {
			// Extract word
			start := i
			for i < len(runes) && (isLetter(runes[i]) || isDigit(runes[i]) || runes[i] == '_') {
				i++
			}
			word := string(runes[start:i])

			if keywords[strings.ToUpper(word)] {
				ide.screen.ColorOn(2) // Yellow for keywords
				ide.screen.MovePrint(y, col, word)
				ide.screen.ColorOff(2)
			} else {
				ide.screen.ColorOn(1) // White for identifiers
				ide.screen.MovePrint(y, col, word)
				ide.screen.ColorOff(1)
			}
			col += len([]rune(word))
			continue
		}

		// Handle numbers
		if isDigit(ch) {
			start := i
			for i < len(runes) && (isDigit(runes[i]) || runes[i] == '.') {
				i++
			}
			number := string(runes[start:i])
			ide.screen.ColorOn(8) // Magenta for numbers
			ide.screen.MovePrint(y, col, number)
			ide.screen.ColorOff(8)
			col += len(number)
			continue
		}

		// Everything else (operators, punctuation)
		ide.screen.ColorOn(1) // White
		ide.screen.MovePrint(y, col, string(ch))
		ide.screen.ColorOff(1)
		col++
		i++
	}
}
