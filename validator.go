package main

type ValidatorI interface {
}

type Validator struct {
}

func InitValidator() ValidatorI {
	return &Validator{}
}
