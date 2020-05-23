package reviewers

import (
	"bufio"
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
	result := []Fault{}

	reader := bufio.NewReader(page)
	firstLine, _, err := reader.ReadLine()
	if err != nil {
		return result, err
	}

	if !strings.Contains(string(firstLine), "<!DOCTYPE html>") {
		result = append(result, Fault{
			ReviewerName: doc.ReviewerName(),
			LineNumber:   1,
			LineContent:  string(firstLine),

			Rule: Rule{
				Name: "Missing Doctype",
				Code: "0001",
			},
		})
	}

	return result, nil
}
