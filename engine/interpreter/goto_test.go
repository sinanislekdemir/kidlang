package interpreter

import (
	"testing"
)

func TestGoto(t *testing.T) {
	type args struct {
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		errMessage   string
		expectedJump string
	}{
		{
			name: "goto expects one argument",
			args: args{
				memory:    KLMemory{},
				stack:     &KLStack{},
				arguments: []VariableBox{},
			},
			wantErr:    true,
			errMessage: "goto expects one argument",
		},
		{
			name: "goto expects a string argument",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{},
				arguments: []VariableBox{
					{
						VariableType: TYPE_INTEGER,
						Integer:      1,
					},
				},
			},
			wantErr:    true,
			errMessage: "goto expects a string argument",
		},
		{
			name: "goto sets jump label",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{},
				arguments: []VariableBox{
					{
						VariableType: TYPE_STRING,
						String:       "label",
					},
				},
			},
			wantErr:      false,
			expectedJump: "label",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Goto(tt.args.memory, tt.args.stack, tt.args.arguments)
			if tt.wantErr && err == nil {
				t.Errorf("Goto() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMessage {
				t.Errorf("Goto() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && tt.args.stack.JumpLabel != nil && *tt.args.stack.JumpLabel != tt.expectedJump {
				t.Errorf("Goto() label = %v, expected %v", *tt.args.stack.JumpLabel, tt.expectedJump)
			}
		})
	}
}
