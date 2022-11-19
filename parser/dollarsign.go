package parser

import (
	"go-sass/ast"
	"go-sass/token"
)

func (p *Parser) parseDollarSignStatement() *ast.VariableStatement {
	stmt := &ast.VariableStatement{Token: p.curToken}

	for p.expectedPeek(token.WHITESPACE) {
		p.nextToken()
	}

	if !p.expectedPeek(token.STRING) {
		return nil
	}

	// TODO per ora skippiamo tutto
	for !p.expectedPeek(token.SEMICOLON) {
		p.nextToken()
	}

	// TODO capisci cosa fa
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return stmt
}
