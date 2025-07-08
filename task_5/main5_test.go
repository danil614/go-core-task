package task_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectSlicesDiff(t *testing.T) {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}

	ok, intersect := intersectSlices(slice1, slice2)

	assert.True(t, ok)
	assert.Equal(t, []int{3, 64}, intersect)
}

func TestIntersectSlicesEqual(t *testing.T) {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{65, 3, 58, 678, 64}

	ok, intersect := intersectSlices(slice1, slice2)

	assert.True(t, ok)
	assert.Equal(t, []int{65, 3, 58, 678, 64}, intersect)
}

func TestIntersectSlicesNo(t *testing.T) {
	slice1 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ok, intersect := intersectSlices(slice1, slice2)

	assert.False(t, ok)
	assert.Equal(t, []int{}, intersect)
}

func TestIntersectSlicesEmpty(t *testing.T) {
	slice1 := make([]int, 0)
	slice2 := []int{1, 2, 3}

	ok, intersect := intersectSlices(slice1, slice2)

	assert.False(t, ok)
	assert.Equal(t, []int{}, intersect)
}

func TestIntersectSlicesWithDuplicates(t *testing.T) {
	slice1 := []int{1, 1, 2, 2, 3}
	slice2 := []int{2, 2, 3, 3, 4}

	ok, intersect := intersectSlices(slice1, slice2)

	assert.True(t, ok)
	assert.Equal(t, []int{2, 3}, intersect)
}
