package service

import (
	"errors"
	"time"
)

var (
	ErrTaskAlreadyStarterd error = errors.New("task already started")
	ErrTaskNotStarted      error = errors.New("task not started")
)

type TaskStopSignal struct{}

type Task struct {
	Name     string
	Func     func()        // Function to be executed
	stop     chan TaskStopSignal
	isRun    bool
	Ticker   *time.Ticker
}

func (t *Task) run() {
	for {
		select {
		case <-t.stop:
			t.Ticker.Stop()
			t.isRun = false
			return
		case <-t.Ticker.C:
			t.Func()
		}
	}
}
func (t *Task) Start() error {
	if t.isRun {
		return ErrTaskAlreadyStarterd
	}
	t.stop = make(chan TaskStopSignal, 1)
	t.isRun = true

	go t.run()

	return nil
}

func (t *Task) Stop() error {
	if !t.isRun {
		return ErrTaskNotStarted
	}

	t.stop <- TaskStopSignal{}

	t.isRun = false

	return nil
}
