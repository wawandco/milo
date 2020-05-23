package milo

import (
	"io"
	"wawandco/milo/reviewers"
)

var _ Reviewer = (*reviewers.Doctype)(nil)

type Reviewer interface {
	ReviewerName() string
	Accepts(fileName string) bool
	Review(io.Reader) ([]reviewers.Fault, error)
}
