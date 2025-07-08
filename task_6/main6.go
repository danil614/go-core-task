package task_6

import "math/rand"

func getRandomNumChan(n int) <-chan int {
	if n <= 0 {
		return nil
	}

	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			randNum := rand.Intn(1001) - 500
			ch <- randNum
		}
	}()

	return ch
}
