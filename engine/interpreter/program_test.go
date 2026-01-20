package interpreter

import (
	"os"
	"testing"
)

func TestProgram_RegisterStdlib(t *testing.T) {
	type fields struct {
		StackTrace  []int
		Stack       KLStack
		Memory      KLMemory
		Breakpoints []int
		Mode        int
		Language    string
		Statements  []KLStatement
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				StackTrace:  tt.fields.StackTrace,
				Stack:       tt.fields.Stack,
				Memory:      tt.fields.Memory,
				Breakpoints: tt.fields.Breakpoints,
				Mode:        tt.fields.Mode,
				Statements:  tt.fields.Statements,
			}
			p.RegisterStdlib()
		})
	}
}

func TestProgram_Init(t *testing.T) {
	p := Program{}
	p.Init()
	if p.StackTrace == nil {
		t.Errorf("StackTrace is nil!")
	}
	if p.Memory == nil {
		t.Errorf("Memory is nil!")
	}
	if p.Breakpoints == nil {
		t.Errorf("Breakpoints are nil!")
	}
	if p.Statements == nil {
		t.Errorf("Statements are nil!")
	}
}

func TestProgram_parseLabel(t *testing.T) {
	type fields struct {
		StackTrace  []int
		Stack       KLStack
		Memory      KLMemory
		Breakpoints []int
		Mode        int
		Language    string
		Statements  []KLStatement
	}
	type args struct {
		tokens     []string
		line       string
		lineNumber int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Can parse a label",
			fields: fields{
				Statements: []KLStatement{},
			},
			args: args{
				tokens:     []string{"label:"},
				line:       "label:",
				lineNumber: 0,
			},
			want: true,
		},
		{
			name: "Invalid label",
			fields: fields{
				Statements: []KLStatement{},
			},
			args: args{
				tokens:     []string{"label"},
				line:       "label",
				lineNumber: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				StackTrace:  tt.fields.StackTrace,
				Stack:       tt.fields.Stack,
				Memory:      tt.fields.Memory,
				Breakpoints: tt.fields.Breakpoints,
				Mode:        tt.fields.Mode,
				Statements:  tt.fields.Statements,
			}
			if got := p.parseLabel(tt.args.tokens, tt.args.line, tt.args.lineNumber); got != tt.want {
				t.Errorf("Program.parseLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_parseEndOfScope(t *testing.T) {
	type fields struct {
		StackTrace  []int
		Stack       KLStack
		Memory      KLMemory
		Breakpoints []int
		Mode        int
		Language    string
		Statements  []KLStatement
	}
	type args struct {
		line       string
		lineNumber int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Can parse end of scope",
			fields: fields{
				Statements: []KLStatement{},
			},
			args: args{
				line:       "END",
				lineNumber: 0,
			},
			want: true,
		},
		{
			name: "Invalid end of scope",
			fields: fields{
				Statements: []KLStatement{},
			},
			args: args{
				line:       "NOT_END",
				lineNumber: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				StackTrace:  tt.fields.StackTrace,
				Stack:       tt.fields.Stack,
				Memory:      tt.fields.Memory,
				Breakpoints: tt.fields.Breakpoints,
				Mode:        tt.fields.Mode,
				Statements:  tt.fields.Statements,
			}
			if got := p.parseEndOfScope(tt.args.line, tt.args.lineNumber); got != tt.want {
				t.Errorf("Program.parseEndOfScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_Load(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		wantErr   bool
		setupFile func(string)
	}{
		{
			name:     "Can load a valid file",
			filename: "testfile.kid",
			wantErr:  false,
			setupFile: func(filename string) {
				content := `box 2 = 1 * 2
	print hello world and box 3
	if box 2 >= 5 then
	END`
				os.WriteFile(filename, []byte(content), 0644)
			},
		},
		{
			name:     "File does not exist",
			filename: "nonexistent.kid",
			wantErr:  true,
			setupFile: func(filename string) {
				// No setup needed for this test case
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFile != nil {
				tt.setupFile(tt.filename)
				defer os.Remove(tt.filename)
			}
			p := &Program{}
			p.Init()
			if err := p.Load(tt.filename); (err != nil) != tt.wantErr {
				t.Errorf("Program.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProgram_Step(t *testing.T) {
	tests := []struct {
		name           string
		statements     []KLStatement
		memory         KLMemory
		cursor         int
		expectedCursor int
		wantErr        bool
	}{
		{
			name: "Can step through a simple assignment",
			statements: []KLStatement{
				{
					Type:            ST_ASSIGNMENT,
					CommandFunction: Assign,
					FullLine:        "BOX_2 = 1",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_REFERENCE, String: "BOX_2"},
						{VariableType: TYPE_INTEGER, Integer: 1},
					},
				},
			},
			memory:         KLMemory{},
			cursor:         0,
			wantErr:        false,
			expectedCursor: 1,
		},
		{
			name: "Can step through a print statement",
			statements: []KLStatement{
				{
					Type:            ST_COMMAND,
					CommandFunction: Print,
					FullLine:        "print hello world",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_STRING, String: "hello"},
						{VariableType: TYPE_STRING, String: "world"},
					},
				},
			},
			memory:         KLMemory{},
			cursor:         0,
			wantErr:        false,
			expectedCursor: 1,
		},
		{
			name: "Can step through a condition",
			statements: []KLStatement{
				{
					Type:            ST_CONDITION,
					CommandFunction: If,
					FullLine:        "if BOX_2 >= 1 then",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_REFERENCE, String: "BOX_2"},
						{VariableType: TYPE_STRING, String: ">"},
						{VariableType: TYPE_STRING, String: "="},
						{VariableType: TYPE_INTEGER, Integer: 1},
					},
				},
				{
					Type:            ST_SCOPE_END,
					CommandFunction: nil,
					FullLine:        "END",
					LineNumber:      1,
					Arguments:       []VariableBox{},
				},
			},
			memory: KLMemory{
				"BOX_2": VariableBox{VariableType: TYPE_INTEGER, Integer: 0},
			},
			cursor:         0,
			wantErr:        false,
			expectedCursor: 2,
		},
		{
			name: "Can step through a jump",
			statements: []KLStatement{
				{
					Type:            ST_LABEL,
					CommandFunction: nil,
					FullLine:        "label:",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_STRING, String: "label"},
					},
				},
				{
					Type:            ST_COMMAND,
					CommandFunction: Goto,
					FullLine:        "goto label",
					LineNumber:      1,
					Arguments: []VariableBox{
						{VariableType: TYPE_STRING, String: "label"},
					},
				},
			},
			memory:         KLMemory{},
			cursor:         1,
			wantErr:        false,
			expectedCursor: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{}
			p.Init()
			p.Statements = tt.statements
			p.Stack.Cursor = tt.cursor
			p.Memory = tt.memory
			if err := p.Step(); (err != nil) != tt.wantErr {
				t.Errorf("Program.Step() error = %v, wantErr %v", err, tt.wantErr)
			}
			if p.Stack.Cursor != tt.expectedCursor {
				t.Errorf("Program.Step() cursor = %d, expected %d", p.Stack.Cursor, tt.expectedCursor)
			}
		})
	}
}

func TestProgram_parseIfThen(t *testing.T) {
	type fields struct {
		StackTrace  []int
		Stack       KLStack
		Memory      KLMemory
		Breakpoints []int
		Mode        int
		Language    string
		Statements  []KLStatement
	}
	type args struct {
		tokens     []string
		lineNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Can parse if-then statement",
			fields: fields{
				Statements: []KLStatement{},
			},
			args: args{
				tokens:     []string{"if", "box", "2", ">=", "5", "then"},
				lineNumber: 0,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Invalid if-then statement",
			fields: fields{
				Statements: []KLStatement{},
			},
			args: args{
				tokens:     []string{"if", "box", "2", ">=", "5"},
				lineNumber: 0,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				StackTrace:  tt.fields.StackTrace,
				Stack:       tt.fields.Stack,
				Memory:      tt.fields.Memory,
				Breakpoints: tt.fields.Breakpoints,
				Mode:        tt.fields.Mode,
				Statements:  tt.fields.Statements,
			}
			got, err := p.parseIfThen(tt.args.tokens, tt.args.lineNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Program.parseIfThen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Program.parseIfThen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_findNextEnd(t *testing.T) {
	type fields struct {
		StackTrace  []int
		Stack       KLStack
		Memory      KLMemory
		Breakpoints []int
		Mode        int
		Language    string
		Statements  []KLStatement
		Debug       bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Can find next end",
			fields: fields{
				Stack: KLStack{
					Cursor: 0,
				},
				Statements: []KLStatement{
					{
						Type:            ST_CONDITION,
						CommandFunction: If,
						FullLine:        "if BOX_2 >= 5 then",
						LineNumber:      0,
						Arguments: []VariableBox{
							{VariableType: TYPE_REFERENCE, String: "BOX_2"},
							{VariableType: TYPE_STRING, String: ">"},
							{VariableType: TYPE_STRING, String: "="},
							{VariableType: TYPE_INTEGER, Integer: 5},
						},
					},
					{
						Type:            ST_SCOPE_END,
						CommandFunction: nil,
						FullLine:        "END",
						LineNumber:      1,
						Arguments:       []VariableBox{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Cannot find next end",
			fields: fields{
				Stack: KLStack{
					Cursor: 0,
				},
				Statements: []KLStatement{
					{
						Type:            ST_CONDITION,
						CommandFunction: If,
						FullLine:        "if BOX_2 >= 5 then",
						LineNumber:      0,
						Arguments: []VariableBox{
							{VariableType: TYPE_REFERENCE, String: "BOX_2"},
							{VariableType: TYPE_STRING, String: ">"},
							{VariableType: TYPE_STRING, String: "="},
							{VariableType: TYPE_INTEGER, Integer: 5},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				StackTrace:  tt.fields.StackTrace,
				Stack:       tt.fields.Stack,
				Memory:      tt.fields.Memory,
				Breakpoints: tt.fields.Breakpoints,
				Mode:        tt.fields.Mode,
				Statements:  tt.fields.Statements,
				Debug:       tt.fields.Debug,
			}
			if err := p.findNextEnd(); (err != nil) != tt.wantErr {
				t.Errorf("Program.findNextEnd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProgram_Run(t *testing.T) {
	tests := []struct {
		name       string
		statements []KLStatement
		wantErr    bool
	}{
		{
			name: "Can run a simple program",
			statements: []KLStatement{
				{
					Type:            ST_ASSIGNMENT,
					CommandFunction: Assign,
					FullLine:        "BOX_2 = 1",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_REFERENCE, String: "BOX_2"},
						{VariableType: TYPE_INTEGER, Integer: 1},
					},
				},
				{
					Type:            ST_COMMAND,
					CommandFunction: Print,
					FullLine:        "print hello world",
					LineNumber:      1,
					Arguments: []VariableBox{
						{VariableType: TYPE_STRING, String: "hello"},
						{VariableType: TYPE_STRING, String: "world"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Can run a program with a condition",
			statements: []KLStatement{
				{
					Type:            ST_ASSIGNMENT,
					CommandFunction: Assign,
					FullLine:        "BOX_2 = 1",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_REFERENCE, String: "BOX_2"},
						{VariableType: TYPE_INTEGER, Integer: 1},
					},
				},
				{
					Type:            ST_CONDITION,
					CommandFunction: If,
					FullLine:        "if BOX_2 >= 1 then",
					LineNumber:      1,
					Arguments: []VariableBox{
						{VariableType: TYPE_REFERENCE, String: "BOX_2"},
						{VariableType: TYPE_STRING, String: ">"},
						{VariableType: TYPE_STRING, String: "="},
						{VariableType: TYPE_INTEGER, Integer: 1},
					},
				},
				{
					Type:            ST_SCOPE_END,
					CommandFunction: nil,
					FullLine:        "END",
					LineNumber:      2,
					Arguments:       []VariableBox{},
				},
			},
			wantErr: false,
		},
		{
			name: "Can run a program with a jump",
			statements: []KLStatement{
				{
					Type:            ST_LABEL,
					CommandFunction: nil,
					FullLine:        "label:",
					LineNumber:      0,
					Arguments: []VariableBox{
						{VariableType: TYPE_STRING, String: "label"},
					},
				},
				{
					Type:            ST_COMMAND,
					CommandFunction: Goto,
					FullLine:        "goto label",
					LineNumber:      1,
					Arguments: []VariableBox{
						{VariableType: TYPE_STRING, String: "label"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Statements: tt.statements,
			}
			p.Init()
			if err := p.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Program.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProgram_ParseLine(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		lineNumber int
		wantErr    bool
	}{
		{
			name:       "Parse a valid assignment",
			line:       "BOX_2 = 1",
			lineNumber: 0,
			wantErr:    false,
		},
		{
			name:       "Parse a valid if-then statement",
			line:       "if BOX_2 >= 1 then",
			lineNumber: 1,
			wantErr:    false,
		},
		{
			name:       "Parse a valid label",
			line:       "label:",
			lineNumber: 2,
			wantErr:    false,
		},
		{
			name:       "Parse a valid end of scope",
			line:       "END",
			lineNumber: 3,
			wantErr:    false,
		},
		{
			name:       "Parse a valid print command",
			line:       "print hello world",
			lineNumber: 4,
			wantErr:    false,
		},
		{
			name:       "Parse a comment line",
			line:       "// this is a comment",
			lineNumber: 7,
			wantErr:    false,
		},
		{
			name:       "Parse an empty line",
			line:       "",
			lineNumber: 8,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{}
			p.Init()
			p.Memory = KLMemory{
				"BOX_2": VariableBox{VariableType: TYPE_INTEGER, Integer: 0},
			}
			if err := p.ParseLine(tt.line, tt.lineNumber); (err != nil) != tt.wantErr {
				t.Errorf("Program.ParseLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProgram_SetLanguage(t *testing.T) {
	p := &Program{}
	p.Init()

	p.SetLanguage("tr")
	if activeLanguage != "tr" {
		t.Errorf("SetLanguage() failed to set language to 'tr'")
	}

	p.SetLanguage("en")
	if activeLanguage != "en" {
		t.Errorf("SetLanguage() failed to set language to 'en'")
	}
}

func TestProgram_Cleanup(t *testing.T) {
	tmpFile := "test_cleanup.txt"
	file, _ := os.Create(tmpFile)
	defer os.Remove(tmpFile)

	p := &Program{}
	p.Init()

	p.Memory["TESTFILE"] = VariableBox{
		VariableType: TYPE_FILE,
		fileHandler:  file,
	}

	p.Cleanup()

	if p.Memory["TESTFILE"].fileHandler != nil {
		t.Error("Cleanup() should close all file handlers")
	}
}
