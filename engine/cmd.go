package main

import (
	"flag"
	"fmt"

	"github.com/sinanislekdemir/kidlang/interpreter"
)

func main() {
	p := interpreter.Program{}
	p.Init()

	debug := flag.Bool("debug", false, "debug mode")
	dump := flag.Bool("dump", false, "dump statements")
	lang := flag.String("lang", "en", "language")
	wait := flag.Bool("wait", false, "wait for key press")
	build := flag.String("build", "", "build embedded binary from script (e.g., --build script.kid)")

	flag.Parse()

	if *build != "" {
		outputPath := *build + ".bin" + interpreter.GetEmbeddedBinaryExtension()
		fmt.Printf("Building embedded binary: %s -> %s\n", *build, outputPath)
		if err := interpreter.BuildEmbeddedBinary(*build, outputPath); err != nil {
			fmt.Println("Error building binary:", err)
			return
		}
		fmt.Printf("Successfully created: %s\n", outputPath)
		return
	}

	if interpreter.HasEmbeddedScript() {
		scriptName, scriptData, err := interpreter.ExtractEmbeddedScript()
		if err != nil {
			fmt.Println("Error extracting embedded script:", err)
			return
		}
		fmt.Printf("Running embedded script: %s\n", scriptName)
		p.SetLanguage(*lang)
		p.Debug = *debug
		if err := p.LoadFromString(string(scriptData)); err != nil {
			fmt.Println("Error loading embedded script:", err)
			return
		}
		if *dump {
			for _, s := range p.Statements {
				fmt.Printf("Index %d: %s\n", s.LineNumber, s.FullLine)
			}
		}
		if err := p.Run(); err != nil {
			fmt.Println("Error running embedded script:", err)
		}
		if *wait {
			fmt.Println("Press any key to exit...")
			fmt.Scanln()
		}
		return
	}

	if flag.NArg() == 0 {
		if err := interpreter.StartIDE(); err != nil {
			fmt.Println("Error starting IDE:", err)
			return
		}
		return
	}

	p.SetLanguage(*lang)
	p.Debug = *debug
	filename := flag.Args()[0]
	err := p.Load(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	if *dump {
		for _, s := range p.Statements {
			fmt.Printf("Index %d: %s\n", s.LineNumber, s.FullLine)
		}
	}
	p.Run()
	if *wait {
		fmt.Println("Press any key to exit...")
		fmt.Scanln()
	}
}
