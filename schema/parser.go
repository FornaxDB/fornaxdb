package schema

import (
	"github.com/FornaxDB/fornaxdb/errors"
	"github.com/FornaxDB/fornaxdb/logger"
)

type Parser struct {
	schema *Schema
	tokenizer *Tokenizer
	currToken *Token
	prevToken *Token
	logger *logger.Logger
}

func NewParser() Parser {
	l := logger.New()
	return Parser{
		schema: nil,
		tokenizer: nil,
		currToken: nil,
		prevToken: nil,
		logger: &l,
	}
}

func (p *Parser) next() {
	p.prevToken = p.currToken
	p.currToken = p.tokenizer.NextToken()
}

func (p *Parser) parseType() {
	p.logger.Trace("Bla", nil)
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
	p.next()

	for p.currToken.Type != EOF_TOKEN {
		switch p.currToken.Type {
		case TYPE_TOKEN:
			p.parseType()
		case RELATION_TOKEN:
			p.parseRelation()
		default:
			return nil, errors.InvalidToken.New("Token should be type or relation")
		}
		p.next()
	}
	
	return p.schema, nil
}