//go:build !windows
// +build !windows

package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gbin/goncurses"
)

// LoadFile loads a file into the editor
func (ide *UnixIDE) LoadFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	ide.lines = strings.Split(string(content), "\n")
	if len(ide.lines) == 0 {
		ide.lines = []string{""}
	}
	ide.filename = filename
	ide.modified = false
	ide.cursorX = 0
	ide.cursorY = 0
	ide.scrollY = 0
	return nil
}

// SaveFile saves the current file
func (ide *UnixIDE) SaveFile(filename string) error {
	content := strings.Join(ide.lines, "\n")
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	ide.filename = filename
	ide.modified = false
	return nil
}

// Execute runs the current program
func (ide *UnixIDE) Execute() error {
	// Properly end ncurses mode
	goncurses.End()

	// Save program to temporary file
	tmpFile, err := os.CreateTemp("", "kidlang-*.kid")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		fmt.Println("Press any key to continue...")
		var dummy string
		fmt.Scanln(&dummy)
		ide.Init()
		return err
	}
	defer os.Remove(tmpFile.Name())

	programText := strings.Join(ide.lines, "\n")
	if _, err := tmpFile.WriteString(programText); err != nil {
		fmt.Printf("Error writing temp file: %v\n", err)
		fmt.Println("Press any key to continue...")
		var dummy string
		fmt.Scanln(&dummy)
		ide.Init()
		return err
	}
	tmpFile.Close()

	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		exePath = "kidlang" // fallback
	}

	// Full terminal reset and clear
	fmt.Print("\033c")

	msgs := getExecutionMessages()
	fmt.Println(msgs.Header)
	fmt.Println(msgs.StopInfo + "\n")

	// Run kidlang in a subprocess with clean terminal
	cmd := exec.Command(exePath, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Printf("\nError starting program: %v\n", err)
		fmt.Println("\nPress Enter to continue...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Print("\033c")
		ide.Init()
		return err
	}

	// Set up channel to catch Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for program to finish or Ctrl+C
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		// Program finished normally
		signal.Stop(sigChan)
		if err != nil {
			fmt.Printf("\nProgram error: %v\n", err)
		}
	case <-sigChan:
		// Ctrl+C pressed - kill the subprocess
		fmt.Println("\n\n" + msgs.Interrupted)
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		// Wait a bit for process to die
		<-done
	}

	fmt.Println("\n" + msgs.Finished)
	fmt.Print(msgs.PressEnter)

	// Read a single line with bufio to ensure it works
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Full reset before returning to IDE
	fmt.Print("\033c")

	// Re-initialize ncurses
	return ide.Init()
}

// showFileBrowser shows a file browser dialog for opening or saving files
func (ide *UnixIDE) showFileBrowser(openMode bool) {
	currentDir, _ := os.Getwd()
	selectedFile := ""
	selectedIndex := 0
	scrollOffset := 0

	for {
		ide.screen.Clear()

		// Draw the editor background
		ide.screen.ColorOn(1)
		for row := 0; row < ide.maxY; row++ {
			ide.screen.MovePrint(row, 0, strings.Repeat(" ", ide.maxX))
		}
		ide.screen.ColorOff(1)

		// Calculate dialog dimensions
		dialogWidth := 60
		dialogHeight := 20
		dialogX := (ide.maxX - dialogWidth) / 2
		dialogY := (ide.maxY - dialogHeight) / 2

		// Draw dialog box
		ide.screen.ColorOn(5) // Black on white

		// Title
		var title string
		if openMode {
			title = "Open File"
		} else {
			title = "Save File As"
		}

		ide.screen.MovePrint(dialogY, dialogX, strings.Repeat(" ", dialogWidth))
		ide.screen.MovePrint(dialogY, dialogX+(dialogWidth-len(title))/2, title)

		// Draw dialog content area
		for row := 1; row < dialogHeight; row++ {
			ide.screen.MovePrint(dialogY+row, dialogX, strings.Repeat(" ", dialogWidth))
		}

		// Name field
		ide.screen.MovePrint(dialogY+2, dialogX+2, "Name")
		ide.screen.ColorOn(4) // Cyan background for input
		nameField := fmt.Sprintf(" %-*s ", dialogWidth-10, selectedFile)
		ide.screen.MovePrint(dialogY+3, dialogX+2, nameField)
		ide.screen.ColorOff(4)
		ide.screen.ColorOn(5)

		// OK button
		ide.screen.ColorOn(6) // Green button
		ide.screen.MovePrint(dialogY+3, dialogX+dialogWidth-8, "  Ok  ")
		ide.screen.ColorOff(6)
		ide.screen.ColorOn(5)

		// Files label
		ide.screen.MovePrint(dialogY+5, dialogX+2, "Files")

		// List files and directories
		fileList := buildFileList(currentDir)

		// Draw file list
		listHeight := dialogHeight - 10
		ide.screen.ColorOn(4) // Cyan background for list
		for row := 0; row < listHeight; row++ {
			ide.screen.MovePrint(dialogY+6+row, dialogX+2, strings.Repeat(" ", dialogWidth-14))

			fileIdx := scrollOffset + row
			if fileIdx < len(fileList) {
				displayName := fileList[fileIdx].Name
				if len(displayName) > dialogWidth-16 {
					displayName = displayName[:dialogWidth-16]
				}

				if fileIdx == selectedIndex {
					ide.screen.ColorOff(4)
					ide.screen.ColorOn(1) // Blue background for selection
					ide.screen.MovePrint(dialogY+6+row, dialogX+2, fmt.Sprintf(" %-*s ", dialogWidth-15, displayName))
					ide.screen.ColorOff(1)
					ide.screen.ColorOn(4)
				} else {
					ide.screen.MovePrint(dialogY+6+row, dialogX+3, displayName)
				}
			}
		}
		ide.screen.ColorOff(4)
		ide.screen.ColorOn(5)

		// Cancel button
		ide.screen.ColorOn(6) // Green button
		ide.screen.MovePrint(dialogY+6, dialogX+dialogWidth-10, " Cancel ")
		ide.screen.ColorOff(6)
		ide.screen.ColorOn(5)

		// Current directory path at bottom
		ide.screen.MovePrint(dialogY+dialogHeight-2, dialogX+2, fmt.Sprintf("%-*s", dialogWidth-4, currentDir))

		ide.screen.ColorOff(5)
		ide.screen.Refresh()

		// Handle input
		key := ide.screen.GetChar()

		switch key {
		case goncurses.KEY_UP:
			if selectedIndex > 0 {
				selectedIndex--
				if selectedIndex < scrollOffset {
					scrollOffset = selectedIndex
				}
			}
		case goncurses.KEY_DOWN:
			if selectedIndex < len(fileList)-1 {
				selectedIndex++
				if selectedIndex >= scrollOffset+listHeight {
					scrollOffset = selectedIndex - listHeight + 1
				}
			}
		case goncurses.KEY_RETURN, '\r':
			// If user has typed a filename in the Name field, save/open with that
			if selectedFile != "" {
				fullPath := currentDir + "/" + selectedFile
				if openMode {
					// Load the file
					if err := ide.LoadFile(fullPath); err != nil {
						// Show error but continue - user can see it didn't load
					}
				} else {
					// Save the file
					ide.SaveFile(fullPath)
				}
				return
			}

			// Otherwise, handle list selection
			if selectedIndex < len(fileList) {
				selected := fileList[selectedIndex]
				if selected.IsDir {
					// Directory selected - navigate
					currentDir = navigateDirectory(currentDir, selected)
					os.Chdir(currentDir)
					selectedIndex = 0
					scrollOffset = 0
					selectedFile = ""
				} else {
					// File selected from list
					selectedFile = selected.Name
					if openMode {
						// Load the file immediately
						fullPath := currentDir + "/" + selectedFile
						if err := ide.LoadFile(fullPath); err != nil {
							// Show error but continue - user can see it didn't load
						}
						return
					}
					// For save mode, just populate the name field, user can press Enter again
				}
			}
		case 27: // ESC - Cancel
			return
		default:
			// Typing to filter/select file
			if key >= 32 && key <= 126 {
				selectedFile += string(rune(key))
			} else if key == goncurses.KEY_BACKSPACE || key == 127 {
				if len(selectedFile) > 0 {
					selectedFile = selectedFile[:len(selectedFile)-1]
				}
			}
		}
	}
}
