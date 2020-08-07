package reviewers

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type StyleTag struct{}

func (css StyleTag) ReviewerName() string {
	return "style/tag-present"
}

func (css StyleTag) Accepts(path string) bool {
	return true
}

func (css StyleTag) Review(path string, reader io.Reader) ([]Fault, error) {
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
		re := regexp.MustCompile(`.*<style[^>]*>`)
		if !re.MatchString(lineLower) {
			continue
		}

		result = append(result, Fault{
			Reviewer: css.ReviewerName(),
			Line:     number,
			Path:     path,
			Rule:     Rules[css.ReviewerName()],
		})
	}

	return result, nil
}
