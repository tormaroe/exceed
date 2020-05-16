package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/tormaroe/exceed/picol"
)

func commandPuts(i *picol.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", fmt.Errorf("Wrong number of args for %s %s", argv[0], argv)
	}
	fmt.Println(argv[1])
	return "", nil
}

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

	interp := picol.InitInterp()
	interp.RegisterCoreCommands()
	interp.RegisterCommand("puts", commandPuts, nil)

	result, err := interp.Eval(*evalScript)
	if err != nil {
		fmt.Println("ERROR", result, err)
	}

	if *interactive {
		// TODO: Print header
		for {
			fmt.Print("exceed> ")
			scanner := bufio.NewReader(os.Stdin)
			clibuf, _ := scanner.ReadString('\n')
			result, err := interp.Eval(clibuf[:len(clibuf)-1])
			if err != nil {
				fmt.Println("ERROR", result, err)
			}
		}
	}

	// TODO: Save
}
