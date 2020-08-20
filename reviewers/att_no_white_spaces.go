package reviewers

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type AttrNoWhiteSpaces struct{}

func (at AttrNoWhiteSpaces) ReviewerName() string {
	return "attribute/no-white-spaces"
}

func (at AttrNoWhiteSpaces) Review(path string, page io.Reader) ([]Fault, error) {
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

		exp := regexp.MustCompile(`\S+(\s+=\s*|\s*=\s+)`)
		if exp.MatchString(line) {
			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     number,
				Rule:     Rules[at.ReviewerName()],
				Path:     path,
			})
		}
	}

	return result, nil
}
