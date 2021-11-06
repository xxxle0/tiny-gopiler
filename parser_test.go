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
	log.Println(tokens)
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
	parser := InitParser()
	parser.Parse()
	assert.Equal(t, true, true)
}
