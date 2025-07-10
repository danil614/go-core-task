package task_8

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type waitGroup struct {
	addCh  chan int      // сюда приходят Add / Done
	doneCh chan struct{} // сигнал о том, что counter опустился до 0
}

func NewWaitGroup() WaitGroup {
	wg := &waitGroup{
		addCh:  make(chan int),
		doneCh: make(chan struct{}),
	}
	go wg.loop()
	return wg
}

func (wg *waitGroup) loop() {
	var counter int
	for delta := range wg.addCh {
		counter += delta
		if counter < 0 {
			panic("WaitGroup: negative counter")
		}
		if counter == 0 {
			// Разбудить всех, кто ждёт
			close(wg.doneCh)
			// doneCh надо создать заново, чтобы WG можно было использовать ещё раз
			wg.doneCh = make(chan struct{})
		}
	}
}

func (wg *waitGroup) Add(delta int) {
	if delta == 0 {
		panic("WaitGroup: Add with delta 0")
	}
	wg.addCh <- delta
}

func (wg *waitGroup) Done() {
	wg.Add(-1)
}

func (wg *waitGroup) Wait() {
	<-wg.doneCh
}
