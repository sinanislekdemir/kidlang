package interpreter

import (
	"reflect"
	"testing"
)

func TestResolve(t *testing.T) {
	type args struct {
		memory  KLMemory
		varname string
	}
	tests := []struct {
		name    string
		args    args
		want    *VariableBox
		wantErr bool
	}{
		{
			name: "Test Resolve existing variable",
			args: args{
				memory: KLMemory{
					"TEST": VariableBox{
						VariableType: TYPE_INTEGER,
						Integer:      1,
					},
				},
				varname: "test",
			},
			want: &VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      1,
			},
			wantErr: false,
		},
		{
			name: "Test Resolve non-existing variable",
			args: args{
				memory: KLMemory{
					"test": VariableBox{
						VariableType: TYPE_INTEGER,
						Integer:      1,
					},
				},
				varname: "test2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Resolve reference, nil memory",
			args: args{
				memory:  nil,
				varname: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Resolve reference, pointer reference",
			args: args{
				memory: KLMemory{
					"TEST": VariableBox{
						VariableType: TYPE_REFERENCE,
						String:       "test2",
					},
					"TEST2": VariableBox{
						VariableType: TYPE_INTEGER,
						Integer:      1,
					},
				},
				varname: "test",
			},
			want: &VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Resolve(tt.args.memory, tt.args.varname)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKLMemory_GetMode(t *testing.T) {
	tests := []struct {
		name string
		km   KLMemory
		want int
	}{
		{
			name: "Test GetMode",
			km: KLMemory{
				ADDRESS_MODE: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      1,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.km.GetMode(); got != tt.want {
				t.Errorf("KLMemory.GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolve_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		memory  KLMemory
		varname string
		wantErr bool
	}{
		{
			name: "Resolve with box prefix",
			memory: KLMemory{
				"BOX TEST": VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      99,
				},
			},
			varname: "BOX TEST",
			wantErr: false,
		},
		{
			name: "Resolve stack reference",
			memory: KLMemory{
				"STACK1": VariableBox{
					VariableType: TYPE_STACK,
					StackData:    map[string]VariableBox{"1": {VariableType: TYPE_INTEGER, Integer: 10}},
				},
			},
			varname: "STACK1",
			wantErr: false,
		},
		{
			name:    "Resolve from empty memory",
			memory:  KLMemory{},
			varname: "NOTEXIST",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Resolve(tt.memory, tt.varname)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
