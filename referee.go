package milo

import (
	"bytes"
	"io"
	"io/ioutil"
	"wawandco/milo/reviewers"
)

type Referee struct {
	Reviewers []reviewers.Reviewer
}

func (r *Referee) Review(path string, reader io.Reader) ([]reviewers.Fault, error) {
	faults := []reviewers.Fault{}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return faults, err
	}

	for _, reviewer := range r.Reviewers {
		if !reviewer.Accepts(path) {
			continue
		}

		reader := bytes.NewReader(content)
		reviewerFaults, err := reviewer.Review(path, reader)
		if err != nil {
			return faults, err
		}

		faults = append(faults, reviewerFaults...)
	}

	return faults, nil
}

func NewReferee() *Referee {
	return &Referee{}
}
