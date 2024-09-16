/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed by
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

// This package is responsible to synchronise starting up the different components of the "scheduler".
// In particular, it is responsible for making sure that the time between the scheduler starts and while
// the different model servers connect that the data plane (inferences) is not affected.
package synchroniser

import (
	"sync"
	"sync/atomic"
	"time"
)

type Synchroniser interface {
	IsReady() bool
	WaitReady()
	Signals(uint)
}

type SimpleSynchroniser struct {
	isReady atomic.Bool
	wg      sync.WaitGroup
	timeout time.Duration
}

func NewSimpleSynchroniser(timeout time.Duration) *SimpleSynchroniser {
	s := &SimpleSynchroniser{
		isReady: atomic.Bool{},
		wg:      sync.WaitGroup{},
		timeout: timeout,
	}
	s.isReady.Store(false)
	return s
}

func (s *SimpleSynchroniser) IsReady() bool {
	return s.isReady.Load()
}

func (s *SimpleSynchroniser) Signals(numSignals uint) {
	s.wg.Add(int(numSignals))
	time.AfterFunc(s.timeout, s.done)
}

func (s *SimpleSynchroniser) WaitReady() {
	s.wg.Wait()
}

func (s *SimpleSynchroniser) done() {
	s.isReady.Store(true)
	s.wg.Done()
}
