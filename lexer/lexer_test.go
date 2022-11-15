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
		{token.VARIABLE, "$variable"},
		{token.ASSIGN, "="},
		{token.COLOR_HEX, "#000"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q, %q", i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
