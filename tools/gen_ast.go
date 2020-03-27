package main

import (
	"fmt"
	"os"
	"strings"
)

func generateAst() {
	outputDir := "../ast"

	defineAst(outputDir, "Expr", []string{
		"Assign   : Name *token.Token, Value Expr",
		"Binary	  : Left Expr, Operator *token.Token, Right Expr",
		"Grouping : Expression Expr",
		"Literal  : Value interface{}",
		"Unary    : Operator *token.Token, Right Expr",
		"Variable : Name *token.Token",
	})

	defineAst(outputDir, "Stmt", []string{
		"Expression : Expr Expr",
		"Print      : Expr Expr",
		"Var        : Name *token.Token, Initializer Expr",
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeLine(f *os.File, line string) {
	f.WriteString(line + "\n")
}

func makeDelimited(args []string, delim string) string {
	var ret string
	for i, arg := range args {
		ret += arg
		if i < len(args)-1 {
			ret += delim
		}
	}
	return ret
}

func writeImports(f *os.File, imports []string) {
	writeLine(f, "import (")
	for _, im := range imports {
		writeLine(f, im)
	}
	writeLine(f, ")")
}

func defineAst(outputDir, baseName string, types []string) {
	path := outputDir + "/" + strings.ToLower(baseName) + ".go"
	f, err := os.Create(path)
	check(err)

	defer f.Close()

	writeLine(f, "package ast")

	switch baseName {
	case "Expr":
		writeImports(f, []string{`"github.com/iCiaran/golox/token"`})
	case "Stmt":
		writeImports(f, []string{`"github.com/iCiaran/golox/token"`})
	}

	defineVisitor(f, baseName, types)

	writeLine(f, fmt.Sprintf("type %s interface {", baseName))
	writeLine(f, fmt.Sprintf("Accept(v %sVisitor) interface{}", baseName))
	writeLine(f, "}")

	defineTypes(f, baseName, types)
}

func defineVisitor(f *os.File, baseName string, types []string) {
	writeLine(f, "type "+baseName+"Visitor interface {")
	for _, t := range types {
		className := strings.TrimSpace(strings.Split(t, ":")[0])
		writeLine(f, fmt.Sprintf("Visit%s%s(expr %s) interface{}", className, baseName, className))
	}
	writeLine(f, "}")
}

func defineTypes(f *os.File, baseName string, types []string) {
	for _, t := range types {
		split := strings.Split(t, ":")
		className := strings.TrimSpace(split[0])
		args := strings.Split(split[1], ",")
		defineStruct(f, className, args)
		defineNew(f, className, args)
		defineAccept(f, baseName, className)
	}
}

func defineStruct(f *os.File, className string, args []string) {
	writeLine(f, fmt.Sprintf("type %s struct {", className))
	for _, arg := range args {
		writeLine(f, arg)
	}
	writeLine(f, "}")
}

func defineNew(f *os.File, className string, args []string) {
	argsList := make([]string, 0)
	returnList := make([]string, 0)
	for _, arg := range args {
		split := strings.Split(strings.TrimSpace(arg), " ")
		i := strings.TrimSpace(split[0])
		t := strings.TrimSpace(split[1])
		argsList = append(argsList, fmt.Sprintf("%s %s", strings.ToLower(i), t))
		returnList = append(returnList, fmt.Sprintf("%s: %s", i, strings.ToLower(i)))
	}

	writeLine(f, fmt.Sprintf("func New%s(%s) *%s {", className, makeDelimited(argsList, ","), className))
	writeLine(f, fmt.Sprintf("return &%s{%s}", className, makeDelimited(returnList, ",")))
	writeLine(f, "}")
}

func defineAccept(f *os.File, baseName, className string) {
	c := strings.ToLower(className[0:1])
	writeLine(f, fmt.Sprintf("func (%s *%s) Accept(vis %sVisitor) interface{} {", c, className, baseName))
	writeLine(f, fmt.Sprintf("return vis.Visit%s%s(*%s)", className, baseName, c))
	writeLine(f, "}")
}
