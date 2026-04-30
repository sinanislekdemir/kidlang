package interpreter

import (
	"fmt"
	"math"
	"math/rand"
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
	Run      func(arguments []VariableBox) (VariableBox, error)
}

var builtinFunctions = []BuiltinFunction{
	{Name: "RANDOM", Arity: 0, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_INTEGER, Integer: rand.Int63()}, nil
	}},
	{Name: "NOW", Arity: 0, Run: func(arguments []VariableBox) (VariableBox, error) {
		now := time.Now()
		return VariableBox{VariableType: TYPE_STRING, String: now.Format("Monday, January 2, 2006 15:03:05")}, nil
	}},
	{Name: "SQRT", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Sqrt(arguments[0].ToFloat())}, nil
	}},
	{Name: "ABS", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Abs(arguments[0].ToFloat())}, nil
	}},
	{Name: "SQR", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		value := arguments[0].ToFloat()
		return VariableBox{VariableType: TYPE_FLOAT, Float: value * value}, nil
	}},
	{Name: "SIN", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Sin(arguments[0].ToFloat())}, nil
	}},
	{Name: "COS", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Cos(arguments[0].ToFloat())}, nil
	}},
	{Name: "TAN", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Tan(arguments[0].ToFloat())}, nil
	}},
	{Name: "LOG", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Log(arguments[0].ToFloat())}, nil
	}},
	{Name: "ASIN", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Asin(arguments[0].ToFloat())}, nil
	}},
	{Name: "ACOS", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_FLOAT, Float: math.Acos(arguments[0].ToFloat())}, nil
	}},
	{Name: "LEN", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_INTEGER, Integer: int64(len(arguments[0].ToString()))}, nil
	}},
	{Name: "LOWER", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_STRING, String: strings.ToLower(arguments[0].ToString())}, nil
	}},
	{Name: "UPPER", Arity: 1, Run: func(arguments []VariableBox) (VariableBox, error) {
		return VariableBox{VariableType: TYPE_STRING, String: strings.ToUpper(arguments[0].ToString())}, nil
	}},
	{Name: "MIN", Arity: 2, Run: func(arguments []VariableBox) (VariableBox, error) {
		if arguments[0].ToFloat() <= arguments[1].ToFloat() {
			return arguments[0], nil
		}
		return arguments[1], nil
	}},
	{Name: "MAX", Arity: 2, Run: func(arguments []VariableBox) (VariableBox, error) {
		if arguments[0].ToFloat() >= arguments[1].ToFloat() {
			return arguments[0], nil
		}
		return arguments[1], nil
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
		return builtin.Run(resolvedArguments)
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
