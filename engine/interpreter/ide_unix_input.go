//go:build !windows
// +build !windows

package interpreter

import (
	"os"

	"github.com/gbin/goncurses"
)

// handleInput processes keyboard input
func (ide *UnixIDE) handleInput(key goncurses.Key) {
	// Disable welcome on first keystroke
	if ide.showWelcome && key != goncurses.KEY_MOUSE && key > 0 {
		ide.showWelcome = false
	}
	
	// Handle mouse events
	if key == goncurses.KEY_MOUSE {
		ide.handleMouse()
		return
	}

	// Menu navigation
	if ide.menuActive {
		if ide.submenuActive {
			// Submenu is open - navigate items
			switch key {
			case goncurses.KEY_UP:
				ide.submenuSelected--
				if ide.submenuSelected < 0 {
					ide.submenuSelected = ide.getSubmenuItemCount() - 1
				}
				return
			case goncurses.KEY_DOWN:
				ide.submenuSelected++
				if ide.submenuSelected >= ide.getSubmenuItemCount() {
					ide.submenuSelected = 0
				}
				return
			case goncurses.KEY_LEFT:
				// Move to previous menu and keep submenu open
				ide.submenuSelected = 0
				ide.menuSelected--
				if ide.menuSelected < 0 {
					ide.menuSelected = 4
				}
				return
			case goncurses.KEY_RIGHT:
				// Move to next menu and keep submenu open
				ide.submenuSelected = 0
				ide.menuSelected++
				if ide.menuSelected > 4 {
					ide.menuSelected = 0
				}
				return
			case 27: // ESC
				ide.menuActive = false
				ide.submenuActive = false
				ide.submenuSelected = 0
				return
			case goncurses.KEY_RETURN, '\r':
				ide.handleMenuSelection()
				return
			}
		} else {
			// Top-level menu navigation
			switch key {
			case goncurses.KEY_LEFT:
				ide.menuSelected--
				if ide.menuSelected < 0 {
					ide.menuSelected = 4
				}
				return
			case goncurses.KEY_RIGHT:
				ide.menuSelected++
				if ide.menuSelected > 4 {
					ide.menuSelected = 0
				}
				return
			case goncurses.KEY_DOWN, goncurses.KEY_RETURN, '\r':
				// Open submenu
				ide.submenuActive = true
				ide.submenuSelected = 0
				return
			case 27: // ESC
				ide.menuActive = false
				return
			}
		}
	}

	// Global keys
	switch key {
	case goncurses.KEY_F10:
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
		return
	case goncurses.KEY_F5:
		ide.Execute()
		return
	}

	// Editor keys (only when menu is not active)
	if ide.menuActive {
		return
	}

	// Ctrl key combinations
	switch key {
	case 14: // Ctrl+N - New
		ide.lines = []string{""}
		ide.cursorX = 0
		ide.cursorY = 0
		ide.scrollY = 0
		ide.filename = ""
		ide.modified = false
		return
	case 15: // Ctrl+O - Open
		ide.showFileBrowser(true)
		return
	case 19: // Ctrl+S - Save
		if ide.filename == "" {
			ide.showFileBrowser(false)
		} else {
			ide.SaveFile(ide.filename)
		}
		return
	}

	switch key {
	case goncurses.KEY_UP:
		if ide.cursorY > 0 {
			ide.cursorY--
			ide.adjustScroll()
		}
	case goncurses.KEY_DOWN:
		if ide.cursorY < len(ide.lines)-1 {
			ide.cursorY++
			ide.adjustScroll()
		}
	case goncurses.KEY_LEFT:
		if ide.cursorX > 0 {
			ide.cursorX--
		}
	case goncurses.KEY_RIGHT:
		if ide.cursorX < len(ide.lines[ide.cursorY]) {
			ide.cursorX++
		}
	case goncurses.KEY_HOME:
		ide.cursorX = 0
	case goncurses.KEY_END:
		ide.cursorX = len([]rune(ide.lines[ide.cursorY]))
	case goncurses.KEY_BACKSPACE, 127:
		ide.handleBackspace()
	case goncurses.KEY_DC: // Delete key
		ide.handleDelete()
	case goncurses.KEY_RETURN, '\r':
		ide.handleEnter()
	default:
		if key >= 32 && key <= 126 { // Printable characters only
			ide.insertChar(rune(key))
		}
	}
}

// handleMenuSelection handles Enter key press in menu
func (ide *UnixIDE) handleMenuSelection() {
	switch ide.menuSelected {
	case 0: // File menu
		switch ide.submenuSelected {
		case 0: // New
			ide.lines = []string{""}
			ide.cursorX = 0
			ide.cursorY = 0
			ide.scrollY = 0
			ide.filename = ""
			ide.modified = false
			ide.menuActive = false
			ide.submenuActive = false
		case 1: // Open
			ide.menuActive = false
			ide.submenuActive = false
			ide.showFileBrowser(true) // true = open mode
		case 2: // Save
			if ide.filename == "" {
				ide.showFileBrowser(false) // false = save mode
			} else {
				ide.SaveFile(ide.filename)
			}
			ide.menuActive = false
			ide.submenuActive = false
		case 3: // Save As
			ide.menuActive = false
			ide.submenuActive = false
			ide.showFileBrowser(false) // false = save mode
		case 5: // Exit (skip separator at index 4)
			goncurses.End()
			os.Exit(0)
		}
	case 1: // Run menu
		switch ide.submenuSelected {
		case 0: // Run (F5)
			ide.menuActive = false
			ide.submenuActive = false
			ide.Execute()
		}
	case 2: // Examples menu
		switch ide.submenuSelected {
		case 0: // Browse Examples
			ide.menuActive = false
			ide.submenuActive = false
			ide.showExamples()
		}
	case 3: // Help menu
		switch ide.submenuSelected {
		case 0: // About
			ide.menuActive = false
			ide.submenuActive = false
			ide.showAbout()
		case 1: // Documentation
			ide.menuActive = false
			ide.submenuActive = false
			ide.showDocumentation()
		case 2: // Keyboard Shortcuts
			ide.menuActive = false
			ide.submenuActive = false
			ide.showKeyboardShortcuts()
		}
	case 4: // Language menu
		ide.menuActive = false
		ide.submenuActive = false
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
	default:
		ide.menuActive = false
		ide.submenuActive = false
	}
}

// handleMouse processes mouse events
// handleMouse processes mouse events
func (ide *UnixIDE) handleMouse() {
	mouseEvent := goncurses.GetMouse()
	if mouseEvent == nil {
		return
	}

	x := int(mouseEvent.X)
	y := int(mouseEvent.Y)

	// Check for any button 1 event (pressed, released, or clicked)
	isButton1 := (mouseEvent.State&goncurses.M_B1_PRESSED) != 0 ||
		(mouseEvent.State&goncurses.M_B1_RELEASED) != 0 ||
		(mouseEvent.State&goncurses.M_B1_CLICKED) != 0

	if !isButton1 {
		return
	}

	// Only process on button release for cleaner interaction
	if (mouseEvent.State & goncurses.M_B1_RELEASED) == 0 {
		return
	}

	// Check if clicked on menu bar (row 0)
	if y == 0 {
		ide.handleMenuBarClick(x)
		return
	}

	// Check if clicked on dropdown menu when active
	if ide.menuActive && ide.submenuActive {
		ide.handleSubmenuClick(x, y)
		return
	}

	// Check if click was outside menu to close it
	if ide.menuActive {
		ide.menuActive = false
		ide.submenuActive = false
		ide.submenuSelected = 0
	}
}

// handleMenuBarClick handles clicks on the menu bar
func (ide *UnixIDE) handleMenuBarClick(x int) {
	t := getIDETranslation()
	menuItems := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}

	menuX := 1
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
func (ide *UnixIDE) handleSubmenuClick(x, y int) {
	t := getIDETranslation()

	// Calculate menu position
	menuX := 1
	menuItems := []string{t.MenuFile, t.MenuRun, t.MenuExamples, t.MenuHelp, t.MenuLanguage}
	for i := 0; i < ide.menuSelected && i < len(menuItems); i++ {
		menuX += len([]rune(menuItems[i])) + 3
	}

	// Get submenu items
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

	// Find max width
	maxWidth := 0
	for _, item := range items {
		if len(item) > maxWidth {
			maxWidth = len([]rune(item))
		}
	}
	maxWidth += 4

	menuY := 1

	// Check if click is within submenu bounds
	if x >= menuX && x < menuX+maxWidth && y >= menuY && y < menuY+len(items) {
		itemIndex := y - menuY
		if itemIndex >= 0 && itemIndex < len(items) {
			// Skip separator lines
			if items[itemIndex] != "---" {
				ide.submenuSelected = itemIndex
				ide.handleMenuSelection()
			}
		}
	}
}
