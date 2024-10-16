/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed by
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

package experiment

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

func TestSaveAndRestore(t *testing.T) {
	g := NewGomegaWithT(t)
	type test struct {
		name        string
		experiments []*Experiment
	}

	tests := []test{
		{
			name: "basic model experiment",
			experiments: []*Experiment{
				{
					Name: "test1",
					Candidates: []*Candidate{
						{
							Name:   "model1",
							Weight: 50,
						},
						{
							Name:   "model2",
							Weight: 50,
						},
					},
					Mirror: &Mirror{
						Name:    "model3",
						Percent: 90,
					},
					Config: &Config{
						StickySessions: true,
					},
					KubernetesMeta: &KubernetesMeta{
						Namespace:  "default",
						Generation: 2,
					},
				},
			},
		},
		{
			name: "basic pipeline experiment",
			experiments: []*Experiment{
				{
					Name:         "test1",
					ResourceType: PipelineResourceType,
					Candidates: []*Candidate{
						{
							Name:   "pipeline1",
							Weight: 50,
						},
						{
							Name:   "pipeline2",
							Weight: 50,
						},
					},
					Mirror: &Mirror{
						Name:    "pipeline3",
						Percent: 90,
					},
					Config: &Config{
						StickySessions: true,
					},
					KubernetesMeta: &KubernetesMeta{
						Namespace:  "default",
						Generation: 2,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := fmt.Sprintf("%s/db", t.TempDir())
			logger := log.New()
			db, err := newExperimentDbManager(getExperimentDbFolder(path), logger)
			g.Expect(err).To(BeNil())
			for _, p := range test.experiments {
				err := db.save(p)
				g.Expect(err).To(BeNil())
			}
			err = db.Stop()
			g.Expect(err).To(BeNil())

			es := NewExperimentServer(log.New(), nil, nil, nil)
			err = es.InitialiseOrRestoreDB(path)
			g.Expect(err).To(BeNil())
			for _, p := range test.experiments {
				g.Expect(cmp.Equal(p, es.experiments[p.Name])).To(BeTrue())
			}
		})
	}
}
