package schema

import (
	"github.com/FornaxDB/fornaxdb/schema/util"
)

type TokenType string

const (
	//single character tokens
	LEFT_BRACE_TOKEN  = "{"
	RIGHT_BRACE_TOKEN = "}"
	COMMA_TOKEN       = ","
	COLON_TOKEN       = ":"
	SEMICOLON_TOKEN   = ";"
	//keywords
	TYPE_TOKEN     = "type"
	RELATION_TOKEN = "relation"
	// __src is a special keyword used to indicate the source of a token
	__SRC_TOKEN = "__src"
	// __des is a special keyword used to indicate the destination of a token
	__DES_TOKEN = "__des"
	//literal types
	STRING_TOKEN = "string"
	INT_TOKEN    = "int"
	FLOAT_TOKEN  = "float"
	BOOL_TOKEN   = "boolean"
	//operators and delimiters
	REQUIRED_TOKEN = "!"
	OPTIONAL_TOKEN = "?"
	LIST_TOKEN     = "[]"
	UNION_TOKEN    = "|"
	//other
	EOF_TOKEN     = "EOF"
	IDENT_TOKEN   = "ident"
	ILLEGAL_TOKEN = "illegal"
)

// Token struct
type Token struct {
	Type    TokenType
	Literal string
}

// Tokenizer struct
type Tokenizer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// NewTokenizer returns a new Tokenizer instance
func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input:        input,
		position:     0,
		readPosition: 0,
	}
}

//readChar reads the next character from the input
func (t *Tokenizer) readChar() {
	if t.readPosition >= len(t.input) {
		t.ch = 0
	} else {
		t.ch = t.input[t.readPosition]
	}
	t.position = t.readPosition
	t.readPosition++
}

// readIdentifier reads an identifier from the input
func (t *Tokenizer) readIdentifier() string {
	position := t.position
	for util.IsLetter(t.ch) {
		t.readChar()
	}
	return t.input[position:t.position]
}

// readNumber reads a number from the input
func (t *Tokenizer) readNumber() string {
	position := t.position
	for util.IsDigit(t.ch) {
		t.readChar()
	}
	return t.input[position:t.position]
}

// readOperator reads an operator from the input
func (t *Tokenizer) readOperator() string {
	position := t.position
	for util.IsOperator(t.ch) {
		t.readChar()
	}
	return t.input[position:t.position]
}

// skipWhitespace skips whitespace in the input
func (t *Tokenizer) skipWhitespace() {
	for t.ch == ' ' || t.ch == '\t' || t.ch == '\n' || t.ch == '\r' {
		t.readChar()
	}
}

// LookupIdent returns the token type for an identifier
func LookupIdent(ident string) TokenType {
	if ident == "type" {
		return TYPE_TOKEN
	}
	if ident == "relation" {
		return RELATION_TOKEN
	}
	if ident == "string" {
		return STRING_TOKEN
	}
	if ident == "int" {
		return INT_TOKEN
	}
	if ident == "float" {
		return FLOAT_TOKEN
	}
	if ident == "bool" {
		return BOOL_TOKEN
	}
	if ident == "__src" {
		return __SRC_TOKEN
	}
	if ident == "__des" {
		return __DES_TOKEN
	}
	return IDENT_TOKEN
}

// NextToken returns the next token in the input
func (t *Tokenizer) NextToken() *Token {
	var tok Token

	// skip whitespace
	t.skipWhitespace()

	// get the next character
	ch := t.ch
	t.readChar()

	switch ch {
	case '{':
		tok = Token{
			Type:    LEFT_BRACE_TOKEN,
			Literal: string(ch),
		}
	case '}':
		tok = Token{
			Type:    RIGHT_BRACE_TOKEN,
			Literal: string(ch),
		}
	case ',':
		tok = Token{
			Type:    COMMA_TOKEN,
			Literal: string(ch),
		}
	case ':':
		tok = Token{
			Type:    COLON_TOKEN,
			Literal: string(ch),
		}
	case ';':
		tok = Token{
			Type:    SEMICOLON_TOKEN,
			Literal: string(ch),
		}
	case '!':
		tok = Token{
			Type:    REQUIRED_TOKEN,
			Literal: t.readOperator(),
		}
	case '?':
		tok = Token{
			Type:    OPTIONAL_TOKEN,
			Literal: t.readOperator(),
		}
	case '[':
		tok = Token{
			Type:    LIST_TOKEN,
			Literal: t.readOperator(),
		}
	case '|':
		tok = Token{
			Type:    UNION_TOKEN,
			Literal: t.readOperator(),
		}
	case 'i':
		tok.Type = INT_TOKEN
		tok.Literal = t.readIdentifier()
	case 'f':
		tok.Type = FLOAT_TOKEN
		tok.Literal = t.readIdentifier()
	case 'b':
		tok.Type = BOOL_TOKEN
		tok.Literal = t.readIdentifier()
	case 't':
		tok.Type = TYPE_TOKEN
		tok.Literal = t.readIdentifier()
	case 'r':
		tok.Type = RELATION_TOKEN
		tok.Literal = t.readIdentifier()
	default:
		if util.IsLetter(ch) {
			tok.Literal = t.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return &tok
		}
		if util.IsDigit(ch) {
			tok.Type = INT_TOKEN
			tok.Literal = t.readNumber()
			return &tok
		}
		tok = Token{
			Type:    ILLEGAL_TOKEN,
			Literal: string(ch),
		}
	}
	return &tok
}