//go:build windows
// +build windows

package interpreter

import (
	"fmt"
	"strings"
	"unsafe"
)

// Windows console color attributes
const (
	colorWhite         = 0x07 // White on black
	colorYellow        = 0x0E // Yellow on black
	colorCyan          = 0x0B // Cyan on black
	colorGreen         = 0x0A // Green on black
	colorMagenta       = 0x0D // Magenta on black
	colorBlackOnYellow = 0x60 // Black on yellow
	colorBlackOnWhite  = 0x70 // Black on white
	colorYellowOnBlack = 0x0E // Yellow on black
)

// screenBuffer holds the screen contents to batch writes
type screenBuffer struct {
	buffer []charInfo
	width  int
	height int
}

// newScreenBuffer creates a new screen buffer
func newScreenBuffer(width, height int) *screenBuffer {
	size := width * height
	return &screenBuffer{
		buffer: make([]charInfo, size),
		width:  width,
		height: height,
	}
}

// set sets a character and attribute at x, y
func (sb *screenBuffer) set(x, y int, ch rune, attr uint16) {
	if x >= 0 && x < sb.width && y >= 0 && y < sb.height {
		idx := y*sb.width + x
		sb.buffer[idx].Char = uint16(ch)
		sb.buffer[idx].Attributes = attr
	}
}

// writeString writes a string at x, y with given attribute
func (sb *screenBuffer) writeString(x, y int, s string, attr uint16) {
	col := 0
	for _, ch := range s {
		if x+col >= sb.width {
			break
		}
		sb.set(x+col, y, ch, attr)
		col++
	}
}

// clear clears the buffer with spaces and default attribute
func (sb *screenBuffer) clear(attr uint16) {
	for i := range sb.buffer {
		sb.buffer[i].Char = ' '
		sb.buffer[i].Attributes = attr
	}
}

var screenBuf *screenBuffer

// clearScreen clears the entire console using Windows API
func (ide *WindowsIDE) clearScreen() {
	if screenBuf == nil {
		screenBuf = newScreenBuffer(ide.maxX, ide.maxY)
	}
	screenBuf.clear(colorWhite)
}

// flushScreen writes the screen buffer to console in one call
func (ide *WindowsIDE) flushScreen() {
	if screenBuf == nil {
		return
	}

	// Use WriteConsoleOutput to write entire buffer at once
	rect := smallRect{
		Left:   0,
		Top:    0,
		Right:  int16(ide.maxX - 1),
		Bottom: int16(ide.maxY - 1),
	}

	bufSize := coord{
		X: int16(ide.maxX),
		Y: int16(ide.maxY),
	}

	bufCoord := coord{X: 0, Y: 0}

	procWriteConsoleOutput.Call(
		uintptr(ide.stdout),
		uintptr(unsafe.Pointer(&screenBuf.buffer[0])),
		*(*uintptr)(unsafe.Pointer(&bufSize)),
		*(*uintptr)(unsafe.Pointer(&bufCoord)),
		uintptr(unsafe.Pointer(&rect)),
	)
}

// setCursor sets the cursor position using Windows API
func (ide *WindowsIDE) setCursor(x, y int) {
	pos := coord{X: int16(x), Y: int16(y)}
	procSetConsoleCursorPosition.Call(uintptr(ide.stdout), *(*uintptr)(unsafe.Pointer(&pos)))
}

// updateCursorPosition updates only the cursor position and status bar (fast update)
func (ide *WindowsIDE) updateCursorPosition() {
	// Update status bar
	ide.drawStatusBar()
	ide.flushScreen()

	// Position cursor in editor area
	editorHeight := ide.maxY - 4

	// Check if we need to scroll
	if ide.cursorY < ide.scrollY {
		ide.scrollY = ide.cursorY
		ide.draw() // Need full redraw when scrolling
		return
	}
	if ide.cursorY >= ide.scrollY+editorHeight {
		ide.scrollY = ide.cursorY - editorHeight + 1
		ide.draw() // Need full redraw when scrolling
		return
	}

	screenY := ide.cursorY - ide.scrollY + 1
	screenX := ide.cursorX + 5 // Account for line numbers
	ide.setCursor(screenX, screenY)
}

// draw renders the entire IDE screen
func (ide *WindowsIDE) draw() {
	ide.clearScreen()

	// Adjust scroll if cursor is off screen
	editorHeight := ide.maxY - 4 // Leave space for menu, status, and info bars
	if ide.cursorY < ide.scrollY {
		ide.scrollY = ide.cursorY
	}
	if ide.cursorY >= ide.scrollY+editorHeight {
		ide.scrollY = ide.cursorY - editorHeight + 1
	}

	ide.drawMenuBar()
	ide.drawEditor()
	ide.drawStatusBar()
	ide.drawInfoBar()

	if ide.submenuActive {
		ide.drawSubmenu()
	}

	// Flush everything at once
	ide.flushScreen()

	// Position cursor in editor area
	screenY := ide.cursorY - ide.scrollY + 1
	screenX := ide.cursorX + 5 // Account for line numbers
	ide.setCursor(screenX, screenY)
}

// drawMenuBar draws the top menu bar
func (ide *WindowsIDE) drawMenuBar() {
	t := getIDETranslation()
	menus := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}
	x := 0

	for i, menu := range menus {
		var color uint16
		if ide.menuActive && i == ide.menuSelected {
			color = colorBlackOnYellow
		} else {
			color = colorBlackOnWhite
		}
		menuText := " " + menu + " "
		screenBuf.writeString(x, 0, menuText, color)
		x += len([]rune(menuText))
		
		// Add space between menus
		if i < len(menus)-1 {
			screenBuf.set(x, 0, ' ', colorBlackOnWhite)
			x++
		}
	}

	// Fill rest of line with menu bar color
	for ; x < ide.maxX; x++ {
		screenBuf.set(x, 0, ' ', colorBlackOnWhite)
	}
}

// drawEditor draws the main editor area with line numbers
func (ide *WindowsIDE) drawEditor() {
	editorHeight := ide.maxY - 4
	
	// Show welcome text if enabled and no content
	if ide.showWelcome && len(ide.lines) == 1 && ide.lines[0] == "" {
		t := getIDETranslation()
		for i := 0; i < editorHeight; i++ {
			y := i + 1
			if i < len(t.WelcomeText) {
				// Center the welcome text
				line := t.WelcomeText[i]
				startX := (ide.maxX - len([]rune(line))) / 2
				if startX < 0 {
					startX = 0
				}
				ide.drawHighlightedLineAt(startX, y, line)
			} else {
				// Clear rest of screen
				for x := 0; x < ide.maxX; x++ {
					screenBuf.set(x, y, ' ', colorWhite)
				}
			}
		}
		return
	}

	for i := 0; i < editorHeight; i++ {
		lineNum := ide.scrollY + i
		y := i + 1

		if lineNum < len(ide.lines) {
			// Line number in yellow
			lineNumStr := fmt.Sprintf("%4d ", lineNum+1)
			screenBuf.writeString(0, y, lineNumStr, colorYellow)

			// Syntax highlighted line content
			line := ide.lines[lineNum]
			ide.drawHighlightedLineAt(5, y, line)
		} else {
			screenBuf.writeString(0, y, "    ~", colorWhite)
			// Clear rest of line
			for x := 5; x < ide.maxX; x++ {
				screenBuf.set(x, y, ' ', colorWhite)
			}
		}
	}
}

// drawHighlightedLineAt draws a line with syntax highlighting at a specific position
func (ide *WindowsIDE) drawHighlightedLineAt(startX, y int, line string) {
t := getIDETranslation()
keywords := t.SyntaxKeywords

x := startX

// Check if line starts with REM (comment)
if len(line) >= 3 && strings.ToUpper(line[:3]) == "REM" {
screenBuf.writeString(x, y, line, colorGreen)
x += len([]rune(line))
// Clear rest of line
for ; x < ide.maxX; x++ {
screenBuf.set(x, y, ' ', colorWhite)
}
return
}

words := splitPreservingSpaces(line)

for _, word := range words {
if x >= ide.maxX {
break
}

// Check if it's a keyword
wordUpper := strings.ToUpper(word)
if keywords[wordUpper] {
screenBuf.writeString(x, y, word, colorYellow)
} else {
var color uint16
// Check if it's a string (starts with quote)
if len(word) > 0 && word[0] == '"' {
color = colorCyan
} else if len(word) > 0 && word[0] >= '0' && word[0] <= '9' {
// Number
color = colorMagenta
} else {
// Default
color = colorWhite
}
screenBuf.writeString(x, y, word, color)
}

x += len([]rune(word))
}

// Clear rest of line
for ; x < ide.maxX; x++ {
screenBuf.set(x, y, ' ', colorWhite)
}
}

// drawStatusBar draws the status bar
func (ide *WindowsIDE) drawStatusBar() {
	y := ide.maxY - 2
	t := getIDETranslation()

	state := editorState{
		Filename: ide.filename,
		Modified: ide.modified,
	}
	status := fmt.Sprintf(" %s%s | %s %d/%d | %s %d ",
		state.getFilenameDisplay(),
		state.getModifiedIndicator(),
		t.StatusLine,
		ide.cursorY+1,
		len(ide.lines),
		t.StatusCol,
		ide.cursorX+1)

	screenBuf.writeString(0, y, status, colorBlackOnWhite)

	// Fill rest of line
	for x := len(status); x < ide.maxX; x++ {
		screenBuf.set(x, y, ' ', colorBlackOnWhite)
	}
}

// drawInfoBar draws the bottom info bar
func (ide *WindowsIDE) drawInfoBar() {
	y := ide.maxY - 1

	info := " F1:Help F2:Save F3:Open F5:Run F10:Menu "
	screenBuf.writeString(0, y, info, colorYellowOnBlack)

	// Fill rest of line
	for x := len(info); x < ide.maxX; x++ {
		screenBuf.set(x, y, ' ', colorYellowOnBlack)
	}
}

// drawSubmenu draws the active submenu
func (ide *WindowsIDE) drawSubmenu() {
	t := getIDETranslation()
	var items []string

	switch ide.menuSelected {
	case 0: // File
		items = []string{t.FileNew, t.FileOpen + "...", t.FileSave, t.FileSaveAs + "...", t.FileClose, t.FileExit}
	case 1: // Run
		items = []string{t.RunRun, t.RunStop, t.RunDebug}
	case 2: // Examples
		items = []string{t.ExamplesBrowse}
	case 3: // Help
		items = []string{t.HelpShortcuts, t.HelpLangRef, t.HelpAbout}
	case 4: // Language
		items = []string{t.LangEnglish, t.LangTurkish, t.LangFinnish, t.LangGerman}
	}

	// Calculate position under menu item dynamically
	menuItems := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}
	x := 0
	for i := 0; i < ide.menuSelected && i < len(menuItems); i++ {
		x += len([]rune(menuItems[i])) + 2 // " Item "
		x += 1                              // Space after each menu (matching menu bar drawing)
	}
	y := 1

	// Find max width using rune length for proper UTF-8 handling
	maxWidth := 0
	for _, item := range items {
		runeLen := len([]rune(item))
		if runeLen > maxWidth {
			maxWidth = runeLen
		}
	}
	maxWidth += 4

	// Draw submenu box
	for i, item := range items {
		var color uint16
		if i == ide.submenuSelected {
			color = colorBlackOnYellow
		} else {
			color = colorBlackOnWhite
		}

		itemRuneLen := len([]rune(item))
		itemStr := " " + item + strings.Repeat(" ", maxWidth-itemRuneLen-2)
		screenBuf.writeString(x, y+i, itemStr, color)
	}
}
