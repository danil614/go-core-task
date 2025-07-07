package task_3

import (
	"fmt"
	"maps"
)

type StringIntMap interface {
	Add(key string, value int)
	Remove(key string)
	Copy() map[string]int
	Exists(key string) bool
	Get(key string) (int, bool)
}

type stringIntMap struct {
	data map[string]int
}

func NewStringIntMap() StringIntMap {
	return &stringIntMap{data: make(map[string]int)}
}

// Add добавляет новую пару "ключ-значение" в карту.
func (m *stringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove удаляет элемент по ключу из карты.
func (m *stringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает новую карту, содержащую все элементы текущей карты.
func (m *stringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int, len(m.data))
	maps.Copy(copyMap, m.data)
	return copyMap
}

// Exists проверяет, существует ли ключ в карте.
func (m *stringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

// Get возвращает значение по ключу и булевый флаг, указывающий на успешность операции.
func (m *stringIntMap) Get(key string) (int, bool) {
	v, ok := m.data[key]
	return v, ok
}

func Main() {
	m := NewStringIntMap()
	m.Add("key", 10)
	m.Add("key2", 22)
	fmt.Println(m.Get("key2"))

	m.Remove("key2")
	fmt.Println(m.Exists("key2"))

	m.Add("key3", 33)
	m2 := m.Copy()
	m.Remove("key3")
	fmt.Println(m2["key3"])
}
