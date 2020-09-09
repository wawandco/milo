package reviewers

import (
	"io"

	"github.com/wawandco/milo/external/html"
)

// AttrIDUnique is a reviewer that checks that all tags have only an id tag.
type AttrIDUnique struct{}

// ReviewerName returns the reviewer name.
func (at AttrIDUnique) ReviewerName() string {
	return "attribute/id-unique"
}

// Accepts checks if the file can be reviewed.
func (at AttrIDUnique) Accepts(filePath string) bool {
	return true
}

// Review return a fault for each tag that has 2 or more id tags.
// For example, tags like <div id="..."> is correct.
// For tags like <div id="..." id="...">, will return 1 fault.
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
			if ID == "" {
				continue
			}

			if ID != "" && !IDs[ID] {
				IDs[ID] = true

				continue
			}

			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     token.Line,
				Col:      token.Col,
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
