package interpreter

import (
	"testing"
)

func TestIf(t *testing.T) {
	type args struct {
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		exitScope bool
	}{
		// Integer - Integer
		{
			name: "Test If two integers are equal",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If two integers are equal, no match",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 2},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If two integers are NOT equal",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If two integers are NOT equal, no match",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 2},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If an integer is bigger than the other",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 3},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If an integer is bigger than the other, no match",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If an integer is smaller than the other",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If an integer is smaller than the other, no match",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 4},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		// Integer - Float
		{
			name: "Test If an integer is equal to a float",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If an integer is equal to a float, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If an integer is NOT equal to a float",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.1},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If an integer is NOT equal to a float, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If an integer is bigger than a float",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 2},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If an integer is bigger than a float, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If an integer is smaller than a float",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 1},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_FLOAT, Float: 2.0},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If an integer is smaller than a float, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 2},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		// Float - Float
		{
			name: "Test If two floats are equal",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If two floats are equal, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If two floats are NOT equal",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.1},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If two floats are NOT equal, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is bigger than the other",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 2.0},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If a float is bigger than the other, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_FLOAT, Float: 2.0},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is smaller than the other",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_FLOAT, Float: 2.0},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If a float is smaller than the other, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 2.0},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_FLOAT, Float: 1.0},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		// Float - Integer
		{
			name: "Test If a float is equal to an integer",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If a float is equal to an integer, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is NOT equal to an integer",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If a float is NOT equal to an integer, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is bigger than an integer",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 2.0},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If a float is bigger than an integer, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is smaller than an integer",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 1.0},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If a float is smaller than an integer, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 2.0},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is smaller than an integer, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 2.0},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If a float is smaller than an integer, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_FLOAT, Float: 2.0},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_INTEGER, Integer: 1},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		// String - String
		{
			name: "Test If two strings are equal",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "hello"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_STRING, String: "hello"},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If two strings are equal, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "hello"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_STRING, String: "world"},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test If two strings are NOT equal",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "hello"},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_STRING, String: "world"},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test If two strings are NOT equal, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "hello"},
					{VariableType: TYPE_STRING, String: "!"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_STRING, String: "hello"},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test if two strings are equal, case insensitive",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "hello"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_STRING, String: "HELLO"},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test if first string is alphabetically smaller than the second",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "apple"},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_STRING, String: "banana"},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test if first string is alphabetically smaller than the second, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "banana"},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_STRING, String: "apple"},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		{
			name: "Test if first string is alphabetically bigger than the second",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "banana"},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_STRING, String: "apple"},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test if first string is alphabetically bigger than the second, fail",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_STRING, String: "apple"},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_STRING, String: "banana"},
				},
			},
			wantErr:   false,
			exitScope: true,
		},
		// Memory resolve test
		{
			name: "Test if function can resolve memory items (BOX es)",
			args: args{
				memory: KLMemory{
					"BOX_2": VariableBox{VariableType: TYPE_INTEGER, Integer: 1},
					"BOX_3": VariableBox{VariableType: TYPE_INTEGER, Integer: 1},
				},
				stack: &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_REFERENCE, String: "BOX_2"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_REFERENCE, String: "BOX_3"},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test if >= or => work",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 2},
					{VariableType: TYPE_STRING, String: ">"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
		{
			name: "Test if <= or =< work",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{Cursor: 5},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 2},
					{VariableType: TYPE_STRING, String: "<"},
					{VariableType: TYPE_STRING, String: "="},
					{VariableType: TYPE_INTEGER, Integer: 2},
				},
			},
			wantErr:   false,
			exitScope: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := If(tt.args.memory, tt.args.stack, tt.args.arguments); (err != nil) != tt.wantErr {
				t.Errorf("If() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.exitScope != tt.args.stack.ExitScope {
				t.Errorf("If() exitScope = %v, want %v", tt.args.stack.ExitScope, tt.exitScope)
			}
		})
	}
}
