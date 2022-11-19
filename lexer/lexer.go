package lexer

import (
	"go-sass/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case ' ':
		l.skipWhitespace()
		tok = newToken(token.WHITESPACE, ' ')
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LSQUARE, l.ch)
	case ']':
		tok = newToken(token.RSQUARE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ':':
		tok = newToken(token.COLONS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '.':
		tok = newToken(token.PERIOD, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '-':
		tok = newToken(token.SUB, l.ch)
	case '%':
		tok = newToken(token.MUL, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)
	case '@':
		tok = token.LookupAtRule(l.readStringIdentifier())
	case '$':
		identifier := l.readStringIdentifier()
		tok.Type = token.LookupVariable(identifier)
		tok.Literal = identifier
	case '#':
		tok = newToken(token.HASH, l.ch)
	case 0:
		tok = newToken(token.EOF, 0)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readStringIdentifier()
			tok.Type = token.STRING
		} else if isDigit(l.ch) {
			tok.Literal = l.readDigitIdentifier()
			tok.Type = token.NUMBER
		} else {
			// UHMMM....
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	if tok.Type != token.ILLEGAL && tok.Type != token.EOF {
		l.readChar()
	}

	return tok
}

func (l *Lexer) readStringIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readDigitIdentifier() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
