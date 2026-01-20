//go:build windows
// +build windows

package interpreter

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32                       = windows.NewLazyDLL("kernel32.dll")
	procSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procWriteConsoleOutput         = kernel32.NewProc("WriteConsoleOutputW")
	procReadConsoleInput           = kernel32.NewProc("ReadConsoleInputW")
)

// Windows console color attributes
const (
	fgBlack   = 0
	fgBlue    = 1
	fgGreen   = 2
	fgCyan    = 3
	fgRed     = 4
	fgMagenta = 5
	fgYellow  = 6
	fgWhite   = 7
	fgGray    = 8

	bgBlack   = 0
	bgBlue    = 0x10
	bgGreen   = 0x20
	bgCyan    = 0x30
	bgRed     = 0x40
	bgMagenta = 0x50
	bgYellow  = 0x60
	bgWhite   = 0x70
)

type charInfo struct {
	Char       uint16
	Attributes uint16
}

type coord struct {
	X int16
	Y int16
}

type smallRect struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

type consoleScreenBufferInfo struct {
	Size              coord
	CursorPosition    coord
	Attributes        uint16
	Window            smallRect
	MaximumWindowSize coord
}

type inputRecord struct {
	EventType uint16
	_         [2]byte
	Event     [16]byte
}

type keyEventRecord struct {
	KeyDown         int32
	RepeatCount     uint16
	VirtualKeyCode  uint16
	VirtualScanCode uint16
	UnicodeChar     uint16
	ControlKeyState uint32
}

type mouseEventRecord struct {
	MousePosition   coord
	ButtonState     uint32
	ControlKeyState uint32
	EventFlags      uint32
}

// WindowsIDE implements the IDE interface for Windows
type WindowsIDE struct {
	stdin           windows.Handle
	stdout          windows.Handle
	originalMode    uint32
	lines           []string
	cursorX         int
	cursorY         int
	scrollY         int
	filename        string
	modified        bool
	maxY            int
	maxX            int
	menuActive      bool
	menuSelected    int
	submenuActive   bool
	submenuSelected int
}

// NewIDE creates a new platform-specific IDE instance
func NewIDE() IDE {
	return &WindowsIDE{
		lines: []string{""},
	}
}

// Init initializes the Windows IDE
func (ide *WindowsIDE) Init() error {
	var err error

	ide.stdin = windows.Handle(os.Stdin.Fd())
	ide.stdout = windows.Handle(os.Stdout.Fd())

	// Get original console mode
	err = windows.GetConsoleMode(ide.stdin, &ide.originalMode)
	if err != nil {
		return fmt.Errorf("failed to get console mode: %w", err)
	}

	// Set console mode for raw input with mouse support
	newMode := ide.originalMode &^ (windows.ENABLE_ECHO_INPUT | windows.ENABLE_LINE_INPUT | windows.ENABLE_PROCESSED_INPUT)
	newMode |= windows.ENABLE_MOUSE_INPUT    // Enable mouse
	newMode |= windows.ENABLE_EXTENDED_FLAGS // Enable extended flags

	err = windows.SetConsoleMode(ide.stdin, newMode)
	if err != nil {
		return fmt.Errorf("failed to set console mode: %w", err)
	}

	// Get console size
	var csbi consoleScreenBufferInfo
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(uintptr(ide.stdout), uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		ide.maxX = 80
		ide.maxY = 24
	} else {
		ide.maxX = int(csbi.Window.Right - csbi.Window.Left + 1)
		ide.maxY = int(csbi.Window.Bottom - csbi.Window.Top + 1)
	}

	// Set cursor position to top-left to avoid scroll issues
	coord := coord{X: 0, Y: 0}
	procSetConsoleCursorPosition.Call(uintptr(ide.stdout), uintptr(*(*int32)(unsafe.Pointer(&coord))))

	return nil
}

// Run starts the IDE main loop
func (ide *WindowsIDE) Run() error {
	// Initial draw
	ide.draw()

	for {
		record := ide.readInput()

		if record.EventType == KEY_EVENT {
			keyEvent := *(*keyEventRecord)(unsafe.Pointer(&record.Event[0]))
			if !ide.handleInput(keyEvent) {
				break
			}
		} else if record.EventType == MOUSE_EVENT {
			mouseEvent := *(*mouseEventRecord)(unsafe.Pointer(&record.Event[0]))
			ide.handleMouse(mouseEvent)
		}
		// Only redraw after input is handled - the handler will set needsRedraw if needed
	}
	return nil
}

// Close cleans up the IDE
func (ide *WindowsIDE) Close() {
	// Restore original console mode first
	windows.SetConsoleMode(ide.stdin, ide.originalMode)

	// Clear screen using cls command (simpler and more reliable)
	fmt.Print("\033[2J\033[H") // ANSI clear screen and move cursor home

	// Alternative: use Windows API to clear
	var csbi consoleScreenBufferInfo
	procGetConsoleScreenBufferInfo.Call(uintptr(ide.stdout), uintptr(unsafe.Pointer(&csbi)))

	homePos := coord{X: 0, Y: 0}
	var numWritten uint32
	consoleSize := uint32(csbi.Window.Right-csbi.Window.Left+1) * uint32(csbi.Window.Bottom-csbi.Window.Top+1)

	procFillChar := kernel32.NewProc("FillConsoleOutputCharacterW")
	procFillChar.Call(
		uintptr(ide.stdout),
		uintptr(' '),
		uintptr(consoleSize),
		*(*uintptr)(unsafe.Pointer(&homePos)),
		uintptr(unsafe.Pointer(&numWritten)),
	)

	procFillAttr := kernel32.NewProc("FillConsoleOutputAttribute")
	procFillAttr.Call(
		uintptr(ide.stdout),
		uintptr(0x07), // White on black
		uintptr(consoleSize),
		*(*uintptr)(unsafe.Pointer(&homePos)),
		uintptr(unsafe.Pointer(&numWritten)),
	)

	procSetConsoleCursorPosition.Call(uintptr(ide.stdout), *(*uintptr)(unsafe.Pointer(&homePos)))
}
