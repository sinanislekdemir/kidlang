package interpreter

import (
	"fmt"
	"time"
)

func Sleep(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	localArguments, err := processArguments(memory, arguments)
	if err != nil {
		return err
	}

	if len(localArguments) != 1 {
		return fmt.Errorf("we expect one argument")
	}

	if localArguments[0].VariableType != TYPE_INTEGER {
		return fmt.Errorf("we expect an integer")
	}

	time.Sleep(time.Duration(localArguments[0].Integer) * time.Millisecond)

	return nil
}
