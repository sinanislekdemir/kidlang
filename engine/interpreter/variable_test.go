package interpreter

import (
	"reflect"
	"testing"
)

func TestVariableBox_ToString(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Integer to String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			want: "42",
		},
		{
			name: "Test Float to String",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			want: "42.42",
		},
		{
			name: "Test String to String",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "Hello, World!",
			},
			want: "Hello, World!",
		},
		{
			name: "Test Bool to String: True",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			want: "True",
		},
		{
			name: "Test Bool to String: False",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
			want: "False",
		},
		{
			name: "Test Unknown to String",
			fields: fields{
				VariableType: TYPE_UNKNOWN,
			},
			want: "",
		},
		{
			name: "Test Bool to String: False",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
			want: "False",
		},
		{
			name: "Test Reference Variable",
			fields: fields{
				VariableType: TYPE_REFERENCE,
				String:       "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			if got := v.ToString(); got != tt.want {
				t.Errorf("VariableBox.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_Sum(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VariableBox
	}{
		{
			name: "Test Integer Sum",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      84,
			},
		},
		{
			name: "Test Integer Sum with Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        84.42,
			},
		},
		{
			name: "Test Integer Sum with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "42",
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "4242",
			},
		},
		{
			name: "Test Integer Sum with Bool",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      43,
			},
		},
		{
			name: "Test Float Sum",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        84.84,
			},
		},
		{
			name: "Test Float Sum with Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        84.42,
			},
		},
		{
			name: "Test Float Sum with String",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "42.42hello",
			},
		},
		{
			name: "Test Float Sum with Bool",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
		},
		{
			name: "Test String Sum",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "world",
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "helloworld",
			},
		},
		{
			name: "Test String Sum with Integer",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "hello42",
			},
		},
		{
			name: "Test String Sum with Float",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "hello42.42",
			},
		},
		{
			name: "Test String Sum with Bool",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "helloTrue",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			if got := v.Sum(tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VariableBox.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_Sub(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		{
			name: "Test Integer Sub",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      29,
			},
		},
		{
			name: "Test Integer Sub with Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        28.87,
			},
		},
		{
			name: "Test Integer Sub with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "13",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Integer Sub with Bool",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      41,
			},
			wantErr: false,
		},
		// Float
		{
			name: "Test Float Sub",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        29.29,
			},
		},
		{
			name: "Test Float Sub with Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        29.42,
			},
		},
		{
			name: "Test Float Sub with String",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Float Sub with Bool",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        41.42,
			},
			wantErr: false,
		},
		{
			name: "Test String Sub",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "sinan islekdemir",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "islek",
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "sinan demir",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.Sub(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, _ := got.EqualTo(tt.want)
			if !eq.Bool {
				t.Errorf("VariableBox.Sub() = %v, want %v", got, tt.want)
			}
			if got.VariableType != tt.want.VariableType {
				t.Errorf("VariableBox.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_Mul(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		{
			name: "Test Integer Mul",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      546,
			},
		},
		{
			name: "Test Integer Mul with Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        551.46,
			},
		},
		{
			name: "Test Integer Mul with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "13",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Integer Mul with Bool",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      0,
			},
			wantErr: false,
		},
		// Float
		{
			name: "Test Float Mul",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        556.9746,
			},
		},
		{
			name: "Test Float Mul with Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        551.46,
			},
		},
		{
			name: "Test Float Mul with String",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Float Mul with Bool",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			wantErr: false,
		},
		{
			name: "Test String Mul",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "sinan",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "islek",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test String Mul with Integer",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "sinan",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      3,
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "sinansinansinan",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.Mul(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, _ := got.EqualTo(tt.want)
			if !eq.Bool {
				t.Errorf("VariableBox.Mul() = %v, want %v", got, tt.want)
			}
			if got.VariableType != tt.want.VariableType {
				t.Errorf("VariableBox.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_Div(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		// Integer
		{
			name: "Test Integer Div",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        3.230769230769231,
			},
		},
		{
			name: "Test Integer Div with Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        3.1987814166031985,
			},
		},
		{
			name: "Test Integer Div with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Integer Div with Bool: True",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
		},
		{
			name: "Test Integer Div with Bool: False",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Integer Div with Zero",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      0,
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Integer Div with Zero Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        0,
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		// Float
		{
			name: "Test Float Div",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        3.230769230769231,
			},
		},
		{
			name: "Test Float Div with Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        3.263076923076923,
			},
		},
		{
			name: "Test Float Div with String",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Float Div with Bool: True",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
		},
		{
			name: "Test Float Div with Bool: False",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Float Div with Zero",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        0,
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		{
			name: "Test Float Div with Zero Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      0,
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		// String
		{
			name: "Test String Div",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "sinanislekdemir",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "islek",
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "sinanislekdemir/islek",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.Div(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, _ := got.EqualTo(tt.want)
			if !eq.Bool {
				t.Errorf("VariableBox.Div() = %v, want %v", got, tt.want)
			}
			if got.VariableType != tt.want.VariableType {
				t.Errorf("VariableBox.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_Mod(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		// Integer
		{
			name: "Test Integer Mod",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      12,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      6,
			},
		},
		{
			name: "Test Integer Mod with Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        12.12,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        5.64,
			},
		},
		{
			name: "Test Integer Mod with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "sinan",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		// Float
		{
			name: "Test Float Mod",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        12.12,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        6.06,
			},
		},
		{
			name: "Test Float Mod with Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      12,
				},
			},
			want: VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        6.42,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.Mod(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.Mod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, _ := got.EqualTo(tt.want)
			if !eq.Bool {
				t.Errorf("VariableBox.Mod() = %v, want %v", got, tt.want)
			}
			if got.VariableType != tt.want.VariableType {
				t.Errorf("VariableBox.Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_GreaterThan(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		// Integer
		{
			name: "Test Integer GreaterThan: Correct",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer GreaterThan: Incorrect",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Integer GreaterThan with Float: Correct",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer GreaterThan with Float: Incorrect",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Integer GreaterThan with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "13",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			wantErr: false,
		},
		{
			name: "Test Integer GreaterThan with Bool: True",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
			wantErr: false,
		},
		// Float
		{
			name: "Test Float GreaterThan: Correct",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Float GreaterThan: Incorrect",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        13.13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Float GreaterThan with Integer: Correct",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Float GreaterThan with Integer: Incorrect",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        13.13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		// String
		{
			name: "Test String GreaterThan: Correct",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "banana",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "apple",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test String GreaterThan: Incorrect",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "apple",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "banana",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.GreaterThan(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.GreaterThan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, _ := got.EqualTo(tt.want)
			if !eq.Bool {
				t.Errorf("VariableBox.GreaterThan() = %v, want %v", got, tt.want)
			}
			if got.VariableType != tt.want.VariableType {
				t.Errorf("VariableBox.GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_LessThan(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		// Integer
		{
			name: "Test Integer LessThan: Correct",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer LessThan: Incorrect",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Integer LessThan with Float: Correct",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer LessThan with Float: Incorrect",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Integer LessThan with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "42",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			wantErr: false,
		},
		{
			name: "Test Integer LessThan with Bool: True",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
			wantErr: false,
		},
		// Float
		{
			name: "Test Float LessThan: Correct",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        13.13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Float LessThan: Incorrect",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Float LessThan with Integer: Correct",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        13.13,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Float LessThan with Integer: Incorrect",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		// String
		{
			name: "Test String LessThan: Correct",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "apple",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "banana",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test String LessThan: Incorrect",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "banana",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "apple",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.LessThan(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.LessThan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			eq, _ := got.EqualTo(tt.want)
			if !eq.Bool {
				t.Errorf("VariableBox.LessThan() = %v, want %v", got, tt.want)
			}
			if got.VariableType != tt.want.VariableType {
				t.Errorf("VariableBox.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_EqualTo(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		// Integer
		{
			name: "Test Integer EqualTo: Correct",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer EqualTo: Incorrect",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Integer EqualTo with Float: Correct",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.0,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer EqualTo with Float: Incorrect",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Integer EqualTo with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
			wantErr: false,
		},
		{
			name: "Test Integer EqualTo with Bool: True",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      1,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Integer EqualTo with Bool: False",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      1,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		// Float
		{
			name: "Test Float EqualTo: Correct",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        42.42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Float EqualTo: Incorrect",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Float EqualTo with Integer: Correct",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Float EqualTo with Integer: Incorrect",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      42,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		// String
		{
			name: "Test String EqualTo: Correct",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test String EqualTo: Incorrect",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "world",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		// Bool
		{
			name: "Test Bool EqualTo: Correct",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         true,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		{
			name: "Test Bool EqualTo: Incorrect",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         false,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
		},
		{
			name: "Test Bool EqualTo with Integer",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      1,
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
		},
		// Type mismatch
		{
			name: "Test Bool EqualTo with String",
			fields: fields{
				VariableType: TYPE_BOOL,
				Bool:         true,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "hello",
				},
			},
			want: VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.EqualTo(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.EqualTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VariableBox.EqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xor_strings(t *testing.T) {
	type args struct {
		text     string
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test xor_strings",
			args: args{
				text:     "hello",
				password: "world",
			},
			want: "rkAan",
		},
		{
			name: "Test xor_strings reverse",
			args: args{
				text:     "rkAan",
				password: "world",
			},
			want: "hello",
		},
		{
			name: "Test xor_strings short password",
			args: args{
				text:     "rsDDy",
				password: "w",
			},
			want: "hello",
		},
		{
			name: "Test xor_strings short password reverse",
			args: args{
				text:     "hello",
				password: "w",
			},
			want: "rsDDy",
		},
		{
			name: "Test xor_strings empty password",
			args: args{
				text:     "hello",
				password: "",
			},
			want: "hello",
		},
		{
			name: "Test xor_strings short text",
			args: args{
				text:     "hello",
				password: "hello world",
			},
			want: "aaaaa",
		},
		{
			name: "Test xor_strings short text reverse",
			args: args{
				text:     "aaaaa",
				password: "hello world",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xor_strings(tt.args.text, tt.args.password); got != tt.want {
				t.Errorf("xor_strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableBox_Xor(t *testing.T) {
	type fields struct {
		VariableType int8
		Integer      int64
		Float        float64
		String       string
		Bool         bool
		Processed    bool
	}
	type args struct {
		second VariableBox
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VariableBox
		wantErr bool
	}{
		// Integer
		{
			name: "Test Integer Xor",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      39,
			},
		},
		{
			name: "Test Integer Xor with Float",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      39,
			},
		},
		{
			name: "Test Integer Xor with String",
			fields: fields{
				VariableType: TYPE_INTEGER,
				Integer:      42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "13",
				},
			},
			want:    VariableBox{},
			wantErr: true,
		},
		// Float
		{
			name: "Test Float Xor",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_FLOAT,
					Float:        13.13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      39,
			},
		},
		{
			name: "Test Float Xor with Integer",
			fields: fields{
				VariableType: TYPE_FLOAT,
				Float:        42.42,
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_INTEGER,
					Integer:      13,
				},
			},
			want: VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      39,
			},
		},
		// String
		{
			name: "Test String Xor",
			fields: fields{
				VariableType: TYPE_STRING,
				String:       "hello",
			},
			args: args{
				second: VariableBox{
					VariableType: TYPE_STRING,
					String:       "world",
				},
			},
			want: VariableBox{
				VariableType: TYPE_STRING,
				String:       "rkAan",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VariableBox{
				VariableType: tt.fields.VariableType,
				Integer:      tt.fields.Integer,
				Float:        tt.fields.Float,
				String:       tt.fields.String,
				Bool:         tt.fields.Bool,
				Processed:    tt.fields.Processed,
			}
			got, err := v.Xor(tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("VariableBox.Xor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VariableBox.Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}
