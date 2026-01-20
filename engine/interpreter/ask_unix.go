//go:build !windows
// +build !windows

package interpreter

import (
	"bufio"
)

// readLineWindows is the Unix version (just calls ReadLine directly)
func readLineWindows(reader *bufio.Reader) (string, error) {
	buffer, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
