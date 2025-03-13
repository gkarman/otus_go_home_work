package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	tasksCh := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	tasksResultCh := make(chan error)
	stopWorkerFlagCh := make(chan struct{})
	var once sync.Once
	go func() {
		taskResultMonitor(tasksResultCh, stopWorkerFlagCh, m, &once)
	}()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				println("goroutine stopped")
				wg.Done()
			}()
			worker(tasksCh, tasksResultCh, stopWorkerFlagCh)
		}()
	}
	wg.Wait()
	close(tasksResultCh)
	return nil
}

func worker(tasksCh <-chan Task, tasksResultCh chan<- error, stopWorkerFlagCh <-chan struct{}) {
	for {
		select {
		case <-stopWorkerFlagCh:
			return
		case task, ok := <-tasksCh:
			if !ok {
				return
			}
			result := task()
			tasksResultCh <- result
		}
	}
}

func taskResultMonitor(tasksResultCh <-chan error, stopWorkerFlagCh chan<- struct{}, maxErrCount int, once *sync.Once) {
	ErrCount := 0
	for result := range tasksResultCh {
		if result != nil {
			ErrCount++
		}
		if ErrCount >= maxErrCount {
			once.Do(func() { close(stopWorkerFlagCh) })
		}
	}
}
