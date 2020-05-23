package milo

import "io"

type Reviewer interface {
	ReviewerName() string
	Review(io.ReadCloser) ([]Fault, error)
}
