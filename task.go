package tasker

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrTaskAlreadyStarterd error = errors.New("task already started")
	ErrTaskNotStarted      error = errors.New("task not started")
)

type TaskStopSignal struct{}

type Task struct {
	Name     string
	Func     func() // Function to be executed
	stop     chan TaskStopSignal
	Interval time.Duration
	m        *sync.Mutex
}

func NewTask(name string, interval time.Duration, function func()) *Task {
	return &Task{
		stop:     make(chan TaskStopSignal, 1),
		m:        &sync.Mutex{},
		Name:     name,
		Interval: interval,
		Func:     function,
	}
}

func (t *Task) run() {
	defer t.m.Unlock()

	for {
		select {
		case <-t.stop:
			return
		default:
			t.Func()
			time.Sleep(t.Interval)
		}
	}
}

func (t *Task) Start() error {
	if !t.m.TryLock() {
		return ErrTaskAlreadyStarterd
	}

	go t.run()

	return nil
}

func (t *Task) Stop() {
	t.stop <- TaskStopSignal{}
}
