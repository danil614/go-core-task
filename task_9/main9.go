package task_9

import (
	"fmt"
	"math"
	"sync"
)

func NumberConveyor(in <-chan uint8) <-chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for v := range in {
			out <- math.Pow(float64(v), 3)
		}
	}()

	return out
}

func Main() {
	fmt.Println("\nTask 9")
	in := make(chan uint8)
	out := NumberConveyor(in)

	wg := sync.WaitGroup{}
	wg.Add(2)

	n := 100
	go func() {
		defer wg.Done()
		defer close(in)
		for i := 0; i < n; i++ {
			in <- uint8(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			v := <-out
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
