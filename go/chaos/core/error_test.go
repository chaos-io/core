package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewErrorFrom(t *testing.T) {
	err1 := NewErrorFrom(500, "this is an error")
	assert.Error(t, err1)
}
