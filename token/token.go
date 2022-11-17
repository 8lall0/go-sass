package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
	AT
	DOLLARSIGN
	WHITESPACE
	HASH

	ADD
	SUB
	MUL
	DIV
	SEMICOLON
	COLONS
	PERIOD
	COMMA

	LBRACE
	RBRACE
	LSQUARE
	RSQUARE
	LPAREN
	RPAREN

	ASSIGN

	NUMBER
	STRING
)
