package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllUnique(t *testing.T) {
	assert.Equal(t, true, allUnique("abcdлор"))
	assert.Equal(t, true, allUnique("abcdл орL"))
	assert.Equal(t, false, allUnique("abcdлорЛ"))
	assert.Equal(t, false, allUnique("abcdлорD"))
	assert.Equal(t, false, allUnique("abcdлорa"))
	assert.Equal(t, false, allUnique("abcd лорл"))
}
