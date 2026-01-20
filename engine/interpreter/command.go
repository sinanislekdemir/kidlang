package interpreter

// Command function definition
// Pass the program memory, stack, function arguments and program mode
type CommandFunc func(memory KLMemory, stack *KLStack, arguments []VariableBox) error

type KLCommand struct {
	Command         string
	CommandFunction CommandFunc
	Type            StatementType
}
