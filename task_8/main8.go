package task_8

import "sync"

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type waitGroup struct {
	doneCh  chan struct{} // сигнал о том, что counter опустился до 0
	mu      sync.Mutex
	counter int
}

func NewWaitGroup() WaitGroup {
	wg := &waitGroup{
		doneCh: make(chan struct{}),
	}
	return wg
}

func (wg *waitGroup) Add(delta int) {
	if delta == 0 {
		panic("WaitGroup: Add with delta 0")
	}

	wg.mu.Lock()

	wg.counter += delta
	if wg.counter < 0 {
		panic("WaitGroup: negative counter")
	}
	if wg.counter == 0 {
		// Разбудить всех, кто ждёт
		close(wg.doneCh)
		// doneCh надо создать заново, чтобы WG можно было использовать ещё раз
		wg.doneCh = make(chan struct{})
	}

	wg.mu.Unlock()
}

func (wg *waitGroup) Done() {
	wg.Add(-1)
}

func (wg *waitGroup) Wait() {
	<-wg.doneCh
}
