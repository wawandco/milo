package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type AttrLowercase struct{}

func (a AttrLowercase) ReviewerName() string {
	return "attribute/lowercase"
}

func (a AttrLowercase) Accepts(path string) bool {
	return true
}

func (a AttrLowercase) Review(path string, page io.Reader) ([]Fault, error) {
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

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			tok := z.Token()
			for _, attr := range tok.Attr {
				if attr.Name != attr.Key {
					result = append(result, Fault{
						Reviewer: a.ReviewerName(),
						Line:     tok.Line,
						Rule:     Rules["0013"],
						Path:     path,
					})
				}
			}
		}
	}

	return result, nil
}
