package reviewers

import (
	"io"
	"regexp"

	"github.com/wawandco/milo/external/html"
)

type AttrNoWhiteSpaces struct{}

func (at AttrNoWhiteSpaces) ReviewerName() string {
	return "attribute/no-white-spaces"
}

func (at AttrNoWhiteSpaces) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	exp := regexp.MustCompile(`\S+(\s+=\s*|\s*=\s+)`)

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}
			return []Fault{}, err
		}

		if tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		if exp.MatchString(string(z.Raw())) {
			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     tok.Line,
				Col:      tok.Col,
				Path:     path,
				Rule:     Rules[at.ReviewerName()],
			})
		}
	}

	return result, nil
}
