//go:build windows
// +build windows

package interpreter

import (
	"bufio"
	"os"

	"golang.org/x/sys/windows"
)

// ensureConsoleInputMode ensures the console is in the right mode for reading input
func ensureConsoleInputMode() {
	stdin := windows.Handle(os.Stdin.Fd())
	var mode uint32
	err := windows.GetConsoleMode(stdin, &mode)
	if err != nil {
		return // Not a console, probably redirected
	}

	// Enable line input and echo for normal Ask operation
	mode |= windows.ENABLE_LINE_INPUT
	mode |= windows.ENABLE_ECHO_INPUT
	mode |= windows.ENABLE_PROCESSED_INPUT

	windows.SetConsoleMode(stdin, mode)
}

// readLineWindows reads a line with proper console handling on Windows
func readLineWindows(reader *bufio.Reader) (string, error) {
	ensureConsoleInputMode()
	buffer, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
