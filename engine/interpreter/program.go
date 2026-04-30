// Package kidlang is the main program handler.
package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type StatementType int

// KLStatement is any kind of operation in a program.
// AKA Executable line
type KLStatement struct {
	Type            StatementType `json:"type"`
	CommandFunction CommandFunc   `json:"-"`
	FullLine        string        `json:"full_line"`
	LineNumber      int           `json:"line_number"`
	Arguments       []VariableBox `json:"arguments"`
}

type Program struct {
	StackTrace      []int                 `json:"stack_trace"` // Executed statement indexes
	Stack           KLStack               `json:"stack"`
	Memory          KLMemory              `json:"memory"`
	Breakpoints     []int                 `json:"breakpoints"`
	Mode            int                   `json:"program_mode"`
	Statements      []KLStatement         `json:"statements"`
	Functions       map[string]KLFunction `json:"functions"`
	Debug           bool                  `json:"debug"`
	currentFunction *KLFunction
	functionIfDepth int
}

func (p *Program) RegisterStdlib() {
	p.Statements = make([]KLStatement, 0)
	p.Functions = make(map[string]KLFunction)
}

func (p *Program) Init() {
	activeLanguage = LANG_EN
	p.StackTrace = make([]int, 0)
	p.Stack = KLStack{Cursor: 0, Error: nil, JumpLabel: nil, Program: p}
	p.Memory = make(KLMemory)
	p.Breakpoints = make([]int, 0)
	p.Mode = PROG_CLI
	p.Statements = make([]KLStatement, 0)
	p.Functions = make(map[string]KLFunction)
	p.Memory[ADDRESS_LANGUAGE] = VariableBox{
		VariableType: TYPE_STRING,
	}
	p.Memory[ADDRESS_MODE] = VariableBox{
		VariableType: TYPE_INTEGER,
		Integer:      int64(p.Mode),
	}
}

func (p *Program) SetLanguage(lang string) {
	activeLanguage = lang
}

func (p *Program) appendStatement(statement KLStatement) {
	if p.currentFunction != nil {
		p.currentFunction.Statements = append(p.currentFunction.Statements, statement)
		return
	}
	p.Statements = append(p.Statements, statement)
}

func (p *Program) parseFunctionDeclaration(tokens []string, lineNumber int) (bool, error) {
	functionKeyword := getTranslation(FUNCTION)
	if !strings.EqualFold(tokens[0], functionKeyword) {
		return false, nil
	}
	if p.currentFunction != nil {
		return false, fmt.Errorf("nested function declarations are not allowed")
	}
	if len(tokens) < 2 {
		return false, fmt.Errorf("function declaration requires a name")
	}

	signatureText := strings.Join(tokens[1:], " ")
	call, ok := parseFunctionCall(signatureText)
	if !ok {
		return false, fmt.Errorf("invalid function declaration")
	}
	if isReservedKeyword(call.Name) {
		return false, fmt.Errorf("function name %s is reserved", call.Name)
	}
	if findBuiltinFunction(call.Name) != nil {
		return false, fmt.Errorf("function name %s is reserved for a builtin", call.Name)
	}
	if _, exists := p.Functions[strings.ToUpper(call.Name)]; exists {
		return false, fmt.Errorf("function %s already exists", call.Name)
	}

	parameters := make([]FunctionParameter, 0, len(call.Arguments))
	boxKeyword := getTranslation(BOX)
	for _, rawParameter := range call.Arguments {
		parameterTokens := tokenizer(rawParameter)
		if len(parameterTokens) != 2 || !strings.EqualFold(parameterTokens[0], boxKeyword) {
			return false, fmt.Errorf("function parameters must use box declarations")
		}
		parameters = append(parameters, FunctionParameter{
			Name:       parameterTokens[1],
			BindingKey: strings.ToUpper(parameterTokens[0] + " " + parameterTokens[1]),
		})
	}

	p.currentFunction = &KLFunction{
		Name:       call.Name,
		Parameters: parameters,
		Statements: make([]KLStatement, 0),
	}
	p.functionIfDepth = 0
	_ = lineNumber
	return true, nil
}

func (p *Program) parseAssignment(tokens []string, lineNumber int) (bool, error) {
	// Check if this is an assignment expression
	// Special handling for "box varname =", "file varname =", and "stack varname ="
	varName := ""
	equalPos := -1
	isFile := false
	isStack := false
	isStackIndexed := false

	// Check for "box X = ...", "file X = ...", or "stack X = ..."
	boxName := getTranslation(BOX)
	fileName := getTranslation(FILE)
	stackName := getTranslation("STACK")

	if len(tokens) > 2 && (strings.ToUpper(tokens[0]) == boxName || strings.ToUpper(tokens[0]) == fileName || strings.ToUpper(tokens[0]) == stackName) {
		// Check for "stack toys(3) = ..." or "stack toys (3) = ..."
		if tokens[2] == "=" {
			varName = tokens[0] + " " + tokens[1]
			equalPos = 2
			isFile = strings.ToUpper(tokens[0]) == fileName
			isStack = strings.ToUpper(tokens[0]) == stackName

			// Check if tokens[1] has indexing (e.g., name[index])
			if _, _, indexed := parseStackIndex(tokens[1]); indexed {
				isStackIndexed = true
			}
		} else if len(tokens) > 3 && tokens[3] == "=" {
			// Handle "stack toys (3) = ..." where (3) is separate token
			// Check if tokens[2] looks like an index: starts with ( or [
			if (strings.HasPrefix(tokens[2], "(") || strings.HasPrefix(tokens[2], "[")) &&
				(strings.HasSuffix(tokens[2], ")") || strings.HasSuffix(tokens[2], "]")) {
				// Combine tokens[1] and tokens[2] as "toys(3)"
				varName = tokens[0] + " " + tokens[1] + tokens[2]
				equalPos = 3
				isFile = strings.ToUpper(tokens[0]) == fileName
				isStack = strings.ToUpper(tokens[0]) == stackName
				isStackIndexed = true
			}
		}
	} else if len(tokens) > 1 && tokens[1] == "=" {
		varName = tokens[0]
		equalPos = 1
	}

	if equalPos == -1 {
		return false, nil
	}

	// For FILE variables, create in memory as TYPE_FILE but don't open yet
	if isFile {
		// Extract just the variable name part (without "file " prefix)
		justName := tokens[1] // The part after "file"
		p.Memory[strings.ToUpper(justName)] = VariableBox{
			VariableType: TYPE_FILE,
			String:       justName,
		}
		// Don't create a statement, just return
		return true, nil
	}

	// For STACK variables without indexing, create empty stack in memory
	if isStack && !isStackIndexed {
		// Extract just the variable name part (without "stack " prefix)
		justName := tokens[1] // The part after "stack"
		p.Memory[strings.ToUpper(justName)] = NewStack()
		// Don't create a statement, just return
		return true, nil
	}

	// For stack with indexing OR any other assignment, create assignment statement
	arguments := make([]VariableBox, 0)
	arguments = append(arguments, VariableBox{
		VariableType: TYPE_REFERENCE,
		String:       varName,
	})

	rightArgs := stringsToArguments(p.Memory, tokens[equalPos+1:])
	arguments = append(arguments, rightArgs...)
	p.appendStatement(KLStatement{
		Type:            ST_ASSIGNMENT,
		CommandFunction: Assign,
		FullLine:        strings.Join(tokens, " "),
		LineNumber:      lineNumber,
		Arguments:       arguments,
	})
	return true, nil
}

func (p *Program) parseIfThen(tokens []string, lineNumber int) (bool, error) {
	// Check if this is an if statement
	ifKey := getTranslation(IF)
	thenKey := getTranslation(THEN)
	gotoKey := getTranslation("GOTO")

	if strings.ToUpper(tokens[0]) != ifKey {
		return false, nil
	}

	// Find THEN position
	thenPos := -1
	for i, token := range tokens {
		if strings.ToUpper(token) == thenKey {
			thenPos = i
			break
		}
	}

	if thenPos == -1 {
		return false, nil
	}

	// Extract condition tokens (between IF and THEN)
	conditionTokens := tokens[1:thenPos]
	arguments := stringsToArguments(p.Memory, conditionTokens)

	// Check for "THEN GOTO label" pattern
	if thenPos+2 < len(tokens) && strings.ToUpper(tokens[thenPos+1]) == gotoKey {
		// Add the label as last argument
		arguments = append(arguments, VariableBox{
			VariableType: TYPE_STRING,
			String:       tokens[thenPos+2],
		})
	}

	p.appendStatement(KLStatement{
		Type:            ST_CONDITION,
		CommandFunction: If,
		FullLine:        strings.Join(tokens, " "),
		LineNumber:      lineNumber,
		Arguments:       arguments,
	})
	if p.currentFunction != nil {
		p.functionIfDepth++
	}
	return true, nil
}

func (p *Program) parseLabel(tokens []string, line string, lineNumber int) bool {
	// Check if this is a label
	if len(tokens) == 1 && line[len(line)-1] == ':' {
		p.appendStatement(KLStatement{
			Type:     ST_LABEL,
			FullLine: line,
			Arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       line[:len(line)-1],
				},
			},
			LineNumber: lineNumber,
		})
		return true
	}

	return false
}

func (p *Program) parseEndOfScope(line string, lineNumber int) bool {
	endName := getTranslation(END)
	if strings.ToUpper(line) == endName {
		if p.currentFunction != nil && p.functionIfDepth == 0 {
			p.Functions[strings.ToUpper(p.currentFunction.Name)] = *p.currentFunction
			p.currentFunction = nil
			return true
		}
		p.appendStatement(KLStatement{
			Type:       ST_SCOPE_END,
			FullLine:   line,
			LineNumber: lineNumber,
			Arguments:  []VariableBox{},
		})
		if p.currentFunction != nil && p.functionIfDepth > 0 {
			p.functionIfDepth--
		}
		return true
	}
	return false
}

func (p *Program) parseCommand(tokens []string, lineNumber int) (bool, error) {
	// Check if this is a command.
	for _, cmd := range Commands {
		if cmd.Command == "" {
			continue
		}

		pattern := getTranslation(cmd.Command)
		if strings.EqualFold(tokens[0], pattern) {
			// Special handling for READ command with stack target
			if cmd.Command == "READ" && len(tokens) >= 3 {
				stackKeyword := strings.ToUpper(getTranslation("STACK"))
				if strings.ToUpper(tokens[2]) == stackKeyword && len(tokens) >= 4 {
					// This is "read <file> stack <name>" - pre-create the stack
					stackName := strings.ToUpper(tokens[3])
					if _, exists := p.Memory[stackName]; !exists {
						p.Memory[stackName] = NewStack()
					}
				}
			}

			// Special handling for WRITE command with stack source
			if cmd.Command == "WRITE" && len(tokens) >= 3 {
				stackKeyword := strings.ToUpper(getTranslation("STACK"))
				if strings.ToUpper(tokens[2]) == stackKeyword && len(tokens) >= 4 {
					// This is "write <file> stack <name>" - pre-create the stack if needed
					stackName := strings.ToUpper(tokens[3])
					if _, exists := p.Memory[stackName]; !exists {
						p.Memory[stackName] = NewStack()
					}
				}
			}

			arguments := stringsToArguments(p.Memory, tokens[1:])
			p.appendStatement(KLStatement{
				Type:            StatementType(cmd.Type),
				FullLine:        strings.Join(tokens, " "),
				CommandFunction: cmd.CommandFunction,
				LineNumber:      lineNumber,
				Arguments:       arguments,
			})
			return true, nil
		}
	}
	return false, nil
}

func (p *Program) parseReturnStatement(tokens []string, lineNumber int) (bool, error) {
	returnKeyword := getTranslation(RETURN)
	if !strings.EqualFold(tokens[0], returnKeyword) {
		return false, nil
	}
	if p.currentFunction == nil {
		return false, fmt.Errorf("return can only be used inside functions")
	}

	p.appendStatement(KLStatement{
		Type:       ST_RETURN,
		FullLine:   strings.Join(tokens, " "),
		LineNumber: lineNumber,
		Arguments:  stringsToArguments(p.Memory, tokens[1:]),
	})
	return true, nil
}

func (p *Program) parseExpressionStatement(tokens []string, line string, lineNumber int) {
	p.appendStatement(KLStatement{
		Type:       ST_EXPRESSION,
		FullLine:   line,
		LineNumber: lineNumber,
		Arguments:  stringsToArguments(p.Memory, tokens),
	})
}

func (p *Program) parseStackDeclaration(tokens []string) bool {
	stackName := getTranslation("STACK")
	if len(tokens) == 2 && strings.ToUpper(tokens[0]) == stackName {
		justName := tokens[1]
		if p.currentFunction == nil {
			p.Memory[strings.ToUpper(justName)] = NewStack()
		}
		return true
	}
	return false
}

func (p *Program) parseFunctionBodyLine(line string, tokens []string, newLine string, lineNumber int) error {
	if pass, err := p.parseReturnStatement(tokens, lineNumber); err != nil || pass {
		return err
	}
	if pass, err := p.parseIfThen(tokens, lineNumber); err != nil || pass {
		return err
	}
	if pass, err := p.parseCommand(tokens, lineNumber); err != nil || pass {
		return err
	}
	if pass, err := p.parseAssignment(tokens, lineNumber); err != nil || pass {
		return err
	}
	if pass := p.parseLabel(tokens, line, lineNumber); pass {
		return nil
	}
	if pass := p.parseEndOfScope(newLine, lineNumber); pass {
		return nil
	}

	p.parseExpressionStatement(tokens, line, lineNumber)
	return nil
}

// ParseLine to process a single programming line
// This turns the given line into a statement in the program
func (p *Program) ParseLine(line string, lineNumber int) error {
	if p.Debug {
		fmt.Printf("Parsing line: %s\n", line)
	}

	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil
	}
	if strings.HasPrefix(line, "//") {
		return nil
	}
	// Check if this is a language change
	langUpper := strings.ToUpper(line)
	if langUpper == "EN" {
		activeLanguage = LANG_EN
		return nil
	}
	if langUpper == "TR" {
		activeLanguage = LANG_TR
		return nil
	}
	if langUpper == "FI" {
		activeLanguage = LANG_FI
		return nil
	}
	if langUpper == "DE" {
		activeLanguage = LANG_DE
		return nil
	}
	tokens := tokenizer(line)
	if len(tokens) == 0 {
		return nil
	}

	// We will no longer use the line, instead we will use the cleaned newLine
	newLine := strings.Join(tokens, " ")

	if p.currentFunction == nil {
		if pass, err := p.parseFunctionDeclaration(tokens, lineNumber); err != nil || pass {
			return err
		}
	}

	if p.currentFunction != nil {
		return p.parseFunctionBodyLine(line, tokens, newLine, lineNumber)
	}

	// Check for "stack name" declaration (before other parsing)
	if p.parseStackDeclaration(tokens) {
		return nil
	}

	if pass, err := p.parseIfThen(tokens, lineNumber); err != nil || pass {
		return err
	}

	if pass, err := p.parseCommand(tokens, lineNumber); err != nil || pass {
		return err
	}

	// Check if this is a mathematical expression
	if pass, err := p.parseAssignment(tokens, lineNumber); (err != nil) || pass {
		return err
	}

	if pass := p.parseLabel(tokens, line, lineNumber); pass {
		return nil
	}

	if pass := p.parseEndOfScope(newLine, lineNumber); pass {
		return nil
	}

	// get arguments
	arguments := stringsToArguments(p.Memory, tokens)

	// No match, consider this a print statement
	p.appendStatement(KLStatement{
		Type:            ST_COMMAND,
		CommandFunction: Print,
		FullLine:        line,
		LineNumber:      lineNumber,
		Arguments:       arguments,
	})

	return nil
}

func (p *Program) Load(filename string) error {
	source, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer source.Close()
	reader := bufio.NewReader(source)
	lineNumber := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					if parseErr := p.ParseLine(line, lineNumber); parseErr != nil {
						fmt.Println("Error line: ", line)
						return parseErr
					}
				}
				break
			}
			return err
		}
		if err := p.ParseLine(line, lineNumber); err != nil {
			fmt.Println("Error line: ", line)
			return err
		}
		lineNumber++
	}
	if p.currentFunction != nil {
		return fmt.Errorf("function %s is missing END", p.currentFunction.Name)
	}
	return nil
}

// LoadFromString loads a program from a string
func (p *Program) LoadFromString(programText string) error {
	lines := strings.Split(programText, "\n")
	for lineNumber, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			if err := p.ParseLine(line, lineNumber); err != nil {
				fmt.Println("Error line: ", line)
				return err
			}
		}
	}
	if p.currentFunction != nil {
		return fmt.Errorf("function %s is missing END", p.currentFunction.Name)
	}
	return nil
}

func (p *Program) findNextEnd() error {
	nextEnd, err := findNextScopeEnd(p.Statements, p.Stack.Cursor)
	if err != nil {
		return err
	}
	p.Stack.Cursor = nextEnd
	return nil
}

func (p *Program) Run() error {
	defer p.Cleanup()
	iterations := 0
	for {
		err := p.Step()
		if p.Stack.Cursor >= len(p.Statements) {
			break
		}
		if err != nil {
			return err
		}
		iterations++
		if iterations > MAX_ITERATIONS {
			return fmt.Errorf("program exceeded maximum iterations (%d) - possible infinite loop", MAX_ITERATIONS)
		}
	}
	return nil
}

func (p *Program) Cleanup() {
	for key, val := range p.Memory {
		if val.VariableType == TYPE_FILE && val.fileHandler != nil {
			val.fileHandler.Close()
			val.fileHandler = nil
			p.Memory[key] = val
		}
	}
}

func (p *Program) Jump() error {
	labelIndex, err := findLabelIndex(p.Statements, *p.Stack.JumpLabel)
	if err != nil {
		return err
	}
	p.Stack.Cursor = labelIndex
	p.Stack.JumpLabel = nil
	return nil
}

func (p *Program) Step() error {
	if p.Stack.Cursor >= len(p.Statements) {
		return nil
	}
	statement := p.Statements[p.Stack.Cursor]
	if p.Debug {
		// Print cursor position and the statement string
		fmt.Printf("Cursor: %d, Line: %d, Statement: %s\n", p.Stack.Cursor, statement.LineNumber, statement.FullLine)
	}

	switch statement.Type {
	case ST_COMMAND, ST_CONDITION, ST_ASSIGNMENT:
		err := statement.CommandFunction(p.Memory, &p.Stack, statement.Arguments)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
	}
	if p.Stack.Error != nil {
		return fmt.Errorf("error in line %d: %s", p.Stack.Cursor, p.Stack.Error.ErrorMessage)
	}
	if p.Stack.ExitScope {
		p.Stack.ExitScope = false
		p.findNextEnd()
	}
	if p.Stack.JumpLabel != nil {
		p.Jump()
	}
	p.Stack.Cursor++

	return nil
}

func (p *Program) Stop() error {
	return nil
}
