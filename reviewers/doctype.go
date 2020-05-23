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
	var prevLine, line string
	var number int

	scanner := bufio.NewScanner(page)
	for scanner.Scan() {
		number++

		line = scanner.Text()

		if strings.Contains(strings.ToLower(line), "<html") && strings.Contains(prevLine, "<!DOCTYPE html>") {
			break
		}

		if strings.Contains(strings.ToLower(line), "<html") {
			result = append(result, Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   number,

				Rule: Rule{
					Name: "Missing Doctype",
					Code: "0001",
				},
			})
			break
		}

		prevLine = line
	}

	return result, nil
}
