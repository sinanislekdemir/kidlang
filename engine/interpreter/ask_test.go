package interpreter

import (
	"io"
	"os"
	"testing"
)

func TestAsk(t *testing.T) {
	type args struct {
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
	}
	tests := []struct {
		name    string
		input   string
		args    args
		wantErr bool
		answer  VariableBox
	}{
		{
			name:  "Asks for input",
			input: "hello world",
			args: args{
				memory:    make(KLMemory),
				stack:     &KLStack{},
				arguments: []VariableBox{},
			},
			wantErr: false,
			answer: VariableBox{
				VariableType: TYPE_STRING,
				String:       "hello world",
			},
		},
		{
			name:  "Asks for input, empty string",
			input: "",
			args: args{
				memory:    make(KLMemory),
				stack:     &KLStack{},
				arguments: []VariableBox{},
			},
			wantErr: false,
			answer: VariableBox{
				VariableType: TYPE_STRING,
				String:       "",
			},
		},
		{
			name:  "Asks for input, single token",
			input: "hello",
			args: args{
				memory:    make(KLMemory),
				stack:     &KLStack{},
				arguments: []VariableBox{},
			},
			wantErr: false,
			answer: VariableBox{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
		},
		{
			name:  "Asks for input, single token, integer",
			input: "123",
			args: args{
				memory:    make(KLMemory),
				stack:     &KLStack{},
				arguments: []VariableBox{},
			},
			wantErr: false,
			answer: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.CreateTemp("", "")
			if err != nil {
				t.Error(err)
			}
			tt.args.stack.IN = f
			n, err := io.WriteString(f, tt.input)
			if err != nil {
				t.Error(err)
			}
			if n != len(tt.input) {
				t.Errorf("Print() error = %v, wantErr %v", err, tt.wantErr)
			}

			f.Seek(0, io.SeekStart)

			if err := Ask(tt.args.memory, tt.args.stack, tt.args.arguments); (err != nil) != tt.wantErr {
				t.Errorf("Ask() error = %v, wantErr %v", err, tt.wantErr)
			}

			f.Close()
			if !tt.wantErr {
				got := tt.args.memory[ADDRESS_ANSWER]
				if got.VariableType != tt.answer.VariableType ||
					got.Integer != tt.answer.Integer ||
					got.Float != tt.answer.Float ||
					got.String != tt.answer.String {
					t.Errorf("Ask() = %v, want %v", got, tt.answer)
				}
			}
		})
	}
}
