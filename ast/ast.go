package ast

import "jade/token"

//every node in our AST has to implement this interface
// has to provide a TokenLiteral method that returns the literal value of the token associated with it
type Node interface {
	TokenLiteral() string
}

// defined by Thorsten as something that does not produce a value
type Statement interface {
	Node
	statementNode() //good for error catching by signifying to the compiler the difference between statement and expression
}

// defined by Thorsten as something that does produce a value
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement // every program is just a series of statements i.e. our root node
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//Concrete type that is tied to Statement
type LetStatement struct {
	Token token.Token // will have the token.LET token struct
	Name  *Identifier
	Value Expression
}

//the functions with recievers that tie LetStatement to Statement interface
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// a concerete type that technically implements Expression as it has its own expressionNode funcntion and token Literal
type Identifier struct {
	Token token.Token // will hold the token.IDENT token struct
	Value string
}

//Functions that tie identifier to Expression
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
