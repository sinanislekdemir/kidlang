//go:build windows
// +build windows

package interpreter

import (
	"unicode/utf8"
)

// runeIndexToByteIndex converts a rune index to byte index in a string
func runeIndexToByteIndex(s string, runeIndex int) int {
	if runeIndex <= 0 {
		return 0
	}

	byteIndex := 0
	for i := 0; i < runeIndex && byteIndex < len(s); i++ {
		_, size := utf8.DecodeRuneInString(s[byteIndex:])
		byteIndex += size
	}
	return byteIndex
}

// byteIndexToRuneIndex converts a byte index to rune index in a string
func byteIndexToRuneIndex(s string, byteIndex int) int {
	return len([]rune(s[:byteIndex]))
}

// insertChar inserts a character at the current cursor position
func (ide *WindowsIDE) insertChar(ch rune) {
	line := ide.lines[ide.cursorY]
	// Convert cursorX (rune index) to byte index
	bytePos := runeIndexToByteIndex(line, ide.cursorX)
	ide.lines[ide.cursorY] = line[:bytePos] + string(ch) + line[bytePos:]
	ide.cursorX++
	ide.modified = true
}

// deleteChar deletes the character before the cursor (backspace)
func (ide *WindowsIDE) deleteChar() {
	if ide.cursorX > 0 {
		line := ide.lines[ide.cursorY]
		// Convert cursorX to byte positions
		bytePos := runeIndexToByteIndex(line, ide.cursorX)
		prevBytePos := runeIndexToByteIndex(line, ide.cursorX-1)
		ide.lines[ide.cursorY] = line[:prevBytePos] + line[bytePos:]
		ide.cursorX--
		ide.modified = true
	} else if ide.cursorY > 0 {
		// Join with previous line
		prevLine := ide.lines[ide.cursorY-1]
		ide.cursorX = len([]rune(prevLine)) // Count runes, not bytes
		ide.lines[ide.cursorY-1] = prevLine + ide.lines[ide.cursorY]
		ide.lines = append(ide.lines[:ide.cursorY], ide.lines[ide.cursorY+1:]...)
		ide.cursorY--
		ide.modified = true
	}
}

// insertNewline inserts a new line at the cursor position
func (ide *WindowsIDE) insertNewline() {
	line := ide.lines[ide.cursorY]
	// Convert cursorX to byte position
	bytePos := runeIndexToByteIndex(line, ide.cursorX)
	ide.lines[ide.cursorY] = line[:bytePos]
	newLine := line[bytePos:]

	// Insert new line
	ide.lines = append(ide.lines[:ide.cursorY+1], append([]string{newLine}, ide.lines[ide.cursorY+1:]...)...)
	ide.cursorY++
	ide.cursorX = 0
	ide.modified = true
}

// cutLine cuts the current line
func (ide *WindowsIDE) cutLine() {
	// Simple implementation - just delete the line
	ide.deleteLine()
}

// copyLine copies the current line
func (ide *WindowsIDE) copyLine() {
	// TODO: Implement clipboard support
}

// pasteLine pastes a line
func (ide *WindowsIDE) pasteLine() {
	// TODO: Implement clipboard support
}

// deleteLine deletes the current line
func (ide *WindowsIDE) deleteLine() {
	if len(ide.lines) > 1 {
		ide.lines = append(ide.lines[:ide.cursorY], ide.lines[ide.cursorY+1:]...)
		if ide.cursorY >= len(ide.lines) {
			ide.cursorY = len(ide.lines) - 1
		}
		ide.modified = true
	} else {
		ide.lines[0] = ""
		ide.cursorX = 0
		ide.modified = true
	}
}

// clearAll clears all content
func (ide *WindowsIDE) clearAll() {
	ide.lines = []string{""}
	ide.cursorX = 0
	ide.cursorY = 0
	ide.scrollY = 0
	ide.modified = true
}
