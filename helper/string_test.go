package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input  string
	Output bool
}

func TestIsKeyword(t *testing.T) {
	testTableDriven := []TestCase{
		TestCase{
			Input:  "a",
			Output: false,
		},
		TestCase{
			Input:  "9",
			Output: false,
		},
		TestCase{
			Input:  "-",
			Output: false,
		},
		TestCase{
			Input:  "let",
			Output: true,
		},
	}
	for _, testCase := range testTableDriven {
		assert.Equal(t, IsKeyword(testCase.Input), testCase.Output)
	}
}

func TestIsNumber(t *testing.T) {
	testTableDriven := []TestCase{
		TestCase{
			Input:  "a",
			Output: false,
		},
		TestCase{
			Input:  "9",
			Output: true,
		},
		TestCase{
			Input:  "-",
			Output: false,
		},
		TestCase{
			Input:  "let",
			Output: false,
		},
	}
	for _, testCase := range testTableDriven {
		assert.Equal(t, IsNumber(testCase.Input), testCase.Output)
	}
}

func TestIsIdentifier(t *testing.T) {
	testTableDriven := []TestCase{
		TestCase{
			Input:  "a",
			Output: true,
		},
		TestCase{
			Input:  "9",
			Output: false,
		},
		TestCase{
			Input:  "-",
			Output: false,
		},
		TestCase{
			Input:  "let",
			Output: false,
		},
	}
	for _, testCase := range testTableDriven {
		assert.Equal(t, IsIdentifier(testCase.Input), testCase.Output)
	}
}

func TestIsOperand(t *testing.T) {
	testTableDriven := []TestCase{
		TestCase{
			Input:  "a",
			Output: false,
		},
		TestCase{
			Input:  "9",
			Output: false,
		},
		TestCase{
			Input:  "-",
			Output: true,
		},
		TestCase{
			Input:  "let",
			Output: false,
		},
	}
	for _, testCase := range testTableDriven {
		assert.Equal(t, IsOperand(testCase.Input), testCase.Output)
	}
}
