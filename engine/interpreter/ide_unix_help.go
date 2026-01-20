//go:build !windows
// +build !windows

package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/gbin/goncurses"
)

func (ide *UnixIDE) showAbout() {
	ide.screen.Clear()

	t := getIDETranslation()

	// Draw black background
	ide.screen.ColorOn(1)
	for row := 0; row < ide.maxY; row++ {
		ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
	}

	// Get about info
	about := getAboutInfo()
	centerY := ide.maxY/2 - 3

	// Draw title in yellow
	ide.screen.ColorOff(1)
	ide.screen.ColorOn(2)
	ide.screen.MovePrint(centerY, (ide.maxX-len(about.Title))/2, about.Title)
	ide.screen.ColorOff(2)

	ide.screen.ColorOn(1)
	ide.screen.MovePrint(centerY+1, (ide.maxX-len(about.Subtitle))/2, about.Subtitle)
	ide.screen.MovePrint(centerY+3, (ide.maxX-len(about.Author))/2, about.Author)
	ide.screen.MovePrint(centerY+4, (ide.maxX-len(about.Email))/2, about.Email)
	ide.screen.MovePrint(centerY+5, (ide.maxX-len(about.Year))/2, about.Year)

	pressKey := t.MsgPressKey
	ide.screen.MovePrint(ide.maxY-2, (ide.maxX-len(pressKey))/2, pressKey)
	ide.screen.ColorOff(1)

	ide.screen.Refresh()
	ide.screen.GetChar()
}

// showDocumentation shows a file picker for documentation files
func (ide *UnixIDE) showDocumentation() {
	docsDir, err := getDocsDir()
	if err != nil {
		return
	}

	docFiles := listMarkdownFiles(docsDir)
	if len(docFiles) == 0 {
		return
	}

	scrollState := listScrollState{}

	t := getIDETranslation()

	for {
		ide.screen.Clear()

		// Draw black background
		ide.screen.ColorOn(1)
		for row := 0; row < ide.maxY; row++ {
			ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
		}

		// Title
		title := t.TitleDocumentation
		ide.screen.ColorOff(1)
		ide.screen.ColorOn(2)
		ide.screen.MovePrint(0, (ide.maxX-len(title))/2, title)
		ide.screen.ColorOff(2)
		ide.screen.ColorOn(1)

		// Draw file list
		listHeight := ide.maxY - 4
		listY := 2

		for row := 0; row < listHeight && scrollState.ScrollOffset+row < len(docFiles); row++ {
			fileIdx := scrollState.ScrollOffset + row
			filename := docFiles[fileIdx]

			if fileIdx == scrollState.SelectedIndex {
				ide.screen.ColorOff(1)
				ide.screen.ColorOn(4) // Yellow background for selection
				ide.screen.MovePrint(listY+row, 2, fmt.Sprintf(" %-*s ", ide.maxX-5, filename))
				ide.screen.ColorOff(4)
				ide.screen.ColorOn(1)
			} else {
				ide.screen.MovePrint(listY+row, 3, filename)
			}
		}

		// Help text
		helpText := t.HelpFileSelect
		ide.screen.MovePrint(ide.maxY-1, (ide.maxX-len(helpText))/2, helpText)
		ide.screen.ColorOff(1)

		ide.screen.Refresh()

		key := ide.screen.GetChar()
		switch key {
		case goncurses.KEY_UP:
			scrollState.moveUp()
		case goncurses.KEY_DOWN:
			scrollState.moveDown(len(docFiles), listHeight)
		case goncurses.KEY_RETURN, '\r':
			// View the selected document
			ide.viewMarkdownFile(docsDir + "/" + docFiles[scrollState.SelectedIndex])
		case 27: // ESC
			return
		}
	}
}

// viewMarkdownFile displays a markdown file with glamour rendering
func (ide *UnixIDE) viewMarkdownFile(filepath string) {
	// Exit ncurses temporarily
	goncurses.End()

	// Clear screen
	fmt.Print("\033[2J\033[H")

	// Read the markdown file
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		ide.Init()
		return
	}

	// Use glamour to render the markdown
	rendered, err := renderMarkdown(string(content), ide.maxX-4)
	if err != nil {
		// Fallback to plain text if glamour fails
		rendered = string(content)
	}

	// Use less for scrolling if available, otherwise just print
	useLess := true

	// Create a temporary lesskey file that maps ESC to quit
	tmpLesskey, err := os.CreateTemp("", "kidlang-lesskey-*")
	if err == nil {
		defer os.Remove(tmpLesskey.Name())
		// Lesskey format: command key
		// \e is ESC, quit is the quit command
		tmpLesskey.WriteString("\\e quit\n")
		tmpLesskey.WriteString("q quit\n")
		tmpLesskey.Close()
	}

	lessCmd := exec.Command("less", "-R", "-K") // -R for color, -K for Ctrl+C exits
	lessCmd.Stdin = strings.NewReader(rendered)
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr
	// Add lesskey file if created successfully
	if err == nil {
		lessCmd.Env = append(os.Environ(), "LESSKEY="+tmpLesskey.Name())
	}

	if err := lessCmd.Run(); err != nil {
		// less not available or failed, just print and wait
		useLess = false
		fmt.Print(rendered)
		fmt.Println("\n\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	// Clear screen before returning
	if !useLess {
		fmt.Print("\033[2J\033[H")
	}

	// Re-initialize ncurses
	ide.Init()
}

// renderMarkdown renders markdown content using glamour
func renderMarkdown(content string, width int) (string, error) {
	// Create a custom dark style renderer that matches our Batman theme
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return "", err
	}

	rendered, err := r.Render(content)
	if err != nil {
		return "", err
	}

	return rendered, nil
}

// showKeyboardShortcuts displays keyboard shortcuts
func (ide *UnixIDE) showKeyboardShortcuts() {
	ide.screen.Clear()

	t := getIDETranslation()

	// Draw black background
	ide.screen.ColorOn(1)
	for row := 0; row < ide.maxY; row++ {
		ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
	}

	// Title
	title := t.TitleKeyboardShort
	ide.screen.ColorOff(1)
	ide.screen.ColorOn(2)
	ide.screen.MovePrint(1, (ide.maxX-len(title))/2, title)
	ide.screen.ColorOff(2)
	ide.screen.ColorOn(1)

	shortcuts := getKeyboardShortcuts()

	startY := 3
	for i, line := range shortcuts {
		ide.screen.MovePrint(startY+i, 4, line)
	}

	pressKey := t.MsgPressKey
	ide.screen.MovePrint(ide.maxY-2, (ide.maxX-len(pressKey))/2, pressKey)
	ide.screen.ColorOff(1)

	ide.screen.Refresh()
	ide.screen.GetChar()
}

// showFileBrowser shows a file browser dialog for opening or saving files

// showExamples shows a file picker for example files
func (ide *UnixIDE) showExamples() {
	examplesDir, err := getExamplesDir()
	if err != nil {
		return
	}

	subdirs := listSubdirectories(examplesDir)
	if len(subdirs) == 0 {
		return
	}

	scrollState := listScrollState{}
	t := getIDETranslation()

	for {
		ide.screen.Clear()

		// Draw black background
		ide.screen.ColorOn(1)
		for row := 0; row < ide.maxY; row++ {
			ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
		}

		// Title
		title := t.TitleExamplesSelect
		ide.screen.ColorOff(1)
		ide.screen.ColorOn(2)
		ide.screen.MovePrint(0, (ide.maxX-len(title))/2, title)
		ide.screen.ColorOff(2)
		ide.screen.ColorOn(1)

		// Draw subdirectory list
		listHeight := ide.maxY - 4
		listY := 2

		for row := 0; row < listHeight && scrollState.ScrollOffset+row < len(subdirs); row++ {
			fileIdx := scrollState.ScrollOffset + row
			dirname := subdirs[fileIdx]

			if fileIdx == scrollState.SelectedIndex {
				ide.screen.ColorOff(1)
				ide.screen.ColorOn(4) // Yellow background for selection
				ide.screen.MovePrint(listY+row, 2, fmt.Sprintf(" %-*s ", ide.maxX-5, dirname))
				ide.screen.ColorOff(4)
				ide.screen.ColorOn(1)
			} else {
				ide.screen.MovePrint(listY+row, 3, dirname)
			}
		}

		// Help text
		helpText := t.HelpFileSelect
		ide.screen.MovePrint(ide.maxY-1, (ide.maxX-len(helpText))/2, helpText)
		ide.screen.ColorOff(1)

		ide.screen.Refresh()

		key := ide.screen.GetChar()
		switch key {
		case goncurses.KEY_UP:
			scrollState.moveUp()
		case goncurses.KEY_DOWN:
			scrollState.moveDown(len(subdirs), listHeight)
		case goncurses.KEY_RETURN, '\r':
			// Show files in selected subdirectory
			subdir := examplesDir + "/" + subdirs[scrollState.SelectedIndex]
			if ide.showExampleFiles(subdir) {
				// File was loaded, exit
				return
			}
		// User pressed ESC, continue showing categories
		case 27: // ESC
			return
		}
	}
}

// showExampleFiles shows files in a specific examples subdirectory
func (ide *UnixIDE) showExampleFiles(dir string) bool {
	exampleFiles := listKidFiles(dir)
	if len(exampleFiles) == 0 {
		return false
	}

	scrollState := listScrollState{}

	t := getIDETranslation()
	// Get category name from path
	categoryName := dir[strings.LastIndex(dir, "/")+1:]

	for {
		ide.screen.Clear()

		// Draw black background
		ide.screen.ColorOn(1)
		for row := 0; row < ide.maxY; row++ {
			ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
		}

		// Title
		title := t.TitleExamplesPrefix + categoryName
		ide.screen.ColorOff(1)
		ide.screen.ColorOn(2)
		ide.screen.MovePrint(0, (ide.maxX-len(title))/2, title)
		ide.screen.ColorOff(2)
		ide.screen.ColorOn(1)

		// Draw file list
		listHeight := ide.maxY - 4
		listY := 2

		for row := 0; row < listHeight && scrollState.ScrollOffset+row < len(exampleFiles); row++ {
			fileIdx := scrollState.ScrollOffset + row
			filename := exampleFiles[fileIdx]

			if fileIdx == scrollState.SelectedIndex {
				ide.screen.ColorOff(1)
				ide.screen.ColorOn(4) // Yellow background for selection
				ide.screen.MovePrint(listY+row, 2, fmt.Sprintf(" %-*s ", ide.maxX-5, filename))
				ide.screen.ColorOff(4)
				ide.screen.ColorOn(1)
			} else {
				ide.screen.MovePrint(listY+row, 3, filename)
			}
		}

		// Help text
		helpText := t.HelpFileSelect
		ide.screen.MovePrint(ide.maxY-1, (ide.maxX-len(helpText))/2, helpText)
		ide.screen.ColorOff(1)

		ide.screen.Refresh()

		key := ide.screen.GetChar()
		switch key {
		case goncurses.KEY_UP:
			scrollState.moveUp()
		case goncurses.KEY_DOWN:
			scrollState.moveDown(len(exampleFiles), listHeight)
		case goncurses.KEY_RETURN, '\r':
			// Load the selected example
			fullPath := dir + "/" + exampleFiles[scrollState.SelectedIndex]
			if err := ide.LoadFile(fullPath); err == nil {
				// Successfully loaded - clear modified flag since it's an example
				ide.modified = false
			}
			return true // File was loaded (or attempted), exit to editor
		case 27: // ESC
			return false // Go back to category selection
		}
	}
}

// promptSaveChanges shows a dialog asking if user wants to save changes
// Returns: true if should continue with action, false if cancelled
func (ide *UnixIDE) promptSaveChanges() bool {
	t := getIDETranslation()
	
	ide.screen.Clear()
	ide.screen.ColorOn(1)
	for row := 0; row < ide.maxY; row++ {
		ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
	}
	
	centerY := ide.maxY / 2
	centerX := (ide.maxX - len([]rune(t.MsgSaveChanges))) / 2
	
	ide.screen.ColorOn(2)
	ide.screen.MovePrint(centerY, centerX, t.MsgSaveChanges)
	ide.screen.ColorOff(2)
	
	// Show options: Yes (Y) / No (N) / Cancel (ESC)
	options := "Y = Yes, N = No, ESC = Cancel"
	optX := (ide.maxX - len([]rune(options))) / 2
	ide.screen.MovePrint(centerY+2, optX, options)
	
	ide.screen.ColorOff(1)
	ide.screen.Refresh()
	
	for {
		ch := ide.screen.GetChar()
		switch ch {
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
		case 27: // ESC
			// Cancel action
			return false
		}
	}
}
