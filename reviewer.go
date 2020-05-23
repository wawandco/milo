package milo

import (
	"io"
	"wawandco/milo/reviewers"
)

var _ Reviewer = (*reviewers.DoctypePresent)(nil)

type Reviewer interface {
	ReviewerName() string
	Accepts(fileName string) bool
	Review(path string, content io.Reader) ([]reviewers.Fault, error)
}
