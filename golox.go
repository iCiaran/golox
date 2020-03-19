package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/iCiaran/golox/ast"
	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/parser"
	"github.com/iCiaran/golox/scanner"
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
	if loxerror.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		run(line)
		loxerror.HadError = false
	}
}

func run(source string) {
	sc := scanner.New(source)
	tokens := sc.ScanTokens()

	if loxerror.HadError {
		return
	}

	pa := parser.NewParser(tokens)
	ex := pa.Parse()

	if loxerror.HadError {
		return
	}

	pr := ast.NewPrinter()
	fmt.Println(pr.Print(ex))
}
