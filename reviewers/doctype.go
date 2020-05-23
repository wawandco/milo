package reviewers

import "io"

type Doctype struct{}

func (doc Doctype) ReviewerName() string {
	return "Doctype Reviewer"
}

func (doc Doctype) Accepts(fileName string) bool {
	return false
}

func (doc Doctype) Review(page io.Reader) ([]Fault, error) {
	return []Fault{}, nil
}
