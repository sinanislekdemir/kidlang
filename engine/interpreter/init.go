package interpreter

var Commands = []KLCommand{
	{Command: "ASSIGN", CommandFunction: Assign, Type: ST_ASSIGNMENT}, // Index should not be changed!
	{Command: "PRINT", CommandFunction: Print, Type: ST_COMMAND},
	{Command: "GOTO", CommandFunction: Goto, Type: ST_COMMAND},
	{Command: "IF", CommandFunction: If, Type: ST_CONDITION},
	{Command: "ASK", CommandFunction: Ask, Type: ST_COMMAND},
	{Command: "EXEC", CommandFunction: Exec, Type: ST_COMMAND},
	{Command: "CLOSE", CommandFunction: CloseFile, Type: ST_COMMAND},
	{Command: "OPEN", CommandFunction: OpenFile, Type: ST_COMMAND},
	{Command: "READ", CommandFunction: ReadFile, Type: ST_COMMAND},
	{Command: "WRITE", CommandFunction: WriteFile, Type: ST_COMMAND},
	{Command: "READLINE", CommandFunction: ReadLine, Type: ST_COMMAND},
	{Command: "SEEK", CommandFunction: SeekLine, Type: ST_COMMAND},
	{Command: "SLEEP", CommandFunction: Sleep, Type: ST_COMMAND},
}
