package main

import (
	"flag"
	"fmt"

	"github.com/sinanislekdemir/kidlang/interpreter"
)

func main() {
	p := interpreter.Program{}
	p.Init()

	// Get the first command line argument as the filename and run the program
	// And get the others as flags. I have a flag as debug, which is a boolean.
	debug := flag.Bool("debug", false, "debug mode")
	dump := flag.Bool("dump", false, "dump statements")
	lang := flag.String("lang", "en", "language")
	wait := flag.Bool("wait", false, "wait for key press")

	flag.Parse()

	// If no arguments, launch the IDE
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
	// wait for key press
	if *wait {
		fmt.Println("Press any key to exit...")
		fmt.Scanln()
	}
}
