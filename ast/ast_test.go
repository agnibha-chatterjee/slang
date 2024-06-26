package ast

import (
	"slang/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "a"}, Value: "a"},
			},
		},
	}

	if program.String() != "let x = a;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
