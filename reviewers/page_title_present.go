package reviewers

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/wawandco/milo/internal/html"
)

// PageTitlePresent is a reviewer that checks if the HTML file the title tag.
type PageTitlePresent struct{}

// ReviewerName returns the reviewer name.
func (doc PageTitlePresent) ReviewerName() string {
	return "page-title-present"
}

// Accepts checks if the file can be reviewed.
func (doc PageTitlePresent) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")

	return !isPartial
}

// Review return a fault if file does not have a <title> tag.
func (doc PageTitlePresent) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	var htmlTag *html.Token
	var startTag *html.Token
	var endTag *html.Token
	var content string

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()
		tag := token.DataAtom.String()

		switch tt {
		case html.StartTagToken:
			if htmlTag == nil && tag == "html" {
				htmlTag = &token

				continue
			}

			if tag == "title" {
				startTag = &token
			}
		case html.EndTagToken:
			if tag == "title" {
				endTag = &token
			}

		case html.TextToken:
			if htmlTag == nil {
				continue
			}

			if startTag == nil {
				continue
			}

			if content != "" {
				continue
			}

			content = strings.TrimSpace(string(z.Raw()))
		}
	}

	if htmlTag == nil || (content != "" && endTag != nil) {
		return result, nil
	}

	result = append(result, Fault{
		Reviewer: doc.ReviewerName(),
		Line:     1,
		Col:      1,
		Path:     path,
		Rule:     Rules[doc.ReviewerName()],
	})

	return result, nil
}
