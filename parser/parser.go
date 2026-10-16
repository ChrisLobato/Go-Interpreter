package parser

import (
	"fmt"
	"jade/ast"
	"jade/lexer"
	"jade/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	//init parser with passed in lexer and empty error array
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens so cur token and peek token are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// appends error msg to errors array if the peekToken was not what was expected
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

// advance tokens checks what to do next and call another parsing function or throw an error
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil // i.e. if the next token is an identifier, move onto next token and return true, skip if clause if type matches
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil //i.e. if the next token is an =, move onto next token and return true, skip if clause if type matches
	}

	//TODO skipping expressions until we have to use semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		//while we cur tok isnt a semicolon keep reading skips over expression for now will make test pass
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t //check type of the cur tok same as passed in type
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t // check type of next tok same as passed in type
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}

}
