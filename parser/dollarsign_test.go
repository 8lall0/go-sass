package parser

import (
	"go-sass/lexer"
	"testing"
)

func TestVariableStatement(t *testing.T) {

	l := lexer.New(`$variable = #000;`)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}
}
