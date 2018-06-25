package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EFO"

	// identifer + literal
	IDENT = "IDENT" // add, foobar, x, y, ....
	INT   = "INT"

	// calcurator
	ASSIGN = "="
	PLUS   = "+"

	// delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
