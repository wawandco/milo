package reviewers

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type SrcEmpty struct{}

func (css SrcEmpty) ReviewerName() string {
	return "tag/src-empty"
}

func (css SrcEmpty) Accepts(path string) bool {
	return true
}

func (css SrcEmpty) Review(path string, reader io.Reader) ([]Fault, error) {
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

		rh := regexp.MustCompile(`.*<.*(src|href|data)`)
		re := regexp.MustCompile(`.*<.*(src|href|data)="[^"]+"`)
		if !rh.MatchString(line) || re.MatchString(line) {
			continue
		}

		result = append(result, Fault{
			Reviewer: css.ReviewerName(),
			Line:     number,
			Path:     path,
			Rule:     Rules["0007"],
		})
	}

	return result, nil
}
