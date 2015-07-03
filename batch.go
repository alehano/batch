/*
Package batch runs many funcs asynchronously with multiple workers
You have to call Start() in the beginning and Close() in the end.
*/
package batch

import (
	"sync"
	"time"
)

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

func (sb batch) Start() {
	for i := 0; i < sb.workers; i++ {
		go func() {
			for job := range sb.jobChan {
				err := job()
				if err != nil {
					sb.errFn(err)
				}
				sb.wg.Done()
			}
		}()
	}
}

func (sb batch) Add(fn func() error) {
	sb.wg.Add(1)
	sb.jobChan <- fn
	time.Sleep(time.Millisecond * 1)
}

func (sb batch) Close() {
	sb.wg.Wait()
	close(sb.jobChan)
}

func (sb batch) ForceClose() {
	close(sb.jobChan)
}
