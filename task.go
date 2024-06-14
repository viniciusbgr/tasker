package tasker

import (
	"errors"
	"time"
)

var (
	TaskAlreadyStarterd error = errors.New("task already started")
	TaskNotStarted      error = errors.New("task not started")
)

type Task struct {
	Name     string
	Interval time.Duration
	stopChan chan struct{}
	Func     func(stop chan<- struct{})
}

func (t *Task) run() {
	ticker := time.NewTicker(t.Interval)

	for {
		select {
		case <-t.stopChan:
			ticker.Stop()
			break
		case <-ticker.C:
			t.Func(t.stopChan)
		}
	}
}

func (t *Task) Start() error {
	if t.stopChan != nil {
		_, isClosed := <-t.stopChan

		if !isClosed {
			return TaskAlreadyStarterd

		}
	}

	t.stopChan = make(chan struct{})

	go t.run()

	return nil
}

func (t *Task) Stop() error {
	if t.stopChan == nil {
		return TaskNotStarted
	}
	close(t.stopChan)

	return nil
}
