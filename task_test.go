package tasker_test

import (
	"github.com/viniciusbgr/tasker"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	task := tasker.NewTask(
		"Test task",
		time.Millisecond*500,
		func() {
			t.Log("Hi everyone!")
		},
	)

	task.Start()

	time.Sleep(time.Second * 3)

	task.Stop()
}

func TestErrTask(t *testing.T) {
	task := tasker.NewTask(
		"Test task for expect err on Start",
		time.Millisecond*500,
		func() {
			t.Log("Hi everyone!")
		},
	)

	task.Start()

	time.Sleep(time.Second * 2)

	if err := task.Start(); err == nil {
		t.Fatal("Expected error: ", tasker.ErrTaskAlreadyStarterd.Error())
	}
}
