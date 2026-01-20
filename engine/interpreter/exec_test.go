package interpreter

import (
	"os"
	"testing"
)

func TestExec(t *testing.T) {
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
			name: "Executes a command",
			args: args{
				memory: make(KLMemory),
				stack:  &KLStack{},
				arguments: []VariableBox{
					{
						VariableType: TYPE_STRING,
						String:       "echo",
					},
					{
						VariableType: TYPE_STRING,
						String:       "'hello world'",
					},
					{
						VariableType: TYPE_STRING,
						String:       ">",
					},
					{
						VariableType: TYPE_STRING,
						String:       "'/tmp/output.txt'",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Exec(tt.args.memory, tt.args.stack, tt.args.arguments); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}

			// check the output file
			source, err := os.ReadFile("/tmp/output.txt")
			if err != nil {
				t.Errorf("Error reading output file %v", err)
			}
			if string(source) != "hello world\n" {
				t.Errorf("Unexpected output %v", string(source))
			}
		})
	}
}
