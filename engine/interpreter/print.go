package interpreter

import (
	"fmt"
	"os"
	"strings"
)

func Print(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	outIo := stack.OUT
	if outIo == nil {
		outIo = os.Stdout
	}

	localArguments, err := processArguments(memory, arguments)
	if err != nil {
		return err
	}

	out := make([]string, 0, len(localArguments))
	for _, arg := range localArguments {
		out = append(out, arg.ToString())
	}
	outStr := strings.Join(out, " ")

	switch memory.GetMode() {
	case PROG_CLI:
		fmt.Fprintln(outIo, outStr)
	case PROG_WEB:
		// TODO: Implement web output
	}

	return nil
}
