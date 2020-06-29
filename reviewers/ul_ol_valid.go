package reviewers

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type OlUlValid struct{}

func (ol OlUlValid) ReviewerName() string {
	return "ol/valid"
}

func (ol OlUlValid) Accepts(filePath string) bool {
	return true
}

func (ol OlUlValid) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return result, err
	}

	olulSelection := doc.Find("ol").AddSelection(doc.Find("ul"))
	notolul := olulSelection.Children().Not("li")

	for i := 0; i < notolul.Length(); i++ {
		result = append(result, Fault{
			Reviewer: ol.ReviewerName(),
			Line:     i + 1,
			Path:     path,
			Rule:     Rules["0008"],
		})
	}
	return result, nil
}
