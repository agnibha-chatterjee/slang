package parser

import (
	"slang/ast"
	"slang/lexer"
	"slang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // points to current token
	peekToken token.Token // points to next token

	/*
		Think of a single line only containing 5;.
		Then curToken is a token.INT and we need peekToken to decide whether we are at the end of the line or if we are at just the start of an arithmetic expression.
	*/
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
