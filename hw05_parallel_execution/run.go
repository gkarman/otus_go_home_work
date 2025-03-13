package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	successRunCh := make(chan bool)
	tasksCh := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	tasksResultCh := make(chan error)
	stopWorkerFlagCh := make(chan struct{})
	go func() {
		taskResultMonitor(tasksResultCh, stopWorkerFlagCh, successRunCh, m)
	}()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(tasksCh, tasksResultCh, stopWorkerFlagCh)
		}()
	}
	wg.Wait()
	close(tasksResultCh)

	success := <-successRunCh
	close(successRunCh)
	if !success {
		return ErrErrorsLimitExceeded
	}
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
			select {
			case <-stopWorkerFlagCh:
				return
			default:
				result := task()
				tasksResultCh <- result
			}
		}
	}
}

func taskResultMonitor(tasksResultCh <-chan error, stopFlagCh chan<- struct{}, successCh chan<- bool, maxErr int) {
	errCount := 0
	success := true
	for result := range tasksResultCh {
		if maxErr <= 0 {
			continue
		}
		if result != nil {
			errCount++
		}
		if errCount == maxErr {
			close(stopFlagCh)
			success = false
		}
	}
	successCh <- success
}
