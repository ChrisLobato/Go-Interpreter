package parser

import (
	"jade/ast"
	"jade/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i] //current statement to eval
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	//double check that the literal string on the statements token is let
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement) //ok if let statement s could exist as an ast.LetStatement pointer
	//comma ok is similar to when we did it in the map lookup, in this case it serves as a way of asserting that the type conversion from statement -> letstatement worked
	if !ok {
		t.Errorf("s not *ast.LetStatement. got =%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got=%s", name, letStmt.Name.Value) //checks that the value was properly parsed and stored in the Value field of Identifier
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Name not %s. got=%s", name, letStmt.Name) //checks that identifier was correctly linked to the token it contains, i.e. checks the token for the name rather than the value
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return //return if no errors
	}

	t.Errorf("parser has %d errors", len(errors))
	//wild card since we dont really care for the counter that gets init from range in Go
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
