package main

import (
	"tiny-gopiler/helper"
	"tiny-gopiler/internal"
)

type ParserI interface {
	Tokenize(string) []internal.Token
	Parse([]internal.Token) internal.AST
	TransformTokenToNode(token internal.Token) internal.Node
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
					Type:  "NumberLiteral",
					Value: tempString,
				}
				tokens = append(tokens, token)
			} else if helper.IsIdentifier(tempString) {
				token := internal.Token{
					Type:  "Identifier",
					Value: tempString,
				}
				tokens = append(tokens, token)
			} else if helper.IsEndExpression(tempString) {
				token := internal.Token{
					Type:  "EndExpression",
					Value: ";",
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

func (p *Parser) Parse(tokens []internal.Token) internal.AST {
	ast := internal.AST{
		Type: "Program",
		Body: []internal.Node{},
	}
	for _, token := range tokens {
		node := p.TransformTokenToNode(token)
		ast.Body = append(ast.Body, node)
	}
	return ast
}

func (p *Parser) TransformTokenToNode(token internal.Token) internal.Node {
	if token.Type == "Keyword" {
		if token.Value == "let" || token.Value == "var" || token.Value == "const" {
			return internal.Node{
				Type:   "DeclarationExpression",
				Params: []internal.Node{},
			}
		}
	} else if token.Type == "Identifier" {

	} else if token.Type == "Operand" {

	} else if token.Type == "Number" {

	} else {
		return internal.Node{}
	}
	return internal.Node{}
}
