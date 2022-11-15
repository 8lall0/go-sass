package parser

import (
	"go-sass/ast"
	"go-sass/token"
)

func (p *Parser) parseVariableStatement() *ast.VariableStatement {
	stmt := &ast.VariableStatement{Token: p.curToken}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// TODO per ora skippiamo tutto
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	// In seguito mi aspetterò robe di questo tipo
	//if !p.expectPeek(token.COLOR_HEX) {
	//	return nil
	//}

	return stmt
}
