package main

import (
	"tiny-gopiler/helper"
	"tiny-gopiler/internal"
)

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
	rawCode = rawCode + " "
	runes := []rune(rawCode)
	tokens := []internal.Token{}
	tempString := ""
	for i := 0; i < len(runes); i++ {
		ch := string([]rune(runes)[i])
		if ch == " " {
			if helper.IsKeyword(tempString) {
				token := internal.Token{
					Type:  "Keyword",
					Value: tempString,
				}
				tokens = append(tokens, token)
			} else if helper.IsOperand(tempString) {
				token := internal.Token{
					Type:  "Operand",
					Value: tempString,
				}
				tokens = append(tokens, token)
			} else if helper.IsNumber(tempString) {
				token := internal.Token{
					Type:  "Number",
					Value: tempString,
				}
				tokens = append(tokens, token)
			} else if helper.IsIdentifier(tempString) {
				token := internal.Token{
					Type:  "Identifier",
					Value: tempString,
				}
				tokens = append(tokens, token)
			}
			tempString = ""
		} else {
			tempString = tempString + ch
		}
	}
	return tokens
}

func (p *Parser) Parse() {

}
