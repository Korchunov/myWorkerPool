package myWorker

import (
	"sync"
	"time"
)

type Job struct {
	job string
}

type Result struct {
	StatusCode   int
	Url          string
	ResponseTime time.Duration
	err          error
}

type Pool struct {
	worker       *worker
	workersCount int
	jobs         chan Job
	results      chan Result
	wg           *sync.WaitGroup
	stopped      bool
}

func NewPool(workerCount int, result Result) *Pool {
	return &Pool{
		worker:       NewWorkerV2(),
		workersCount: workerCount,
		jobs:         make(chan Job),
		results:      make(chan Result),
		wg:           new(sync.WaitGroup),
	}
}

func (p *Pool) PushJobsChan(j Job) {
	for {
		p.jobs <- j
		p.wg.Add(1)
	}
}

/*func (p *Pool) PushResultsChan(r Result) {
		p.results <- r
		p.wg.Done()
}*/
