package task_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringIntMap_Add_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	mCopy := m.Copy()

	assert.Equal(t, map[string]int{"key1": 1, "key2": 2, "key3": 3}, mCopy)
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	m.Remove("key2")

	mCopy := m.Copy()

	assert.Equal(t, map[string]int{"key1": 1, "key3": 3}, mCopy)
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	assert.True(t, m.Exists("key3"))
	assert.False(t, m.Exists("abc"))
}

func TestStringIntMap_Get(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	v, ok := m.Get("key3")
	assert.Equal(t, 3, v)
	assert.True(t, ok)

	v, ok = m.Get("abc")
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}
