package interpreter

import (
	"bufio"
	"os"
	"strings"
)

func Ask(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	Print(memory, stack, arguments)
	inIo := stack.IN
	if inIo == nil {
		inIo = os.Stdin
	}

	answerKey := ADDRESS_ANSWER
	if _, exists := memory[ADDRESS_LANGUAGE]; exists {
		answerKey = getTranslation(ADDRESS_ANSWER)
	}

	switch memory.GetMode() {
	case PROG_CLI:
		// Get user input from the keyboard
		// Reuse the same reader to avoid buffering issues
		if stack.Reader == nil {
			if inIo == nil {
				inIo = os.Stdin
			}
			stack.Reader = bufio.NewReader(inIo)
		}
		bufferStr, err := readLineWindows(stack.Reader)
		if err != nil {
			memory[answerKey] = VariableBox{
				VariableType: TYPE_STRING,
				String:       "",
			}
			return nil
		}

		// Split the input into arguments
		tokens := strings.Split(bufferStr, " ")
		if len(tokens) == 0 {
			return nil
		}
		if len(tokens) == 1 {
			// If there is only one token, it is the command
			arguments := stringsToArguments(memory, []string{tokens[0]})
			if len(arguments) == 0 {
				memory[answerKey] = VariableBox{
					VariableType: TYPE_STRING,
					String:       bufferStr,
				}
				return nil
			}
			memory[answerKey] = arguments[0]
			return nil
		}
		// If there are more than one token, assume it string
		memory[answerKey] = VariableBox{
			VariableType: TYPE_STRING,
			String:       bufferStr,
		}
	}
	return nil
}
