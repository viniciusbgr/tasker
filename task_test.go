package tasker_test

import (
	"github.com/viniciusbgr/tasker"
	"testing"
	"time"
)

func Benchmark_Allocs_task(b *testing.B) {
	task := &tasker.Task{
		Name:     "Task bentchmark Allocs",
		Interval: time.Second,
		Func: func(stop chan<- struct{}, _ ...any) {
			b.Log("Bingo!!")
		},
	}
	go task.Run()

	time.Sleep(2 * time.Second)

	if err := task.Stop(); err != nil {
		panic(err)
	}
}
