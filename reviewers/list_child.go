package reviewers

import (
	"io"
)

type ListChild struct{}

func (doc ListChild) ReviewerName() string {
	return "list/child"
}

func (doc ListChild) Accepts(filePath string) bool {
	return true
}

func (doc ListChild) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	return result, nil
}
