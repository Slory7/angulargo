package worker

import (
	"sync"
)

type WkPool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *WkPool {
	if size <= 0 {
		size = 1
	}
	return &WkPool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *WkPool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *WkPool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *WkPool) Wait() {
	p.wg.Wait()
}
