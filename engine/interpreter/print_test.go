package interpreter

import (
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	type args struct {
		memory    KLMemory
		stack     *KLStack
		arguments []VariableBox
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		output  string
	}{
		{
			name: "Prints hello world",
			args: args{
				memory: make(KLMemory),
				stack:  &KLStack{},
				arguments: []VariableBox{
					{
						VariableType: TYPE_STRING,
						String:       "hello",
					},
					{
						VariableType: TYPE_STRING,
						String:       "world",
					},
				},
			},
			wantErr: false,
			output:  "hello world\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.CreateTemp("", "")
			if err != nil {
				t.Error(err)
			}

			tt.args.stack.OUT = f

			if err := Print(tt.args.memory, tt.args.stack, tt.args.arguments); (err != nil) != tt.wantErr {
				t.Errorf("Print() error = %v, wantErr %v", err, tt.wantErr)
			}
			f.Seek(0, io.SeekStart)
			content, err := io.ReadAll(f)
			if err != nil {
				t.Error(err)
			}

			if string(content) != tt.output {
				t.Errorf("Print mismatch, expected [%v], received [%v]", tt.output, string(content))
			}
			f.Close()
		})
	}
}

func TestPrintRandom(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
	} else {
		defer f.Close()
	}
	memory := make(KLMemory)
	stack := &KLStack{
		OUT: f,
	}
	arguments := []VariableBox{
		{
			VariableType: TYPE_STRING,
			String:       "Hello",
		}, {
			VariableType: TYPE_STRING,
			String:       "RaNdoM",
		},
	}
	err = Print(memory, stack, arguments)
	if err != nil {
		t.Errorf("Unexpected error [%v]", err)
	}

	f.Seek(0, io.SeekStart)
	content, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	parts := strings.Split(string(content), " ")
	if parts[0] != "Hello" {
		t.Errorf("First variable mismatch")
	}
	randomStr := strings.TrimSpace(parts[1])
	if _, err := strconv.Atoi(randomStr); err != nil {
		t.Error(err)
	}
	t.Logf("Received [%s]", parts[1])
}
