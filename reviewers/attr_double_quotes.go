package reviewers

import (
	"bufio"
	"io"
	"strings"
)

type AttrDoubleQuotes struct{}

func (at AttrDoubleQuotes) ReviewerName() string {
	return "attribute/value-double-quotes"
}

func (at AttrDoubleQuotes) Accepts(filePath string) bool {
	return true
}

func (at AttrDoubleQuotes) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	var number int
	var line string

	scanner := bufio.NewScanner(page)
	for scanner.Scan() {
		number++

		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		// if strings.Contains(lineLower, "<html") {
		// 	result = append(result, Fault{
		// 		Reviewer: at.ReviewerName(),
		// 		Line:     number,
		// 		Path:     path,

		// 		Rule: Rules["0001"],
		// 	})
		// 	break
		// }
	}

	return result, nil
}
