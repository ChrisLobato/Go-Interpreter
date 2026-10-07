package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// define possible token types as constants
const (
	ILLEGAL = "ILLEGAL" //token we dont know about or cant identify
	EOF     = "EOF"     // end of file to tell parser to stop

	// Identifiers and literals
	IDENT = "IDENT" // add, function names, x, y
	INT   = "INT"   // 123

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// a go dictionary that maps the literals to token types
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Identifier lookup
// comma ok is included because Go naturally returns two things on lookup, the value and a bool if the key existed
func LookupIdent(ident string) TokenType {
	//left of semi is initialization, right side is the actual boolean check for ok
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
