package main

import "tiny-gopiler/internal"

type ParserI interface {
	Tokenize(rawCode string) []internal.Token
	Parse()
}

type Parser struct {
}

func InitParser() ParserI {
	return &Parser{}
}

func (p *Parser) Tokenize(rawCode string) []internal.Token {
	tokens := []internal.Token{}
	return tokens
}

func (p *Parser) Parse() {

}
