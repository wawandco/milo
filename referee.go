package milo

import "io"

type Referee struct {
	Reviewers []Reviewer
}

func (r *Referee) Review(io.Reader) ([]Fault, error) {
	return []Fault{}, nil
}

func NewReferee() *Referee {
	return &Referee{}
}
