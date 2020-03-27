package parser

import (
	"github.com/iCiaran/golox/ast"
	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/token"
)

type Parser struct {
	Tokens  []*token.Token
	Current int
}

func NewParser(tokens []*token.Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() []ast.Stmt {
	statements := make([]ast.Stmt, 0)
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
}

func (p *Parser) expression() ast.Expr {
	return p.assignment()
}

func (p *Parser) assignment() ast.Expr {
	expr := p.equality()
	if p.match(token.EQUAL) {
		equals := p.previous()
		value := p.assignment()

		if val, ok := expr.(*ast.Variable); ok {
			return ast.NewAssign(val.Name, value)
		}
		loxerror.ParseError(equals, "Invalid assignment target.")
	}
	return expr
}

func (p *Parser) declaration() ast.Stmt {
	defer func() {
		if r := recover(); r != nil {
			p.synchronise()
		}
	}()

	if p.match(token.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() ast.Expr {
	expr := p.addition()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.addition()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) addition() ast.Expr {
	expr := p.multiplication()
	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.multiplication()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) multiplication() ast.Expr {
	expr := p.unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = ast.NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.NewUnary(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	switch {
	case p.match(token.FALSE):
		return ast.NewLiteral(false)
	case p.match(token.TRUE):
		return ast.NewLiteral(true)
	case p.match(token.NIL):
		return ast.NewLiteral(nil)
	case p.match(token.NUMBER) || p.match(token.STRING):
		return ast.NewLiteral(p.previous().Literal)
	case p.match(token.LEFT_PAREN):
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return ast.NewGrouping(expr)
	case p.match(token.IDENTIFIER):
		return ast.NewVariable(p.previous())
	}

	loxerror.ParseError(p.peek(), "Expect expression.")
	return nil
}

func (p *Parser) statement() ast.Stmt {
	if p.match(token.PRINT) {
		return p.printStatement()
	}
	if p.match(token.LEFT_BRACE) {
		return ast.NewBlock(p.block())
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() ast.Stmt {
	value := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after value.")
	return ast.NewPrint(value)
}

func (p *Parser) block() []ast.Stmt {
	statements := make([]ast.Stmt, 0)
	for !p.check(token.RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	p.consume(token.RIGHT_BRACE, "Expect '}' after block.")
	return statements
}

func (p *Parser) expressionStatement() ast.Stmt {
	expr := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after value.")
	return ast.NewExpression(expr)
}

func (p *Parser) varDeclaration() ast.Stmt {
	name := p.consume(token.IDENTIFIER, "Expect variable name.")

	var initializer ast.Expr = nil

	if p.match(token.EQUAL) {
		initializer = p.expression()
	}

	p.consume(token.SEMICOLON, "Expect ';' after variable declaration.")
	return ast.NewVar(name, initializer)
}

func (p *Parser) advance() *token.Token {
	if !p.isAtEnd() {
		p.Current++
	}
	return p.previous()
}

func (p *Parser) check(t token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) consume(tokenType token.Type, message string) *token.Token {
	if p.check(tokenType) {
		return p.advance()
	}

	loxerror.ParseError(p.peek(), message)
	return nil
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) match(tokenTypes ...token.Type) bool {
	for _, t := range tokenTypes {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) peek() *token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() *token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) synchronise() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case token.CLASS:
			fallthrough
		case token.FUN:
			fallthrough
		case token.VAR:
			fallthrough
		case token.FOR:
			fallthrough
		case token.IF:
			fallthrough
		case token.WHILE:
			fallthrough
		case token.PRINT:
			fallthrough
		case token.RETURN:
			return
		}

		p.advance()
	}
}
