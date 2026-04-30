package interpreter

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

type FunctionParameter struct {
	Name       string `json:"name"`
	BindingKey string `json:"binding_key"`
}

type KLFunction struct {
	Name       string              `json:"name"`
	Parameters []FunctionParameter `json:"parameters"`
	Statements []KLStatement       `json:"statements"`
}

type FunctionCall struct {
	Name      string
	Arguments []string
}

type BuiltinFunction struct {
	Name     string
	Arity    int
	Variadic bool
	Run      func(invocation BuiltinInvocation) (VariableBox, error)
}

var builtinFunctions = []BuiltinFunction{
	{Name: "RANDOM", Arity: 0, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinInteger(rand.Int63()), nil
	}},
	{Name: "NOW", Arity: 0, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		now := time.Now()
		return builtinString(now.Format("Monday, January 2, 2006 15:03:05")), nil
	}},
	{Name: "SQRT", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Sqrt(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "ABS", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Abs(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "SQR", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		value := invocation.Arguments[0].ToFloat()
		return builtinFloat(value * value), nil
	}},
	{Name: "SIN", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Sin(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "COS", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Cos(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "TAN", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Tan(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "LOG", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Log(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "ASIN", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Asin(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "ACOS", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinFloat(math.Acos(invocation.Arguments[0].ToFloat())), nil
	}},
	{Name: "LEN", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinInteger(int64(len(invocation.Arguments[0].ToString()))), nil
	}},
	{Name: "LOWER", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinString(strings.ToLower(invocation.Arguments[0].ToString())), nil
	}},
	{Name: "UPPER", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinString(strings.ToUpper(invocation.Arguments[0].ToString())), nil
	}},
	{Name: "TRIM", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinString(strings.TrimSpace(invocation.Arguments[0].ToString())), nil
	}},
	{Name: "CONTAINS", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinBool(strings.Contains(invocation.Arguments[0].ToString(), invocation.Arguments[1].ToString())), nil
	}},
	{Name: "STARTSWITH", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinBool(strings.HasPrefix(invocation.Arguments[0].ToString(), invocation.Arguments[1].ToString())), nil
	}},
	{Name: "ENDSWITH", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinBool(strings.HasSuffix(invocation.Arguments[0].ToString(), invocation.Arguments[1].ToString())), nil
	}},
	{Name: "REPLACE", Arity: 3, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		return builtinString(strings.ReplaceAll(invocation.Arguments[0].ToString(), invocation.Arguments[1].ToString(), invocation.Arguments[2].ToString())), nil
	}},
	{Name: "SUBSTRING", Arity: 3, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		text := invocation.Arguments[0].ToString()
		start := int(invocation.Arguments[1].ToFloat())
		length := int(invocation.Arguments[2].ToFloat())

		if start < 1 {
			return VariableBox{}, fmt.Errorf("substring start must be at least 1")
		}
		if length < 0 {
			return VariableBox{}, fmt.Errorf("substring length must not be negative")
		}
		if length == 0 || start > len(text) {
			return builtinString(""), nil
		}

		startIndex := start - 1
		endIndex := startIndex + length
		if endIndex > len(text) {
			endIndex = len(text)
		}

		return builtinString(text[startIndex:endIndex]), nil
	}},
	{Name: "INDEXOF", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		index := strings.Index(invocation.Arguments[0].ToString(), invocation.Arguments[1].ToString())
		if index < 0 {
			return builtinInteger(0), nil
		}
		return builtinInteger(int64(index + 1)), nil
	}},
	{Name: "SPLIT", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		text := invocation.Arguments[0].ToString()
		delimiter := invocation.Arguments[1].ToString()
		parts := strings.Split(text, delimiter)
		stackValue := NewStack()
		for index, part := range parts {
			stackValue.SetInStack(fmt.Sprintf("%d", index+1), builtinString(part))
		}
		return stackValue, nil
	}},
	{Name: "JOIN", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackValue, err := expectStack(invocation.Arguments[0], "join")
		if err != nil {
			return VariableBox{}, err
		}
		keys := sortedStackKeys(stackValue.StackData)
		parts := make([]string, 0, len(keys))
		for _, key := range keys {
			parts = append(parts, stackValue.StackData[key].ToString())
		}
		return builtinString(strings.Join(parts, invocation.Arguments[1].ToString())), nil
	}},
	{Name: "FILEEXISTS", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		path, err := resolveBuiltinFilePath(invocation.Arguments[0])
		if err != nil {
			return VariableBox{}, err
		}
		_, err = os.Stat(path)
		if err == nil {
			return builtinBool(true), nil
		}
		if os.IsNotExist(err) {
			return builtinBool(false), nil
		}
		return VariableBox{}, err
	}},
	{Name: "FILEREAD", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		content, err := readBuiltinFileText(invocation.Arguments[0])
		if err != nil {
			return VariableBox{}, err
		}
		return builtinString(content), nil
	}},
	{Name: "FILESIZE", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		path, err := resolveBuiltinFilePath(invocation.Arguments[0])
		if err != nil {
			return VariableBox{}, err
		}
		stat, err := os.Stat(path)
		if err != nil {
			return VariableBox{}, err
		}
		return builtinInteger(stat.Size()), nil
	}},
	{Name: "STACKHAS", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackValue, err := expectStack(invocation.Arguments[0], "stackhas")
		if err != nil {
			return VariableBox{}, err
		}
		_, exists := stackValue.StackData[invocation.Arguments[1].ToString()]
		return builtinBool(exists), nil
	}},
	{Name: "STACKGET", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackValue, err := expectStack(invocation.Arguments[0], "stackget")
		if err != nil {
			return VariableBox{}, err
		}
		return cloneVariableBox(stackValue.GetFromStack(invocation.Arguments[1].ToString())), nil
	}},
	{Name: "STACKSET", Arity: 3, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackKey, stackValue, err := stackReferenceKey(invocation.Memory, invocation.RawArguments[0], true)
		if err != nil {
			return VariableBox{}, err
		}
		valueToStore := cloneVariableBox(invocation.Arguments[2])
		stackValue.SetInStack(invocation.Arguments[1].ToString(), valueToStore)
		invocation.Memory[stackKey] = stackValue
		return cloneVariableBox(valueToStore), nil
	}},
	{Name: "STACKDELETE", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackKey, stackValue, err := stackReferenceKey(invocation.Memory, invocation.RawArguments[0], false)
		if err != nil {
			return VariableBox{}, err
		}
		lookupKey := invocation.Arguments[1].ToString()
		_, exists := stackValue.StackData[lookupKey]
		if exists {
			delete(stackValue.StackData, lookupKey)
			invocation.Memory[stackKey] = stackValue
		}
		return builtinBool(exists), nil
	}},
	{Name: "STACKKEYS", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackValue, err := expectStack(invocation.Arguments[0], "stackkeys")
		if err != nil {
			return VariableBox{}, err
		}
		keys := sortedStackKeys(stackValue.StackData)
		keyStack := NewStack()
		for index, key := range keys {
			keyStack.SetInStack(fmt.Sprintf("%d", index+1), builtinString(key))
		}
		return keyStack, nil
	}},
	{Name: "STACKLEN", Arity: 1, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		stackValue, err := expectStack(invocation.Arguments[0], "stacklen")
		if err != nil {
			return VariableBox{}, err
		}
		return builtinInteger(int64(len(stackValue.StackData))), nil
	}},
	{Name: "MIN", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		if invocation.Arguments[0].ToFloat() <= invocation.Arguments[1].ToFloat() {
			return invocation.Arguments[0], nil
		}
		return invocation.Arguments[1], nil
	}},
	{Name: "MAX", Arity: 2, Run: func(invocation BuiltinInvocation) (VariableBox, error) {
		if invocation.Arguments[0].ToFloat() >= invocation.Arguments[1].ToFloat() {
			return invocation.Arguments[0], nil
		}
		return invocation.Arguments[1], nil
	}},
}

func parseFunctionCall(text string) (*FunctionCall, bool) {
	text = strings.TrimSpace(text)
	openIndex := strings.Index(text, "(")
	if openIndex <= 0 || !strings.HasSuffix(text, ")") {
		return nil, false
	}

	name := strings.TrimSpace(text[:openIndex])
	if name == "" || strings.Contains(name, " ") {
		return nil, false
	}

	inner := strings.TrimSpace(text[openIndex+1 : len(text)-1])
	return &FunctionCall{
		Name:      name,
		Arguments: splitFunctionArguments(inner),
	}, true
}

func splitFunctionArguments(text string) []string {
	if strings.TrimSpace(text) == "" {
		return []string{}
	}

	parts := make([]string, 0)
	current := strings.Builder{}
	parensDepth := 0
	bracketDepth := 0
	inQuote := rune(0)

	for _, char := range text {
		switch char {
		case '\'', '"':
			if inQuote == 0 {
				inQuote = char
			} else if inQuote == char {
				inQuote = 0
			}
			current.WriteRune(char)
		case '(':
			if inQuote == 0 {
				parensDepth++
			}
			current.WriteRune(char)
		case ')':
			if inQuote == 0 {
				parensDepth--
			}
			current.WriteRune(char)
		case '[':
			if inQuote == 0 {
				bracketDepth++
			}
			current.WriteRune(char)
		case ']':
			if inQuote == 0 {
				bracketDepth--
			}
			current.WriteRune(char)
		case ',':
			if inQuote == 0 && parensDepth == 0 && bracketDepth == 0 {
				parts = append(parts, strings.TrimSpace(current.String()))
				current.Reset()
				continue
			}
			current.WriteRune(char)
		default:
			current.WriteRune(char)
		}
	}

	if strings.TrimSpace(current.String()) != "" {
		parts = append(parts, strings.TrimSpace(current.String()))
	}

	return parts
}

func findBuiltinFunction(name string) *BuiltinFunction {
	for index := range builtinFunctions {
		if strings.EqualFold(name, builtinFunctions[index].Name) {
			return &builtinFunctions[index]
		}
		if strings.EqualFold(name, getTranslation(builtinFunctions[index].Name)) {
			return &builtinFunctions[index]
		}
		for _, languageTranslations := range translations {
			if strings.EqualFold(name, languageTranslations[builtinFunctions[index].Name]) {
				return &builtinFunctions[index]
			}
		}
	}
	return nil
}

func cloneMemory(memory KLMemory) KLMemory {
	cloned := make(KLMemory, len(memory))
	for key, value := range memory {
		cloned[key] = cloneVariableBox(value)
	}
	return cloned
}

func cloneVariableBox(value VariableBox) VariableBox {
	cloned := value
	if value.VariableType == TYPE_STACK && value.StackData != nil {
		cloned.StackData = make(map[string]VariableBox, len(value.StackData))
		for key, item := range value.StackData {
			cloned.StackData[key] = cloneVariableBox(item)
		}
	}
	return cloned
}

func validateFunctionArity(name string, arguments []VariableBox, arity int, variadic bool) error {
	if variadic {
		if len(arguments) < arity {
			return fmt.Errorf("function %s expects at least %d arguments", name, arity)
		}
		return nil
	}
	if len(arguments) != arity {
		return fmt.Errorf("function %s expects %d arguments", name, arity)
	}
	return nil
}

func evaluateFunctionCall(memory KLMemory, stack *KLStack, rawCall string) (VariableBox, error) {
	call, ok := parseFunctionCall(rawCall)
	if !ok {
		return VariableBox{}, fmt.Errorf("invalid function call: %s", rawCall)
	}

	resolvedArguments := make([]VariableBox, 0, len(call.Arguments))
	for _, argumentExpression := range call.Arguments {
		argumentTokens := tokenizer(argumentExpression)
		argumentValues := stringsToArguments(memory, argumentTokens)
		resolvedValue, err := processArguments(memory, stack, argumentValues)
		if err != nil {
			return VariableBox{}, err
		}
		if len(resolvedValue) != 1 {
			return VariableBox{}, fmt.Errorf("function argument must resolve to a single value")
		}
		resolvedArguments = append(resolvedArguments, resolvedValue[0])
	}

	if builtin := findBuiltinFunction(call.Name); builtin != nil {
		if err := validateFunctionArity(call.Name, resolvedArguments, builtin.Arity, builtin.Variadic); err != nil {
			return VariableBox{}, err
		}
		return builtin.Run(BuiltinInvocation{
			Memory:       memory,
			Stack:        stack,
			RawArguments: call.Arguments,
			Arguments:    resolvedArguments,
		})
	}

	if stack == nil || stack.Program == nil {
		return VariableBox{}, fmt.Errorf("user-defined function %s is not available in this context", call.Name)
	}

	return stack.Program.executeUserFunction(memory, stack, call.Name, resolvedArguments)
}

func findLabelIndex(statements []KLStatement, label string) (int, error) {
	for index, statement := range statements {
		if statement.Type == ST_LABEL && strings.EqualFold(statement.Arguments[0].String, label) {
			return index, nil
		}
	}
	return 0, fmt.Errorf("label not found: %s", label)
}

func findNextScopeEnd(statements []KLStatement, cursor int) (int, error) {
	depth := 0
	for index := cursor + 1; index < len(statements); index++ {
		if statements[index].Type == ST_CONDITION {
			depth++
		}
		if statements[index].Type == ST_SCOPE_END {
			if depth == 0 {
				return index, nil
			}
			depth--
		}
	}
	return 0, fmt.Errorf("no matching END found for conditional statement")
}

func (p *Program) executeUserFunction(memory KLMemory, stack *KLStack, name string, arguments []VariableBox) (VariableBox, error) {
	function, exists := p.Functions[strings.ToUpper(name)]
	if !exists {
		return VariableBox{}, fmt.Errorf("function %s not found", name)
	}
	if len(arguments) != len(function.Parameters) {
		return VariableBox{}, fmt.Errorf("function %s expects %d arguments", name, len(function.Parameters))
	}

	localMemory := cloneMemory(memory)
	for index, parameter := range function.Parameters {
		localMemory[parameter.BindingKey] = cloneVariableBox(arguments[index])
	}

	functionStack := &KLStack{
		Program: p,
		IN:      stack.IN,
		OUT:     stack.OUT,
		ERR:     stack.ERR,
		Reader:  stack.Reader,
	}

	return p.runFunctionStatements(localMemory, functionStack, function.Statements)
}

func (p *Program) runFunctionStatements(memory KLMemory, stack *KLStack, statements []KLStatement) (VariableBox, error) {
	lastValue := VariableBox{VariableType: TYPE_UNKNOWN}
	iterations := 0

	for stack.Cursor < len(statements) {
		statement := statements[stack.Cursor]

		switch statement.Type {
		case ST_COMMAND, ST_CONDITION, ST_ASSIGNMENT:
			if err := statement.CommandFunction(memory, stack, statement.Arguments); err != nil {
				return VariableBox{}, err
			}
			if statement.Type == ST_ASSIGNMENT && len(statement.Arguments) > 0 {
				assignedValue, err := Resolve(memory, statement.Arguments[0].String)
				if err == nil {
					lastValue = cloneVariableBox(*assignedValue)
				}
			}
		case ST_EXPRESSION:
			values, err := processArguments(memory, stack, statement.Arguments)
			if err != nil {
				return VariableBox{}, err
			}
			if len(values) == 1 {
				lastValue = cloneVariableBox(values[0])
			}
		case ST_RETURN:
			if len(statement.Arguments) == 0 {
				return VariableBox{VariableType: TYPE_UNKNOWN}, nil
			}
			values, err := processArguments(memory, stack, statement.Arguments)
			if err != nil {
				return VariableBox{}, err
			}
			if len(values) != 1 {
				return VariableBox{}, fmt.Errorf("return statement must resolve to a single value")
			}
			return cloneVariableBox(values[0]), nil
		}

		if stack.Error != nil {
			return VariableBox{}, fmt.Errorf("error in function line %d: %s", statement.LineNumber, stack.Error.ErrorMessage)
		}
		if stack.ExitScope {
			stack.ExitScope = false
			nextEnd, err := findNextScopeEnd(statements, stack.Cursor)
			if err != nil {
				return VariableBox{}, err
			}
			stack.Cursor = nextEnd
		}
		if stack.JumpLabel != nil {
			labelIndex, err := findLabelIndex(statements, *stack.JumpLabel)
			if err != nil {
				return VariableBox{}, err
			}
			stack.Cursor = labelIndex
			stack.JumpLabel = nil
		}

		stack.Cursor++
		iterations++
		if iterations > MAX_ITERATIONS {
			return VariableBox{}, fmt.Errorf("function exceeded maximum iterations (%d)", MAX_ITERATIONS)
		}
	}

	return lastValue, nil
}
