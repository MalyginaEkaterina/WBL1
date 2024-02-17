package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 6, 8, 11, 13, 15, 17, 48, 59, 100, 114, 150}
	assert.Equal(t, -1, binarySearch(nil, 10))
	assert.Equal(t, 5, binarySearch(arr, 11))
	assert.Equal(t, 0, binarySearch(arr, 1))
	assert.Equal(t, 13, binarySearch(arr, 150))
	assert.Equal(t, -1, binarySearch(arr, 16))
	arr2 := []int{1, 5}
	assert.Equal(t, 0, binarySearch(arr2, 1))
	assert.Equal(t, 1, binarySearch(arr2, 5))
	assert.Equal(t, -1, binarySearch(arr2, 0))
	arr3 := []int{3}
	assert.Equal(t, 0, binarySearch(arr3, 3))
}
