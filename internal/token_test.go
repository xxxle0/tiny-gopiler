package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenStruct(t *testing.T) {
	token := Token{
		Type:  "Identifier",
		Value: "x",
	}
	assert.Equal(t, token.Type, "Identifier")
	assert.Equal(t, token.Value, "x")
	assert.NotEqual(t, token.Value, "y")
}
