//go:build windows
// +build windows

package interpreter

import (
	"os"
	"unsafe"
)

// Key codes for Windows
const (
	VK_BACK   = 0x08
	VK_TAB    = 0x09
	VK_RETURN = 0x0D
	VK_ESCAPE = 0x1B
	VK_SPACE  = 0x20
	VK_PRIOR  = 0x21 // Page Up
	VK_NEXT   = 0x22 // Page Down
	VK_END    = 0x23
	VK_HOME   = 0x24
	VK_LEFT   = 0x25
	VK_UP     = 0x26
	VK_RIGHT  = 0x27
	VK_DOWN   = 0x28
	VK_DELETE = 0x2E
	VK_F1     = 0x70
	VK_F2     = 0x71
	VK_F3     = 0x72
	VK_F5     = 0x74
	VK_F10    = 0x79
)

// Console input event types
const (
	KEY_EVENT   = 0x0001
	MOUSE_EVENT = 0x0002
)

// Mouse button states
const (
	FROM_LEFT_1ST_BUTTON_PRESSED = 0x0001
	RIGHTMOST_BUTTON_PRESSED     = 0x0002
)

// Mouse event flags
const (
	MOUSE_MOVED   = 0x0001
	DOUBLE_CLICK  = 0x0002
	MOUSE_WHEELED = 0x0004
)

// readInput reads an input event (key or mouse) from the console
func (ide *WindowsIDE) readInput() inputRecord {
	var record inputRecord
	var numRead uint32

	for {
		ret, _, _ := procReadConsoleInput.Call(
			uintptr(ide.stdin),
			uintptr(unsafe.Pointer(&record)),
			1,
			uintptr(unsafe.Pointer(&numRead)),
		)

		if ret == 0 || numRead == 0 {
			continue
		}

		// Return key events (type 1) when key is down
		if record.EventType == KEY_EVENT {
			keyEvent := *(*keyEventRecord)(unsafe.Pointer(&record.Event[0]))
			if keyEvent.KeyDown != 0 {
				return record
			}
		}

		// Return mouse events (type 2)
		if record.EventType == MOUSE_EVENT {
			return record
		}
	}
}

// readKey reads a key press from the console (legacy compatibility)
func (ide *WindowsIDE) readKey() keyEventRecord {
	for {
		record := ide.readInput()
		if record.EventType == KEY_EVENT {
			return *(*keyEventRecord)(unsafe.Pointer(&record.Event[0]))
		}
	}
}

// handleInput processes keyboard input
func (ide *WindowsIDE) handleInput(key keyEventRecord) bool {
	// Disable welcome on first keystroke
	if ide.showWelcome && key.UnicodeChar > 0 {
		ide.showWelcome = false
	}
	
	// Check for Ctrl+C to quit
	if key.UnicodeChar == 3 || (key.VirtualKeyCode == 'C' && (key.ControlKeyState&0x0008) != 0) { // Ctrl pressed
		return false // Exit
	}

	needsRedraw := false

	// Check for Ctrl key combinations (when menu is not active)
	if !ide.menuActive && (key.ControlKeyState&0x0008) != 0 {
		switch key.VirtualKeyCode {
		case 'N': // Ctrl+N - New
			ide.newFile()
			needsRedraw = true
		case 'O': // Ctrl+O - Open
			ide.promptOpenFile()
			needsRedraw = true
		case 'S': // Ctrl+S - Save
			ide.saveCurrentFile()
			needsRedraw = true
		default:
			return true
		}
		if needsRedraw {
			ide.draw()
		}
		return true
	}

	// Menu navigation
	if ide.menuActive {
		if ide.submenuActive {
			switch key.VirtualKeyCode {
			case VK_UP:
				ide.submenuSelected--
				if ide.submenuSelected < 0 {
					ide.submenuSelected = ide.getSubmenuItemCount() - 1
				}
				needsRedraw = true
			case VK_DOWN:
				ide.submenuSelected++
				if ide.submenuSelected >= ide.getSubmenuItemCount() {
					ide.submenuSelected = 0
				}
				needsRedraw = true
			case VK_LEFT:
				ide.submenuSelected = 0
				ide.menuSelected--
				if ide.menuSelected < 0 {
					ide.menuSelected = 4
				}
				needsRedraw = true
			case VK_RIGHT:
				ide.submenuSelected = 0
				ide.menuSelected++
				if ide.menuSelected > 4 {
					ide.menuSelected = 0
				}
				needsRedraw = true
			case VK_ESCAPE:
				ide.menuActive = false
				ide.submenuActive = false
				ide.submenuSelected = 0
				needsRedraw = true
			case VK_RETURN:
				ide.executeSubmenuAction()
				needsRedraw = true
			}
		} else {
			switch key.VirtualKeyCode {
			case VK_LEFT:
				ide.menuSelected--
				if ide.menuSelected < 0 {
					ide.menuSelected = 4
				}
				needsRedraw = true
			case VK_RIGHT:
				ide.menuSelected++
				if ide.menuSelected > 4 {
					ide.menuSelected = 0
				}
				needsRedraw = true
			case VK_DOWN, VK_RETURN:
				ide.submenuActive = true
				ide.submenuSelected = 0
				needsRedraw = true
			case VK_ESCAPE:
				ide.menuActive = false
				needsRedraw = true
			}
		}
		if needsRedraw {
			ide.draw()
		}
		return true
	}

	// Normal editing mode
	switch key.VirtualKeyCode {
	case VK_F10:
		if !ide.menuActive {
			// Opening menu - show submenu immediately
			ide.menuActive = true
			ide.submenuActive = true
			ide.submenuSelected = 0
		} else {
			// Closing menu
			ide.menuActive = false
			ide.submenuActive = false
			ide.submenuSelected = 0
		}
		ide.draw()
		return true

	case VK_F1:
		ide.menuActive = true
		ide.menuSelected = 3 // Help
		ide.submenuActive = true
		ide.submenuSelected = 0
		ide.draw()
		return true

	case VK_F2:
		ide.saveCurrentFile()
		ide.draw()
		return true

	case VK_F3:
		ide.promptOpenFile()
		ide.draw()
		return true

	case VK_F5:
		ide.Execute()
		ide.draw()
		return true

	case VK_UP:
		if ide.cursorY > 0 {
			ide.cursorY--
			lineLen := len([]rune(ide.lines[ide.cursorY]))
			if ide.cursorX > lineLen {
				ide.cursorX = lineLen
			}
			ide.updateCursorPosition()
		}
		return true

	case VK_DOWN:
		if ide.cursorY < len(ide.lines)-1 {
			ide.cursorY++
			lineLen := len([]rune(ide.lines[ide.cursorY]))
			if ide.cursorX > lineLen {
				ide.cursorX = lineLen
			}
			ide.updateCursorPosition()
		}
		return true

	case VK_LEFT:
		if ide.cursorX > 0 {
			ide.cursorX--
			ide.updateCursorPosition()
		}
		return true

	case VK_RIGHT:
		lineLen := len([]rune(ide.lines[ide.cursorY]))
		if ide.cursorX < lineLen {
			ide.cursorX++
			ide.updateCursorPosition()
		}
		return true

	case VK_HOME:
		ide.cursorX = 0
		ide.updateCursorPosition()
		return true

	case VK_END:
		ide.cursorX = len([]rune(ide.lines[ide.cursorY]))
		ide.updateCursorPosition()
		return true

	case VK_PRIOR: // Page Up
		ide.cursorY -= (ide.maxY - 4)
		if ide.cursorY < 0 {
			ide.cursorY = 0
		}
		if ide.cursorX > len(ide.lines[ide.cursorY]) {
			ide.cursorX = len(ide.lines[ide.cursorY])
		}
		ide.draw()
		return true

	case VK_NEXT: // Page Down
		ide.cursorY += (ide.maxY - 4)
		if ide.cursorY >= len(ide.lines) {
			ide.cursorY = len(ide.lines) - 1
		}
		if ide.cursorX > len(ide.lines[ide.cursorY]) {
			ide.cursorX = len(ide.lines[ide.cursorY])
		}
		ide.draw()
		return true

	case VK_RETURN:
		ide.insertNewline()
		ide.draw()
		return true

	case VK_BACK:
		ide.deleteChar()
		ide.draw()
		return true

	case VK_DELETE:
		line := ide.lines[ide.cursorY]
		lineLen := len([]rune(line))
		if ide.cursorX < lineLen {
			// Delete character at cursor
			bytePos := runeIndexToByteIndex(line, ide.cursorX)
			nextBytePos := runeIndexToByteIndex(line, ide.cursorX+1)
			ide.lines[ide.cursorY] = line[:bytePos] + line[nextBytePos:]
			ide.modified = true
			ide.draw()
		} else if ide.cursorY < len(ide.lines)-1 {
			// Join with next line
			ide.lines[ide.cursorY] += ide.lines[ide.cursorY+1]
			ide.lines = append(ide.lines[:ide.cursorY+1], ide.lines[ide.cursorY+2:]...)
			ide.modified = true
			ide.draw()
		}
		return true

	case VK_TAB:
		ide.insertChar('\t')
		ide.draw()
		return true

	default:
		// Handle regular character input - only if it's not a control key
		if key.UnicodeChar >= 32 && key.UnicodeChar < 127 && key.VirtualKeyCode >= 32 {
			ide.insertChar(rune(key.UnicodeChar))
			ide.draw()
		}
		return true
	}
}

// executeSubmenuAction executes the selected menu action
func (ide *WindowsIDE) executeSubmenuAction() {
	switch ide.menuSelected {
	case 0: // File menu
		switch ide.submenuSelected {
		case 0: // New
			ide.newFile()
		case 1: // Open
			ide.promptOpenFile()
		case 2: // Save
			ide.saveCurrentFile()
		case 3: // Save As
			ide.promptSaveAs()
		case 4: // Close
			ide.closeCurrentFile()
		case 5: // Exit
			os.Exit(0)
		}
	case 1: // Run menu
		switch ide.submenuSelected {
		case 0: // Run (F5)
			ide.Execute()
		case 1: // Stop
			// TODO: implement stop
		case 2: // Debug
			// TODO: implement debug
		}
	case 2: // Examples menu
		switch ide.submenuSelected {
		case 0: // Browse Examples
			ide.showExamplesBrowser()
		}
	case 3: // Help menu
		switch ide.submenuSelected {
		case 0: // Keyboard Shortcuts
			ide.showKeyboardShortcuts()
		case 1: // Language Reference
			ide.showLanguageReference()
		case 2: // About
			ide.showAbout()
		}
	case 4: // Language menu
		switch ide.submenuSelected {
		case 0: // English
			setIDELanguage(IDE_LANG_EN)
			saveIDELanguageToConfig()
		case 1: // Turkish
			setIDELanguage(IDE_LANG_TR)
			saveIDELanguageToConfig()
		case 2: // Finnish
			setIDELanguage(IDE_LANG_FI)
			saveIDELanguageToConfig()
		case 3: // German
			setIDELanguage(IDE_LANG_DE)
			saveIDELanguageToConfig()
		}
	}

	ide.menuActive = false
	ide.submenuActive = false
	ide.submenuSelected = 0
}

// getSubmenuItemCount returns the number of items in the current submenu
func (ide *WindowsIDE) getSubmenuItemCount() int {
	return getSubmenuItemCount(ide.menuSelected)
}

// handleMouse processes mouse events
func (ide *WindowsIDE) handleMouse(mouseEvent mouseEventRecord) {
	x := int(mouseEvent.MousePosition.X)
	y := int(mouseEvent.MousePosition.Y)

	// Only handle left button clicks
	if (mouseEvent.ButtonState & FROM_LEFT_1ST_BUTTON_PRESSED) == 0 {
		return
	}

	// Check if clicked on menu bar (row 0)
	if y == 0 {
		ide.handleMenuBarClick(x)
		ide.draw()
		return
	}

	// Check if clicked on dropdown menu when active
	if ide.menuActive && ide.submenuActive {
		ide.handleSubmenuClick(x, y)
		ide.draw()
		return
	}

	// Check if click was outside menu to close it
	if ide.menuActive {
		ide.menuActive = false
		ide.submenuActive = false
		ide.submenuSelected = 0
		ide.draw()
	}
}

// handleMenuBarClick handles clicks on the menu bar
func (ide *WindowsIDE) handleMenuBarClick(x int) {
	t := getIDETranslation()
	menuItems := []string{t.MenuFile, t.MenuEdit, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}

	menuX := 0
	for i, item := range menuItems {
		menuWidth := len([]rune(item)) + 2 // " Item "
		if x >= menuX && x < menuX+menuWidth {
			// Clicked on this menu
			if ide.menuActive && ide.menuSelected == i {
				// Clicking same menu again - close it
				ide.menuActive = false
				ide.submenuActive = false
				ide.submenuSelected = 0
			} else {
				// Open this menu
				ide.menuActive = true
				ide.submenuActive = true
				ide.menuSelected = i
				ide.submenuSelected = 0
			}
			return
		}
		menuX += menuWidth + 1
	}
}

// handleSubmenuClick handles clicks on submenu items
func (ide *WindowsIDE) handleSubmenuClick(x, y int) {
	t := getIDETranslation()

	// Calculate menu position dynamically
	menuItems := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}
	menuX := 0
	for i := 0; i < ide.menuSelected && i < len(menuItems); i++ {
		menuX += len([]rune(menuItems[i])) + 2 // " Item "
		menuX += 1                             // Space between menus
	}
	menuY := 1

	// Get submenu items
	var items []string
	switch ide.menuSelected {
	case 0: // File
		items = []string{t.FileNew, t.FileOpen + "...", t.FileSave, t.FileSaveAs + "...", "Close", t.FileExit}
	case 1: // Run
		items = []string{t.RunRun, t.RunStop, "Debug"}
	case 2: // Examples
		items = []string{t.ExamplesBrowse}
	case 3: // Help
		items = []string{t.HelpShortcuts, "Language Reference", t.HelpAbout}
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
	maxWidth += 4

	// Check if click is within submenu bounds
	if x >= menuX && x < menuX+maxWidth && y >= menuY && y < menuY+len(items) {
		itemIndex := y - menuY
		if itemIndex >= 0 && itemIndex < len(items) {
			ide.submenuSelected = itemIndex
			ide.executeSubmenuAction()
		}
	}
}
