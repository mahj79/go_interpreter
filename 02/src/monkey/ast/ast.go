package ast

import (
	"monkey/token"
	"bytes"	
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type ReturnStatement struct {
	Token token.Token //the 'return' token
	ReturnValue Expression
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (ex *ExpressionStatement) TokenLiteral() string {return es.Token.Literal}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (p *Program) tokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}


func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	
	out.WriteString(";")
 	
	return out.String()
 }

 func (rs *ReturnStatement) String() string {
 	var out bytes.Buffer
 	
	out.WriteString(rs.TokenLiteral() + " ")
 	if rs.ReturnValue != nil {
 		out.WriteString(rs.ReturnValue.String())
 	}
 
	out.WriteString(";")
 
	return out.String()
 }

 func (es *ExpressionStatement) String() string {
 	if es.Expression != nil {
 		return es.Expression.String()
 	}
 	return ""
 }

 func (i *Identifier) String() string { return i.Value }