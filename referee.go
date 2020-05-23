package milo

import (
	"io"
	"wawandco/milo/reviewers"
)

type Referee struct {
	Reviewers []Reviewer
}

func (r *Referee) Review(io.Reader) ([]reviewers.Fault, error) {
	return []reviewers.Fault{}, nil
}

func NewReferee() *Referee {
	return &Referee{}
}
