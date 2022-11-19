package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
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

	// AT RULES
	FUNCTION
	MIXIN
	MEDIA
	USE
	INCLUDE
	FORWARD
	IMPORT
	EXTEND
	ERROR
	WARN
	DEBUG
	ATROOT
	IF
	ELSE
	ELSEIF
	EACH
	FOR
	WHILE

	VARIABLE
)

var reservedAtRules = map[string]TokenType{
	"function": FUNCTION,
	"mixin":    MIXIN,
	"media":    MEDIA,
	"use":      USE,
	"include":  INCLUDE,
	"forward":  FORWARD,
	"import":   IMPORT,
	"extend":   EXTEND,
	"error":    ERROR,
	"warn":     WARN,
	"debug":    DEBUG,
	"at-root":  ATROOT,
	"if":       IF,
	"else if":  ELSEIF,
	"else":     ELSE,
	"each":     EACH,
	"for":      FOR,
	"while":    WHILE,
}

func LookupAtRule(identifier string) Token {
	var tok Token
	tok.Literal = identifier

	if tokType, ok := reservedAtRules[identifier]; ok {
		tok.Type = tokType
	} else {
		tok.Type = STRING
	}

	return tok
}

func LookupHash(identifier string) TokenType {
	// TODO check if is identifier or HEX CODE
	return STRING
}

func LookupVariable(identifier string) TokenType {
	// TODO check if is a valid variable
	return ILLEGAL
}
