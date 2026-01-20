package interpreter

const (
	// Program modes
	PROG_CLI    = 0
	PROG_IDE    = 1
	PROG_WEB    = 2
	PROG_TEST   = 3
	_PROG_ROBOT = 4 // Reserved

	// Security modes
	SECURITY_MODE_SANDBOX    = 0
	SECURITY_MODE_PRIVILEDGE = 1

	// Statement types
	ST_ASSIGNMENT  = 0
	ST_COMMAND     = 1
	ST_LABEL       = 2
	ST_FUNCTION    = 3
	ST_SCOPE_BEGIN = 4
	ST_SCOPE_END   = 5
	ST_CONDITION   = 6
	ST_LOOP        = 7

	BOX  = "BOX"
	END  = "END"
	IF   = "IF"
	THEN = "THEN"
	FILE = "FILE"

	// Maximum number of statements to execute before considering it an infinite loop
	MAX_ITERATIONS = 1000000
)
