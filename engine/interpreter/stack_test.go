package interpreter

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()

	if stack.VariableType != TYPE_STACK {
		t.Errorf("NewStack() type = %v, want TYPE_STACK", stack.VariableType)
	}

	if stack.StackData == nil {
		t.Error("NewStack() StackData should not be nil")
	}
}

func TestSetInStack(t *testing.T) {
	stack := NewStack()

	value := VariableBox{
		VariableType: TYPE_INTEGER,
		Integer:      42,
	}

	stack.SetInStack("key1", value)

	result := stack.GetFromStack("key1")
	if result.VariableType != TYPE_INTEGER || result.Integer != 42 {
		t.Errorf("SetInStack/GetFromStack() = %v, want 42", result.Integer)
	}
}

func TestGetFromStack(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() VariableBox
		key      string
		wantType int8
	}{
		{
			name: "Get existing key",
			setup: func() VariableBox {
				stack := NewStack()
				stack.SetInStack("test", VariableBox{VariableType: TYPE_STRING, String: "value"})
				return stack
			},
			key:      "test",
			wantType: TYPE_STRING,
		},
		{
			name: "Get non-existing key",
			setup: func() VariableBox {
				return NewStack()
			},
			key:      "nonexistent",
			wantType: TYPE_UNKNOWN,
		},
		{
			name: "Get from non-stack variable",
			setup: func() VariableBox {
				return VariableBox{VariableType: TYPE_INTEGER, Integer: 42}
			},
			key:      "any",
			wantType: TYPE_UNKNOWN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.setup()
			result := stack.GetFromStack(tt.key)
			if result.VariableType != tt.wantType {
				t.Errorf("GetFromStack() type = %v, want %v", result.VariableType, tt.wantType)
			}
		})
	}
}

func TestSetInStackNonStack(t *testing.T) {
	// Test that SetInStack on non-stack variable does nothing (no panic)
	v := VariableBox{VariableType: TYPE_INTEGER, Integer: 42}
	v.SetInStack("key", VariableBox{VariableType: TYPE_STRING, String: "value"})

	// Verify it didn't corrupt the variable
	if v.VariableType != TYPE_INTEGER || v.Integer != 42 {
		t.Error("SetInStack on non-stack should not modify the variable")
	}
}

func TestSetFileHandler(t *testing.T) {
	v := VariableBox{VariableType: TYPE_FILE}
	// Just test that the method exists and can be called
	v.SetFileHandler(nil)
}
