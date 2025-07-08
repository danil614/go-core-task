package task_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRandomNumChan(t *testing.T) {
	n := 100
	ch := getRandomNumChan(n)
	arr := make([]int, 0, n)

	for v := range ch {
		arr = append(arr, v)
	}

	assert.Equal(t, n, len(arr))
}

func TestGetRandomNumChan_Nil(t *testing.T) {
	n := 0
	ch := getRandomNumChan(n)

	assert.Nil(t, ch)
}
