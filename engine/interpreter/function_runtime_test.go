package interpreter

import (
	"os"
	"testing"
)

func TestProgramRunBuiltinFunctionCalls(t *testing.T) {
	p := &Program{}
	p.Init()

	program := `
box root = sqrt(16)
box letters = len(hello)
box loud = upper(kidlang)
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX ROOT"].VariableType != TYPE_FLOAT || !almostEqual(p.Memory["BOX ROOT"].Float, 4) {
		t.Fatalf("BOX ROOT = %+v, want 4", p.Memory["BOX ROOT"])
	}
	if p.Memory["BOX LETTERS"].VariableType != TYPE_INTEGER || p.Memory["BOX LETTERS"].Integer != 5 {
		t.Fatalf("BOX LETTERS = %+v, want 5", p.Memory["BOX LETTERS"])
	}
	if p.Memory["BOX LOUD"].VariableType != TYPE_STRING || p.Memory["BOX LOUD"].String != "KIDLANG" {
		t.Fatalf("BOX LOUD = %+v, want KIDLANG", p.Memory["BOX LOUD"])
	}
}

func TestProgramRunUserDefinedFunctionImplicitReturn(t *testing.T) {
	p := &Program{}
	p.Init()

	program := `
function add(box a, box b)
box a + box b
end

box result = add(4, 5)
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX RESULT"].VariableType != TYPE_INTEGER || p.Memory["BOX RESULT"].Integer != 9 {
		t.Fatalf("BOX RESULT = %+v, want 9", p.Memory["BOX RESULT"])
	}
}

func TestProgramRunUserDefinedFunctionExplicitReturn(t *testing.T) {
	p := &Program{}
	p.Init()

	program := `
function double_and_add_one(box value)
box doubled = box value * 2
return box doubled + 1
end

box result = double_and_add_one(6)
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX RESULT"].VariableType != TYPE_INTEGER || p.Memory["BOX RESULT"].Integer != 13 {
		t.Fatalf("BOX RESULT = %+v, want 13", p.Memory["BOX RESULT"])
	}
}

func TestProgramRunNestedFunctionCalls(t *testing.T) {
	p := &Program{}
	p.Init()

	program := `
function add(box a, box b)
box a + box b
end

box answer = max(add(2, 3), min(10, 4))
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX ANSWER"].VariableType != TYPE_INTEGER || p.Memory["BOX ANSWER"].Integer != 5 {
		t.Fatalf("BOX ANSWER = %+v, want 5", p.Memory["BOX ANSWER"])
	}
}

func TestBuiltinFunctionStringLibrary(t *testing.T) {
	trimmed, err := findBuiltinFunction("TRIM").Run(BuiltinInvocation{
		Arguments: []VariableBox{
			{VariableType: TYPE_STRING, String: "  hello  "},
		},
	})
	if err != nil {
		t.Fatalf("TRIM() error = %v", err)
	}
	if trimmed.VariableType != TYPE_STRING || trimmed.String != "hello" {
		t.Fatalf("TRIM() = %+v, want hello", trimmed)
	}

	p := &Program{}
	p.Init()

	program := `
box found = contains(kidlang, lang)
box prefix = startswith(kidlang, kid)
box suffix = endswith(kidlang, lang)
box changed = replace(kidlang, kid, play)
box part = substring(kidlang, 2, 3)
box where = indexof(kidlang, lang)
box pieces = split(red|blue|green, |)
box joined = join(box pieces, :)
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX FOUND"].VariableType != TYPE_BOOL || !p.Memory["BOX FOUND"].Bool {
		t.Fatalf("BOX FOUND = %+v, want true", p.Memory["BOX FOUND"])
	}
	if p.Memory["BOX PREFIX"].VariableType != TYPE_BOOL || !p.Memory["BOX PREFIX"].Bool {
		t.Fatalf("BOX PREFIX = %+v, want true", p.Memory["BOX PREFIX"])
	}
	if p.Memory["BOX SUFFIX"].VariableType != TYPE_BOOL || !p.Memory["BOX SUFFIX"].Bool {
		t.Fatalf("BOX SUFFIX = %+v, want true", p.Memory["BOX SUFFIX"])
	}
	if p.Memory["BOX CHANGED"].VariableType != TYPE_STRING || p.Memory["BOX CHANGED"].String != "playlang" {
		t.Fatalf("BOX CHANGED = %+v, want playlang", p.Memory["BOX CHANGED"])
	}
	if p.Memory["BOX PART"].VariableType != TYPE_STRING || p.Memory["BOX PART"].String != "idl" {
		t.Fatalf("BOX PART = %+v, want idl", p.Memory["BOX PART"])
	}
	if p.Memory["BOX WHERE"].VariableType != TYPE_INTEGER || p.Memory["BOX WHERE"].Integer != 4 {
		t.Fatalf("BOX WHERE = %+v, want 4", p.Memory["BOX WHERE"])
	}
	if p.Memory["BOX PIECES"].VariableType != TYPE_STACK {
		t.Fatalf("BOX PIECES = %+v, want stack", p.Memory["BOX PIECES"])
	}
	if p.Memory["BOX JOINED"].VariableType != TYPE_STRING || p.Memory["BOX JOINED"].String != "red:blue:green" {
		t.Fatalf("BOX JOINED = %+v, want red:blue:green", p.Memory["BOX JOINED"])
	}
}

func TestBuiltinFunctionFileLibrary(t *testing.T) {
	p := &Program{}
	p.Init()

	filename := "/tmp/kidlang_builtin_file_library.txt"
	defer os.Remove(filename)
	content := "alpha\nbeta"
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	program := `
file sample = source
open sample /tmp/kidlang_builtin_file_library.txt
box exists = fileexists(sample)
box text = fileread(sample)
box bytes = filesize(sample)
close sample
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX EXISTS"].VariableType != TYPE_BOOL || !p.Memory["BOX EXISTS"].Bool {
		t.Fatalf("BOX EXISTS = %+v, want true", p.Memory["BOX EXISTS"])
	}
	if p.Memory["BOX TEXT"].VariableType != TYPE_STRING || p.Memory["BOX TEXT"].String != content {
		t.Fatalf("BOX TEXT = %+v, want %q", p.Memory["BOX TEXT"], content)
	}
	if p.Memory["BOX BYTES"].VariableType != TYPE_INTEGER || p.Memory["BOX BYTES"].Integer != int64(len(content)) {
		t.Fatalf("BOX BYTES = %+v, want %d", p.Memory["BOX BYTES"], len(content))
	}
}

func TestBuiltinFunctionStackLibrary(t *testing.T) {
	p := &Program{}
	p.Init()

	program := `
stack toys
box first = stackset(toys, 1, robot)
box second = stackset(toys, 2, kite)
box has_first = stackhas(toys, 1)
box fetched = stackget(toys, 2)
box keys = stackkeys(toys)
box key_line = join(box keys, :)
box count_before = stacklen(toys)
box removed = stackdelete(toys, 1)
box count_after = stacklen(toys)
`

	if err := p.LoadFromString(program); err != nil {
		t.Fatalf("LoadFromString() error = %v", err)
	}
	if err := p.Run(); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if p.Memory["BOX FIRST"].VariableType != TYPE_STRING || p.Memory["BOX FIRST"].String != "robot" {
		t.Fatalf("BOX FIRST = %+v, want robot", p.Memory["BOX FIRST"])
	}
	if p.Memory["BOX SECOND"].VariableType != TYPE_STRING || p.Memory["BOX SECOND"].String != "kite" {
		t.Fatalf("BOX SECOND = %+v, want kite", p.Memory["BOX SECOND"])
	}
	if p.Memory["BOX HAS_FIRST"].VariableType != TYPE_BOOL || !p.Memory["BOX HAS_FIRST"].Bool {
		t.Fatalf("BOX HAS_FIRST = %+v, want true", p.Memory["BOX HAS_FIRST"])
	}
	if p.Memory["BOX FETCHED"].VariableType != TYPE_STRING || p.Memory["BOX FETCHED"].String != "kite" {
		t.Fatalf("BOX FETCHED = %+v, want kite", p.Memory["BOX FETCHED"])
	}
	if p.Memory["BOX COUNT_BEFORE"].VariableType != TYPE_INTEGER || p.Memory["BOX COUNT_BEFORE"].Integer != 2 {
		t.Fatalf("BOX COUNT_BEFORE = %+v, want 2", p.Memory["BOX COUNT_BEFORE"])
	}
	if p.Memory["BOX REMOVED"].VariableType != TYPE_BOOL || !p.Memory["BOX REMOVED"].Bool {
		t.Fatalf("BOX REMOVED = %+v, want true", p.Memory["BOX REMOVED"])
	}
	if p.Memory["BOX COUNT_AFTER"].VariableType != TYPE_INTEGER || p.Memory["BOX COUNT_AFTER"].Integer != 1 {
		t.Fatalf("BOX COUNT_AFTER = %+v, want 1", p.Memory["BOX COUNT_AFTER"])
	}
	if p.Memory["BOX KEY_LINE"].VariableType != TYPE_STRING || p.Memory["BOX KEY_LINE"].String != "1:2" {
		t.Fatalf("BOX KEY_LINE = %+v, want 1:2", p.Memory["BOX KEY_LINE"])
	}
	if p.Memory["TOYS"].VariableType != TYPE_STACK {
		t.Fatalf("TOYS = %+v, want stack", p.Memory["TOYS"])
	}
	if _, exists := p.Memory["TOYS"].StackData["1"]; exists {
		t.Fatalf("TOYS should not contain key 1 after delete: %+v", p.Memory["TOYS"])
	}
	if value, exists := p.Memory["TOYS"].StackData["2"]; !exists || value.String != "kite" {
		t.Fatalf("TOYS[2] = %+v, want kite", p.Memory["TOYS"].StackData["2"])
	}
}

func TestBuiltinFunctionLocalizedNames(t *testing.T) {
	tests := []struct {
		name      string
		language  string
		program   string
		memoryKey string
		expected  VariableBox
	}{
		{
			name:     "Turkish contains translation",
			language: LANG_TR,
			program: `
TR
kutu sonuc = icindevar(kidlang, lang)
`,
			memoryKey: "KUTU SONUC",
			expected:  VariableBox{VariableType: TYPE_BOOL, Bool: true},
		},
		{
			name:     "German replace translation",
			language: LANG_DE,
			program: `
DE
kiste sonuc = ersetze(katze, ze, zchen)
`,
			memoryKey: "KISTE SONUC",
			expected:  VariableBox{VariableType: TYPE_STRING, String: "katzchen"},
		},
		{
			name:     "Finnish stack length translation",
			language: LANG_FI,
			program: `
FI
lista lelut
laatikko eka = laitalistaan(lelut, 1, robotti)
laatikko maara = listanpituus(lelut)
`,
			memoryKey: "LAATIKKO MAARA",
			expected:  VariableBox{VariableType: TYPE_INTEGER, Integer: 1},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			p := &Program{}
			p.Init()
			p.SetLanguage(testCase.language)

			if err := p.LoadFromString(testCase.program); err != nil {
				t.Fatalf("LoadFromString() error = %v", err)
			}
			if err := p.Run(); err != nil {
				t.Fatalf("Run() error = %v", err)
			}

			actual := p.Memory[testCase.memoryKey]
			if actual.VariableType != testCase.expected.VariableType {
				t.Fatalf("%s type = %+v, want %+v", testCase.memoryKey, actual, testCase.expected)
			}
			switch actual.VariableType {
			case TYPE_BOOL:
				if actual.Bool != testCase.expected.Bool {
					t.Fatalf("%s = %+v, want %+v", testCase.memoryKey, actual, testCase.expected)
				}
			case TYPE_STRING:
				if actual.String != testCase.expected.String {
					t.Fatalf("%s = %+v, want %+v", testCase.memoryKey, actual, testCase.expected)
				}
			case TYPE_INTEGER:
				if actual.Integer != testCase.expected.Integer {
					t.Fatalf("%s = %+v, want %+v", testCase.memoryKey, actual, testCase.expected)
				}
			}
		})
	}
}
