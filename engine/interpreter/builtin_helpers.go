package interpreter

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BuiltinInvocation struct {
	Memory       KLMemory
	Stack        *KLStack
	RawArguments []string
	Arguments    []VariableBox
}

func builtinInteger(value int64) VariableBox {
	return VariableBox{
		VariableType: TYPE_INTEGER,
		Integer:      value,
	}
}

func builtinFloat(value float64) VariableBox {
	return VariableBox{
		VariableType: TYPE_FLOAT,
		Float:        value,
	}
}

func builtinString(value string) VariableBox {
	return VariableBox{
		VariableType: TYPE_STRING,
		String:       value,
	}
}

func builtinBool(value bool) VariableBox {
	return VariableBox{
		VariableType: TYPE_BOOL,
		Bool:         value,
	}
}

func sortedStackKeys(items map[string]VariableBox) []string {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		leftNumber, leftErr := strconv.ParseInt(keys[i], 10, 64)
		rightNumber, rightErr := strconv.ParseInt(keys[j], 10, 64)

		if leftErr == nil && rightErr == nil {
			return leftNumber < rightNumber
		}

		return keys[i] < keys[j]
	})

	return keys
}

func expectStack(value VariableBox, functionName string) (VariableBox, error) {
	if value.VariableType != TYPE_STACK {
		return VariableBox{}, fmt.Errorf("%s expects a stack argument", functionName)
	}
	if value.StackData == nil {
		value.StackData = make(map[string]VariableBox)
	}
	return value, nil
}

func stackReferenceKey(memory KLMemory, rawArgument string, createIfMissing bool) (string, VariableBox, error) {
	stackKeyword := strings.ToUpper(getTranslation("STACK"))
	trimmedArgument := strings.TrimSpace(rawArgument)
	normalizedArgument := strings.ToUpper(trimmedArgument)

	if strings.HasPrefix(normalizedArgument, stackKeyword+" ") {
		trimmedArgument = strings.TrimSpace(trimmedArgument[len(stackKeyword)+1:])
	}

	if trimmedArgument == "" {
		return "", VariableBox{}, fmt.Errorf("stack reference is required")
	}

	stackKey := strings.ToUpper(trimmedArgument)
	stackValue, exists := memory[stackKey]
	if !exists {
		if !createIfMissing {
			return "", VariableBox{}, fmt.Errorf("stack %s not found", trimmedArgument)
		}
		stackValue = NewStack()
		memory[stackKey] = stackValue
	}

	if stackValue.VariableType != TYPE_STACK {
		return "", VariableBox{}, fmt.Errorf("%s is not a stack", trimmedArgument)
	}
	if stackValue.StackData == nil {
		stackValue.StackData = make(map[string]VariableBox)
	}

	return stackKey, stackValue, nil
}

func resolveBuiltinFilePath(value VariableBox) (string, error) {
	if value.VariableType == TYPE_FILE {
		if value.Filename == "" {
			return "", fmt.Errorf("file has no filename")
		}
		return value.Filename, nil
	}

	path := strings.TrimSpace(value.ToString())
	if path == "" {
		return "", fmt.Errorf("filename is required")
	}

	return path, nil
}

func readBuiltinFileText(value VariableBox) (string, error) {
	if value.VariableType == TYPE_FILE {
		if value.Filename == "" {
			return "", fmt.Errorf("file has no filename")
		}
		content, err := os.ReadFile(value.Filename)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}

	path, err := resolveBuiltinFilePath(value)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
