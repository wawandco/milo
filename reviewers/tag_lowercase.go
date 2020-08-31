package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type TagLowercase struct{}

func (t TagLowercase) ReviewerName() string {
	return "tag/lowercase"
}

func (t TagLowercase) Accepts(path string) bool {
	return true
}

func (t TagLowercase) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}
			return []Fault{}, err
		}

		if tt != html.EndTagToken && tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		if tok.Data == tok.Name {
			continue
		}

		result = append(result, Fault{
			Reviewer: t.ReviewerName(),
			Line:     tok.Line,
			Col:      tok.Col,
			Path:     path,
			Rule:     Rules[t.ReviewerName()],
		})
	}

	return result, nil
}
