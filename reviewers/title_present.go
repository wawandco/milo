package reviewers

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"

	"github.com/wawandco/milo/external/html"
)

type TitlePresent struct{}

func (doc TitlePresent) ReviewerName() string {
	return "title/present"
}

func (doc TitlePresent) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc TitlePresent) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	var htmlTag *html.Token
	var found bool
	var startBuff []byte

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

		if token.DataAtom.String() != "title" {
			continue
		}

		if tt == html.StartTagToken {
			startBuff = z.Buffered()
			continue
		}

		between := bytes.ReplaceAll(startBuff, z.Buffered(), []byte{})
		between = bytes.ReplaceAll(between, z.Raw(), []byte{})

		if len(between) == 0 {
			continue
		}

		found = true
	}

	if htmlTag == nil || found {
		return result, nil
	}

	result = append(result, Fault{
		Reviewer: doc.ReviewerName(),
		Line:     1,
		Path:     path,
		Rule:     Rules["0004"],
	})

	return result, nil
}
