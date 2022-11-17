package lexer

import (
	"go-sass/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `$variable = #000;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DOLLARSIGN, "$"},
		{token.STRING, "variable"},
		{token.WHITESPACE, " "},
		{token.ASSIGN, "="},
		{token.WHITESPACE, " "},
		{token.HASH, "#"},
		{token.NUMBER, "000"},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %d, got: %d", i, int(tt.expectedType), int(tok.Type))
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
