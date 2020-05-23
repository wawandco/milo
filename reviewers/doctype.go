package reviewers

import (
	"bufio"
	"io"
	"path/filepath"
	"strings"
)

const (
	DoctypeExpression = "<!DOCTYPE html>"
)

type Doctype struct{}

func (doc Doctype) ReviewerName() string {
	return "Doctype Reviewer"
}

func (doc Doctype) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc Doctype) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	var prevLine, line string
	var number int

	scanner := bufio.NewScanner(page)
	for scanner.Scan() {
		number++

		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(strings.ToLower(line), "<html") && strings.Contains(line, DoctypeExpression) {
			break
		}

		if strings.Contains(strings.ToLower(line), "<html") && strings.Contains(prevLine, DoctypeExpression) {
			break
		}

		if strings.Contains(strings.ToLower(line), "<html") {
			result = append(result, Fault{
				ReviewerName: doc.ReviewerName(),
				LineNumber:   number,
				Path:         path,

				Rule: Rule{
					Name:        "Missing Doctype",
					Description: "HTML pages must have a Doctype declaration",
					Code:        "0001",
				},
			})
			break
		}

		prevLine = line
	}

	return result, nil
}
