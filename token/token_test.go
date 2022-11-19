package token

import (
	"testing"
)

func TestLookupAtRule(t *testing.T) {
	input := []string{"function", "media", "pippo"}
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{expectedType: FUNCTION, expectedLiteral: "function"},
		{expectedType: MEDIA, expectedLiteral: "media"},
		{expectedType: IDENTIFIER, expectedLiteral: "pippo"},
	}
	for i, tt := range tests {
		tok := LookupAtRule(input[i])

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %d, got: %d", i, int(tt.expectedType), int(tok.Type))
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
