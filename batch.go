/*
Package batch runs many funcs asynchronously with multiple workers
You have to call Start() in the beginning and Close() in the end.
*/
package batch

import "sync"

type batch struct {
	workers int
	jobChan chan func() error
	errFn   func(error)
	wg      *sync.WaitGroup
}

// New creates new Batch instance
func New(workers int, errCallback func(error)) *batch {
	return &batch{
		workers: workers,
		jobChan: make(chan func() error),
		errFn:   errCallback,
		wg:      new(sync.WaitGroup),
	}
}

func (b batch) Start() {
	for i := 0; i < b.workers; i++ {
		go func() {
			for job := range b.jobChan {
				err := job()
				if err != nil {
					b.errFn(err)
				}
				b.wg.Done()
			}
		}()
	}
}

func (b batch) Add(fn func() error) {
	b.wg.Add(1)
	b.jobChan <- fn
}

func (b batch) Close() {
	b.wg.Wait()
	close(b.jobChan)
}

func (b batch) ForceClose() {
	close(b.jobChan)
}
