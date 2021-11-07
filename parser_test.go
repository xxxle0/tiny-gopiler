package main

import (
	"log"
	"testing"
	"tiny-gopiler/internal"

	"github.com/stretchr/testify/assert"
)

func TestParserTokenize(t *testing.T) {
	parser := InitParser()
	stringTest := "let a = 0"
	tokens := parser.Tokenize(stringTest)
	expectTokens := []internal.Token{
		internal.Token{
			Type:  "Keyword",
			Value: "let",
		},
		internal.Token{
			Type:  "Identifier",
			Value: "a",
		},
		internal.Token{
			Type:  "Operand",
			Value: "=",
		},
		internal.Token{
			Type:  "Number",
			Value: "0",
		},
	}
	assert.Equal(t, tokens, expectTokens)
}

func TestParserParse(t *testing.T) {
	stringTest := "let a = 0"
	parser := InitParser()
	tokens := parser.Tokenize(stringTest)
	ast := parser.Parse(tokens)
	log.Println(ast)
	assert.Equal(t, ast, true)
}
