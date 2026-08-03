package reviewers

import (
	"io"

	"github.com/wawandco/milo/internal/html"
)

type AttributeStyleForbidden struct{}

func (sa AttributeStyleForbidden) ReviewerName() string {
	return "attribute-style-forbidden"
}

func (sa AttributeStyleForbidden) Accepts(filePath string) bool {
	return true
}

func (sa AttributeStyleForbidden) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	var svg svgTracker
	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()
		if svg.inSVG(tt, token) {
			continue
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			for _, attr := range token.Attr {
				if attr.Key != "style" {
					continue
				}

				result = append(result, Fault{
					Reviewer: sa.ReviewerName(),
					Line:     token.Line,
					Col:      attr.Col,
					Path:     path,
					Rule:     Rules[sa.ReviewerName()],
				})

				break
			}
		}
	}

	return result, nil
}
