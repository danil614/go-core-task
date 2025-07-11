package task_8

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup_Add_1(t *testing.T) {
	wg := NewWaitGroup()
	count := atomic.Int32{}

	wg.Add(1)
	go func() {
		count.Add(1)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		wg.Done()
	}()

	wg.Wait()

	assert.Equal(t, int32(1), count.Load())
}

func TestWaitGroup_Add_N(t *testing.T) {
	n := 1000
	wg := NewWaitGroup()
	count := atomic.Int32{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count.Add(1)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, int32(n), count.Load())
}

func TestWaitGroup_Add_1_NTimes(t *testing.T) {
	n := 100
	wg := NewWaitGroup()
	count := atomic.Int32{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			count.Add(1)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, int32(n), count.Load())
}

func TestWaitGroup_Add_Negative_NTimes(t *testing.T) {
	wg := NewWaitGroup()
	wgOrig := sync.WaitGroup{}
	count := atomic.Int32{}

	for i := 0; i < 100; i++ {
		wg.Add(2)
		wgOrig.Add(2)
		go func() {
			count.Add(1)
			wgOrig.Done()
		}()
		go func() {
			count.Add(1)
			wgOrig.Done()
		}()
	}

	go func() {
		wgOrig.Wait()
		wg.Add(-200)
	}()

	wg.Wait()

	assert.Equal(t, int32(200), count.Load())
}

func TestWaitGroup_Add_Delta0(t *testing.T) {
	wg := NewWaitGroup()
	assert.Panics(t, func() {
		wg.Add(0)
	})
}

func TestWaitGroup_Add_NegativeDelta(t *testing.T) {
	wg := NewWaitGroup()
	assert.Panics(t, func() {
		wg.Add(-12)
	})
}
