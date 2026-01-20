package interpreter

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestPrepareArgumentsMemoryReference(t *testing.T) {
	type args struct {
		memory    KLMemory
		arguments []VariableBox
	}
	tests := []struct {
		name            string
		args            args
		resultArguments []VariableBox
		wantErr         bool
	}{
		{
			name: "Test ProcessArguments, memory reference",
			args: args{
				memory: KLMemory{
					"BOX_39": VariableBox{
						VariableType: TYPE_INTEGER,
						Integer:      42,
					},
				},
				arguments: []VariableBox{
					{
						VariableType: TYPE_REFERENCE,
						String:       "BOX_39",
					},
				},
			},
			resultArguments: []VariableBox{
				{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
		},
		{
			name: "Test ProcessArguments, memory reference with multiple arguments",
			args: args{
				memory: KLMemory{
					"BOX_39": VariableBox{
						VariableType: TYPE_INTEGER,
						Integer:      42,
					},
					"BOX_40": VariableBox{
						VariableType: TYPE_FLOAT,
						Float:        6.06,
					},
				},
				arguments: []VariableBox{
					{
						VariableType: TYPE_REFERENCE,
						String:       "BOX_39",
					},
					{
						VariableType: TYPE_REFERENCE,
						String:       "BOX_40",
					},
				},
			},
			resultArguments: []VariableBox{
				{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        6.06,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, err := prepareArguments(tt.args.memory, tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(args, tt.resultArguments) {
				t.Errorf("ProcessArguments() = %v, want %v", tt.args.arguments, tt.resultArguments)
			}
		})
	}
}

func TestPrepareArgumentsRandom(t *testing.T) {
	memory := KLMemory{}
	arguments := []VariableBox{
		{
			VariableType: TYPE_STRING,
			String:       "RANDOM",
		},
	}
	args, err := prepareArguments(memory, arguments)
	if err != nil {
		t.Errorf("ProcessArguments() error = %v", err)
	}
	if args[0].VariableType != TYPE_INTEGER {
		t.Errorf("ProcessArguments() = %v, want %v", args[0].VariableType, TYPE_INTEGER)
	}
	args2, err := prepareArguments(memory, arguments)
	if err != nil {
		t.Errorf("ProcessArguments() error = %v", err)
	}
	if args[0].Integer == args2[0].Integer {
		t.Errorf("ProcessArguments() = %v, want different", args[0].Integer)
	}
}

func TestPrepareArgumentsNow(t *testing.T) {
	memory := KLMemory{}
	arguments := []VariableBox{
		{
			VariableType: TYPE_STRING,
			String:       "NOW",
		},
	}
	args, err := prepareArguments(memory, arguments)
	if err != nil {
		t.Errorf("ProcessArguments() error = %v", err)
	}
	if args[0].VariableType != TYPE_STRING {
		t.Errorf("ProcessArguments() = %v, want %v", args[0].VariableType, TYPE_INTEGER)
	}
	// is variable a valid date?
	_, err = time.Parse("Monday, January 2, 2006 15:03:05", args[0].String)
	if err != nil {
		t.Errorf("ProcessArguments() = %v, want valid date", args[0].String)
	}
}

func TestPrepareArgumentsAnswer(t *testing.T) {
	memory := KLMemory{
		ADDRESS_ANSWER: VariableBox{
			VariableType: TYPE_STRING,
			String:       "42",
		},
		"BOX_39": VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      13,
		},
	}
	arguments := []VariableBox{
		{
			VariableType: TYPE_STRING,
			String:       "ANSWER",
		},
	}
	args, err := prepareArguments(memory, arguments)
	if err != nil {
		t.Errorf("ProcessArguments() error = %v", err)
	}
	if args[0].VariableType != TYPE_STRING {
		t.Errorf("ProcessArguments() = %v, want %v", args[0].VariableType, TYPE_INTEGER)
	}
	if args[0].String != "42" {
		t.Errorf("ProcessArguments() = %v, want %v", args[0].String, "42")
	}
}

func TestStringsToArguments(t *testing.T) {
	type args struct {
		args    []string
		boxname string
	}
	tests := []struct {
		name   string
		args   args
		memory KLMemory
		want   []VariableBox
	}{
		{
			name: "Test GetArguments",
			args: args{
				args:    []string{"42", "13.13", "hello", "true", "BOX_5"},
				boxname: "BOX",
			},
			memory: KLMemory{
				"BOX_5": VariableBox{
					VariableType: TYPE_REFERENCE,
					String:       "BOX_5",
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
				{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
				{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
				{
					VariableType: TYPE_REFERENCE,
					String:       "BOX_5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsToArguments(tt.memory, tt.args.args)
			for i, arg := range got {
				eq, err := arg.EqualTo(tt.want[i])
				if err != nil {
					t.Errorf("GetArguments() error = %v", err)
				}
				if !eq.Bool {
					t.Errorf("GetArguments() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
func TestEvaluateInlineFunctions(t *testing.T) {
	tests := []struct {
		name      string
		arguments []VariableBox
		want      []VariableBox
	}{
		{
			name: "Test RANDOM",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       getTranslation("RANDOM"),
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_INTEGER,
				},
			},
		},
		{
			name: "Test NOW",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       getTranslation("NOW"),
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_STRING,
				},
			},
		},
		{
			name: "Test SQRT",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       getTranslation("SQRT"),
				},
				{
					VariableType: TYPE_INTEGER,
					Integer:      16,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        4,
				},
			},
		},
		{
			name: "Test ABS",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       getTranslation("ABS"),
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        -42.42,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
		},
		{
			name: "Test SQR",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       getTranslation("SQR"),
				},
				{
					VariableType: TYPE_INTEGER,
					Integer:      4,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        16,
				},
			},
		},
		{
			name: "Test SIN",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "SIN",
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        math.Pi / 2,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        1,
				},
			},
		},
		{
			name: "Test COS",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "COS",
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        0,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        1,
				},
			},
		},
		{
			name: "Test TAN",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "TAN",
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        math.Pi / 4,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        1,
				},
			},
		},
		{
			name: "Test LOG",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "LOG",
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        math.E,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        1,
				},
			},
		},
		{
			name: "Test ASIN",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "ASIN",
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        1,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        math.Pi / 2,
				},
			},
		},
		{
			name: "Test ACOS",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "ACOS",
				},
				{
					VariableType: TYPE_FLOAT,
					Float:        1,
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_FLOAT,
					Float:        0,
				},
			},
		},
		{
			name: "Test Regular String",
			arguments: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want: []VariableBox{
				{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evaluateInlineFunctions(tt.arguments)
			if len(got) != len(tt.want) {
				t.Errorf("evaluateInlineFunctions() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i].VariableType != tt.want[i].VariableType {
					t.Errorf("evaluateInlineFunctions() = %v, want %v", got[i].VariableType, tt.want[i].VariableType)
				}
				if tt.want[i].VariableType == TYPE_INTEGER && got[i].Integer == 0 {
					t.Errorf("evaluateInlineFunctions() = %v, want non-zero integer", got[i].Integer)
				}
				if tt.want[i].VariableType == TYPE_STRING && tt.want[i].String == "" {
					if _, err := time.Parse("Monday, January 2, 2006 15:03:05", got[i].String); err != nil {
						t.Errorf("evaluateInlineFunctions() = %v, want valid date", got[i].String)
					}
				}
				if tt.want[i].VariableType == TYPE_FLOAT && !almostEqual(got[i].Float, tt.want[i].Float) {
					t.Errorf("evaluateInlineFunctions() = %v, want %v", got[i].Float, tt.want[i].Float)
				}
			}
		})
	}
}
