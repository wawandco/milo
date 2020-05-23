package reviewers

import (
	"bufio"
	"io"
	"path/filepath"
	"strings"
)

const (
	DoctypeExpression = "<!doctype"
)

type DoctypePresent struct{}

func (doc DoctypePresent) ReviewerName() string {
	return "doctype/present"
}

func (doc DoctypePresent) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc DoctypePresent) Review(path string, page io.Reader) ([]Fault, error) {
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

		lineLower := strings.ToLower(line)
		prevLineLower := strings.ToLower(prevLine)

		if strings.Contains(lineLower, "<html") && strings.Contains(lineLower, DoctypeExpression) {
			break
		}

		if strings.Contains(lineLower, "<html") && strings.Contains(prevLineLower, DoctypeExpression) {
			break
		}

		if strings.Contains(lineLower, "<html") {
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
