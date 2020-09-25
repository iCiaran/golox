package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/iCiaran/golox/interpreter"
	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/parser"
	"github.com/iCiaran/golox/scanner"
)

var (
	in = interpreter.NewInterpreter()
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(66)
	}
	run(string(source))
	if loxerror.HadRuntimeError {
		os.Exit(70)
	}
	if loxerror.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		if len(line) > 1 && line[len(line)-2] != ';' {
			line = line[:len(line)-1] + ";\n"
		}
		run(line)
		loxerror.HadError = false
		loxerror.HadRuntimeError = false
	}
}

func run(source string) {
	sc := scanner.New(source)
	tokens := sc.ScanTokens()

	if loxerror.HadError {
		return
	}

	pa := parser.NewParser(tokens)
	st := pa.Parse()

	if loxerror.HadError {
		return
	}

	in.Interpret(st)
}
