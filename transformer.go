package main

type TransformerI interface {
}

type Transformer struct {
}

func InitTransformer() TransformerI {
	return &Transformer{}
}
