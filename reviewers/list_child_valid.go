package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type ListChildValid struct{}

func (ol ListChildValid) ReviewerName() string {
	return "list/child-valid"
}

func (ol ListChildValid) Accepts(filePath string) bool {
	return true
}

func (ol ListChildValid) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	var parents []html.Token
	z := html.NewTokenizer(page)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()
		switch tt {
		case html.StartTagToken:
			if len(parents) > 0 && ol.isList(parents[0]) && token.Data != "li" {
				result = append(result, Fault{
					Reviewer: ol.ReviewerName(),
					Line:     token.Line,
					Col:      token.Col,
					Path:     path,
					Rule:     Rules[ol.ReviewerName()],
				})
			}

			parents = append([]html.Token{token}, parents...)

		case html.EndTagToken:
			if len(parents) == 0 {
				continue
			}

			parents = parents[1:]
		}
	}

	return result, nil
}

func (ol ListChildValid) isList(token html.Token) bool {
	if token.Data != "ul" && token.Data != "ol" {
		return false
	}

	return true
}
