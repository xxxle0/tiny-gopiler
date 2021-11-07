package main

import (
	"log"
)

func main() {
	expressionInput := "let x = 0;"
	parser := InitParser()
	tokens := parser.Tokenize(expressionInput)
	ast := parser.Parse(tokens)
	log.Println(tokens, ast)
}
