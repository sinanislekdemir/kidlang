//go:build windows
// +build windows

package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// LoadFile loads a file into the editor
func (ide *WindowsIDE) LoadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	ide.lines = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ide.lines = append(ide.lines, scanner.Text())
	}

	if len(ide.lines) == 0 {
		ide.lines = []string{""}
	}

	ide.filename = filename
	ide.modified = false
	ide.cursorX = 0
	ide.cursorY = 0
	ide.scrollY = 0

	return scanner.Err()
}

// SaveFile saves the current file
func (ide *WindowsIDE) SaveFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i, line := range ide.lines {
		if i > 0 {
			writer.WriteString("\n")
		}
		writer.WriteString(line)
	}

	err = writer.Flush()
	if err == nil {
		ide.filename = filename
		ide.modified = false
	}

	return err
}

// Execute runs the current program
func (ide *WindowsIDE) Execute() error {
	// Save to temp file and execute
	tmpFile := "temp_kidlang_execute.kid"
	content := strings.Join(ide.lines, "\n")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		ide.showMessage("Error saving temporary file: " + err.Error())
		return err
	}
	defer os.Remove(tmpFile)

	// Close IDE and restore console
	ide.Close()

	// Clear screen using Windows API
	fmt.Print("\033c") // Clear screen escape sequence

	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		exePath = "kidlang.exe" // fallback
	}

	msgs := getExecutionMessages()
	fmt.Println(msgs.Header)
	fmt.Println(msgs.StopInfo + "\n")

	// Run kidlang in a subprocess with clean terminal
	cmd := exec.Command(exePath, tmpFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("\nError running program: %v\n", err)
	}

	fmt.Println("\n" + msgs.Finished + " " + msgs.PressEnter)
	bufio.NewReader(os.Stdin).ReadString('\n')

	// Clear screen again before returning to IDE
	fmt.Print("\033c")

	// Reinitialize IDE
	ide.Init()
	ide.draw()

	return nil
}

// newFile creates a new file
func (ide *WindowsIDE) newFile() {
	if ide.modified {
		if !ide.promptSaveChanges() {
			return
		}
	}

	ide.lines = []string{""}
	ide.showWelcome = true
	ide.filename = ""
	ide.modified = false
	ide.cursorX = 0
	ide.cursorY = 0
	ide.scrollY = 0
}

// promptOpenFile prompts for a filename and opens it
func (ide *WindowsIDE) promptOpenFile() {
	if ide.modified {
		if !ide.promptSaveChanges() {
			return
		}
	}
	ide.showFileBrowser(true)
}

// saveCurrentFile saves the current file
func (ide *WindowsIDE) saveCurrentFile() {
	if ide.filename == "" {
		ide.promptSaveAs()
		return
	}

	err := ide.SaveFile(ide.filename)
	if err != nil {
		ide.showMessage("Error saving file: " + err.Error())
	}
}

// promptSaveAs prompts for a filename and saves
func (ide *WindowsIDE) promptSaveAs() {
	ide.showFileBrowser(false)
}

// closeCurrentFile closes the current file
func (ide *WindowsIDE) closeCurrentFile() {
	if ide.modified {
		if !ide.promptSaveChanges() {
			return
		}
	}
	ide.newFile()
}
