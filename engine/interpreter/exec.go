package interpreter

import (
	"os/exec"
	"runtime"
	"strings"
)

func Exec(memory KLMemory, stack *KLStack, arguments []VariableBox) error {
	localArguments, err := processArguments(memory, arguments)
	if err != nil {
		return err
	}

	cmd := make([]string, 0, len(localArguments))
	for _, arg := range localArguments {
		cmd = append(cmd, arg.ToString())
	}
	cmdStr := strings.Join(cmd, " ")

	var command *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows
		command = exec.Command("cmd", "/c", cmdStr)
	} else {
		// Linux
		command = exec.Command("sh", "-c", cmdStr)
	}
	err = command.Run()

	return err
}
