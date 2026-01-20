package interpreter

import (
	"fmt"
	"strings"
)

func Assign(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	var err error
	if !arguments[0].isAssignable() {
		return fmt.Errorf("we expect a reference or file type")
	}

	if len(arguments) == 1 {
		return nil
	}

	localArguments, err := processArguments(memory, arguments[1:])
	if err != nil {
		return err
	}

	if len(localArguments) == 0 {
		return nil
	}

	if len(localArguments) == 1 {
		varName := arguments[0].String

		// Check if this is stack indexing: "stack name[index]" or "name[index]"
		if stackVar, index, isIndexed := parseStackIndex(varName); isIndexed {
			// This is assigning to a stack index
			// Check if it starts with stack keyword (translated)
			stackKey := ""
			actualStackVar := stackVar
			stackKeyword := strings.ToUpper(getTranslation("STACK"))
			if strings.HasPrefix(strings.ToUpper(stackVar), stackKeyword+" ") {
				// Format is "stack name[index]", extract "name"
				actualStackVar = strings.TrimSpace(stackVar[len(stackKeyword)+1:]) // Remove "stack "
				stackKey = strings.ToUpper(actualStackVar)
			} else {
				// Format is "name[index]", check if it exists
				stackKey = strings.ToUpper(stackVar)
			}

			// Get or create the stack
			var stackVal VariableBox
			if val, exists := memory[stackKey]; exists && val.VariableType == TYPE_STACK {
				stackVal = val
			} else {
				// Auto-create stack if it doesn't exist
				stackVal = NewStack()
			}

			// Resolve index
			indexArgs := stringsToArguments(memory, []string{index})
			var indexKey string
			if len(indexArgs) > 0 {
				indexKey = indexArgs[0].ToString()
			} else {
				indexKey = index
			}
			// Set value in stack
			stackVal.SetInStack(indexKey, localArguments[0])
			memory[stackKey] = stackVal
			return nil
		}

		memory[strings.ToUpper(varName)] = localArguments[0]
	}

	if len(localArguments) > 1 {
		return fmt.Errorf("too many arguments, %v", localArguments)
	}

	return nil
}
