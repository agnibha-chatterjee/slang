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

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
