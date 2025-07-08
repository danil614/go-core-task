package task_7

import "sync"

func ChannelsInOne[T any](channels []chan T) <-chan T {
	// Закрытый канал, если входных нет
	if len(channels) == 0 {
		closed := make(chan T)
		close(closed)
		return closed
	}

	out := make(chan T, len(channels))
	wg := sync.WaitGroup{}

	for _, channel := range channels {
		wg.Add(1)
		go func() {
			for v := range channel {
				out <- v
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
