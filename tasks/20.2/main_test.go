package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTurnWords(t *testing.T) {
	assert.Equal(t, "", turnWords(""))
	assert.Equal(t, "sun dog снег", turnWords("снег dog sun"))
	assert.Equal(t, "snow ", turnWords(" snow"))
	assert.Equal(t, " snow   cat", turnWords("cat   snow "))
	assert.Equal(t, "a, b c", turnWords("c b a,"))
	assert.Equal(t, "   ", turnWords("   "))
}
