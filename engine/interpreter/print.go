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

	// Make print forgiving: strip trailing punctuation from variable references
	cleanedArguments := make([]VariableBox, 0, len(arguments))
	for _, arg := range arguments {
		if arg.VariableType == TYPE_REFERENCE {
			// Strip trailing colons, commas, periods from references
			cleaned := strings.TrimRight(arg.String, ":,.")
			suffix := arg.String[len(cleaned):]
			
			cleanedArguments = append(cleanedArguments, VariableBox{
				VariableType: TYPE_REFERENCE,
				String:       cleaned,
			})
			
			// If there was punctuation, add it as a separate string argument
			if suffix != "" {
				cleanedArguments = append(cleanedArguments, VariableBox{
					VariableType: TYPE_STRING,
					String:       suffix,
				})
			}
		} else {
			cleanedArguments = append(cleanedArguments, arg)
		}
	}

	localArguments, err := processArguments(memory, cleanedArguments)
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
