package schema

import (
	"github.com/FornaxDB/fornaxdb/errors"
)

type Parser struct {
	schema *Schema
	tokenizer *Tokenizer
	currToken *Token
	prevToken *Token
}

func NewParser() Parser {
	return Parser{
		schema: nil,
		tokenizer: nil,
		currToken: nil,
		prevToken: nil,
	}
}

func (p *Parser) next() {
	p.prevToken = p.currToken
	p.currToken = p.tokenizer.NextToken()
}

func (p *Parser) parseType() {

}

func (p *Parser) parseRelation() {

}

func (p *Parser) parseField() {}

func (p *Parser) parseFieldKey() {}

func (p *Parser) parseFieldKeyArguments() {}

func (p *Parser) parseFieldReturnType() {}

func (p *Parser) parseFieldScalarReturnType() {}

func (p *Parser) parseFieldVectorReturnType() {}



func (p *Parser) Parse(input string) (*Schema, error) {
	s := NewSchema()
	t := NewTokenizer(input)
	p.schema = s
	p.tokenizer = t

	for p.currToken.Type != EOF_TOKEN {
		p.next()
		switch p.currToken.Type {
		case TYPE_TOKEN:
			p.parseType()
		case RELATION_TOKEN:
			p.parseRelation()
		default:
			return nil, errors.InvalidToken.New("Token should be type or relation")
		}
	}
	
	return p.schema, nil
}