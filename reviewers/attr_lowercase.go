package reviewers

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type AttrLowercase struct{}

func (a AttrLowercase) ReviewerName() string {
	return "attribute/lowercase"
}

func (a AttrLowercase) Accepts(path string) bool {
	return true
}

func (a AttrLowercase) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	var number int
	var line string

	exp := regexp.MustCompile(`<[^>]+\s+(.*[A-Z].*)=`)
	scanner := bufio.NewScanner(page)
	for scanner.Scan() {
		number++

		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if exp.MatchString(line) {
			result = append(result, Fault{
				Reviewer: a.ReviewerName(),
				Line:     number,
				Rule:     Rules["0013"],
				Path:     path,
			})
		}
	}

	return result, nil
}
