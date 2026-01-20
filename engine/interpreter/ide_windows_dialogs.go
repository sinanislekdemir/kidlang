//go:build windows
// +build windows

package interpreter

import (
	"os"
	"path/filepath"
	"strings"
)

// dialogBox represents a reusable dialog component
type dialogBox struct {
	title        string
	width        int
	height       int
	x            int
	y            int
	items        []string
	selected     int
	scrollOffset int
	inputText    string
}

// newDialogBox creates a new dialog box centered on screen
func newDialogBox(title string, width, height int, maxX, maxY int) *dialogBox {
	return &dialogBox{
		title:    title,
		width:    width,
		height:   height,
		x:        (maxX - width) / 2,
		y:        (maxY - height) / 2,
		selected: 0,
	}
}

// draw draws the dialog box on the screen buffer
func (d *dialogBox) draw() {
	// Draw title bar
	screenBuf.fillRect(d.x, d.y, d.width, 1, ' ', colorBlackOnWhite)
	titleX := d.x + (d.width-len(d.title))/2
	screenBuf.writeString(titleX, d.y, d.title, colorBlackOnWhite)

	// Draw dialog content area
	for row := 1; row < d.height; row++ {
		screenBuf.fillRect(d.x, d.y+row, d.width, 1, ' ', colorBlackOnWhite)
	}
}

// drawList draws a list of items in the dialog
func (d *dialogBox) drawList(startY, listHeight int) {
	visibleItems := listHeight
	if d.scrollOffset+visibleItems > len(d.items) {
		d.scrollOffset = len(d.items) - visibleItems
		if d.scrollOffset < 0 {
			d.scrollOffset = 0
		}
	}

	for i := 0; i < visibleItems; i++ {
		itemIdx := d.scrollOffset + i
		y := d.y + startY + i

		if itemIdx < len(d.items) {
			item := d.items[itemIdx]
			if len(item) > d.width-6 {
				item = item[:d.width-6]
			}

			if itemIdx == d.selected {
				// Selected item - yellow background
				screenBuf.writeString(d.x+2, y, " "+item, colorBlackOnYellow)
				for j := len(item) + 1; j < d.width-4; j++ {
					screenBuf.set(d.x+2+j, y, ' ', colorBlackOnYellow)
				}
			} else {
				// Normal item
				screenBuf.writeString(d.x+3, y, item, colorBlackOnWhite)
			}
		}
	}
}

// drawInputField draws an input field with text
func (d *dialogBox) drawInputField(y int, label string) {
	screenBuf.writeString(d.x+2, y, label, colorBlackOnWhite)

	// Input field with cyan background
	fieldY := y + 1
	fieldWidth := d.width - 4
	screenBuf.fillRect(d.x+2, fieldY, fieldWidth, 1, ' ', colorCyan)

	displayText := d.inputText
	if len(displayText) > fieldWidth-2 {
		displayText = displayText[len(displayText)-(fieldWidth-2):]
	}
	screenBuf.writeString(d.x+3, fieldY, displayText, colorCyan)
}

// drawButton draws a button
func (d *dialogBox) drawButton(x, y int, label string, selected bool) {
	var color uint16
	color = colorBlackOnWhite
	if selected {
		color = colorBlackOnYellow
	}
	buttonText := "  " + label + "  "
	screenBuf.writeString(x, y, buttonText, color)
}

// fillRect is a helper to fill a rectangle in the screen buffer
func (sb *screenBuffer) fillRect(x, y, width, height int, ch rune, attr uint16) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			sb.set(x+col, y+row, ch, attr)
		}
	}
}

// showFileBrowser shows a file browser dialog
func (ide *WindowsIDE) showFileBrowser(openMode bool) {
	currentDir, _ := os.Getwd()
	selectedIndex := 0
	scrollOffset := 0
	inputText := ""

	for {
		ide.clearScreen()
		ide.drawMenuBar()
		ide.drawStatusBar()
		ide.drawInfoBar()

		// Create dialog
		dialog := newDialogBox("Open File", 60, 20, ide.maxX, ide.maxY)
		if !openMode {
			dialog.title = "Save File As"
		}

		dialog.draw()

		// Draw input field
		dialog.inputText = inputText
		dialog.drawInputField(dialog.y+2, "Name")

		// Draw OK button
		dialog.drawButton(dialog.x+dialog.width-10, dialog.y+3, "Ok", false)

		// Files label
		screenBuf.writeString(dialog.x+2, dialog.y+5, "Files", colorBlackOnWhite)

		// List files and directories
		entries, _ := os.ReadDir(currentDir)
		fileList := []string{}

		// Add parent directory
		if currentDir != "/" && currentDir != filepath.VolumeName(currentDir)+"\\" {
			fileList = append(fileList, "../")
		}

		// Add directories first
		for _, entry := range entries {
			if entry.IsDir() {
				fileList = append(fileList, entry.Name()+"/")
			}
		}

		// Add .kid files
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".kid") {
				fileList = append(fileList, entry.Name())
			}
		}

		// Draw file list
		dialog.items = fileList
		dialog.selected = selectedIndex
		dialog.scrollOffset = scrollOffset
		dialog.drawList(6, 10)

		// Current directory
		dirText := "Dir: " + currentDir
		if len(dirText) > dialog.width-4 {
			dirText = "..." + dirText[len(dirText)-(dialog.width-7):]
		}
		screenBuf.writeString(dialog.x+2, dialog.y+dialog.height-2, dirText, colorBlackOnWhite)

		ide.flushScreen()

		// Handle input
		key := ide.readKey()

		switch key.VirtualKeyCode {
		case VK_UP:
			if selectedIndex > 0 {
				selectedIndex--
				if selectedIndex < scrollOffset {
					scrollOffset = selectedIndex
				}
			}
		case VK_DOWN:
			if selectedIndex < len(fileList)-1 {
				selectedIndex++
				if selectedIndex >= scrollOffset+10 {
					scrollOffset = selectedIndex - 9
				}
			}
		case VK_RETURN:
			if selectedIndex < len(fileList) {
				selected := fileList[selectedIndex]
				if strings.HasSuffix(selected, "/") {
					// Navigate to directory
					if selected == "../" {
						currentDir = filepath.Dir(currentDir)
					} else {
						currentDir = filepath.Join(currentDir, strings.TrimSuffix(selected, "/"))
					}
					selectedIndex = 0
					scrollOffset = 0
				} else {
					// Select file
					if openMode {
						err := ide.LoadFile(filepath.Join(currentDir, selected))
						if err != nil {
							ide.showMessage("Error opening file: " + err.Error())
						}
					} else {
						err := ide.SaveFile(filepath.Join(currentDir, selected))
						if err != nil {
							ide.showMessage("Error saving file: " + err.Error())
						}
					}
					ide.draw()
					return
				}
			}
		case VK_ESCAPE:
			ide.draw()
			return
		default:
			// Handle text input for filename
			if key.UnicodeChar >= 32 && key.UnicodeChar < 127 {
				inputText += string(rune(key.UnicodeChar))
			} else if key.VirtualKeyCode == VK_BACK && len(inputText) > 0 {
				inputText = inputText[:len(inputText)-1]
			}
		}
	}
}

// showExamplesBrowser shows a browser for example files
func (ide *WindowsIDE) showExamplesBrowser() {
	// Get the directory of the executable
	exePath, err := os.Executable()
	if err != nil {
		ide.showMessage("Cannot find executable path")
		return
	}

	// Get the directory containing the executable
	exeDir := filepath.Dir(exePath)

	// Look for Examples directory (capital E to match Unix)
	examplesDir := filepath.Join(exeDir, "Examples")
	_, err = os.Stat(examplesDir)
	if err != nil {
		ide.showMessage("Examples directory not found")
		return
	}

	// First level: Show subdirectories (categories)
	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		ide.showMessage("Error reading examples: " + err.Error())
		return
	}

	var subdirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			subdirs = append(subdirs, entry.Name())
		}
	}

	if len(subdirs) == 0 {
		ide.showMessage("No example categories found")
		return
	}

	selected := 0
	scrollOffset := 0

	for {
		ide.clearScreen()
		ide.drawMenuBar()
		ide.drawStatusBar()
		ide.drawInfoBar()

		dialog := newDialogBox("Examples - Select Category", 60, 20, ide.maxX, ide.maxY)
		dialog.draw()

		// Instructions
		instrY := dialog.y + 2
		screenBuf.writeString(dialog.x+2, instrY, "Select an example category:", colorBlackOnWhite)

		// Draw subdirectory list
		dialog.items = subdirs
		dialog.selected = selected
		dialog.scrollOffset = scrollOffset
		dialog.drawList(4, 13)

		// Instructions at bottom
		instr := "Press Enter to open, ESC to cancel"
		screenBuf.writeString(dialog.x+2, dialog.y+dialog.height-2, instr, colorBlackOnWhite)

		ide.flushScreen()

		key := ide.readKey()

		switch key.VirtualKeyCode {
		case VK_UP:
			if selected > 0 {
				selected--
				if selected < scrollOffset {
					scrollOffset = selected
				}
			}
		case VK_DOWN:
			if selected < len(subdirs)-1 {
				selected++
				if selected >= scrollOffset+13 {
					scrollOffset = selected - 12
				}
			}
		case VK_RETURN:
			// Show files in selected subdirectory
			subdir := filepath.Join(examplesDir, subdirs[selected])
			if ide.showExampleFiles(subdir) {
				// File was loaded, exit
				ide.draw()
				return
			}
			// User pressed ESC in file browser, continue showing categories
		case VK_ESCAPE:
			ide.draw()
			return
		}
	}
}

// showExampleFiles shows files in a specific examples subdirectory
func (ide *WindowsIDE) showExampleFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}

	var exampleFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".kid") {
			exampleFiles = append(exampleFiles, entry.Name())
		}
	}

	if len(exampleFiles) == 0 {
		return false
	}

	selected := 0
	scrollOffset := 0

	for {
		ide.clearScreen()
		ide.drawMenuBar()
		ide.drawStatusBar()
		ide.drawInfoBar()

		dialog := newDialogBox("Examples - Select File", 60, 20, ide.maxX, ide.maxY)
		dialog.draw()

		// Instructions
		instrY := dialog.y + 2
		screenBuf.writeString(dialog.x+2, instrY, "Select an example to load:", colorBlackOnWhite)

		// Draw file list
		dialog.items = exampleFiles
		dialog.selected = selected
		dialog.scrollOffset = scrollOffset
		dialog.drawList(4, 13)

		// Instructions at bottom
		instr := "Press Enter to load, ESC to go back"
		screenBuf.writeString(dialog.x+2, dialog.y+dialog.height-2, instr, colorBlackOnWhite)

		ide.flushScreen()

		key := ide.readKey()

		switch key.VirtualKeyCode {
		case VK_UP:
			if selected > 0 {
				selected--
				if selected < scrollOffset {
					scrollOffset = selected
				}
			}
		case VK_DOWN:
			if selected < len(exampleFiles)-1 {
				selected++
				if selected >= scrollOffset+13 {
					scrollOffset = selected - 12
				}
			}
		case VK_RETURN:
			filename := filepath.Join(dir, exampleFiles[selected])
			err := ide.LoadFile(filename)
			if err != nil {
				ide.showMessage("Error loading example: " + err.Error())
				return false
			}
			return true
		case VK_ESCAPE:
			return false
		}
	}
}

// showMessage shows a message dialog
func (ide *WindowsIDE) showMessage(msg string) {
	ide.clearScreen()
	ide.drawMenuBar()
	ide.drawStatusBar()
	ide.drawInfoBar()

	// Wrap message into multiple lines if needed
	lines := wrapText(msg, 50)

	dialog := newDialogBox("Message", 56, len(lines)+6, ide.maxX, ide.maxY)
	dialog.draw()

	// Draw message lines
	for i, line := range lines {
		screenBuf.writeString(dialog.x+3, dialog.y+2+i, line, colorBlackOnWhite)
	}

	// OK button
	dialog.drawButton(dialog.x+dialog.width/2-4, dialog.y+dialog.height-2, "OK", true)

	ide.flushScreen()

	// Wait for Enter or ESC
	for {
		key := ide.readKey()
		if key.VirtualKeyCode == VK_RETURN || key.VirtualKeyCode == VK_ESCAPE {
			break
		}
	}
}
