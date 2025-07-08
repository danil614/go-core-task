package task_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	assert.Equal(t, []string{"apple", "cherry", "43", "lead", "gno1"}, subSlice(slice1, slice2))
}

func TestSubSlice_Equal(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}

	assert.Equal(t, []string{}, subSlice(slice1, slice2))
}

func TestSubSlice_Diff(t *testing.T) {
	slice1 := []string{"date", "43", "lead", "gno1"}
	slice2 := []string{"apple", "banana", "cherry"}

	assert.Equal(t, []string{"date", "43", "lead", "gno1"}, subSlice(slice1, slice2))
}

func TestSubSlice_EmptySlices(t *testing.T) {
	assert.Equal(t, []string{}, subSlice([]string{}, []string{}))
}
