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

func (p *Parser) Parse() ast.Expr {
	expr := p.expression()
	if loxerror.HadError {
		return nil
	}
	return expr
}

func (p *Parser) expression() ast.Expr {
	return p.equality()
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
	}

	loxerror.ParseError(p.peek(), "Expect expression.")
	return nil
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
