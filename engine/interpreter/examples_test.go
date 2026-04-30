package interpreter

import (
	"path/filepath"
	"testing"
)

func TestExamplesLoad(t *testing.T) {
	exampleFiles, err := filepath.Glob("../../examples/*.kid")
	if err != nil {
		t.Fatalf("Glob() error = %v", err)
	}
	if len(exampleFiles) == 0 {
		t.Fatal("no example files found")
	}

	for _, exampleFile := range exampleFiles {
		t.Run(filepath.Base(exampleFile), func(t *testing.T) {
			program := &Program{}
			program.Init()
			if err := program.Load(exampleFile); err != nil {
				t.Fatalf("Load(%s) error = %v", exampleFile, err)
			}
		})
	}
}

func TestEngineExamplesLoad(t *testing.T) {
	engineExampleFiles, err := filepath.Glob("../Examples/*/*.kid")
	if err != nil {
		t.Fatalf("Glob() error = %v", err)
	}
	if len(engineExampleFiles) == 0 {
		t.Fatal("no engine example files found")
	}

	for _, exampleFile := range engineExampleFiles {
		t.Run(filepath.Base(filepath.Dir(exampleFile))+"_"+filepath.Base(exampleFile), func(t *testing.T) {
			program := &Program{}
			program.Init()
			if err := program.Load(exampleFile); err != nil {
				t.Fatalf("Load(%s) error = %v", exampleFile, err)
			}
		})
	}
}
