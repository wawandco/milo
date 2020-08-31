package reviewers

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/wawandco/milo/external/html"
)

const (
	DoctypeExpression = "<!doctype"
)

type DoctypePresent struct{}

func (doc DoctypePresent) ReviewerName() string {
	return "doctype/present"
}

func (doc DoctypePresent) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc DoctypePresent) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}
	var (
		found   bool
		htmlTag *html.Token
	)

	z := html.NewTokenizer(page)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()
		if token.DataAtom.String() == "html" {
			htmlTag = &token
			continue
		}

		if tt == html.DoctypeToken {
			found = true
		}
	}

	if htmlTag == nil || found {
		return result, nil
	}

	result = append(result, Fault{
		Reviewer: doc.ReviewerName(),
		Line:     htmlTag.Line,
		Col:      htmlTag.Col,
		Path:     path,
		Rule:     Rules[doc.ReviewerName()],
	})

	return result, nil
}
