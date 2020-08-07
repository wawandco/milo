package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

type AttrIDUnique struct{}

func (at AttrIDUnique) ReviewerName() string {
	return "attribute/alt-required"
}

func (at AttrIDUnique) Accepts(filePath string) bool {
	return true
}

func (at AttrIDUnique) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	IDs := map[string]bool{}

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			token := z.Token()

			ID := at.tagID(token)
			if ID != "" && !IDs[ID] {
				IDs[ID] = true
				continue
			}

			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     token.Line,
				Path:     path,
				Rule:     Rules[at.ReviewerName()],
			})
		}
	}

	return result, nil
}

func (at AttrIDUnique) tagID(token html.Token) string {
	for _, attr := range token.Attr {
		if attr.Key == "id" && attr.Val != "" {
			return attr.Val
		}
	}

	return ""
}
