// Code generated by go generate; DO NOT EDIT.
package types

import (
	"github.com/Xuanwo/navvy"
)

// PathRequirement is the requirement for PathTask.
type PathRequirement interface {
	navvy.Task

	// Value
	PathGetter
	PathSetter
	PathValidator
	StorageGetter
	StorageSetter
	StorageValidator
}

type pathScheduleFunc func(navvy.Task) PathRequirement

// SegmentIDRequirement is the requirement for SegmentIDTask.
type SegmentIDRequirement interface {
	navvy.Task

	// Value
	SegmentIDGetter
	SegmentIDSetter
	SegmentIDValidator
	StorageGetter
	StorageSetter
	StorageValidator
}

type segmentIDScheduleFunc func(navvy.Task) SegmentIDRequirement
