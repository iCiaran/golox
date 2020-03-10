package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		run(line)
	}
}

func run(source string) {
	sc := scanner.New(source)
	tokens := sc.ScanTokens()

	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}
}
