package reviewers

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type InlineCSS struct{}

func (css InlineCSS) ReviewerName() string {
	return "css/inline"
}

func (css InlineCSS) Accepts(path string) bool {
	return true
}

func (css InlineCSS) Review(path string, reader io.Reader) ([]Fault, error) {
	result := []Fault{}
	var number int
	var line string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		number++

		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		lineLower := strings.ToLower(line)
		re := regexp.MustCompile(`<\w+[^>]*style=["'].*["'][^>]*>`)
		if !re.MatchString(lineLower) {
			continue
		}

		result = append(result, Fault{
			Reviewer: css.ReviewerName(),
			Line:     number,
			Path:     path,
			Rule:     Rules["0003"],
		})
	}

	return result, nil
}
