package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type OlUlValid struct{}

func (ol OlUlValid) ReviewerName() string {
	return "ol-ul/valid"
}

func (ol OlUlValid) Accepts(filePath string) bool {
	return true
}

func (ol OlUlValid) Review(path string, page io.Reader) ([]Fault, error) {
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
					Path:     path,
					Rule:     Rules[ol.ReviewerName()],
				})
			}

			parents = append([]html.Token{token}, parents...)

		case html.EndTagToken:
			parents = parents[1:]
		}
	}

	return result, nil
}

func (ol OlUlValid) isList(token html.Token) bool {
	if token.Data != "ul" && token.Data != "ol" {
		return false
	}

	return true
}
