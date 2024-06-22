package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	INT   = "INT"
	IDENT = "IDENT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	RPAREN = ")"
	LPAREN = "("
	RBRACE = "}"
	LBRACE = "{"

	LET      = "LET"
	FUNCTION = "FUNCTION"
)
