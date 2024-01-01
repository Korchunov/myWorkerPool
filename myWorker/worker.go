package myWorker

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func NewWorkerV2() *worker {
	return &worker{
		client: &http.Client{},
	}
}

func (w *worker) WorkProcess(ctx context.Context, p *Pool) {
	ctx, _ = context.WithTimeout(ctx, time.Second*10)
	select {
	case <-ctx.Done():
		//defer close(p.results)
		fmt.Println("Context done with timeout")
	case j := <-p.jobs:
		req, err := http.NewRequest("GET", j.job, bytes.NewBuffer(nil))
		if err != nil {
			fmt.Printf("can't creat a new request: [%v]", err)
		}
		t := time.Now()
		resp, err := w.client.Do(req)
		if err != nil {
			fmt.Printf("can't connention by endpoint: [%v]", err)
		}
		responsTime := time.Since(t)
		p.results <- Result{resp.StatusCode, j.job, responsTime, nil}
		p.wg.Done()
	}
}
