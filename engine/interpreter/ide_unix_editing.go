//go:build !windows
// +build !windows

package interpreter

// insertChar inserts a character at cursor position
func (ide *UnixIDE) insertChar(ch rune) {
	line := ide.lines[ide.cursorY]
	if ide.cursorX >= len(line) {
		ide.lines[ide.cursorY] = line + string(ch)
	} else {
		ide.lines[ide.cursorY] = line[:ide.cursorX] + string(ch) + line[ide.cursorX:]
	}
	ide.cursorX++
	ide.modified = true
}

// handleBackspace handles backspace key
func (ide *UnixIDE) handleBackspace() {
	if ide.cursorX > 0 {
		line := ide.lines[ide.cursorY]
		ide.lines[ide.cursorY] = line[:ide.cursorX-1] + line[ide.cursorX:]
		ide.cursorX--
		ide.modified = true
	} else if ide.cursorY > 0 {
		// Join with previous line
		prevLine := ide.lines[ide.cursorY-1]
		currentLine := ide.lines[ide.cursorY]
		ide.lines[ide.cursorY-1] = prevLine + currentLine
		ide.lines = append(ide.lines[:ide.cursorY], ide.lines[ide.cursorY+1:]...)
		ide.cursorY--
		ide.cursorX = len(prevLine)
		ide.modified = true
		ide.adjustScroll()
	}
}

// handleDelete handles delete key
func (ide *UnixIDE) handleDelete() {
	line := ide.lines[ide.cursorY]
	if ide.cursorX < len(line) {
		// Delete character at cursor
		ide.lines[ide.cursorY] = line[:ide.cursorX] + line[ide.cursorX+1:]
		ide.modified = true
	} else if ide.cursorY < len(ide.lines)-1 {
		// At end of line, join with next line
		nextLine := ide.lines[ide.cursorY+1]
		ide.lines[ide.cursorY] = line + nextLine
		ide.lines = append(ide.lines[:ide.cursorY+1], ide.lines[ide.cursorY+2:]...)
		ide.modified = true
	}
}

// handleEnter handles enter key
func (ide *UnixIDE) handleEnter() {
	line := ide.lines[ide.cursorY]
	newLine := line[ide.cursorX:]
	ide.lines[ide.cursorY] = line[:ide.cursorX]

	// Insert new line
	ide.lines = append(ide.lines[:ide.cursorY+1], append([]string{newLine}, ide.lines[ide.cursorY+1:]...)...)
	ide.cursorY++
	ide.cursorX = 0
	ide.modified = true
	ide.adjustScroll()
}

// adjustScroll adjusts scroll position to keep cursor visible
func (ide *UnixIDE) adjustScroll() {
	editorHeight := ide.maxY - 2
	if ide.cursorY < ide.scrollY {
		ide.scrollY = ide.cursorY
	} else if ide.cursorY >= ide.scrollY+editorHeight {
		ide.scrollY = ide.cursorY - editorHeight + 1
	}
}
