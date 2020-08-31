package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
	"github.com/wawandco/milo/external/html/atom"
)

type StyleTag struct{}

func (css StyleTag) ReviewerName() string {
	return "style/tag-present"
}

func (css StyleTag) Accepts(path string) bool {
	return true
}

func (css StyleTag) Review(path string, page io.Reader) ([]Fault, error) {
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

		if tt != html.StartTagToken && tt != html.SelfClosingTagToken {
			continue
		}

		tok := z.Token()
		if tok.DataAtom != atom.Style {
			continue
		}

		result = append(result, Fault{
			Reviewer: css.ReviewerName(),
			Line:     tok.Line,
			Col:      tok.Col,
			Path:     path,
			Rule:     Rules[css.ReviewerName()],
		})
	}

	return result, nil
}
