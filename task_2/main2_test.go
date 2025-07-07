package task_2

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetRandomNumSlice(t *testing.T) {
	assert.Equal(t, 10, len(getRandomNumSlice()))
}

func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, []int{2, 4, 6, 8, 10}, sliceExample(originalSlice))
}

func TestAddElements(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100}, addElements(originalSlice, 100))
}

func TestCopySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copySlice := copySlice(originalSlice)

	assert.Equal(t, originalSlice, copySlice)
	assert.NotEqual(t, reflect.ValueOf(originalSlice).Pointer(), reflect.ValueOf(copySlice).Pointer())
	originalSlice[2] = 9999
	assert.NotEqual(t, originalSlice, copySlice)
}

func TestRemoveElement(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Panics(t, func() { removeElement(originalSlice, -1) })
	assert.Panics(t, func() { removeElement(originalSlice, 85) })

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, removeElement(originalSlice, 9))
}
