package main

import (
	"testing"
	"tiny-gopiler/internal"

	"github.com/stretchr/testify/assert"
)

func TestParserTokenize(t *testing.T) {
	parser := InitParser()
	stringTest := "let a = 0 ;"
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
			Type:  "NumberLiteral",
			Value: "0",
		},
		internal.Token{
			Type:  "EndExpression",
			Value: ";",
		},
	}
	assert.Equal(t, expectTokens, tokens)
}

func TestParserParse(t *testing.T) {
	stringTest := "let a = 0 ;"
	parser := InitParser()
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
			Type:  "NumberLiteral",
			Value: "0",
		},
		internal.Token{
			Type:  "EndExpression",
			Value: ";",
		},
	}
	assert.Equal(t, tokens, expectTokens)
	ast := parser.Parse(tokens)
	expectAST := internal.AST{
		Type: "Program",
		Body: []internal.Node{
			internal.Node{
				Type:   "DeclarationExpression",
				Params: []internal.Node{},
			},
			internal.Node{},
			internal.Node{},
			internal.Node{},
			internal.Node{},
		},
	}
	assert.Equal(t, expectAST, ast)
}

func TestTransformTokenToNode(t *testing.T) {
	parser := InitParser()
	token := internal.Token{}
	expectNode := internal.Node{}
	node := parser.TransformTokenToNode(token)
	assert.Equal(t, expectNode, node)
}
