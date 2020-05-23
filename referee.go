package milo

import (
	"io"
	"wawandco/milo/reviewers"
)

type Referee struct {
	Reviewers []Reviewer
}

func (r *Referee) Review(path string, reader io.Reader) ([]reviewers.Fault, error) {
	faults := []reviewers.Fault{}
	for _, reviewer := range r.Reviewers {
		if !reviewer.Accepts(path) {
			continue
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
