package task_9

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNumberConveyor(t *testing.T) {
	in := make(chan uint8)
	n := 100

	want := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		want = append(want, math.Pow(float64(i), 3))
	}

	got := make([]float64, 0, n)

	go func() {
		defer close(in)
		for i := 0; i < n; i++ {
			in <- uint8(i)
		}
	}()

	for v := range NumberConveyor(in) {
		got = append(got, v)
	}

	assert.Equal(t, want, got)
}
