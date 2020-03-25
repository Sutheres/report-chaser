package worker

import (
	"context"
	"fmt"
	"github.com/Sutheres/report-chaser/internal/edgar"
)

type worker struct {
	ctx context.Context
	e edgar.Edgar
}

type option func(w *worker)

func NewWorker(ctx context.Context, opts... option) *worker {
	w := &worker{
		ctx: ctx,
	}
	w.WithOptions(opts...)
	return w
}

func (w *worker) WithOptions(opts... option) {
	for _, opt := range opts {
		opt(w)
	}
}

func (w *worker) Start() {
	fmt.Println("starting worker...")
	defer fmt.Println("stopping worker...")
}

func WithSEC(e edgar.Edgar) option {
	return func(w *worker) {
		w.e = e
	}
}