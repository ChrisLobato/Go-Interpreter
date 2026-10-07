package lexer

import "jade/token"

// Differentiate with two pointers because position always corresponds to whats in ch, sometimes we might want to know what comes next...
type Lexer struct {
	input        string
	position     int  // current position in input, points to current char
	readPosition int  // current reading position in input, after current char
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	//below instantiation is legal in GO but not C we create and go knows to keep it on heapbc of &
	l := &Lexer{input: input} // note here we are init a Lexer with just the input, this means the other fields get init with zero values in this case ints set to 0
	l.readChar()
	return l
}

// read in the next character and advance pointers in input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // if readposition is at end of string set character to null terminator
	} else {
		l.ch = l.input[l.readPosition] // otherwise read in the char at read position
	}

	l.position = l.readPosition //move over the current pointer to obtained char, position always points to the thing we just read

	l.readPosition += 1 // move over readPosition to next character
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	//depending on last read character create a corresponding token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		//case where the single character isnt a token on its own keyword/identifier
		if isLetter(l.ch) {
			tok.Literal = l.readIdenitfier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			//case is illegal
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // reads in the next character and moves over pointers
	return tok   // returns newly created token
}

// note the type on the left meaning function doesnt stand alone i.e. lexer is the reciever
// lets us do stuff like l.property works similar to *this* or *self* keyword
// read in the current identifier
func (l *Lexer) readIdenitfier() string {
	position := l.position //starting index of our identifier
	for isLetter(l.ch) {
		l.readChar() // for each letter keep moving a long the reading position
	}
	return l.input[position:l.position] //slice input from starting point to the last position that was not a letter (non inclusive)
}

// check if current character is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' //condition to return if a letter
}

// skip whitespace characters to set up positions at the next token to read
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// read in number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check is digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// essentially a constructor for tokens that takes in its type and the character to hold; converts the char to a string literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
