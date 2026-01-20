//go:build windows
// +build windows

package interpreter

// showAbout displays about information
func (ide *WindowsIDE) showAbout() {
	ide.clearScreen()

	t := getIDETranslation()

	// Fill screen with black
	for y := 0; y < ide.maxY; y++ {
		for x := 0; x < ide.maxX; x++ {
			screenBuf.set(x, y, ' ', colorWhite)
		}
	}

	// Get about info
	about := getAboutInfo()
	centerY := ide.maxY/2 - 3

	// Draw title in yellow
	screenBuf.writeString((ide.maxX-len(about.Title))/2, centerY, about.Title, colorYellow)
	screenBuf.writeString((ide.maxX-len(about.Subtitle))/2, centerY+1, about.Subtitle, colorWhite)
	screenBuf.writeString((ide.maxX-len(about.Author))/2, centerY+3, about.Author, colorWhite)
	screenBuf.writeString((ide.maxX-len(about.Email))/2, centerY+4, about.Email, colorWhite)
	screenBuf.writeString((ide.maxX-len(about.Year))/2, centerY+5, about.Year, colorWhite)

	pressKey := t.MsgPressKey
	screenBuf.writeString((ide.maxX-len(pressKey))/2, ide.maxY-2, pressKey, colorWhite)

	ide.flushScreen()
	ide.readKey()
	ide.draw()
}

// showKeyboardShortcuts displays keyboard shortcuts help
func (ide *WindowsIDE) showKeyboardShortcuts() {
	ide.clearScreen()
	ide.drawMenuBar()
	ide.drawStatusBar()
	ide.drawInfoBar()

	t := getIDETranslation()
	dialog := newDialogBox(t.TitleKeyboardShort, 60, 20, ide.maxX, ide.maxY)
	dialog.draw()

	shortcuts := getKeyboardShortcuts()

	for i, shortcut := range shortcuts {
		if i < 15 {
			screenBuf.writeString(dialog.x+3, dialog.y+2+i, shortcut, colorBlackOnWhite)
		}
	}

	pressKey := t.MsgPressKey
	screenBuf.writeString(dialog.x+3, dialog.y+dialog.height-2, pressKey, colorBlackOnWhite)

	ide.flushScreen()
	ide.readKey()
	ide.draw()
}

// showLanguageReference displays language reference
func (ide *WindowsIDE) showLanguageReference() {
	ide.clearScreen()
	ide.drawMenuBar()
	ide.drawStatusBar()
	ide.drawInfoBar()

	dialog := newDialogBox("KidLang Language Reference", 70, 22, ide.maxX, ide.maxY)
	dialog.draw()

	commands := getLanguageReferenceCommands()

	for i, cmd := range commands {
		if i < 17 {
			screenBuf.writeString(dialog.x+3, dialog.y+2+i, cmd, colorBlackOnWhite)
		}
	}

	pressKey := "Press any key to return..."
	screenBuf.writeString(dialog.x+3, dialog.y+dialog.height-2, pressKey, colorBlackOnWhite)

	ide.flushScreen()
	ide.readKey()
	ide.draw()
}

// promptSaveChanges shows a dialog asking if user wants to save changes
// Returns: true if should continue with action, false if cancelled
func (ide *WindowsIDE) promptSaveChanges() bool {
	t := getIDETranslation()

	screenBuf.clear(colorWhite)

	centerY := ide.maxY / 2
	centerX := (ide.maxX - len([]rune(t.MsgSaveChanges))) / 2

	screenBuf.writeString(centerX, centerY, t.MsgSaveChanges, colorYellow)

	// Show options: Yes (Y) / No (N) / Cancel (ESC)
	options := "Y = Yes, N = No, ESC = Cancel"
	optX := (ide.maxX - len([]rune(options))) / 2
	screenBuf.writeString(optX, centerY+2, options, colorWhite)

	ide.flushScreen()

	for {
		event := ide.readKey()
		if event.KeyDown != 0 {
			switch event.UnicodeChar {
			case 'y', 'Y':
				// Save
				if ide.filename == "" {
					ide.showFileBrowser(false) // Show save dialog
				} else {
					ide.SaveFile(ide.filename)
				}
				return true
			case 'n', 'N':
				// Don't save, continue
				return true
			}
			// Check for ESC
			if event.VirtualKeyCode == VK_ESCAPE {
				// Cancel action
				return false
			}
		}
	}
}
