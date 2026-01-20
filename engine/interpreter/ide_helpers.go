package interpreter

import (
	"os"
	"strings"
)

// aboutInfo contains the about dialog information
type aboutInfo struct {
	Title    string
	Subtitle string
	Author   string
	Email    string
	Year     string
}

// getAboutInfo returns the about information
func getAboutInfo() aboutInfo {
	t := getIDETranslation()
	return aboutInfo{
		Title:    t.AboutTitle,
		Subtitle: t.AboutSubtitle,
		Author:   t.AboutAuthor,
		Email:    "<sinan@islekdemir.com>",
		Year:     "2026",
	}
}

// keyboardShortcuts returns the list of keyboard shortcuts
func getKeyboardShortcuts() []string {
	t := getIDETranslation()
	return []string{
		"",
		t.ShortcutRun,
		t.ShortcutMenu,
		t.ShortcutEsc,
		"",
		t.ShortcutNew,
		t.ShortcutOpen,
		t.ShortcutSave,
		"",
		t.ShortcutArrows,
		t.ShortcutEnter,
		t.ShortcutBackspace,
		t.ShortcutDelete,
		"",
		t.ShortcutInMenu,
		t.ShortcutMenuLR,
		t.ShortcutMenuUD,
		t.ShortcutMenuEnter,
	}
}

// languageReferenceCommands returns the language reference commands
func getLanguageReferenceCommands() []string {
	return []string{
		"PRINT <text>           - Display text or variable",
		"INPUT <var>            - Get user input",
		"LET <var> = <value>    - Assign variable",
		"IF <cond> THEN         - Conditional statement",
		"ELSE                   - Alternative condition",
		"ENDIF                  - End if statement",
		"FOR <var>=<s> TO <e>   - Loop with counter",
		"STEP <n>               - Loop increment (optional)",
		"NEXT                   - End for loop",
		"WHILE <condition>      - While loop",
		"WEND                   - End while loop",
		"GOTO <line>            - Jump to line",
		"GOSUB <line>           - Call subroutine",
		"RETURN                 - Return from subroutine",
		"END                    - End program",
		"REM <comment>          - Comment line",
		"CLS                    - Clear screen",
	}
}

// fileListEntry represents a file or directory entry
type fileListEntry struct {
	Name  string
	IsDir bool
}

// buildFileList creates a list of files and directories for the file browser
func buildFileList(currentDir string) []fileListEntry {
	entries, err := os.ReadDir(currentDir)
	if err != nil {
		return []fileListEntry{}
	}

	var fileList []fileListEntry

	// Add parent directory
	if currentDir != "/" {
		fileList = append(fileList, fileListEntry{Name: "../", IsDir: true})
	}

	// Add directories first
	for _, entry := range entries {
		if entry.IsDir() {
			fileList = append(fileList, fileListEntry{Name: entry.Name() + "/", IsDir: true})
		}
	}

	// Add .kid files
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".kid") {
			fileList = append(fileList, fileListEntry{Name: entry.Name(), IsDir: false})
		}
	}

	return fileList
}

// navigateDirectory handles directory navigation
func navigateDirectory(currentDir string, selectedEntry fileListEntry) string {
	if selectedEntry.Name == "../" {
		// Go to parent directory
		currentDir = strings.TrimSuffix(currentDir, "/")
		lastSlash := strings.LastIndex(currentDir, "/")
		if lastSlash > 0 {
			return currentDir[:lastSlash]
		}
		return "/"
	}
	// Go to subdirectory
	return currentDir + "/" + strings.TrimSuffix(selectedEntry.Name, "/")
}

// getExamplesDir returns the path to the Examples directory
func getExamplesDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	exeDir := exePath[:strings.LastIndex(exePath, "/")]
	examplesDir := exeDir + "/Examples"

	_, err = os.Stat(examplesDir)
	if err != nil {
		return "", err
	}

	return examplesDir, nil
}

// getDocsDir returns the path to the docs directory
func getDocsDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	exeDir := exePath[:strings.LastIndex(exePath, "/")]
	docsDir := exeDir + "/docs"

	_, err = os.Stat(docsDir)
	if err != nil {
		return "", err
	}

	return docsDir, nil
}

// listSubdirectories returns a list of subdirectory names in the given directory
func listSubdirectories(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	var subdirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			subdirs = append(subdirs, entry.Name())
		}
	}

	return subdirs
}

// listMarkdownFiles returns a list of .md files in the given directory
func listMarkdownFiles(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	var mdFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			mdFiles = append(mdFiles, entry.Name())
		}
	}

	return mdFiles
}

// listKidFiles returns a list of .kid files in the given directory
func listKidFiles(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	var kidFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".kid") {
			kidFiles = append(kidFiles, entry.Name())
		}
	}

	return kidFiles
}

// scrollListSelection handles scrolling logic for list selections
type listScrollState struct {
	SelectedIndex int
	ScrollOffset  int
}

// moveUp moves the selection up in a list
func (state *listScrollState) moveUp() {
	if state.SelectedIndex > 0 {
		state.SelectedIndex--
		if state.SelectedIndex < state.ScrollOffset {
			state.ScrollOffset = state.SelectedIndex
		}
	}
}

// moveDown moves the selection down in a list
func (state *listScrollState) moveDown(listLength, listHeight int) {
	if state.SelectedIndex < listLength-1 {
		state.SelectedIndex++
		if state.SelectedIndex >= state.ScrollOffset+listHeight {
			state.ScrollOffset = state.SelectedIndex - listHeight + 1
		}
	}
}

// resetScroll resets the scroll state
func (state *listScrollState) reset() {
	state.SelectedIndex = 0
	state.ScrollOffset = 0
}

// programExecutionMessages returns standard messages for program execution
type execMessages struct {
	Header      string
	StopInfo    string
	Finished    string
	PressEnter  string
	Interrupted string
}

// getExecutionMessages returns the standard execution messages
func getExecutionMessages() execMessages {
	t := getIDETranslation()
	return execMessages{
		Header:      t.MsgRunning,
		StopInfo:    t.HelpStopProgram,
		Finished:    t.MsgFinished,
		PressEnter:  t.MsgPressEnter,
		Interrupted: t.MsgInterrupted,
	}
}

// commonEditorState represents common editor state operations
type editorState struct {
	Lines    []string
	CursorX  int
	CursorY  int
	ScrollY  int
	Filename string
	Modified bool
}

// newEditorState creates a new empty editor state
func newEditorState() editorState {
	return editorState{
		Lines:    []string{""},
		CursorX:  0,
		CursorY:  0,
		ScrollY:  0,
		Filename: "",
		Modified: false,
	}
}

// resetState resets the editor to an empty state
func (state *editorState) resetState() {
	state.Lines = []string{""}
	state.CursorX = 0
	state.CursorY = 0
	state.ScrollY = 0
	state.Filename = ""
	state.Modified = false
}

// getFilenameDisplay returns "Untitled" if filename is empty
func (state *editorState) getFilenameDisplay() string {
	if state.Filename == "" {
		t := getIDETranslation()
		return t.StatusUntitled
	}
	return state.Filename
}

// getModifiedIndicator returns "*" if modified, empty string otherwise
func (state *editorState) getModifiedIndicator() string {
	if state.Modified {
		return "*"
	}
	return ""
}
