package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type InlineCSS struct{}

func (css InlineCSS) ReviewerName() string {
	return "css/inline"
}

func (css InlineCSS) Accepts(path string) bool {
	return true
}

func (css InlineCSS) Review(path string, page io.Reader) ([]Fault, error) {
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
		for _, attr := range tok.Attr {
			if attr.Key == "style" {
				result = append(result, Fault{
					Reviewer: css.ReviewerName(),
					Line:     attr.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[css.ReviewerName()],
				})
			}
		}
	}

	return result, nil
}
