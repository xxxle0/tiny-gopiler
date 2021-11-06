package main

type ParserI interface {
	Tokenize()
	Parse()
}

type Parser struct {
}

func InitParser() ParserI {
	return &Parser{}
}

func (p *Parser) Tokenize() {

}

func (p *Parser) Parse() {

}
