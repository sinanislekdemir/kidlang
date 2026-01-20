package interpreter

// IDE represents the cross-platform IDE interface
type IDE interface {
	// Init initializes the IDE
	Init() error

	// Run starts the IDE main loop
	Run() error

	// Close cleans up and exits the IDE
	Close()

	// LoadFile loads a file into the editor
	LoadFile(filename string) error

	// SaveFile saves the current file
	SaveFile(filename string) error

	// Execute runs the current program
	Execute() error
}

// StartIDE creates and starts the platform-specific IDE
func StartIDE() error {
	// Load language preference from config
	loadIDELanguageFromConfig()

	ide := NewIDE()
	if err := ide.Init(); err != nil {
		return err
	}
	defer ide.Close()

	return ide.Run()
}
