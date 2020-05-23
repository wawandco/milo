package reviewers

import (
	"io"
	"strings"
)

type Doctype struct{}

func (doc Doctype) ReviewerName() string {
	return "Doctype Reviewer"
}

func (doc Doctype) Accepts(fileName string) bool {
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc Doctype) Review(page io.Reader) ([]Fault, error) {
	return []Fault{}, nil
}
