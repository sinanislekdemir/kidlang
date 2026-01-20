//go:build !windows
// +build !windows

package interpreter

import (
	"fmt"
	"os"

	"github.com/gbin/goncurses"
)

/*
#cgo LDFLAGS: -lncursesw
#include <locale.h>
#include <stdlib.h>
*/
import "C"

// UnixIDE implements the IDE interface for Unix systems using ncurses
type UnixIDE struct {
	screen          *goncurses.Window
	lines           []string
	cursorX         int
	cursorY         int
	scrollY         int
	filename        string
	modified        bool
	maxY            int
	maxX            int
	menuActive      bool
	menuSelected    int // Which menu item is selected (0=File, 1=Edit, 2=Run, 3=Help)
	submenuActive   bool
	submenuSelected int
}

// NewIDE creates a new platform-specific IDE instance
func NewIDE() IDE {
	return &UnixIDE{
		lines: []string{""},
	}
}

// Init initializes the ncurses-based IDE
func (ide *UnixIDE) Init() error {
	// Set locale for UTF-8 support using C locale functions
	C.setlocale(C.LC_ALL, C.CString(""))

	os.Setenv("LANG", "en_US.UTF-8")
	os.Setenv("LC_ALL", "en_US.UTF-8")
	os.Setenv("NCURSES_NO_UTF8_ACS", "1")

	var err error
	ide.screen, err = goncurses.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize ncurses: %w", err)
	}

	goncurses.Echo(false)
	goncurses.CBreak(true)
	goncurses.Raw(true) // Enable raw mode to disable flow control (Ctrl+S/Ctrl+Q)
	goncurses.Cursor(1)
	ide.screen.Keypad(true)
	ide.screen.Timeout(-1) // Block until input

	// Enable mouse support
	goncurses.MouseMask(goncurses.M_ALL, nil)
	goncurses.MouseInterval(0)

	// Initialize colors for Batman theme (dark with yellow highlights)
	if goncurses.HasColors() {
		goncurses.StartColor()
		goncurses.InitPair(1, goncurses.C_WHITE, goncurses.C_BLACK)   // Normal text - white on black
		goncurses.InitPair(2, goncurses.C_YELLOW, goncurses.C_BLACK)  // Keywords - yellow on black
		goncurses.InitPair(3, goncurses.C_CYAN, goncurses.C_BLACK)    // Strings - cyan on black
		goncurses.InitPair(4, goncurses.C_BLACK, goncurses.C_YELLOW)  // Menu bar - black on yellow
		goncurses.InitPair(5, goncurses.C_BLACK, goncurses.C_WHITE)   // Status bar - black on white
		goncurses.InitPair(6, goncurses.C_GREEN, goncurses.C_BLACK)   // Comments - green on black
		goncurses.InitPair(7, goncurses.C_YELLOW, goncurses.C_BLACK)  // Line numbers - yellow on black
		goncurses.InitPair(8, goncurses.C_MAGENTA, goncurses.C_BLACK) // Numbers - magenta on black
	}

	ide.maxY, ide.maxX = ide.screen.MaxYX()
	return nil
}

// Run starts the IDE main loop
func (ide *UnixIDE) Run() error {
	for {
		ide.draw()

		key := ide.screen.GetChar()

		ide.handleInput(key)
	}
	return nil
}

// Close cleans up ncurses
func (ide *UnixIDE) Close() {
	goncurses.End()
}

// getSubmenuItemCount returns the number of items in the current submenu
func (ide *UnixIDE) getSubmenuItemCount() int {
	return getSubmenuItemCount(ide.menuSelected)
}
