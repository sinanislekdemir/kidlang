package interpreter

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() (KLMemory, *KLStack, []VariableBox, string)
		cleanup   func(string)
		wantErr   bool
		errString string
	}{
		{
			name: "Open file successfully",
			setup: func() (KLMemory, *KLStack, []VariableBox, string) {
				tmpFile := "test_open.txt"
				memory := make(KLMemory)
				stack := &KLStack{}
				args := []VariableBox{
					{VariableType: TYPE_STRING, String: tmpFile},
					{VariableType: TYPE_FILE, String: "TESTFILE"},
				}
				return memory, stack, args, tmpFile
			},
			cleanup: func(filename string) {
				os.Remove(filename)
			},
			wantErr: false,
		},
		{
			name: "Missing filename",
			setup: func() (KLMemory, *KLStack, []VariableBox, string) {
				memory := make(KLMemory)
				stack := &KLStack{}
				args := []VariableBox{
					{VariableType: TYPE_FILE, String: "TESTFILE"},
				}
				return memory, stack, args, ""
			},
			cleanup:   func(string) {},
			wantErr:   true,
			errString: "we expect a filename argument",
		},
		{
			name: "Missing file variable",
			setup: func() (KLMemory, *KLStack, []VariableBox, string) {
				memory := make(KLMemory)
				stack := &KLStack{}
				args := []VariableBox{
					{VariableType: TYPE_STRING, String: "test.txt"},
				}
				return memory, stack, args, ""
			},
			cleanup:   func(string) {},
			wantErr:   true,
			errString: "we expect a file variable name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memory, stack, args, filename := tt.setup()
			defer tt.cleanup(filename)

			err := OpenFile(memory, stack, args)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && tt.errString != "" && err.Error() != tt.errString {
				t.Errorf("OpenFile() error = %v, want error containing %v", err.Error(), tt.errString)
			}
			if !tt.wantErr {
				// Cleanup opened file
				if val, exists := memory["TESTFILE"]; exists && val.fileHandler != nil {
					val.fileHandler.Close()
				}
			}
		})
	}
}

func TestCloseFile(t *testing.T) {
	tmpFile := "test_close.txt"
	os.WriteFile(tmpFile, []byte("test"), 0644)
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Open a file first
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	// Close the file
	closeArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err = CloseFile(memory, stack, closeArgs)
	if err != nil {
		t.Errorf("CloseFile() error = %v", err)
	}

	// Verify file handler is nil
	if memory["TESTFILE"].fileHandler != nil {
		t.Error("File handler should be nil after closing")
	}
}

func TestWriteFile(t *testing.T) {
	tmpFile := "test_write.txt"
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Open file
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer memory["TESTFILE"].fileHandler.Close()

	// Write to file
	writeArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_STRING, String: "Hello, World!"},
	}
	err = WriteFile(memory, stack, writeArgs)
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
	}

	// Verify content
	content, _ := os.ReadFile(tmpFile)
	if string(content) != "Hello, World!" {
		t.Errorf("File content = %v, want 'Hello, World!'", string(content))
	}
}

func TestReadFile(t *testing.T) {
	tmpFile := "test_read.txt"
	testContent := "Test content"
	os.WriteFile(tmpFile, []byte(testContent), 0644)
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Open file
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer memory["TESTFILE"].fileHandler.Close()

	// Read from file
	readArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_REFERENCE, String: "BOX CONTENT"},
	}
	err = ReadFile(memory, stack, readArgs)
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
	}

	// Verify content
	if memory["CONTENT"].String != testContent {
		t.Errorf("Read content = %v, want %v", memory["CONTENT"].String, testContent)
	}
}

func TestReadLine(t *testing.T) {
	tmpFile := "test_readline.txt"
	os.WriteFile(tmpFile, []byte("Line 1\nLine 2\n"), 0644)
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Open file
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer memory["TESTFILE"].fileHandler.Close()

	// Read first line
	readArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_REFERENCE, String: "BOX LINE"},
	}
	err = ReadLine(memory, stack, readArgs)
	if err != nil {
		t.Errorf("ReadLine() error = %v", err)
	}

	if memory["LINE"].String != "Line 1\n" {
		t.Errorf("First line = %v, want 'Line 1\\n'", memory["LINE"].String)
	}
}

func TestSeekLine(t *testing.T) {
	tmpFile := "test_seek.txt"
	os.WriteFile(tmpFile, []byte("Line 1\nLine 2\nLine 3\n"), 0644)
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Open file
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer memory["TESTFILE"].fileHandler.Close()

	// Seek to line 2
	seekArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_INTEGER, Integer: 2},
	}
	err = SeekLine(memory, stack, seekArgs)
	if err != nil {
		t.Errorf("SeekLine() error = %v", err)
	}

	// Read next line (should be line 3)
	readArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_REFERENCE, String: "BOX LINE"},
	}
	ReadLine(memory, stack, readArgs)
	if memory["LINE"].String != "Line 3\n" {
		t.Errorf("After seek, line = %v, want 'Line 3\\n'", memory["LINE"].String)
	}
}

func TestAutoType(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int8
	}{
		{"Integer", "42", TYPE_INTEGER},
		{"Float", "3.14", TYPE_FLOAT},
		{"String", "hello", TYPE_STRING},
		{"Integer with spaces", "  123  ", TYPE_INTEGER},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := autoType(tt.input)
			if result.VariableType != tt.want {
				t.Errorf("autoType(%v) type = %v, want %v", tt.input, result.VariableType, tt.want)
			}
		})
	}
}

func TestWriteFileWithStack(t *testing.T) {
	tmpFile := "test_write_stack.txt"
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Create a stack
	testStack := NewStack()
	testStack.SetInStack("1", VariableBox{VariableType: TYPE_STRING, String: "First"})
	testStack.SetInStack("2", VariableBox{VariableType: TYPE_STRING, String: "Second"})
	memory["MYSTACK"] = testStack

	// Open file
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer memory["TESTFILE"].fileHandler.Close()

	// Write stack to file
	writeArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_REFERENCE, String: "STACK MYSTACK"},
	}
	err = WriteFile(memory, stack, writeArgs)
	if err != nil {
		t.Errorf("WriteFile() with stack error = %v", err)
	}

	// Verify content
	content, _ := os.ReadFile(tmpFile)
	expected := "First\nSecond\n"
	if string(content) != expected {
		t.Errorf("File content = %v, want %v", string(content), expected)
	}
}

func TestReadFileToStack(t *testing.T) {
	tmpFile := "test_read_stack.txt"
	os.WriteFile(tmpFile, []byte("10\n20\n30\n"), 0644)
	defer os.Remove(tmpFile)

	memory := make(KLMemory)
	stack := &KLStack{}

	// Open file
	openArgs := []VariableBox{
		{VariableType: TYPE_STRING, String: tmpFile},
		{VariableType: TYPE_FILE, String: "TESTFILE"},
	}
	err := OpenFile(memory, stack, openArgs)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer memory["TESTFILE"].fileHandler.Close()

	// Read to stack
	readArgs := []VariableBox{
		{VariableType: TYPE_FILE, String: "TESTFILE"},
		{VariableType: TYPE_REFERENCE, String: "STACK MYSTACK"},
	}
	err = ReadFile(memory, stack, readArgs)
	if err != nil {
		t.Errorf("ReadFile() to stack error = %v", err)
	}

	// Verify stack
	if memory["MYSTACK"].VariableType != TYPE_STACK {
		t.Error("MYSTACK should be a stack")
	}
	stackBox := memory["MYSTACK"]
	val := stackBox.GetFromStack("1")
	if val.VariableType != TYPE_INTEGER || val.Integer != 10 {
		t.Errorf("Stack[1] = %v, want 10", val)
	}
}
