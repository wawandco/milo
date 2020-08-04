package reviewers

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type TagLowercase struct{}

func (css TagLowercase) ReviewerName() string {
	return "tag/lowercase"
}

func (css TagLowercase) Accepts(path string) bool {
	return true
}

func (css TagLowercase) Review(path string, reader io.Reader) ([]Fault, error) {
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

		re := regexp.MustCompile(`.*<[a-zA-Z]?[A-Z]+.*>`)
		if !re.MatchString(line) {
			continue
		}

		result = append(result, Fault{
			Reviewer: css.ReviewerName(),
			Line:     number,
			Path:     path,
			Rule:     Rules["0006"],
		})
	}

	return result, nil
}
