// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter_test

import (
	"fmt"
	"time"

	"github.com/canonical/pebble/client"
	"github.com/juju/clock/testclock"
	"github.com/juju/loggo"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v3"
	"github.com/juju/worker/v3/workertest"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/testing"
	"github.com/juju/juju/worker/uniter"
	"github.com/juju/juju/worker/uniter/container"
)

type pebbleNoticerSuite struct {
	clock             *testclock.Clock
	worker            worker.Worker
	clients           map[string]*fakePebbleClient
	workloadEventChan chan string
	workloadEvents    container.WorkloadEvents
}

var _ = gc.Suite(&pebbleNoticerSuite{})

func (s *pebbleNoticerSuite) setUpWorker(c *gc.C, containerNames []string) {
	s.clock = testclock.NewClock(time.Time{})
	s.workloadEventChan = make(chan string)
	s.workloadEvents = container.NewWorkloadEvents()
	s.clients = make(map[string]*fakePebbleClient)
	for _, name := range containerNames {
		s.clients[name] = &fakePebbleClient{
			clock:       testclock.NewClock(time.Time{}),
			noticeAdded: make(chan *client.Notice, 1), // buffered so AddNotice doesn't block
		}
	}
	newClient := func(cfg *client.Config) (uniter.PebbleClient, error) {
		c.Assert(cfg.Socket, gc.Matches, pebbleSocketPathRegexpString)
		matches := pebbleSocketPathRegexp.FindAllStringSubmatch(cfg.Socket, 1)
		return s.clients[matches[0][1]], nil
	}
	s.worker = uniter.NewPebbleNoticer(loggo.GetLogger("test"), s.clock, containerNames, s.workloadEventChan, s.workloadEvents, newClient)
}

func (s *pebbleNoticerSuite) waitWorkloadEvent(c *gc.C, expected container.WorkloadEvent) {
	select {
	case id := <-s.workloadEventChan:
		event, callback, err := s.workloadEvents.GetWorkloadEvent(id)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(event, gc.DeepEquals, expected)
		callback(nil)
	case <-time.After(testing.LongWait):
		c.Fatalf("timed out waiting for event")
	}
}

func (s *pebbleNoticerSuite) TestWaitNotices(c *gc.C) {
	s.setUpWorker(c, []string{"c1"})
	defer workertest.CleanKill(c, s.worker)

	// Simulate a WaitNotices timeout.
	s.clients["c1"].clock.WaitAdvance(30*time.Second, testing.ShortWait, 1)

	// The first notice will always be handled.
	lastRepeated := time.Now()
	s.clients["c1"].AddNotice(c, &client.Notice{
		ID:           "1",
		Type:         "custom",
		Key:          "a.b/c",
		LastRepeated: lastRepeated,
	})
	s.waitWorkloadEvent(c, container.WorkloadEvent{
		Type:         container.CustomNoticeEvent,
		WorkloadName: "c1",
		NoticeID:     "1",
		NoticeType:   "custom",
		NoticeKey:    "a.b/c",
	})

	// Another notice with an earlier LastRepeated time will be skipped.
	s.clients["c1"].AddNotice(c, &client.Notice{
		ID:           "2",
		Type:         "custom",
		Key:          "a.b/d",
		LastRepeated: lastRepeated.Add(-time.Second),
	})
	select {
	case <-s.workloadEventChan:
		c.Fatalf("shouldn't see this notice")
	case <-time.After(testing.ShortWait):
	}

	// A new notice with a later LastRepeated time will be handled.
	s.clients["c1"].AddNotice(c, &client.Notice{
		ID:           "3",
		Type:         "custom",
		Key:          "a.b/e",
		LastRepeated: lastRepeated.Add(time.Second),
	})
	s.waitWorkloadEvent(c, container.WorkloadEvent{
		Type:         container.CustomNoticeEvent,
		WorkloadName: "c1",
		NoticeID:     "3",
		NoticeType:   "custom",
		NoticeKey:    "a.b/e",
	})
}

func (s *pebbleNoticerSuite) TestWaitNoticesError(c *gc.C) {
	s.setUpWorker(c, []string{"c1"})
	defer workertest.CleanKill(c, s.worker)

	s.clients["c1"].AddNotice(c, &client.Notice{
		ID:   "1",
		Type: "error",
		Key:  "WaitNotices error!",
	})
	s.clock.WaitAdvance(testing.LongWait, time.Second, 1)

	s.clients["c1"].AddNotice(c, &client.Notice{
		ID:   "2",
		Type: "custom",
		Key:  "a.b/c",
	})
	s.waitWorkloadEvent(c, container.WorkloadEvent{
		Type:         container.CustomNoticeEvent,
		WorkloadName: "c1",
		NoticeID:     "2",
		NoticeType:   "custom",
		NoticeKey:    "a.b/c",
	})
}

func (s *pebbleNoticerSuite) TestIgnoreUnhandledType(c *gc.C) {
	s.setUpWorker(c, []string{"c1"})
	defer workertest.CleanKill(c, s.worker)

	s.clients["c1"].AddNotice(c, &client.Notice{
		ID:   "1",
		Type: "unhandled",
		Key:  "some-key",
	})

	select {
	case <-s.workloadEventChan:
		c.Fatalf("should ignore this notice")
	case <-time.After(testing.ShortWait):
	}
}

func (s *pebbleNoticerSuite) TestMultipleContainers(c *gc.C) {
	s.setUpWorker(c, []string{"c1", "c2"})
	defer workertest.CleanKill(c, s.worker)

	for i := 1; i <= 2; i++ {
		name := fmt.Sprintf("c%d", i)
		s.clients[name].AddNotice(c, &client.Notice{
			ID:   "1",
			Type: "custom",
			Key:  "example.com/" + name,
		})
		s.waitWorkloadEvent(c, container.WorkloadEvent{
			Type:         container.CustomNoticeEvent,
			WorkloadName: name,
			NoticeID:     "1",
			NoticeType:   "custom",
			NoticeKey:    "example.com/" + name,
		})
	}
}
