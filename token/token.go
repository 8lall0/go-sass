package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	SEMICOLON = ";"
	COLONS    = ":"

	LBRACE = "{"
	RBRACE = "}"

	SELECTOR = "SELECTOR"
	PROPERTY = "PROPERTY"

	VARIABLE  = "VARIABLE"
	ASSIGN    = "="
	IDENT     = "IDENT"
	COLOR_HEX = "COLOR HEX"

	// SASS stuff
	ATRULE = "AT-RULE"
)

func LookupIdent(ident string) TokenType {
	return IDENT
}
