package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// autoType attempts to parse a string and return the appropriate typed VariableBox
func autoType(value string) VariableBox {
	trimmed := strings.TrimSpace(value)

	// Try integer first
	if intVal, err := strconv.ParseInt(trimmed, 10, 64); err == nil {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      intVal,
		}
	}

	// Try float
	if floatVal, err := strconv.ParseFloat(trimmed, 64); err == nil {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        floatVal,
		}
	}

	// Default to string
	return VariableBox{
		VariableType: TYPE_STRING,
		String:       value,
	}
}

func SeekLine(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	if len(arguments) != 2 {
		return fmt.Errorf("we expect a file argument")
	}

	fileVarName := ""
	lineNumber := 0

	localArguments, err := prepareArguments(memory, arguments)
	if err != nil {
		return err
	}

	for _, arg := range localArguments {
		if arg.VariableType == TYPE_FILE {
			fileVarName = strings.ToUpper(arg.String)
		}
		if arg.VariableType == TYPE_INTEGER {
			lineNumber = int(arg.Integer)
		}
	}

	if fileVarName == "" {
		return fmt.Errorf("we expect a file variable name")
	}

	if lineNumber == 0 {
		return fmt.Errorf("we expect a line number")
	}

	// check if file is open
	if memory[fileVarName].fileHandler == nil {
		return fmt.Errorf("file is not open")
	}

	// First, seek to beginning of file
	_, err = memory[fileVarName].fileHandler.Seek(0, 0)
	if err != nil {
		return err
	}

	// Read byte-by-byte to skip lines without buffering issues
	linesSkipped := 0
	b := make([]byte, 1)
	for linesSkipped < lineNumber {
		n, err := memory[fileVarName].fileHandler.Read(b)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		if n > 0 && b[0] == '\n' {
			linesSkipped++
		}
	}

	return nil
}

func WriteFile(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	if len(arguments) < 2 {
		return fmt.Errorf("we expect a file argument and content to write")
	}

	fileVarName := ""
	content := ""
	contentFound := false
	var stackToWrite *VariableBox = nil

	// Check if second argument is a stack reference (before prepareArguments)
	stackKeyword := strings.ToUpper(getTranslation("STACK"))
	var stackVarName string
	isStackSource := false

	if arguments[1].VariableType == TYPE_REFERENCE {
		refStr := strings.ToUpper(arguments[1].String)
		if strings.HasPrefix(refStr, stackKeyword+" ") {
			isStackSource = true
			stackVarName = strings.TrimSpace(refStr[len(stackKeyword)+1:])
		}
	} else if arguments[1].VariableType == TYPE_STACK {
		// Stack value directly passed
		isStackSource = true
		stackVarName = arguments[1].String
		if stackVarName == "" {
			// Try to find the stack in memory by comparing
			for name, val := range memory {
				if val.VariableType == TYPE_STACK {
					// This might be it, but we can't be sure
					stackVarName = name
					break
				}
			}
		}
	}

	localArguments, err := prepareArguments(memory, arguments)
	if err != nil && !isStackSource {
		return err
	}

	// Extract file variable from first argument
	if arguments[0].VariableType == TYPE_FILE {
		fileVarName = strings.ToUpper(arguments[0].String)
	} else if len(localArguments) > 0 {
		for _, arg := range localArguments {
			if arg.VariableType == TYPE_FILE {
				fileVarName = strings.ToUpper(arg.String)
				break
			}
		}
	}

	for _, arg := range localArguments {
		if !isStackSource && arg.VariableType == TYPE_STRING {
			content = arg.String
			contentFound = true
		}
	}

	// If stack source, get the actual stack from memory
	if isStackSource {
		if val, exists := memory[stackVarName]; exists && val.VariableType == TYPE_STACK {
			stackToWrite = &val
			contentFound = true
		} else {
			return fmt.Errorf("stack %s not found", stackVarName)
		}
	}

	if fileVarName == "" {
		return fmt.Errorf("we expect a file variable name")
	}

	if !contentFound {
		return fmt.Errorf("we expect a content to write")
	}

	// check if file is open
	if memory[fileVarName].fileHandler == nil {
		return fmt.Errorf("file is not open")
	}

	// If writing a stack, write each value as a line
	if stackToWrite != nil {
		writer := bufio.NewWriter(memory[fileVarName].fileHandler)

		// Collect and sort keys
		keys := make([]string, 0, len(stackToWrite.StackData))
		for key := range stackToWrite.StackData {
			keys = append(keys, key)
		}

		// Sort keys: try numeric first, fallback to string
		sort.Slice(keys, func(i, j int) bool {
			// Try to parse as integers
			numI, errI := strconv.ParseInt(keys[i], 10, 64)
			numJ, errJ := strconv.ParseInt(keys[j], 10, 64)

			// If both are numbers, compare numerically
			if errI == nil && errJ == nil {
				return numI < numJ
			}

			// Otherwise, compare as strings
			return keys[i] < keys[j]
		})

		// Write each stack value as a line in sorted order
		for _, key := range keys {
			value := stackToWrite.StackData[key]
			_, err = writer.WriteString(value.ToString() + "\n")
			if err != nil {
				return err
			}
		}

		return writer.Flush()
	}

	// write content to file
	_, err = memory[fileVarName].fileHandler.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func ReadLine(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	if len(arguments) != 2 {
		return fmt.Errorf("we expect a file argument")
	}

	fileVarName := ""
	targetVarName := ""

	// Check if it's a box reference (might not exist yet)
	boxKeyword := strings.ToUpper(getTranslation(BOX))
	if arguments[1].VariableType == TYPE_REFERENCE {
		refStr := strings.ToUpper(arguments[1].String)
		if strings.HasPrefix(refStr, boxKeyword+" ") {
			targetVarName = strings.TrimSpace(refStr[len(boxKeyword)+1:])
		} else {
			targetVarName = refStr
		}
	}

	// Try to prepare arguments (might fail if box doesn't exist, which is OK)
	localArguments, err := prepareArguments(memory, arguments)
	if err == nil {
		for _, arg := range localArguments {
			if arg.VariableType == TYPE_FILE {
				fileVarName = strings.ToUpper(arg.String)
			}
			if targetVarName == "" {
				if v, exists := memory[strings.ToUpper(arg.String)]; exists && v.VariableType != TYPE_FILE {
					targetVarName = strings.ToUpper(arg.String)
				}
			}
		}
	} else {
		// prepareArguments failed, extract file from first argument
		if arguments[0].VariableType == TYPE_FILE {
			fileVarName = strings.ToUpper(arguments[0].String)
		}
	}

	if fileVarName == "" {
		return fmt.Errorf("we expect a file variable name")
	}

	if targetVarName == "" {
		return fmt.Errorf("we expect a target variable name")
	}
	// check if file is open
	if memory[fileVarName].fileHandler == nil {
		return fmt.Errorf("file is not open")
	}
	// read file content line by line - read directly from file to avoid buffering issues
	buf := make([]byte, 0, 4096)
	for {
		b := make([]byte, 1)
		n, err := memory[fileVarName].fileHandler.Read(b)
		if n > 0 {
			buf = append(buf, b[0])
			if b[0] == '\n' {
				break
			}
		}
		if err != nil {
			if err.Error() == "EOF" && len(buf) > 0 {
				break
			}
			if err.Error() == "EOF" {
				buf = []byte{}
				break
			}
			return err
		}
	}

	// store file content in target variable
	lineBox := VariableBox{
		VariableType: TYPE_STRING,
		String:       string(buf),
	}
	memory[targetVarName] = lineBox

	// Also store with BOX prefix for compatibility
	memory[boxKeyword+" "+targetVarName] = lineBox

	return nil
}

func ReadFile(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	if len(arguments) != 2 {
		return fmt.Errorf("we expect a file argument")
	}

	fileVarName := ""
	targetVarName := ""
	isStackTarget := false

	// Check if second argument is a stack reference (before prepareArguments)
	stackKeyword := strings.ToUpper(getTranslation("STACK"))
	if arguments[1].VariableType == TYPE_REFERENCE {
		refStr := strings.ToUpper(arguments[1].String)
		if strings.HasPrefix(refStr, stackKeyword+" ") {
			isStackTarget = true
			targetVarName = strings.TrimSpace(refStr[len(stackKeyword)+1:])
		}
	}

	// For stack targets, we don't call prepareArguments since the stack may not exist yet
	// Instead, manually extract the file variable
	if isStackTarget {
		for _, arg := range arguments {
			if arg.VariableType == TYPE_FILE || arg.VariableType == TYPE_REFERENCE {
				// Check if it's a file in memory
				if val, exists := memory[strings.ToUpper(arg.String)]; exists && val.VariableType == TYPE_FILE {
					fileVarName = strings.ToUpper(arg.String)
					break
				}
			}
		}
	} else {
		// Normal path - but first check if it's a reference to a non-existent box
		// For READ operations, the target box might not exist yet
		boxKeyword := strings.ToUpper(getTranslation(BOX))
		if arguments[1].VariableType == TYPE_REFERENCE {
			refStr := strings.ToUpper(arguments[1].String)
			if strings.HasPrefix(refStr, boxKeyword+" ") {
				// This is "box varname" - extract the variable name
				targetVarName = strings.TrimSpace(refStr[len(boxKeyword)+1:])
			} else {
				// Some other reference, use as-is
				targetVarName = refStr
			}
		}

		// Try to prepare arguments (might fail if box doesn't exist, which is OK for read)
		localArguments, err := prepareArguments(memory, arguments)
		if err == nil {
			for _, arg := range localArguments {
				if arg.VariableType == TYPE_FILE {
					fileVarName = strings.ToUpper(arg.String)
				}
				if targetVarName == "" {
					if v, exists := memory[strings.ToUpper(arg.String)]; exists && v.VariableType != TYPE_FILE {
						targetVarName = strings.ToUpper(arg.String)
					}
				}
			}
		} else {
			// prepareArguments failed, extract file from first argument
			if arguments[0].VariableType == TYPE_FILE {
				fileVarName = strings.ToUpper(arguments[0].String)
			}
		}
	}

	if fileVarName == "" {
		return fmt.Errorf("we expect a file variable name")
	}

	if targetVarName == "" {
		return fmt.Errorf("we expect a target variable name")
	}

	// check if file is open
	if memory[fileVarName].fileHandler == nil {
		return fmt.Errorf("file is not open")
	}

	// If target is a stack, read line-by-line with auto-typing
	if isStackTarget {
		fileHandle := memory[fileVarName].fileHandler
		if fileHandle == nil {
			return fmt.Errorf("file not open")
		}

		fileHandle.Seek(0, 0)
		scanner := bufio.NewScanner(fileHandle)

		// Create new stack
		newStack := NewStack()
		lineNumber := int64(1)

		for scanner.Scan() {
			line := scanner.Text()
			// Auto-type each line
			typedValue := autoType(line)
			newStack.SetInStack(strconv.FormatInt(lineNumber, 10), typedValue)
			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		// Store stack in memory
		memory[targetVarName] = newStack
		return nil
	}

	// read file content
	memory[fileVarName].fileHandler.Seek(0, 0)
	stats, err := memory[fileVarName].fileHandler.Stat()
	if err != nil {
		return err
	}
	size := stats.Size()
	fileContent := make([]byte, size)
	_, err = memory[fileVarName].fileHandler.Read(fileContent)
	if err != nil {
		return err
	}
	// store file content in target variable
	fileContentBox := VariableBox{
		VariableType: TYPE_STRING,
		String:       string(fileContent),
	}
	memory[targetVarName] = fileContentBox

	// Also store with BOX prefix for compatibility with print "box varname"
	boxKeyword := strings.ToUpper(getTranslation(BOX))
	memory[boxKeyword+" "+targetVarName] = fileContentBox

	return nil
}

/**
 * @file file.go
 * Kidlang file handling functions
 */
func OpenFile(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	localArguments, err := prepareArguments(memory, arguments)
	if err != nil {
		return err
	}

	filename := ""
	fileVarName := ""
	filenameparts := []string{}

	for _, arg := range localArguments {
		if arg.VariableType == TYPE_STRING {
			filenameparts = append(filenameparts, arg.String)
		}
		if arg.VariableType == TYPE_FILE {
			fileVarName = arg.String
		}
	}

	// Concatenate all string parts to form the filename
	filename = strings.Join(filenameparts, "")

	if filename == "" {
		return fmt.Errorf("we expect a filename argument")
	}

	if fileVarName == "" {
		return fmt.Errorf("we expect a file variable name")
	}

	if memory[strings.ToUpper(fileVarName)].fileHandler != nil {
		return fmt.Errorf("file already open")
	}

	fileHandler, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	var fileStruct VariableBox
	fileStruct.VariableType = TYPE_FILE
	fileStruct.fileHandler = fileHandler
	fileStruct.String = fileVarName
	fileStruct.Filename = filename
	memory[strings.ToUpper(fileVarName)] = fileStruct

	return nil
}

func CloseFile(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	if len(arguments) != 1 {
		return fmt.Errorf("we expect a file argument")
	}

	fileKey := strings.ToUpper(arguments[0].String)
	if _, exists := memory[fileKey]; exists {
		if memory[fileKey].fileHandler != nil {
			err := memory[fileKey].fileHandler.Close()
			if err != nil {
				return err
			}
			fileStruct := memory[fileKey]
			fileStruct.fileHandler = nil
			memory[fileKey] = fileStruct
		}
	}
	return nil
}
