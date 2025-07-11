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
	// При нулевом счётчике канал сразу закрыт, чтобы Wait() немедленно возвращал управление
	ch := make(chan struct{})
	close(ch)
	return &waitGroup{doneCh: ch}
}

func (wg *waitGroup) Add(delta int) {
	if delta == 0 {
		panic("WaitGroup: Add with delta 0")
	}

	wg.mu.Lock()
	defer wg.mu.Unlock()

	// Переход из 0 в >0 - создаём новый (открытый) канал
	if wg.counter == 0 && delta > 0 {
		wg.doneCh = make(chan struct{})
	}

	wg.counter += delta
	if wg.counter < 0 {
		panic("WaitGroup: negative counter")
	}

	if wg.counter == 0 {
		// Разбудить всех, кто ждёт
		close(wg.doneCh)
	}
}

func (wg *waitGroup) Done() {
	wg.Add(-1)
}

func (wg *waitGroup) Wait() {
	// Берём актуальный канал под тем же мьютексом, чтобы не было гонки
	wg.mu.Lock()
	ch := wg.doneCh
	wg.mu.Unlock()

	<-ch
}
