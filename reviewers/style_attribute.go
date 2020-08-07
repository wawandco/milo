package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type StyleAttribute struct{}

func (sa StyleAttribute) ReviewerName() string {
	return "attribute/style"
}

func (sa StyleAttribute) Accepts(filePath string) bool {
	return true
}

func (sa StyleAttribute) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			token := z.Token()

			for _, attr := range token.Attr {
				if attr.Key != "style" {
					continue
				}

				result = append(result, Fault{
					Reviewer: sa.ReviewerName(),
					Line:     token.Line,
					Path:     path,
					Rule:     Rules["0009"],
				})

				break
			}
		}
	}

	return result, nil
}
