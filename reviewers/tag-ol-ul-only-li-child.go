package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

// TagOlUlOnlyLiChild is a reviewer that checks that the ol/ul tag have only li child tags.
type TagOlUlOnlyLiChild struct{}

// ReviewerName returns the reviewer name.
func (ol TagOlUlOnlyLiChild) ReviewerName() string {
	return "tag-ol-ul-only-li-child"
}

// Accepts checks if the file can be reviewed.
func (ol TagOlUlOnlyLiChild) Accepts(filePath string) bool {
	return true
}

// Review returns a fault for each ul/ol child tag other than li.
func (ol TagOlUlOnlyLiChild) Review(path string, page io.Reader) ([]Fault, error) {
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

			// Checking if its a void tag
			if html.VoidElements[token.Data] {
				continue
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

func (ol TagOlUlOnlyLiChild) isList(token html.Token) bool {
	if token.Data != "ul" && token.Data != "ol" {
		return false
	}

	return true
}
