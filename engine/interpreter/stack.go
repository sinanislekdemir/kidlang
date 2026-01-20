package interpreter

import (
	"bufio"
	"os"
)

type KLError struct {
	Traceback    []int64 `json:"traceback"`
	Cursor       int64   `json:"cursor"`
	ErrorMessage string  `json:"error"`
}

type KLStack struct {
	Cursor    int      `json:"cursor"`
	Error     *KLError `json:"error"`
	ExitScope bool     `json:"exit_scope"`
	JumpLabel *string  `json:"jump_label"`

	// IO Control
	IN     *os.File
	OUT    *os.File
	ERR    *os.File
	Reader *bufio.Reader
}
