package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("exceed", "Manipulates Excel spreadsheet using a DSL")

	sourceFilePath := parser.String("f", "file", &argparse.Options{Required: true, Help: "Path to input spreadsheet"})
	saveAsPath := parser.String("o", "output", &argparse.Options{Required: false, Help: "Optionally save modified spreadsheet to other file location"})
	scriptPath := parser.String("s", "script", &argparse.Options{Required: false, Help: "Path to a script of instructions to execute"})
	evalScript := parser.String("e", "eval", &argparse.Options{Required: false, Help: "A verbatim script of instructions to evaluate"})
	interactive := parser.Flag("i", "interactive", &argparse.Options{Help: "Execute instructions interactively"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	fmt.Println("Source file path: ", *sourceFilePath)
	fmt.Println("Save as path: ", *saveAsPath)
	fmt.Println("Script path: ", *scriptPath)
	fmt.Println("Eval script: ", *evalScript)
	fmt.Println("Interactive: ", *interactive)
}
