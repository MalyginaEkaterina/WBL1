package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestQsort(t *testing.T) {
	testQsortOn(t, []int{})
	testQsortOn(t, []int{1})
	testQsortOn(t, []int{1, 2})
	testQsortOn(t, []int{2, 2})
	testQsortOn(t, []int{2, 1})
	testQsortOn(t, []int{1, 2, 3})
	testQsortOn(t, []int{2, 1, 3})
	testQsortOn(t, []int{3, 2, 1})
	testQsortOn(t, []int{3, 1, 2})
	testQsortOn(t, []int{5, 8, 9, 15, 8, 16, 12, 6, 10, 20, 3, 9, 11, 1})
}

func testQsortOn(t *testing.T, a []int) {
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)
	qsort(a)
	assert.Equal(t, a, b)
}
