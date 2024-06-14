package tasker_test

import (
	"github.com/viniciusbgr/tasker"
	"testing"
	"time"
)

func TestAllocsCreateTask(t *testing.T) {
	task := &tasker.Task{
		Name:     "Task test for allocs on start.",
		Interval: time.Second,
		Func: func(stop chan<- struct{}) {
			t.Log("Bingo!!")
		},
	}
	allocs := testing.AllocsPerRun(100, func() {
		if err := task.Start(); err != nil {
			t.Fatal(err)
		}

		time.Sleep(5 * time.Second)
		task.Stop()
	})

	if allocs != 0 {
		t.Fatalf("expected 0 allocs, got %v", allocs)
	}
}
