package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/external/html"
)

type AttrValueDoubleQuotes struct{}

func (a AttrValueDoubleQuotes) Accepts(path string) bool {
	return true
}

func (a AttrValueDoubleQuotes) ReviewerName() string {
	return "attribute/double-quotes"
}

func (a AttrValueDoubleQuotes) Review(path string, r io.Reader) ([]Fault, error) {
	var fault []Fault

	z := html.NewTokenizer(r)
	for {
		tt := z.Next()

		if err := z.Err(); err != nil {
			if err == io.EOF {
				break
			}
			return []Fault{}, nil
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			token := z.Token()
			for _, attr := range token.Attr {
				if attr.Quote == "" && strings.TrimSpace(attr.Val) == "" {
					continue
				}

				if attr.Quote != "\"" {
					fault = append(fault, Fault{
						Reviewer: a.ReviewerName(),
						Line:     token.Line,
						Path:     path,
						Rule:     Rules[a.ReviewerName()],
					})
				}
			}
		}
	}

	return fault, nil
}
