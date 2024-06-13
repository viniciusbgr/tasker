package tasker

import (
	"errors"
	"log"
	"time"
)

const (
	TaskAlreadyStarterd string = "task already started"
	TaskNotStarted      string = "task not started"
)

type Task struct {
	Name     string
	Interval time.Duration
	stopChan chan struct{}
	Func     func(stop chan<- struct{}, args ...any)
}

func (t *Task) Run(args ...any) error {
	if t.stopChan != nil {
		return errors.New(TaskAlreadyStarterd)
	}

	t.stopChan = make(chan struct{})

	ticker := time.NewTicker(t.Interval)

	for {
		select {
		case <-t.stopChan:
			ticker.Stop()

			return nil
		case <-ticker.C:
			log.Printf("Start execution task \"%s\"", t.Name)
			t.Func(t.stopChan, args...)
		}
	}
}

func (t *Task) Stop() error {
	if t.stopChan == nil {
		return errors.New(TaskNotStarted)
	}
	close(t.stopChan)

	return nil
}
