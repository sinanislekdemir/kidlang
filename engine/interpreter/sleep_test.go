package interpreter

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	type args struct {
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Sleep with valid integer",
			args: args{
				memory:    KLMemory{},
				stack:     &KLStack{},
				arguments: []VariableBox{{VariableType: TYPE_INTEGER, Integer: 100}},
			},
			wantErr: false,
		},
		{
			name: "Test Sleep with non-integer argument",
			args: args{
				memory:    KLMemory{},
				stack:     &KLStack{},
				arguments: []VariableBox{{VariableType: TYPE_STRING, String: "100"}},
			},
			wantErr: true,
		},
		{
			name: "Test Sleep with no arguments",
			args: args{
				memory:    KLMemory{},
				stack:     &KLStack{},
				arguments: []VariableBox{},
			},
			wantErr: true,
		},
		{
			name: "Test Sleep with multiple arguments",
			args: args{
				memory: KLMemory{},
				stack:  &KLStack{},
				arguments: []VariableBox{
					{VariableType: TYPE_INTEGER, Integer: 100},
					{VariableType: TYPE_INTEGER, Integer: 200},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			err := Sleep(tt.args.memory, tt.args.stack, tt.args.arguments)
			duration := time.Since(start)

			if (err != nil) != tt.wantErr {
				t.Errorf("Sleep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && duration < 100*time.Millisecond {
				t.Errorf("Sleep() duration = %v, expected at least 100ms", duration)
			}
		})
	}
}
