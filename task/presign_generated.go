// Code generated by go generate; DO NOT EDIT.
package task

import (
	"fmt"

	"github.com/Xuanwo/navvy"
	"github.com/google/uuid"

	"github.com/yunify/qsctl/v2/pkg/types"
	"github.com/yunify/qsctl/v2/utils"
)

var _ navvy.Pool
var _ types.Pool
var _ = utils.SubmitNextTask
var _ = uuid.New()

// presignTaskRequirement is the requirement for execute PresignTask.
type presignTaskRequirement interface {
	navvy.Task

	// Inherited value
}

// mockPresignTask is the mock task for PresignTask.
type mockPresignTask struct {
	types.Todo
	types.Pool
	types.Fault
	types.ID

	// Inherited value
}

func (t *mockPresignTask) Run() {
	panic("mockPresignTask should not be run.")
}

// PresignTask will will handle presign tasks.
type PresignTask struct {
	presignTaskRequirement

	// Predefined runtime value
	types.Fault
	types.ID
	types.Todo

	// Runtime value
	types.BucketName
	types.Expire
	types.Key
	types.Pool
	types.Storage
	types.URL
}

// Run implement navvy.Task
func (t *PresignTask) Run() {
	if t.ValidateFault() {
		return
	}
	utils.SubmitNextTask(t)
}

func (t *PresignTask) TriggerFault(err error) {
	t.SetFault(fmt.Errorf("Task Presign failed: {%w}", err))
}

// Wait will wait until PresignTask has been finished
func (t *PresignTask) Wait() {
	t.GetPool().Wait()
}
