package interpreter

import "testing"

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
