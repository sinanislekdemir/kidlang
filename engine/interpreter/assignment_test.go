package interpreter

import (
	"fmt"
	"testing"
)

func TestAssign(t *testing.T) {
	type args struct {
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
	}
	tests := []struct {
		name                string
		args                args
		wantErr             bool
		expectedMemoryState KLMemory
	}{
		{
			name: "Not a reference",
			args: args{
				memory: make(KLMemory),
				stack:  &KLStack{},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 5, Float: 0, String: ""},
					{VariableType: TYPE_INTEGER, Integer: 8, Float: 0, String: ""},
				},
			},
			wantErr:             true,
			expectedMemoryState: KLMemory{},
		},
		{
			name: "Only a reference",
			args: args{
				memory: make(KLMemory),
				stack:  &KLStack{},
				arguments: []VariableBox{
					{VariableType: TYPE_REFERENCE, String: "BOX_2"},
				},
			},
			wantErr:             false,
			expectedMemoryState: KLMemory{},
		},
		//
		// Sum/Add operations
		//
		{
			name: "Single variable assignment",
			args: args{
				memory: make(KLMemory),
				stack:  &KLStack{},
				arguments: []VariableBox{
					{VariableType: TYPE_REFERENCE, String: "BOX_2"},
					{VariableType: TYPE_INTEGER, Integer: 5, Float: 0, String: ""},
				},
			},
			expectedMemoryState: KLMemory{
				"BOX_2": VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      5,
				},
			},
		},
		{
			name: "Sum two integers",
			args: args{
				memory: make(KLMemory),
				stack:  &KLStack{},
				arguments: []VariableBox{
					{VariableType: TYPE_REFERENCE, String: "BOX_2"},
					{VariableType: TYPE_INTEGER, Integer: 5, Float: 0, String: ""},
					{VariableType: TYPE_STRING, Integer: 0, Float: 0, String: "+"},
					{VariableType: TYPE_INTEGER, Integer: 8, Float: 0, String: ""},
				},
			},
			expectedMemoryState: KLMemory{
				"BOX_2": VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Assign(tt.args.memory, tt.args.stack, tt.args.arguments); (err != nil) != tt.wantErr {
				t.Errorf("Assign() error = %v, wantErr %v", err, tt.wantErr)
			}
			for key := range tt.args.memory {
				if tt.args.memory[key].VariableType != tt.expectedMemoryState[key].VariableType {
					t.Errorf("Unexpected type")
				}
				if tt.args.memory[key].Bool != tt.expectedMemoryState[key].Bool {
					t.Errorf("Unexpected boolean value")
				}
				if !almostEqual(tt.args.memory[key].Float, tt.expectedMemoryState[key].Float) {
					t.Errorf("Unexpected float value")
				}
				if tt.args.memory[key].String != tt.expectedMemoryState[key].String {
					t.Errorf("Unexpected string value")
				}
			}
		})
	}
}

func TestAssignSpecials(t *testing.T) {
	memory := KLMemory{}
	stack := &KLStack{}
	arguments := []VariableBox{
		{VariableType: TYPE_REFERENCE, String: "BOX_2"},
		{VariableType: TYPE_INTEGER, Integer: 5, Float: 0, String: ""},
		{VariableType: TYPE_STRING, Integer: 0, Float: 0, String: "+"},
		{VariableType: TYPE_STRING, Integer: 0, Float: 0, String: "RANDOM"},
	}
	err := Assign(memory, stack, arguments)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if memory["BOX_2"].VariableType != TYPE_INTEGER {
		t.Errorf("Unexpected type")
	}
	if memory["BOX_2"].Integer <= 5 {
		t.Errorf("Unexpected value")
	}
}

func TestAssignError(t *testing.T) {
	tests := []struct {
		name      string
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
		err       error
	}{
		{
			name:   "Divisino by zero",
			memory: KLMemory{},
			stack:  &KLStack{},
			arguments: []VariableBox{
				{VariableType: TYPE_REFERENCE, String: "BOX_2"},
				{VariableType: TYPE_FLOAT, Integer: 0, Float: 10, String: ""},
				{VariableType: TYPE_STRING, Integer: 0, Float: 0, String: "/"},
				{VariableType: TYPE_FLOAT, Integer: 0, Float: 0, String: ""},
			},
			err: fmt.Errorf("division by zero"),
		},
		{
			name:   "Not a number",
			memory: KLMemory{},
			stack:  &KLStack{},
			arguments: []VariableBox{
				{VariableType: TYPE_REFERENCE, String: "BOX_2"},
				{VariableType: TYPE_FLOAT, Integer: 0, Float: 10, String: ""},
				{VariableType: TYPE_STRING, Integer: 0, Float: 0, String: "/"},
				{VariableType: TYPE_STRING, Integer: 0, Float: 0, String: "dummy"},
			},
			err: fmt.Errorf("NaN"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Assign(tt.memory, tt.stack, tt.arguments)
			if err == nil {
				t.Errorf("expected error, got nil")
			} else if err.Error() != tt.err.Error() {
				t.Errorf("expected error %q, got %q", tt.err.Error(), err.Error())
			}
		})
	}
}

func TestAssignToStackIndex(t *testing.T) {
	tests := []struct {
		name      string
		memory    KLMemory
		arguments []VariableBox
		checkKey  string
		wantValue VariableBox
	}{
		{
			name:   "Assign to new stack index",
			memory: make(KLMemory),
			arguments: []VariableBox{
				{VariableType: TYPE_REFERENCE, String: "MYSTACK[1]"},
				{VariableType: TYPE_INTEGER, Integer: 42},
			},
			checkKey: "MYSTACK",
			wantValue: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
		},
		{
			name: "Assign to existing stack",
			memory: KLMemory{
				"MYSTACK": NewStack(),
			},
			arguments: []VariableBox{
				{VariableType: TYPE_REFERENCE, String: "MYSTACK[2]"},
				{VariableType: TYPE_STRING, String: "test"},
			},
			checkKey: "MYSTACK",
			wantValue: VariableBox{
				VariableType: TYPE_STRING,
				String:       "test",
			},
		},
		{
			name:   "Assign to stack with STACK keyword",
			memory: make(KLMemory),
			arguments: []VariableBox{
				{VariableType: TYPE_REFERENCE, String: "STACK MYSTACK[1]"},
				{VariableType: TYPE_INTEGER, Integer: 100},
			},
			checkKey: "MYSTACK",
			wantValue: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &KLStack{}
			err := Assign(tt.memory, stack, tt.arguments)
			if err != nil {
				t.Errorf("Assign() unexpected error = %v", err)
			}

			// Check that stack was created/updated
			if tt.memory[tt.checkKey].VariableType != TYPE_STACK {
				t.Errorf("Expected TYPE_STACK, got %v", tt.memory[tt.checkKey].VariableType)
			}

			// Extract the index from the arguments
			stackBox := tt.memory[tt.checkKey]
			// Check value exists somewhere in the stack
			found := false
			for _, val := range stackBox.StackData {
				if val.VariableType == tt.wantValue.VariableType {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Stack does not contain value of type %v", tt.wantValue.VariableType)
			}
		})
	}
}
