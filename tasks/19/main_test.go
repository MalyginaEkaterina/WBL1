package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTurnString(t *testing.T) {
	assert.Equal(t, "абырвалг", turnString("главрыба"))
	assert.Equal(t, "", turnString(""))
	assert.Equal(t, "а", turnString("а"))
	assert.Equal(t, "а3б", turnString("б3а"))
	assert.Equal(t, "fd", turnString("df"))
	assert.Equal(t, "ed cba", turnString("abc de"))
}
