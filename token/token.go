package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF     = "EOF"     // Tells the interpreter when to stop
	ILLEGAL = "ILLEGAL" // For ILLEGAL tokens that we don't understand

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
