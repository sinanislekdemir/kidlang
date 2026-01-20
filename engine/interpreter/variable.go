package interpreter

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	TYPE_UNKNOWN   = -1
	TYPE_INTEGER   = 0
	TYPE_FLOAT     = 1
	TYPE_STACK     = 2
	TYPE_STRING    = 3
	TYPE_BOOL      = 4
	TYPE_REFERENCE = 5 // Reference to another variable, for function arguments
	TYPE_FILE      = 6
)

var NumericTypes = []int{
	TYPE_FLOAT, TYPE_INTEGER,
}

type SpecialVariable struct {
	Pattern  string
	Function func() VariableBox
}

var Specials = []SpecialVariable{
	{
		Pattern: "\\n",
		Function: func() VariableBox {
			return VariableBox{
				VariableType: TYPE_STRING,
				String:       "\n",
			}
		},
	},
	{
		Pattern: "RANDOM",
		Function: func() VariableBox {
			return VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      int64(rand.Int63()),
			}
		},
	},
	{
		Pattern: "NOW",
		Function: func() VariableBox {
			now := time.Now()
			return VariableBox{
				VariableType: TYPE_STRING,
				String:       now.Format("Monday, January 2, 2006 15:03:05"),
			}
		},
	},
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.0000001
}

type VariableBox struct {
	VariableType int8    `json:"variable_type"`
	Integer      int64   `json:"integer"`
	Float        float64 `json:"float"`
	String       string  `json:"string"`
	Bool         bool    `json:"bool"`
	Processed    bool    `json:"processed"`
	Filename     string  `json:"filename"`
	fileHandler  *os.File
	StackData    map[string]VariableBox `json:"stack_data"` // For STACK type
}

func (v VariableBox) isString() bool {
	return v.VariableType == TYPE_STRING || v.VariableType == TYPE_FILE
}

func (v VariableBox) isAssignable() bool {
	return v.VariableType == TYPE_REFERENCE || v.VariableType == TYPE_FILE
}

func (v *VariableBox) SetFileHandler(file *os.File) {
	v.fileHandler = file
}

func (v VariableBox) ToString() string {
	switch v.VariableType {
	case TYPE_INTEGER:
		return strconv.FormatInt(v.Integer, 10)
	case TYPE_FLOAT:
		return strconv.FormatFloat(v.Float, 'f', -1, 64)
	case TYPE_STRING:
		return v.String
	case TYPE_REFERENCE:
		return v.String // No memory context here
	case TYPE_BOOL:
		if v.Bool {
			return "True"
		} else {
			return "False"
		}
	case TYPE_FILE:
		if v.fileHandler == nil {
			return ""
		}
		if v.Filename == "" {
			return ""
		}
		_, err := v.fileHandler.Stat()
		if err != nil {
			return ""
		}
		cursor, err := v.fileHandler.Seek(0, 1)
		if err != nil {
			return ""
		}
		v.fileHandler.Seek(0, 0)
		content, err := os.ReadFile(v.Filename)
		if err != nil {
			return ""
		}
		v.fileHandler.Seek(cursor, 0)
		return string(content)
	default:
		return ""
	}
}

func (v VariableBox) ToFloat() float64 {
	switch v.VariableType {
	case TYPE_INTEGER:
		return float64(v.Integer)
	case TYPE_FLOAT:
		return v.Float
	case TYPE_BOOL:
		return float64(boolToInt(v.Bool))
	case TYPE_STRING:
		if val, err := strconv.ParseFloat(v.String, 64); err == nil {
			return val
		}
	}
	return 0.0
}

func boolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func (v VariableBox) Sum(second VariableBox) VariableBox {
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer + second.Integer,
		}
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        float64(v.Integer) + second.Float,
		}
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer + int64(boolToInt(second.Bool)),
		}
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float + float64(second.Integer),
		}
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float + second.Float,
		}
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float + float64(boolToInt(second.Bool)),
		}
	}
	return VariableBox{
		VariableType: TYPE_STRING,
		String:       fmt.Sprintf("%s%s", v.ToString(), second.ToString()),
	}
}

func (v VariableBox) Sub(second VariableBox) (VariableBox, error) {
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer - second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        float64(v.Integer) - second.Float,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer - int64(boolToInt(second.Bool)),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float - float64(second.Integer),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float - second.Float,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float - float64(boolToInt(second.Bool)),
		}, nil
	}
	if v.isString() {
		return VariableBox{
			VariableType: TYPE_STRING,
			String:       strings.Replace(v.ToString(), second.String, "", -1),
		}, nil
	}
	return VariableBox{}, fmt.Errorf("NaN")
}

func (v VariableBox) Mul(second VariableBox) (VariableBox, error) {
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer * second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        float64(v.Integer) * second.Float,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer * int64(boolToInt(second.Bool)),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float * float64(second.Integer),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float * second.Float,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float * float64(boolToInt(second.Bool)),
		}, nil
	}
	if v.isString() && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_STRING,
			String:       strings.Repeat(v.ToString(), int(second.Integer)),
		}, nil
	}
	return VariableBox{}, fmt.Errorf("NaN")
}

func (v VariableBox) Div(second VariableBox) (VariableBox, error) {
	if (second.VariableType == TYPE_INTEGER && second.Integer == 0) || (second.VariableType == TYPE_FLOAT && second.Float == 0.0) || (second.VariableType == TYPE_BOOL && !second.Bool) {
		return VariableBox{}, fmt.Errorf("division by zero")
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		res := float64(v.Integer) / float64(second.Integer)
		if res == math.Trunc(res) {
			return VariableBox{
				VariableType: TYPE_INTEGER,
				Integer:      v.Integer / second.Integer,
			}, nil
		} else {
			return VariableBox{
				VariableType: TYPE_FLOAT,
				Float:        float64(v.Integer) / float64(second.Integer),
			}, nil
		}

	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        float64(v.Integer) / second.Float,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer / int64(boolToInt(second.Bool)),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float / float64(second.Integer),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float / second.Float,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        v.Float / float64(boolToInt(second.Bool)),
		}, nil
	}
	if v.isString() && second.VariableType == TYPE_INTEGER {
		if second.Integer > int64(len(v.ToString())) {
			second.Integer = int64(len(v.ToString()))
		}

		return VariableBox{
			VariableType: TYPE_STRING,
			String:       string(v.ToString()[second.Integer-1]),
		}, nil
	}
	if v.isString() && second.isString() {
		return VariableBox{
			VariableType: TYPE_STRING,
			String:       fmt.Sprintf("%s/%s", v.ToString(), second.ToString()),
		}, nil
	}
	return VariableBox{}, fmt.Errorf("NaN")
}

func (v VariableBox) Mod(second VariableBox) (VariableBox, error) {
	if (second.VariableType == TYPE_INTEGER && second.Integer == 0) || (second.VariableType == TYPE_FLOAT && second.Float == 0.0) {
		return VariableBox{}, fmt.Errorf("modulo by zero")
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer % second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        math.Mod(float64(v.Integer), second.Float),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        math.Mod(v.Float, float64(second.Integer)),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_FLOAT,
			Float:        math.Mod(v.Float, second.Float),
		}, nil
	}
	return VariableBox{}, fmt.Errorf("NaN")
}

func (v VariableBox) GreaterThan(second VariableBox) (VariableBox, error) {
	// Try to convert strings to numbers for comparison
	if v.VariableType == TYPE_STRING && (second.VariableType == TYPE_INTEGER || second.VariableType == TYPE_FLOAT) {
		if val, err := strconv.ParseFloat(v.String, 64); err == nil {
			if second.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         val > float64(second.Integer),
				}, nil
			}
			return VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         val > second.Float,
			}, nil
		}
	}
	if (v.VariableType == TYPE_INTEGER || v.VariableType == TYPE_FLOAT) && second.VariableType == TYPE_STRING {
		if val, err := strconv.ParseFloat(second.String, 64); err == nil {
			if v.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         float64(v.Integer) > val,
				}, nil
			}
			return VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         v.Float > val,
			}, nil
		}
	}

	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Integer > second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         float64(v.Integer) > second.Float,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Float > float64(second.Integer),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Float > second.Float,
		}, nil
	}
	if v.isString() && second.isString() {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         strings.Compare(strings.ToLower(v.ToString()), strings.ToLower(second.ToString())) > 0,
		}, nil
	}
	// do your best shot
	return VariableBox{
		VariableType: TYPE_BOOL,
		Bool:         strings.Compare(v.ToString(), second.ToString()) > 0,
	}, nil
}

func (v VariableBox) LessThan(second VariableBox) (VariableBox, error) {
	// Try to convert strings to numbers for comparison
	if v.VariableType == TYPE_STRING && (second.VariableType == TYPE_INTEGER || second.VariableType == TYPE_FLOAT) {
		if val, err := strconv.ParseFloat(v.String, 64); err == nil {
			if second.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         val < float64(second.Integer),
				}, nil
			}
			return VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         val < second.Float,
			}, nil
		}
	}
	if (v.VariableType == TYPE_INTEGER || v.VariableType == TYPE_FLOAT) && second.VariableType == TYPE_STRING {
		if val, err := strconv.ParseFloat(second.String, 64); err == nil {
			if v.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         float64(v.Integer) < val,
				}, nil
			}
			return VariableBox{
				VariableType: TYPE_BOOL,
				Bool:         v.Float < val,
			}, nil
		}
	}

	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Integer < second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         float64(v.Integer) < second.Float,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Float < float64(second.Integer),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Float < second.Float,
		}, nil
	}
	if v.isString() && second.isString() {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         strings.Compare(strings.ToLower(v.ToString()), strings.ToLower(second.ToString())) < 0,
		}, nil
	}
	return VariableBox{
		VariableType: TYPE_BOOL,
		Bool:         false,
	}, nil
}

func (v VariableBox) EqualTo(second VariableBox) (VariableBox, error) {
	// Try to convert strings to numbers for comparison
	if v.VariableType == TYPE_STRING && (second.VariableType == TYPE_INTEGER || second.VariableType == TYPE_FLOAT) {
		// Try to parse string as number
		if val, err := strconv.ParseInt(v.String, 10, 64); err == nil {
			// String is an integer
			if second.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         val == second.Integer,
				}, nil
			}
			if second.VariableType == TYPE_FLOAT {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         almostEqual(float64(val), second.Float),
				}, nil
			}
		}
		if val, err := strconv.ParseFloat(v.String, 64); err == nil {
			// String is a float
			if second.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         almostEqual(val, float64(second.Integer)),
				}, nil
			}
			if second.VariableType == TYPE_FLOAT {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         almostEqual(val, second.Float),
				}, nil
			}
		}
	}

	if (v.VariableType == TYPE_INTEGER || v.VariableType == TYPE_FLOAT) && second.VariableType == TYPE_STRING {
		// Try to parse string as number
		if val, err := strconv.ParseInt(second.String, 10, 64); err == nil {
			// String is an integer
			if v.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         v.Integer == val,
				}, nil
			}
			if v.VariableType == TYPE_FLOAT {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         almostEqual(v.Float, float64(val)),
				}, nil
			}
		}
		if val, err := strconv.ParseFloat(second.String, 64); err == nil {
			// String is a float
			if v.VariableType == TYPE_INTEGER {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         almostEqual(float64(v.Integer), val),
				}, nil
			}
			if v.VariableType == TYPE_FLOAT {
				return VariableBox{
					VariableType: TYPE_BOOL,
					Bool:         almostEqual(v.Float, val),
				}, nil
			}
		}
	}

	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Integer == second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         almostEqual(float64(v.Integer), second.Float),
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Integer == int64(boolToInt(second.Bool)),
		}, nil
	}
	if v.VariableType == TYPE_BOOL && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         boolToInt(v.Bool) == second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         almostEqual(v.Float, float64(second.Integer)),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         almostEqual(v.Float, second.Float),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         almostEqual(v.Float, float64(boolToInt(second.Bool))),
		}, nil
	}
	if v.isString() && second.isString() {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         strings.Compare(strings.ToLower(v.ToString()), strings.ToLower(second.ToString())) == 0,
		}, nil
	}
	if v.VariableType == TYPE_BOOL && second.VariableType == TYPE_BOOL {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.Bool == second.Bool,
		}, nil
	}
	if v.VariableType == TYPE_REFERENCE && second.VariableType == TYPE_REFERENCE {
		return VariableBox{
			VariableType: TYPE_BOOL,
			Bool:         v.String == second.String,
		}, nil
	}
	return VariableBox{
		VariableType: TYPE_BOOL,
		Bool:         strings.EqualFold(v.ToString(), second.ToString()),
	}, nil
}

func xor_strings(text string, password string) string {
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 _/\\:.,?!@#$%^&*()-=+[]{}|;<>`~"
	textLen := len(text)
	passwordLen := len(password)
	result := make([]byte, textLen)
	if passwordLen == 0 {
		return text
	}

	for i := 0; i < textLen; i++ {
		charIndex := strings.IndexByte(validChars, text[i])
		passIndex := strings.IndexByte(validChars, password[i%passwordLen])
		result[i] = validChars[(charIndex^passIndex)%len(validChars)]
	}

	return string(result)
}

func (v VariableBox) Xor(second VariableBox) (VariableBox, error) {
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer ^ second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_INTEGER && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      v.Integer ^ int64(second.Float),
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_INTEGER {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      int64(v.Float) ^ second.Integer,
		}, nil
	}
	if v.VariableType == TYPE_FLOAT && second.VariableType == TYPE_FLOAT {
		return VariableBox{
			VariableType: TYPE_INTEGER,
			Integer:      int64(v.Float) ^ int64(second.Float),
		}, nil
	}
	if v.isString() {
		return VariableBox{
			VariableType: TYPE_STRING,
			String:       xor_strings(v.ToString(), second.ToString()),
		}, nil
	}
	return VariableBox{}, fmt.Errorf("type mismatch")
}

type Crate struct {
	Array []VariableBox `json:"boxes"`
}

// Stack operations
func (v *VariableBox) GetFromStack(key string) VariableBox {
	if v.VariableType != TYPE_STACK {
		return VariableBox{VariableType: TYPE_UNKNOWN}
	}
	if v.StackData == nil {
		return VariableBox{VariableType: TYPE_UNKNOWN}
	}
	if val, exists := v.StackData[key]; exists {
		return val
	}
	// Non-existent key returns empty/unknown
	return VariableBox{VariableType: TYPE_UNKNOWN}
}

func (v *VariableBox) SetInStack(key string, value VariableBox) {
	if v.VariableType != TYPE_STACK {
		return
	}
	if v.StackData == nil {
		v.StackData = make(map[string]VariableBox)
	}
	v.StackData[key] = value
}

func NewStack() VariableBox {
	return VariableBox{
		VariableType: TYPE_STACK,
		StackData:    make(map[string]VariableBox),
	}
}
