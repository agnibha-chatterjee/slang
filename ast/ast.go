package ast

import "slang/token"

type Node interface {
	TokenType() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

type LetStatement struct {
	Token token.Token // the LET token
	Name  *Identifier
	Value Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenType()
	} else {
		return ""
	}
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
