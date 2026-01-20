package interpreter

import (
	"errors"
	"fmt"
	"strings"
)

const (
	ADDRESS_LANGUAGE = "_language"
	ADDRESS_MODE     = "_mode"
	ADDRESS_ANSWER   = "ANSWER"
)

type KLMemory map[string]VariableBox

func Resolve(memory KLMemory, varname string) (*VariableBox, error) {
	if memory == nil {
		return nil, errors.New("nil memory")
	}

	// Check if this is a stack reference with indexing
	if stackVar, index, isIndexed := parseStackIndex(varname); isIndexed {
		// Extract stack name (might have "stack " prefix with translation)
		actualStackVar := stackVar
		stackKeyword := strings.ToUpper(getTranslation("STACK"))
		if strings.HasPrefix(strings.ToUpper(stackVar), stackKeyword+" ") {
			actualStackVar = strings.TrimSpace(stackVar[len(stackKeyword)+1:])
		}

		// Get the stack from memory
		if stackVal, exists := memory[strings.ToUpper(actualStackVar)]; exists && stackVal.VariableType == TYPE_STACK {
			// Resolve the index (might be a variable reference)
			// Try to resolve as "box index" first, then just "index"
			indexVal, err := Resolve(memory, "BOX "+index)
			if err != nil {
				// Try without BOX prefix
				indexVal, err = Resolve(memory, index)
			}

			var indexKey string
			if err == nil {
				indexKey = indexVal.ToString()
			} else {
				// Not a variable, use as literal
				indexKey = index
			}

			// Get value from stack
			val := stackVal.GetFromStack(indexKey)
			return &val, nil
		}
		return nil, fmt.Errorf("stack variable %s not found", actualStackVar)
	}

	if value, exists := memory[strings.ToUpper(varname)]; exists {
		switch value.VariableType {
		case TYPE_REFERENCE:
			return Resolve(memory, value.String)
		default:
			return &value, nil
		}
	}
	return nil, fmt.Errorf("variable %s not found", varname)
}

func (km KLMemory) GetMode() int {
	if mode, exists := km[ADDRESS_MODE]; exists {
		return int(mode.Integer)
	}
	return PROG_CLI
}
