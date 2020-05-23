package milo

import (
	"os"
	"wawandco/milo/reviewers"
)

type Referee struct {
	Reviewers []Reviewer
}

func (r *Referee) Review(path string) ([]reviewers.Fault, error) {
	faults := []reviewers.Fault{}
	for _, reviewer := range r.Reviewers {
		if !reviewer.Accepts(path) {
			continue
		}

		reader, err := os.Open(path)
		if err != nil {
			return faults, err
		}

		rfaults, err := reviewer.Review(path, reader)
		if err != nil {
			return faults, err
		}

		faults = append(faults, rfaults...)
	}

	return faults, nil
}

func NewReferee() *Referee {
	return &Referee{}
}
