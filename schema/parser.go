package schema

import (
	"github.com/FornaxDB/fornaxdb/errors"
)

type Parser struct {
	schema *Schema
	position int
	tokens []Token
}

func NewParser() Parser {
	s := NewSchema()
	return Parser{
		schema: &s,
		position:  0,
		tokens: nil,
	}
}

func (p *Parser) parseType() (error) {
	t := Type{}
	
	
	
	p.schema.Types[t.Name] = t
	return nil
}

func (p *Parser) parseRelation() (error) {
	r := Relation{}

	p.schema.Relations[r.Name] = r
	return nil
}

func (p *Parser) ParseTokens(tokens []Token) (*Schema, error) {
	// all functions modify that copy of the schema and return a pointer
	p.tokens = tokens
	for {
		if p.position < len(tokens) {
			token := p.tokens[p.position]
			if token.Type == TYPE_TOKEN {
				p.position++
				p.parseType()
			} else if token.Type == RELATION_TOKEN {
				p.position++
				p.parseRelation()
			} else {
				return nil, errors.InvalidToken.New("should start with type or relation")
			}
		} else {
			break
		}
	}

	return p.schema, nil
}