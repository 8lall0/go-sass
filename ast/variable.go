package ast

import "go-sass/token"

type VariableStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *VariableStatement) statementNode() {

}

func (vs *VariableStatement) TokenLiteral() string {
	return vs.Token.Literal
}
