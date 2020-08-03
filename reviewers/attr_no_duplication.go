package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type AttrNoDuplication struct{}

func (a AttrNoDuplication) ReviewerName() string {
	return "attribute/no-duplication"
}

func (a AttrNoDuplication) Accepts(path string) bool {
	return true
}

func (a AttrNoDuplication) Review(path string, r io.Reader) ([]Fault, error) {
	var fault []Fault
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		tok := z.Token()
		keys := map[string]bool{}
		for _, attr := range tok.Attr {
			if keys[attr.Key] {
				fault = append(fault, Fault{
					Reviewer: a.ReviewerName(),
					Line:     tok.Line,
					Path:     path,
					Rule:     Rules["0010"],
				})
			}

			keys[attr.Key] = true
		}
	}

	return fault, nil
}
