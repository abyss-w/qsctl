package task

import (
	"errors"

	"github.com/Xuanwo/storage/pkg/iterator"
	"github.com/Xuanwo/storage/pkg/segment"
	typ "github.com/Xuanwo/storage/types"
	log "github.com/sirupsen/logrus"

	"github.com/yunify/qsctl/v2/pkg/types"
)

func (t *ListFileTask) new() {
	oc := make(chan *typ.Object)
	t.SetObjectChannel(oc)
}

func (t *ListFileTask) run() {
	log.Debugf("Task <%s> for key <%s> started", "ListFileTask", t.GetPath())

	it := t.GetStorage().ListDir(t.GetPath(), typ.WithRecursive(t.GetRecursive()))

	// Always close the object channel.
	defer close(t.GetObjectChannel())

	for {
		o, err := it.Next()
		if err != nil && errors.Is(err, iterator.ErrDone) {
			break
		}
		if err != nil {
			t.TriggerFault(types.NewErrUnhandled(err))
			return
		}
		t.GetObjectChannel() <- o
	}

	log.Debugf("Task <%s> for key <%s> finished", "ListFileTask", t.GetPath())
}

func (t *ListSegmentTask) new() {
	sc := make(chan *segment.Segment)
	t.SetSegmentChannel(sc)
}

func (t *ListSegmentTask) run() {
	log.Debugf("Task <%s> for key <%s> started", "ListSegmentTask", t.GetPath())

	it := t.GetStorage().ListSegments(t.GetPath())

	// Always close the segment channel.
	defer close(t.GetSegmentChannel())

	for {
		o, err := it.Next()
		if err != nil && errors.Is(err, iterator.ErrDone) {
			break
		}
		if err != nil {
			t.TriggerFault(types.NewErrUnhandled(err))
			return
		}
		t.GetSegmentChannel() <- o
	}

	log.Debugf("Task <%s> for key <%s> finished", "ListSegmentTask", t.GetPath())
}

func (t *ListStorageTask) new() {}
func (t *ListStorageTask) run() {
	resp, err := t.GetService().List(typ.WithLocation(t.GetZone()))
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
	buckets := make([]string, 0, len(resp))
	for _, v := range resp {
		b, err := v.Metadata()
		if err != nil {
			t.TriggerFault(types.NewErrUnhandled(err))
			return
		}
		if name, ok := b.GetName(); ok {
			buckets = append(buckets, name)
		}
	}
	t.SetBucketList(buckets)
	log.Debugf("Task <%s> in zone <%s> finished.", "BucketListTask", t.GetZone())
}
