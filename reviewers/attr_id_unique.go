package reviewers

import (
	"io"
	"strings"

	"github.com/wawandco/milo/internal/goquery"
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

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return result, err
	}

	matched := doc.Find("*")
	IDCache := map[string]int{}

	for _, node := range matched.Nodes {
		for _, attr := range node.Attr {
			key := strings.ToLower(attr.Key)
			if key != "id" {
				continue
			}

			id := attr.Val
			if IDCache[id] == 0 {
				IDCache[id] = node.Line
				continue
			}

			result = append(result, Fault{
				Reviewer: at.ReviewerName(),
				Line:     node.Line,
				Path:     path,
				Rule:     Rules["0014"],
			})

			break
		}
	}

	return result, nil
}
