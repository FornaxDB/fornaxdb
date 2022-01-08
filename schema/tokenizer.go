package schema

import (
	"github.com/FornaxDB/fornaxdb/schema/util"
	"strings"
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
	IDENT_TOKEN   = "ident"
	ILLEGAL_TOKEN = "illegal"
	// whitespace
	WHITESPACE_TOKEN = "whitespace"
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
	ch           rune
}


// NewTokenizer returns a new Tokenizer instance
func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input:        strings.TrimSpace(input),
		position:     0,
	}
}

//readChar reads the next character from the input
func (t *Tokenizer) readChar() {
	if t.position <= len(t.input ) {
		c := rune(t.input[t.position])
		t.ch = c
	}
	t.position++
}

//read multiple whitespace characters as one whitespace token
func (t *Tokenizer) readWhitespace()  {
	for util.IsWhitespace(t.ch) {
		t.readChar()
	}
}

// readIdentifier reads an identifier from the input
func (t *Tokenizer) readIdentifier() string {
	position := t.position -1
	for util.IsPartofString(t.ch) {
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

// // skipWhitespace skips whitespace in the input
// func (t *Tokenizer) skipWhitespace() {
// 	for t.ch == ' ' || t.ch == '\t' || t.ch == '\n' || t.ch == '\r' {
// 		t.readChar()
// 	}
// }

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
	if ident == "boolean" {
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
	// t.Trim()

	// get the next character
	t.readChar()
	ch := t.ch

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
			Literal: string(ch),
		}
	case ']':
		tok = Token{
			Type:    LIST_TOKEN,
			Literal: string(ch),
		}
	case '|':
		tok = Token{
			Type:    UNION_TOKEN,
			Literal: t.readOperator(),
		}
	//case for whitespace
	case ' ':
		t.readWhitespace()
	case '\t':
		t.readWhitespace()
	case '\n':
		t.readWhitespace()
	case '\r':
		t.readWhitespace()

	default:
		if util.IsPartofString(ch) {
			tok.Literal = t.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return &tok
		}
		
		tok = Token{
			Type:    ILLEGAL_TOKEN,
			Literal: string(ch),
		}
	}

	return &tok
}

// function to get list of tokens in a string
func Tokenize(input string) []Token {
	tokens := []Token{}
	t := NewTokenizer(input)
	for {
		token := t.NextToken()
		tokens = append(tokens, *token)
		if t.position >= len(t.input) {
			break
		}
	}

	return tokens
}
