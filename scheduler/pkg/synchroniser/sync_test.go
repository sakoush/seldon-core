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
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestSimpleSynchroniser(t *testing.T) {
	g := NewGomegaWithT(t)

	type test struct {
		name    string
		timeout time.Duration
		signals uint
	}

	tests := []test{
		{
			name:    "Simple",
			timeout: 100 * time.Millisecond,
			signals: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewSimpleSynchroniser(test.timeout)
			s.Signals(test.signals)
			g.Expect(s.IsReady()).To(BeFalse())
			s.WaitReady()
			g.Expect(s.IsReady()).To(BeTrue())
		})
	}
}
