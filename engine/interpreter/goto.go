package interpreter

import (
	"fmt"
)

func Goto(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	if len(arguments) != 1 {
		return fmt.Errorf("goto expects one argument")
	}
	if arguments[0].VariableType != TYPE_STRING {
		return fmt.Errorf("goto expects a string argument")
	}
	stack.JumpLabel = &arguments[0].String
	return nil
}
