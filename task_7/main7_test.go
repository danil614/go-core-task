package task_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelsInOne(t *testing.T) {
	n := 1000
	channels := make([]chan int, n)

	for i := 0; i < n; i++ {
		channels[i] = make(chan int)

		go func(channel chan int) {
			defer close(channel)
			for i := 0; i < n; i++ {
				channel <- i + 1
			}
		}(channels[i])
	}

	out := ChannelsInOne(channels)
	arr := make([]int, 0, n*n)

	for v := range out {
		arr = append(arr, v)
	}

	assert.Equal(t, n*n, len(arr))
}

func TestChannelsInOne_Nil(t *testing.T) {
	channels := make([]chan int, 0)
	out := ChannelsInOne(channels)
	_, ok := <-out
	assert.False(t, ok)
}
