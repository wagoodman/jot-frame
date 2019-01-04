package recipe

import (
	"context"
	"github.com/wagoodman/jotframe/pkg/frame"

	"github.com/k0kubun/go-ansi"
	"golang.org/x/sync/semaphore"
)

// Worker is working
type Worker interface {
	Work(*frame.Line)
}

type WorkQueue struct {
	maxConcurrent int64
	queue         []interface{}
}

func NewWorkQueue(maxConcurrent int64) *WorkQueue {
	return &WorkQueue{
		maxConcurrent: maxConcurrent,
	}
}

func (wq *WorkQueue) AddWork(work interface{}) {
	wq.queue = append(wq.queue, work)
}

func (wq *WorkQueue) Work() {
	frames := frame.Factory(frame.Config{
		Lines:         0,
		HasHeader:     false,
		HasFooter:     false,
		TrailOnRemove: true,
	})
	fr := frames[0]
	// worker pool
	ctx := context.TODO()
	sem := semaphore.NewWeighted(wq.maxConcurrent)

	for _, item := range wq.queue {
		worker, _ := item.(Worker)
		sem.Acquire(ctx, 1)
		line, _ := fr.Append()
		jotFunc := func(userFunc func(line *frame.Line), line *frame.Line) {
			defer sem.Release(1)

			userFunc(line)
			fr.Remove(line)
		}
		go jotFunc(worker.Work, line)
	}

	fr.Wait()
	fr.Close()

	ansi.CursorShow()
}